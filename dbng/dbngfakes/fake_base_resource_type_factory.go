// This file was generated by counterfeiter
package dbngfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
)

type FakeBaseResourceTypeFactory struct {
	FindStub        func(name string) (*dbng.UsedBaseResourceType, bool, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		name string
	}
	findReturns struct {
		result1 *dbng.UsedBaseResourceType
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBaseResourceTypeFactory) Find(name string) (*dbng.UsedBaseResourceType, bool, error) {
	fake.findMutex.Lock()
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Find", []interface{}{name})
	fake.findMutex.Unlock()
	if fake.FindStub != nil {
		return fake.FindStub(name)
	}
	return fake.findReturns.result1, fake.findReturns.result2, fake.findReturns.result3
}

func (fake *FakeBaseResourceTypeFactory) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeBaseResourceTypeFactory) FindArgsForCall(i int) string {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return fake.findArgsForCall[i].name
}

func (fake *FakeBaseResourceTypeFactory) FindReturns(result1 *dbng.UsedBaseResourceType, result2 bool, result3 error) {
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 *dbng.UsedBaseResourceType
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBaseResourceTypeFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBaseResourceTypeFactory) recordInvocation(key string, args []interface{}) {
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

var _ dbng.BaseResourceTypeFactory = new(FakeBaseResourceTypeFactory)
