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
	mutex    sync.RWMutex                // mutex lock
	adaptors map[string]adaptors.Adaptor // all adaptors
	refs     map[string]int              // each adaptor's reference count
}

// newAdaptorPool return a adaptor pool
func newAdaptorPool() *adaptorPool {
	return &adaptorPool{
		mutex:    sync.RWMutex{},
		adaptors: make(map[string]adaptors.Adaptor),
		refs:     make(map[string]int),
	}
}

// addAdaptor add a newly created adaptor into pool
func (p *adaptorPool) addAdaptor(adaptor adaptors.Adaptor) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.adaptors[adaptor.Name()] = adaptor
	p.refs[adaptor.Name()] = 1
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

	if adaptor, found := p.adaptors[adaptorID]; found {
		return adaptor
	}
	return nil
}

// getAdaptors return domains's all adaptors
func (p *adaptorPool) getAdaptors(domain string) []adaptors.Adaptor {
	adaptors := []adaptors.Adaptor{}
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, adaptor := range p.adaptors {
		if adaptor.Options().Domain == domain {
			adaptors = append(adaptors, adaptor)
		}
	}
	return adaptors
}

// removeAdaptor remove a adaptor from pool
func (p *adaptorPool) removeAdaptor(adaptor adaptors.Adaptor) {
	adaptorID := adaptor.Name()
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if _, found := p.adaptors[adaptorID]; found {
		delete(p.adaptors, adaptorID)
	}
}

// incAdaptorRef increase ref count for specifed adaptor
func (p *adaptorPool) incAdaptorRef(adaptor adaptors.Adaptor) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.refs[adaptor.Name()]++
}

// decAdaptorRef decrease ref count for specifed adaptor
func (p *adaptorPool) decAdaptorRef(adaptor adaptors.Adaptor) int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.refs[adaptor.Name()]--
	return p.refs[adaptor.Name()]
}
