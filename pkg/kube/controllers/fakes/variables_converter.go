// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/cf-operator/pkg/bosh/manifest"
	"code.cloudfoundry.org/cf-operator/pkg/kube/apis/quarkssecret/v1alpha1"
	"code.cloudfoundry.org/cf-operator/pkg/kube/controllers/boshdeployment"
)

type FakeVariablesConverter struct {
	VariablesStub        func(string, string, []manifest.Variable) ([]v1alpha1.QuarksSecret, error)
	variablesMutex       sync.RWMutex
	variablesArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []manifest.Variable
	}
	variablesReturns struct {
		result1 []v1alpha1.QuarksSecret
		result2 error
	}
	variablesReturnsOnCall map[int]struct {
		result1 []v1alpha1.QuarksSecret
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVariablesConverter) Variables(arg1 string, arg2 string, arg3 []manifest.Variable) ([]v1alpha1.QuarksSecret, error) {
	var arg3Copy []manifest.Variable
	if arg3 != nil {
		arg3Copy = make([]manifest.Variable, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.variablesMutex.Lock()
	ret, specificReturn := fake.variablesReturnsOnCall[len(fake.variablesArgsForCall)]
	fake.variablesArgsForCall = append(fake.variablesArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []manifest.Variable
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("Variables", []interface{}{arg1, arg2, arg3Copy})
	fake.variablesMutex.Unlock()
	if fake.VariablesStub != nil {
		return fake.VariablesStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.variablesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVariablesConverter) VariablesCallCount() int {
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	return len(fake.variablesArgsForCall)
}

func (fake *FakeVariablesConverter) VariablesCalls(stub func(string, string, []manifest.Variable) ([]v1alpha1.QuarksSecret, error)) {
	fake.variablesMutex.Lock()
	defer fake.variablesMutex.Unlock()
	fake.VariablesStub = stub
}

func (fake *FakeVariablesConverter) VariablesArgsForCall(i int) (string, string, []manifest.Variable) {
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	argsForCall := fake.variablesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVariablesConverter) VariablesReturns(result1 []v1alpha1.QuarksSecret, result2 error) {
	fake.variablesMutex.Lock()
	defer fake.variablesMutex.Unlock()
	fake.VariablesStub = nil
	fake.variablesReturns = struct {
		result1 []v1alpha1.QuarksSecret
		result2 error
	}{result1, result2}
}

func (fake *FakeVariablesConverter) VariablesReturnsOnCall(i int, result1 []v1alpha1.QuarksSecret, result2 error) {
	fake.variablesMutex.Lock()
	defer fake.variablesMutex.Unlock()
	fake.VariablesStub = nil
	if fake.variablesReturnsOnCall == nil {
		fake.variablesReturnsOnCall = make(map[int]struct {
			result1 []v1alpha1.QuarksSecret
			result2 error
		})
	}
	fake.variablesReturnsOnCall[i] = struct {
		result1 []v1alpha1.QuarksSecret
		result2 error
	}{result1, result2}
}

func (fake *FakeVariablesConverter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVariablesConverter) recordInvocation(key string, args []interface{}) {
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

var _ boshdeployment.VariablesConverter = new(FakeVariablesConverter)
