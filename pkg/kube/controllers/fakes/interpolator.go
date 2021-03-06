// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/quarks-operator/pkg/kube/util/withops"
)

type FakeInterpolator struct {
	AddOpsStub        func([]byte) error
	addOpsMutex       sync.RWMutex
	addOpsArgsForCall []struct {
		arg1 []byte
	}
	addOpsReturns struct {
		result1 error
	}
	addOpsReturnsOnCall map[int]struct {
		result1 error
	}
	InterpolateStub        func([]byte) ([]byte, error)
	interpolateMutex       sync.RWMutex
	interpolateArgsForCall []struct {
		arg1 []byte
	}
	interpolateReturns struct {
		result1 []byte
		result2 error
	}
	interpolateReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterpolator) AddOps(arg1 []byte) error {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.addOpsMutex.Lock()
	ret, specificReturn := fake.addOpsReturnsOnCall[len(fake.addOpsArgsForCall)]
	fake.addOpsArgsForCall = append(fake.addOpsArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("AddOps", []interface{}{arg1Copy})
	fake.addOpsMutex.Unlock()
	if fake.AddOpsStub != nil {
		return fake.AddOpsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addOpsReturns
	return fakeReturns.result1
}

func (fake *FakeInterpolator) AddOpsCallCount() int {
	fake.addOpsMutex.RLock()
	defer fake.addOpsMutex.RUnlock()
	return len(fake.addOpsArgsForCall)
}

func (fake *FakeInterpolator) AddOpsCalls(stub func([]byte) error) {
	fake.addOpsMutex.Lock()
	defer fake.addOpsMutex.Unlock()
	fake.AddOpsStub = stub
}

func (fake *FakeInterpolator) AddOpsArgsForCall(i int) []byte {
	fake.addOpsMutex.RLock()
	defer fake.addOpsMutex.RUnlock()
	argsForCall := fake.addOpsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInterpolator) AddOpsReturns(result1 error) {
	fake.addOpsMutex.Lock()
	defer fake.addOpsMutex.Unlock()
	fake.AddOpsStub = nil
	fake.addOpsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInterpolator) AddOpsReturnsOnCall(i int, result1 error) {
	fake.addOpsMutex.Lock()
	defer fake.addOpsMutex.Unlock()
	fake.AddOpsStub = nil
	if fake.addOpsReturnsOnCall == nil {
		fake.addOpsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addOpsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInterpolator) Interpolate(arg1 []byte) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.interpolateMutex.Lock()
	ret, specificReturn := fake.interpolateReturnsOnCall[len(fake.interpolateArgsForCall)]
	fake.interpolateArgsForCall = append(fake.interpolateArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Interpolate", []interface{}{arg1Copy})
	fake.interpolateMutex.Unlock()
	if fake.InterpolateStub != nil {
		return fake.InterpolateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.interpolateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInterpolator) InterpolateCallCount() int {
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	return len(fake.interpolateArgsForCall)
}

func (fake *FakeInterpolator) InterpolateCalls(stub func([]byte) ([]byte, error)) {
	fake.interpolateMutex.Lock()
	defer fake.interpolateMutex.Unlock()
	fake.InterpolateStub = stub
}

func (fake *FakeInterpolator) InterpolateArgsForCall(i int) []byte {
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	argsForCall := fake.interpolateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInterpolator) InterpolateReturns(result1 []byte, result2 error) {
	fake.interpolateMutex.Lock()
	defer fake.interpolateMutex.Unlock()
	fake.InterpolateStub = nil
	fake.interpolateReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpolator) InterpolateReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.interpolateMutex.Lock()
	defer fake.interpolateMutex.Unlock()
	fake.InterpolateStub = nil
	if fake.interpolateReturnsOnCall == nil {
		fake.interpolateReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.interpolateReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpolator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addOpsMutex.RLock()
	defer fake.addOpsMutex.RUnlock()
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInterpolator) recordInvocation(key string, args []interface{}) {
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

var _ withops.Interpolator = new(FakeInterpolator)
