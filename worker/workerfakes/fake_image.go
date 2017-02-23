// This file was generated by counterfeiter
package workerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/worker"
)

type FakeImage struct {
	FetchForContainerStub        func(logger lager.Logger, container dbng.CreatingContainer) (worker.FetchedImage, error)
	fetchForContainerMutex       sync.RWMutex
	fetchForContainerArgsForCall []struct {
		logger    lager.Logger
		container dbng.CreatingContainer
	}
	fetchForContainerReturns struct {
		result1 worker.FetchedImage
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImage) FetchForContainer(logger lager.Logger, container dbng.CreatingContainer) (worker.FetchedImage, error) {
	fake.fetchForContainerMutex.Lock()
	fake.fetchForContainerArgsForCall = append(fake.fetchForContainerArgsForCall, struct {
		logger    lager.Logger
		container dbng.CreatingContainer
	}{logger, container})
	fake.recordInvocation("FetchForContainer", []interface{}{logger, container})
	fake.fetchForContainerMutex.Unlock()
	if fake.FetchForContainerStub != nil {
		return fake.FetchForContainerStub(logger, container)
	}
	return fake.fetchForContainerReturns.result1, fake.fetchForContainerReturns.result2
}

func (fake *FakeImage) FetchForContainerCallCount() int {
	fake.fetchForContainerMutex.RLock()
	defer fake.fetchForContainerMutex.RUnlock()
	return len(fake.fetchForContainerArgsForCall)
}

func (fake *FakeImage) FetchForContainerArgsForCall(i int) (lager.Logger, dbng.CreatingContainer) {
	fake.fetchForContainerMutex.RLock()
	defer fake.fetchForContainerMutex.RUnlock()
	return fake.fetchForContainerArgsForCall[i].logger, fake.fetchForContainerArgsForCall[i].container
}

func (fake *FakeImage) FetchForContainerReturns(result1 worker.FetchedImage, result2 error) {
	fake.FetchForContainerStub = nil
	fake.fetchForContainerReturns = struct {
		result1 worker.FetchedImage
		result2 error
	}{result1, result2}
}

func (fake *FakeImage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchForContainerMutex.RLock()
	defer fake.fetchForContainerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeImage) recordInvocation(key string, args []interface{}) {
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

var _ worker.Image = new(FakeImage)
