// This file was generated by counterfeiter
package radarfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/radar"
	"github.com/tedsuo/ifrit"
)

type FakeScanRunnerFactory struct {
	ScanResourceRunnerStub        func(lager.Logger, string) ifrit.Runner
	scanResourceRunnerMutex       sync.RWMutex
	scanResourceRunnerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	scanResourceRunnerReturns struct {
		result1 ifrit.Runner
	}
	ScanResourceTypeRunnerStub        func(lager.Logger, string) ifrit.Runner
	scanResourceTypeRunnerMutex       sync.RWMutex
	scanResourceTypeRunnerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	scanResourceTypeRunnerReturns struct {
		result1 ifrit.Runner
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScanRunnerFactory) ScanResourceRunner(arg1 lager.Logger, arg2 string) ifrit.Runner {
	fake.scanResourceRunnerMutex.Lock()
	fake.scanResourceRunnerArgsForCall = append(fake.scanResourceRunnerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ScanResourceRunner", []interface{}{arg1, arg2})
	fake.scanResourceRunnerMutex.Unlock()
	if fake.ScanResourceRunnerStub != nil {
		return fake.ScanResourceRunnerStub(arg1, arg2)
	}
	return fake.scanResourceRunnerReturns.result1
}

func (fake *FakeScanRunnerFactory) ScanResourceRunnerCallCount() int {
	fake.scanResourceRunnerMutex.RLock()
	defer fake.scanResourceRunnerMutex.RUnlock()
	return len(fake.scanResourceRunnerArgsForCall)
}

func (fake *FakeScanRunnerFactory) ScanResourceRunnerArgsForCall(i int) (lager.Logger, string) {
	fake.scanResourceRunnerMutex.RLock()
	defer fake.scanResourceRunnerMutex.RUnlock()
	return fake.scanResourceRunnerArgsForCall[i].arg1, fake.scanResourceRunnerArgsForCall[i].arg2
}

func (fake *FakeScanRunnerFactory) ScanResourceRunnerReturns(result1 ifrit.Runner) {
	fake.ScanResourceRunnerStub = nil
	fake.scanResourceRunnerReturns = struct {
		result1 ifrit.Runner
	}{result1}
}

func (fake *FakeScanRunnerFactory) ScanResourceTypeRunner(arg1 lager.Logger, arg2 string) ifrit.Runner {
	fake.scanResourceTypeRunnerMutex.Lock()
	fake.scanResourceTypeRunnerArgsForCall = append(fake.scanResourceTypeRunnerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ScanResourceTypeRunner", []interface{}{arg1, arg2})
	fake.scanResourceTypeRunnerMutex.Unlock()
	if fake.ScanResourceTypeRunnerStub != nil {
		return fake.ScanResourceTypeRunnerStub(arg1, arg2)
	}
	return fake.scanResourceTypeRunnerReturns.result1
}

func (fake *FakeScanRunnerFactory) ScanResourceTypeRunnerCallCount() int {
	fake.scanResourceTypeRunnerMutex.RLock()
	defer fake.scanResourceTypeRunnerMutex.RUnlock()
	return len(fake.scanResourceTypeRunnerArgsForCall)
}

func (fake *FakeScanRunnerFactory) ScanResourceTypeRunnerArgsForCall(i int) (lager.Logger, string) {
	fake.scanResourceTypeRunnerMutex.RLock()
	defer fake.scanResourceTypeRunnerMutex.RUnlock()
	return fake.scanResourceTypeRunnerArgsForCall[i].arg1, fake.scanResourceTypeRunnerArgsForCall[i].arg2
}

func (fake *FakeScanRunnerFactory) ScanResourceTypeRunnerReturns(result1 ifrit.Runner) {
	fake.ScanResourceTypeRunnerStub = nil
	fake.scanResourceTypeRunnerReturns = struct {
		result1 ifrit.Runner
	}{result1}
}

func (fake *FakeScanRunnerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.scanResourceRunnerMutex.RLock()
	defer fake.scanResourceRunnerMutex.RUnlock()
	fake.scanResourceTypeRunnerMutex.RLock()
	defer fake.scanResourceTypeRunnerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeScanRunnerFactory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ radar.ScanRunnerFactory = new(FakeScanRunnerFactory)
