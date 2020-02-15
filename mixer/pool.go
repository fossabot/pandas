//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package mixer

import (
	"sync"

	"github.com/cloustone/pandas/mixer/adaptors"
)

// adaptorPool manage all adaptors created by client's request
type adaptorPool struct {
	mutex       sync.RWMutex
	adaptors    []adaptors.Adaptor
	adaptorRefs map[string]int
}

// newAdaptorPool return a adaptor pool
func newAdaptorPool() *adaptorPool {
	return &adaptorPool{
		mutex:       sync.RWMutex{},
		adaptors:    []adaptors.Adaptor{},
		adaptorRefs: make(map[string]int),
	}
}

// addAdaptor add a newly created adaptor into pool
func (p *adaptorPool) addAdaptor(adaptor adaptors.Adaptor) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.adaptors = append(p.adaptors, adaptor)
	p.adaptorRefs[adaptor.Name()] = 1
}

// isAdaptorExist return wether a adaptors already exist
func (p *adaptorPool) getAdaptorWithOptions(adaptorOptions *adaptors.AdaptorOptions) adaptors.Adaptor {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, adaptor := range p.adaptors {
		options := adaptor.Options()
		if options.Protocol == adaptorOptions.Protocol &&
			options.IsProvider == adaptorOptions.IsProvider &&
			options.ServicePort == adaptorOptions.ServicePort &&
			options.IsTLSEnabled == adaptorOptions.IsTLSEnabled &&
			options.ConnectURL == adaptorOptions.ConnectURL {
			return adaptor
		}
	}
	return nil
}

// getAdaptor return specified adaptor
func (p *adaptorPool) getAdaptor(adaptorID string) adaptors.Adaptor {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, adaptor := range p.adaptors {
		if adaptorID == adaptor.Name() {
			return adaptor
		}
	}
	return nil
}

// removeAdaptor remove a adaptor from pool
func (p *adaptorPool) removeAdaptor(adaptorID string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for index, adaptor := range p.adaptors {
		if adaptorID == adaptor.Name() {
			p.adaptors = append(p.adaptors[:index], p.adaptors[index:]...)
			break
		}
	}
}

// incAdaptorRef increase ref count for specifed adaptor
func (p *adaptorPool) incAdaptorRef(adaptorID string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.adaptorRefs[adaptorID]++
}

// decAdaptorRef decrease ref count for specifed adaptor
func (p *adaptorPool) decAdaptorRef(adaptorID string) int {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.adaptorRefs[adaptorID]--
	return p.adaptorRefs[adaptorID]
}
