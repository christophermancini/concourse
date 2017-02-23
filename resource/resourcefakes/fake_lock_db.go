// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/db/lock"
	"github.com/concourse/atc/resource"
)

type FakeLockDB struct {
	GetTaskLockStub        func(logger lager.Logger, lockName string) (lock.Lock, bool, error)
	getTaskLockMutex       sync.RWMutex
	getTaskLockArgsForCall []struct {
		logger   lager.Logger
		lockName string
	}
	getTaskLockReturns struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLockDB) GetTaskLock(logger lager.Logger, lockName string) (lock.Lock, bool, error) {
	fake.getTaskLockMutex.Lock()
	fake.getTaskLockArgsForCall = append(fake.getTaskLockArgsForCall, struct {
		logger   lager.Logger
		lockName string
	}{logger, lockName})
	fake.recordInvocation("GetTaskLock", []interface{}{logger, lockName})
	fake.getTaskLockMutex.Unlock()
	if fake.GetTaskLockStub != nil {
		return fake.GetTaskLockStub(logger, lockName)
	}
	return fake.getTaskLockReturns.result1, fake.getTaskLockReturns.result2, fake.getTaskLockReturns.result3
}

func (fake *FakeLockDB) GetTaskLockCallCount() int {
	fake.getTaskLockMutex.RLock()
	defer fake.getTaskLockMutex.RUnlock()
	return len(fake.getTaskLockArgsForCall)
}

func (fake *FakeLockDB) GetTaskLockArgsForCall(i int) (lager.Logger, string) {
	fake.getTaskLockMutex.RLock()
	defer fake.getTaskLockMutex.RUnlock()
	return fake.getTaskLockArgsForCall[i].logger, fake.getTaskLockArgsForCall[i].lockName
}

func (fake *FakeLockDB) GetTaskLockReturns(result1 lock.Lock, result2 bool, result3 error) {
	fake.GetTaskLockStub = nil
	fake.getTaskLockReturns = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLockDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTaskLockMutex.RLock()
	defer fake.getTaskLockMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLockDB) recordInvocation(key string, args []interface{}) {
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

var _ resource.LockDB = new(FakeLockDB)
