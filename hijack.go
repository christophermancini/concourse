package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudfoundry-incubator/garden/warden"
	"github.com/concourse/glider/api/builds"
	"github.com/concourse/glider/routes"
	TurbineHijack "github.com/concourse/turbine/api/hijack"
	"github.com/kr/pty"
	"github.com/pkg/term"
	"github.com/tedsuo/rata"
)

func hijack(reqGenerator *rata.RequestGenerator) {
	argv := flag.Args()[1:]

	var path string
	var args []string

	switch len(argv) {
	case 0:
		path = "bash"
	case 1:
		path = argv[0]
	default:
		path = argv[0]
		args = argv[1:]
	}

	spec := warden.ProcessSpec{
		Path:       path,
		Args:       args,
		TTY:        true,
		Privileged: true,
	}

	buildsReq, err := reqGenerator.CreateRequest(
		routes.GetBuilds,
		nil,
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}

	buildsResp, err := http.DefaultClient.Do(buildsReq)
	if err != nil {
		log.Fatalln(err)
	}

	var builds []builds.Build
	err = json.NewDecoder(buildsResp.Body).Decode(&builds)
	if err != nil {
		log.Fatalln(err)
	}

	if len(builds) == 0 {
		println("no builds to hijack")
		os.Exit(1)
	}

	build := builds[0]

	payload, err := json.Marshal(spec)
	if err != nil {
		log.Fatalln(err)
	}

	hijackReq, err := reqGenerator.CreateRequest(
		routes.HijackBuild,
		rata.Params{"guid": build.Guid},
		bytes.NewBuffer(payload),
	)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.Dial("tcp", hijackReq.URL.Host)
	if err != nil {
		log.Fatalln("failed to dial glider:", err)
	}

	client := httputil.NewClientConn(conn, nil)

	resp, err := client.Do(hijackReq)
	if err != nil {
		log.Fatalln("failed to hijack:", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("bad response:", resp.Status)
		resp.Body.Close()
		resp.Write(os.Stderr)
		os.Exit(1)
	}

	cconn, cbr := client.Hijack()

	term, err := term.Open(os.Stdin.Name())
	if err != nil {
		log.Fatalln("failed to open terminal:", term)
	}

	err = term.SetRaw()
	if err != nil {
		log.Fatalln("failed to set raw:", term)
	}

	defer term.Restore()

	encoder := gob.NewEncoder(cconn)

	sendSize(encoder)

	resized := make(chan os.Signal, 10)
	signal.Notify(resized, syscall.SIGWINCH)

	go func() {
		for {
			<-resized
			sendSize(encoder)
		}
	}()

	go io.Copy(&stdinWriter{encoder}, term)

	io.Copy(term, cbr)
}

func sendSize(enc *gob.Encoder) {
	rows, cols, err := pty.Getsize(os.Stdin)
	if err == nil {
		enc.Encode(TurbineHijack.ProcessPayload{
			WindowSize: &TurbineHijack.WindowSize{
				Columns: cols,
				Rows:    rows,
			},
		})
	}
}

type stdinWriter struct {
	enc *gob.Encoder
}

func (w *stdinWriter) Write(d []byte) (int, error) {
	err := w.enc.Encode(TurbineHijack.ProcessPayload{
		Stdin: d,
	})
	if err != nil {
		return 0, err
	}

	return len(d), nil
}
