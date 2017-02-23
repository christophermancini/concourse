// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/resource"
)

type FakeResourceOptions struct {
	IOConfigStub        func() resource.IOConfig
	iOConfigMutex       sync.RWMutex
	iOConfigArgsForCall []struct{}
	iOConfigReturns     struct {
		result1 resource.IOConfig
	}
	SourceStub        func() atc.Source
	sourceMutex       sync.RWMutex
	sourceArgsForCall []struct{}
	sourceReturns     struct {
		result1 atc.Source
	}
	ParamsStub        func() atc.Params
	paramsMutex       sync.RWMutex
	paramsArgsForCall []struct{}
	paramsReturns     struct {
		result1 atc.Params
	}
	VersionStub        func() atc.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 atc.Version
	}
	ResourceTypeStub        func() resource.ResourceType
	resourceTypeMutex       sync.RWMutex
	resourceTypeArgsForCall []struct{}
	resourceTypeReturns     struct {
		result1 resource.ResourceType
	}
	LockNameStub        func(workerName string) (string, error)
	lockNameMutex       sync.RWMutex
	lockNameArgsForCall []struct {
		workerName string
	}
	lockNameReturns struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceOptions) IOConfig() resource.IOConfig {
	fake.iOConfigMutex.Lock()
	fake.iOConfigArgsForCall = append(fake.iOConfigArgsForCall, struct{}{})
	fake.recordInvocation("IOConfig", []interface{}{})
	fake.iOConfigMutex.Unlock()
	if fake.IOConfigStub != nil {
		return fake.IOConfigStub()
	}
	return fake.iOConfigReturns.result1
}

func (fake *FakeResourceOptions) IOConfigCallCount() int {
	fake.iOConfigMutex.RLock()
	defer fake.iOConfigMutex.RUnlock()
	return len(fake.iOConfigArgsForCall)
}

func (fake *FakeResourceOptions) IOConfigReturns(result1 resource.IOConfig) {
	fake.IOConfigStub = nil
	fake.iOConfigReturns = struct {
		result1 resource.IOConfig
	}{result1}
}

func (fake *FakeResourceOptions) Source() atc.Source {
	fake.sourceMutex.Lock()
	fake.sourceArgsForCall = append(fake.sourceArgsForCall, struct{}{})
	fake.recordInvocation("Source", []interface{}{})
	fake.sourceMutex.Unlock()
	if fake.SourceStub != nil {
		return fake.SourceStub()
	}
	return fake.sourceReturns.result1
}

func (fake *FakeResourceOptions) SourceCallCount() int {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return len(fake.sourceArgsForCall)
}

func (fake *FakeResourceOptions) SourceReturns(result1 atc.Source) {
	fake.SourceStub = nil
	fake.sourceReturns = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeResourceOptions) Params() atc.Params {
	fake.paramsMutex.Lock()
	fake.paramsArgsForCall = append(fake.paramsArgsForCall, struct{}{})
	fake.recordInvocation("Params", []interface{}{})
	fake.paramsMutex.Unlock()
	if fake.ParamsStub != nil {
		return fake.ParamsStub()
	}
	return fake.paramsReturns.result1
}

func (fake *FakeResourceOptions) ParamsCallCount() int {
	fake.paramsMutex.RLock()
	defer fake.paramsMutex.RUnlock()
	return len(fake.paramsArgsForCall)
}

func (fake *FakeResourceOptions) ParamsReturns(result1 atc.Params) {
	fake.ParamsStub = nil
	fake.paramsReturns = struct {
		result1 atc.Params
	}{result1}
}

func (fake *FakeResourceOptions) Version() atc.Version {
	fake.versionMutex.Lock()
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	return fake.versionReturns.result1
}

func (fake *FakeResourceOptions) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeResourceOptions) VersionReturns(result1 atc.Version) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeResourceOptions) ResourceType() resource.ResourceType {
	fake.resourceTypeMutex.Lock()
	fake.resourceTypeArgsForCall = append(fake.resourceTypeArgsForCall, struct{}{})
	fake.recordInvocation("ResourceType", []interface{}{})
	fake.resourceTypeMutex.Unlock()
	if fake.ResourceTypeStub != nil {
		return fake.ResourceTypeStub()
	}
	return fake.resourceTypeReturns.result1
}

func (fake *FakeResourceOptions) ResourceTypeCallCount() int {
	fake.resourceTypeMutex.RLock()
	defer fake.resourceTypeMutex.RUnlock()
	return len(fake.resourceTypeArgsForCall)
}

func (fake *FakeResourceOptions) ResourceTypeReturns(result1 resource.ResourceType) {
	fake.ResourceTypeStub = nil
	fake.resourceTypeReturns = struct {
		result1 resource.ResourceType
	}{result1}
}

func (fake *FakeResourceOptions) LockName(workerName string) (string, error) {
	fake.lockNameMutex.Lock()
	fake.lockNameArgsForCall = append(fake.lockNameArgsForCall, struct {
		workerName string
	}{workerName})
	fake.recordInvocation("LockName", []interface{}{workerName})
	fake.lockNameMutex.Unlock()
	if fake.LockNameStub != nil {
		return fake.LockNameStub(workerName)
	}
	return fake.lockNameReturns.result1, fake.lockNameReturns.result2
}

func (fake *FakeResourceOptions) LockNameCallCount() int {
	fake.lockNameMutex.RLock()
	defer fake.lockNameMutex.RUnlock()
	return len(fake.lockNameArgsForCall)
}

func (fake *FakeResourceOptions) LockNameArgsForCall(i int) string {
	fake.lockNameMutex.RLock()
	defer fake.lockNameMutex.RUnlock()
	return fake.lockNameArgsForCall[i].workerName
}

func (fake *FakeResourceOptions) LockNameReturns(result1 string, result2 error) {
	fake.LockNameStub = nil
	fake.lockNameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceOptions) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iOConfigMutex.RLock()
	defer fake.iOConfigMutex.RUnlock()
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	fake.paramsMutex.RLock()
	defer fake.paramsMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.resourceTypeMutex.RLock()
	defer fake.resourceTypeMutex.RUnlock()
	fake.lockNameMutex.RLock()
	defer fake.lockNameMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeResourceOptions) recordInvocation(key string, args []interface{}) {
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

var _ resource.ResourceOptions = new(FakeResourceOptions)
