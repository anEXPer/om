// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"encoding/json"
	"sync"

	"github.com/pivotal-cf/om/api"
)

type DirectorService struct {
	UpdateStagedDirectorAvailabilityZonesStub        func(api.AZConfiguration) error
	updateStagedDirectorAvailabilityZonesMutex       sync.RWMutex
	updateStagedDirectorAvailabilityZonesArgsForCall []struct {
		arg1 api.AZConfiguration
	}
	updateStagedDirectorAvailabilityZonesReturns struct {
		result1 error
	}
	updateStagedDirectorAvailabilityZonesReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStagedDirectorNetworksStub        func(json.RawMessage) error
	updateStagedDirectorNetworksMutex       sync.RWMutex
	updateStagedDirectorNetworksArgsForCall []struct {
		arg1 json.RawMessage
	}
	updateStagedDirectorNetworksReturns struct {
		result1 error
	}
	updateStagedDirectorNetworksReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStagedDirectorNetworkAndAZStub        func(api.NetworkAndAZConfiguration) error
	updateStagedDirectorNetworkAndAZMutex       sync.RWMutex
	updateStagedDirectorNetworkAndAZArgsForCall []struct {
		arg1 api.NetworkAndAZConfiguration
	}
	updateStagedDirectorNetworkAndAZReturns struct {
		result1 error
	}
	updateStagedDirectorNetworkAndAZReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStagedDirectorPropertiesStub        func(api.DirectorProperties) error
	updateStagedDirectorPropertiesMutex       sync.RWMutex
	updateStagedDirectorPropertiesArgsForCall []struct {
		arg1 api.DirectorProperties
	}
	updateStagedDirectorPropertiesReturns struct {
		result1 error
	}
	updateStagedDirectorPropertiesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DirectorService) UpdateStagedDirectorAvailabilityZones(arg1 api.AZConfiguration) error {
	fake.updateStagedDirectorAvailabilityZonesMutex.Lock()
	ret, specificReturn := fake.updateStagedDirectorAvailabilityZonesReturnsOnCall[len(fake.updateStagedDirectorAvailabilityZonesArgsForCall)]
	fake.updateStagedDirectorAvailabilityZonesArgsForCall = append(fake.updateStagedDirectorAvailabilityZonesArgsForCall, struct {
		arg1 api.AZConfiguration
	}{arg1})
	fake.recordInvocation("UpdateStagedDirectorAvailabilityZones", []interface{}{arg1})
	fake.updateStagedDirectorAvailabilityZonesMutex.Unlock()
	if fake.UpdateStagedDirectorAvailabilityZonesStub != nil {
		return fake.UpdateStagedDirectorAvailabilityZonesStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateStagedDirectorAvailabilityZonesReturns.result1
}

func (fake *DirectorService) UpdateStagedDirectorAvailabilityZonesCallCount() int {
	fake.updateStagedDirectorAvailabilityZonesMutex.RLock()
	defer fake.updateStagedDirectorAvailabilityZonesMutex.RUnlock()
	return len(fake.updateStagedDirectorAvailabilityZonesArgsForCall)
}

func (fake *DirectorService) UpdateStagedDirectorAvailabilityZonesArgsForCall(i int) api.AZConfiguration {
	fake.updateStagedDirectorAvailabilityZonesMutex.RLock()
	defer fake.updateStagedDirectorAvailabilityZonesMutex.RUnlock()
	return fake.updateStagedDirectorAvailabilityZonesArgsForCall[i].arg1
}

func (fake *DirectorService) UpdateStagedDirectorAvailabilityZonesReturns(result1 error) {
	fake.UpdateStagedDirectorAvailabilityZonesStub = nil
	fake.updateStagedDirectorAvailabilityZonesReturns = struct {
		result1 error
	}{result1}
}

func (fake *DirectorService) UpdateStagedDirectorAvailabilityZonesReturnsOnCall(i int, result1 error) {
	fake.UpdateStagedDirectorAvailabilityZonesStub = nil
	if fake.updateStagedDirectorAvailabilityZonesReturnsOnCall == nil {
		fake.updateStagedDirectorAvailabilityZonesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStagedDirectorAvailabilityZonesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DirectorService) UpdateStagedDirectorNetworks(arg1 json.RawMessage) error {
	fake.updateStagedDirectorNetworksMutex.Lock()
	ret, specificReturn := fake.updateStagedDirectorNetworksReturnsOnCall[len(fake.updateStagedDirectorNetworksArgsForCall)]
	fake.updateStagedDirectorNetworksArgsForCall = append(fake.updateStagedDirectorNetworksArgsForCall, struct {
		arg1 json.RawMessage
	}{arg1})
	fake.recordInvocation("UpdateStagedDirectorNetworks", []interface{}{arg1})
	fake.updateStagedDirectorNetworksMutex.Unlock()
	if fake.UpdateStagedDirectorNetworksStub != nil {
		return fake.UpdateStagedDirectorNetworksStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateStagedDirectorNetworksReturns.result1
}

func (fake *DirectorService) UpdateStagedDirectorNetworksCallCount() int {
	fake.updateStagedDirectorNetworksMutex.RLock()
	defer fake.updateStagedDirectorNetworksMutex.RUnlock()
	return len(fake.updateStagedDirectorNetworksArgsForCall)
}

func (fake *DirectorService) UpdateStagedDirectorNetworksArgsForCall(i int) json.RawMessage {
	fake.updateStagedDirectorNetworksMutex.RLock()
	defer fake.updateStagedDirectorNetworksMutex.RUnlock()
	return fake.updateStagedDirectorNetworksArgsForCall[i].arg1
}

func (fake *DirectorService) UpdateStagedDirectorNetworksReturns(result1 error) {
	fake.UpdateStagedDirectorNetworksStub = nil
	fake.updateStagedDirectorNetworksReturns = struct {
		result1 error
	}{result1}
}

func (fake *DirectorService) UpdateStagedDirectorNetworksReturnsOnCall(i int, result1 error) {
	fake.UpdateStagedDirectorNetworksStub = nil
	if fake.updateStagedDirectorNetworksReturnsOnCall == nil {
		fake.updateStagedDirectorNetworksReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStagedDirectorNetworksReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DirectorService) UpdateStagedDirectorNetworkAndAZ(arg1 api.NetworkAndAZConfiguration) error {
	fake.updateStagedDirectorNetworkAndAZMutex.Lock()
	ret, specificReturn := fake.updateStagedDirectorNetworkAndAZReturnsOnCall[len(fake.updateStagedDirectorNetworkAndAZArgsForCall)]
	fake.updateStagedDirectorNetworkAndAZArgsForCall = append(fake.updateStagedDirectorNetworkAndAZArgsForCall, struct {
		arg1 api.NetworkAndAZConfiguration
	}{arg1})
	fake.recordInvocation("UpdateStagedDirectorNetworkAndAZ", []interface{}{arg1})
	fake.updateStagedDirectorNetworkAndAZMutex.Unlock()
	if fake.UpdateStagedDirectorNetworkAndAZStub != nil {
		return fake.UpdateStagedDirectorNetworkAndAZStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateStagedDirectorNetworkAndAZReturns.result1
}

func (fake *DirectorService) UpdateStagedDirectorNetworkAndAZCallCount() int {
	fake.updateStagedDirectorNetworkAndAZMutex.RLock()
	defer fake.updateStagedDirectorNetworkAndAZMutex.RUnlock()
	return len(fake.updateStagedDirectorNetworkAndAZArgsForCall)
}

func (fake *DirectorService) UpdateStagedDirectorNetworkAndAZArgsForCall(i int) api.NetworkAndAZConfiguration {
	fake.updateStagedDirectorNetworkAndAZMutex.RLock()
	defer fake.updateStagedDirectorNetworkAndAZMutex.RUnlock()
	return fake.updateStagedDirectorNetworkAndAZArgsForCall[i].arg1
}

func (fake *DirectorService) UpdateStagedDirectorNetworkAndAZReturns(result1 error) {
	fake.UpdateStagedDirectorNetworkAndAZStub = nil
	fake.updateStagedDirectorNetworkAndAZReturns = struct {
		result1 error
	}{result1}
}

func (fake *DirectorService) UpdateStagedDirectorNetworkAndAZReturnsOnCall(i int, result1 error) {
	fake.UpdateStagedDirectorNetworkAndAZStub = nil
	if fake.updateStagedDirectorNetworkAndAZReturnsOnCall == nil {
		fake.updateStagedDirectorNetworkAndAZReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStagedDirectorNetworkAndAZReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DirectorService) UpdateStagedDirectorProperties(arg1 api.DirectorProperties) error {
	fake.updateStagedDirectorPropertiesMutex.Lock()
	ret, specificReturn := fake.updateStagedDirectorPropertiesReturnsOnCall[len(fake.updateStagedDirectorPropertiesArgsForCall)]
	fake.updateStagedDirectorPropertiesArgsForCall = append(fake.updateStagedDirectorPropertiesArgsForCall, struct {
		arg1 api.DirectorProperties
	}{arg1})
	fake.recordInvocation("UpdateStagedDirectorProperties", []interface{}{arg1})
	fake.updateStagedDirectorPropertiesMutex.Unlock()
	if fake.UpdateStagedDirectorPropertiesStub != nil {
		return fake.UpdateStagedDirectorPropertiesStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateStagedDirectorPropertiesReturns.result1
}

func (fake *DirectorService) UpdateStagedDirectorPropertiesCallCount() int {
	fake.updateStagedDirectorPropertiesMutex.RLock()
	defer fake.updateStagedDirectorPropertiesMutex.RUnlock()
	return len(fake.updateStagedDirectorPropertiesArgsForCall)
}

func (fake *DirectorService) UpdateStagedDirectorPropertiesArgsForCall(i int) api.DirectorProperties {
	fake.updateStagedDirectorPropertiesMutex.RLock()
	defer fake.updateStagedDirectorPropertiesMutex.RUnlock()
	return fake.updateStagedDirectorPropertiesArgsForCall[i].arg1
}

func (fake *DirectorService) UpdateStagedDirectorPropertiesReturns(result1 error) {
	fake.UpdateStagedDirectorPropertiesStub = nil
	fake.updateStagedDirectorPropertiesReturns = struct {
		result1 error
	}{result1}
}

func (fake *DirectorService) UpdateStagedDirectorPropertiesReturnsOnCall(i int, result1 error) {
	fake.UpdateStagedDirectorPropertiesStub = nil
	if fake.updateStagedDirectorPropertiesReturnsOnCall == nil {
		fake.updateStagedDirectorPropertiesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStagedDirectorPropertiesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DirectorService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updateStagedDirectorAvailabilityZonesMutex.RLock()
	defer fake.updateStagedDirectorAvailabilityZonesMutex.RUnlock()
	fake.updateStagedDirectorNetworksMutex.RLock()
	defer fake.updateStagedDirectorNetworksMutex.RUnlock()
	fake.updateStagedDirectorNetworkAndAZMutex.RLock()
	defer fake.updateStagedDirectorNetworkAndAZMutex.RUnlock()
	fake.updateStagedDirectorPropertiesMutex.RLock()
	defer fake.updateStagedDirectorPropertiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DirectorService) recordInvocation(key string, args []interface{}) {
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
