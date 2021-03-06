// Code generated by counterfeiter. DO NOT EDIT.
package historyfakes

import (
	"sync"

	"github.com/desmondrawls/rock-paper-scissors/history"
	"github.com/desmondrawls/rock-paper-scissors/models"
)

type RepositoryStub struct {
	ListStub        func() ([]models.Record, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct{}
	listReturns     struct {
		result1 []models.Record
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []models.Record
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RepositoryStub) List() ([]models.Record, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct{}{})
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listReturns.result1, fake.listReturns.result2
}

func (fake *RepositoryStub) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *RepositoryStub) ListReturns(result1 []models.Record, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []models.Record
		result2 error
	}{result1, result2}
}

func (fake *RepositoryStub) ListReturnsOnCall(i int, result1 []models.Record, result2 error) {
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []models.Record
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []models.Record
		result2 error
	}{result1, result2}
}

func (fake *RepositoryStub) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RepositoryStub) recordInvocation(key string, args []interface{}) {
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

var _ history.Repository = new(RepositoryStub)
