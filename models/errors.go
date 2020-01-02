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

package models

import "fmt"

// ModelError model error
type ModelError struct {
	content string
	fmt     string
}

func newModelError(fmt string) ModelError {
	return ModelError{fmt: fmt}
}

func (m ModelError) Error() string { return m.content }

// With ...
func (m *ModelError) With(v ...interface{}) error {
	m.content = fmt.Sprintf(m.fmt, v...)
	return m
}

// model errors
var (
	InvalidModelError      = newModelError("invalid model '%s'")
	InvalidBundleKindError = newModelError("invalid model bundle '%s'")
	InvalidBundleError     = newModelError("invalid bundle '%s'")
	InvalidParameterError  = newModelError("invalid parameter '%s'")
)
