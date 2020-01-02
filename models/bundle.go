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

import "strings"

// BundleKind Bundle Kind
type BundleKind string

// BundleKind
const (
	BundleKindPresentation BundleKind = "ModelPresentation"
	BundleKindDefinition   BundleKind = "ModelDefinition"
)

// BundleScheme Bundle Scheme
type BundleScheme string

// BundleScheme
const (
	BundleSchemeYaml    BundleScheme = "yaml"
	BundleSchemeJSON    BundleScheme = "json"
	BundleSchemeInvalid BundleScheme = "invalid"
)

// Bundle Bundle factory
type Bundle interface {
	Scheme() BundleScheme
	Content() []byte
	Kind() BundleKind
	Name() string
}

type simpleBundle struct {
	scheme  BundleScheme
	content []byte
	kind    BundleKind
	name    string
}

// BundleSchemeWithNameSuffix ...
func BundleSchemeWithNameSuffix(name string) BundleScheme {
	suffix := string(BundleSchemeInvalid)
	if strings.Contains(name, ".") {
		suffix = name[:strings.LastIndex(name, ".")]
	}

	switch suffix {
	case "yaml", "yml":
		return BundleSchemeYaml
	case "json":
		return BundleSchemeJSON
	default:
		return BundleSchemeInvalid
	}
}

// NewBundle ...
func NewBundle(name string, kind BundleKind, data []byte, scheme BundleScheme) Bundle {
	return &simpleBundle{
		name:    name,
		kind:    kind,
		content: data,
		scheme:  scheme,
	}
}

func (b *simpleBundle) Scheme() BundleScheme { return b.scheme }
func (b *simpleBundle) Content() []byte      { return b.content }
func (b *simpleBundle) Kind() BundleKind     { return b.kind }
func (b *simpleBundle) Name() string         { return b.name }
