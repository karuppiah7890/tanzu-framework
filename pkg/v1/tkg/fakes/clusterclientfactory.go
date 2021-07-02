// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/clusterclient"
)

type ClusterClientFactory struct {
	NewClientStub        func(string, string, clusterclient.Options) (clusterclient.Client, error)
	newClientMutex       sync.RWMutex
	newClientArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 clusterclient.Options
	}
	newClientReturns struct {
		result1 clusterclient.Client
		result2 error
	}
	newClientReturnsOnCall map[int]struct {
		result1 clusterclient.Client
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ClusterClientFactory) NewClient(arg1 string, arg2 string, arg3 clusterclient.Options) (clusterclient.Client, error) {
	fake.newClientMutex.Lock()
	ret, specificReturn := fake.newClientReturnsOnCall[len(fake.newClientArgsForCall)]
	fake.newClientArgsForCall = append(fake.newClientArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 clusterclient.Options
	}{arg1, arg2, arg3})
	stub := fake.NewClientStub
	fakeReturns := fake.newClientReturns
	fake.recordInvocation("NewClient", []interface{}{arg1, arg2, arg3})
	fake.newClientMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ClusterClientFactory) NewClientCallCount() int {
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	return len(fake.newClientArgsForCall)
}

func (fake *ClusterClientFactory) NewClientCalls(stub func(string, string, clusterclient.Options) (clusterclient.Client, error)) {
	fake.newClientMutex.Lock()
	defer fake.newClientMutex.Unlock()
	fake.NewClientStub = stub
}

func (fake *ClusterClientFactory) NewClientArgsForCall(i int) (string, string, clusterclient.Options) {
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	argsForCall := fake.newClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ClusterClientFactory) NewClientReturns(result1 clusterclient.Client, result2 error) {
	fake.newClientMutex.Lock()
	defer fake.newClientMutex.Unlock()
	fake.NewClientStub = nil
	fake.newClientReturns = struct {
		result1 clusterclient.Client
		result2 error
	}{result1, result2}
}

func (fake *ClusterClientFactory) NewClientReturnsOnCall(i int, result1 clusterclient.Client, result2 error) {
	fake.newClientMutex.Lock()
	defer fake.newClientMutex.Unlock()
	fake.NewClientStub = nil
	if fake.newClientReturnsOnCall == nil {
		fake.newClientReturnsOnCall = make(map[int]struct {
			result1 clusterclient.Client
			result2 error
		})
	}
	fake.newClientReturnsOnCall[i] = struct {
		result1 clusterclient.Client
		result2 error
	}{result1, result2}
}

func (fake *ClusterClientFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ClusterClientFactory) recordInvocation(key string, args []interface{}) {
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

var _ clusterclient.ClusterClientFactory = new(ClusterClientFactory)
