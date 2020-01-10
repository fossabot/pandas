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

import "github.com/jinzhu/gorm"

// Widget Widget
// swagger:model Widget
type Widget struct {
	ModelTypeInfo
	gorm.Model
	Name       string       `json:"name" bson:"name"`
	ID         string       `json:"id" bson:"id"`
	Domain     string       `json:"domain" bson:"domain"`
	Version    string       `json:"version" bson:"version"`
	Endpoints  []*Endpoint  `json:"endpoints" bson:"endpoints"`
	DataModels []*DataModel `json:"dataModel" bson:"dataModels"`
	IsLogical  bool         `json:"isLogical" bson:"isLogical"`
	IsCompound bool         `json:"isCompound" bson:"isCompound"`
	IsPreset   bool         `json:"isPreset" bson:"isPreset"`
	Icon       string       `json:"icon" bson:"icon"`
}

// WidgetList ...
type WidgetList struct {
	Widgets map[string][]*Widget `json:"widgets" bson:"widgets"`
}

// NewWidgetWithDeviceModel ...
func NewWidgetWithDeviceModel(m *DeviceModel) *Widget {
	return &Widget{
		Name:       m.Name,
		ID:         m.ID,
		Domain:     m.Domain,
		Icon:       m.Icon,
		Endpoints:  m.Endpoints,
		DataModels: m.DataModels,
		IsLogical:  m.IsLogical,
		IsCompound: m.IsCompound,
		IsPreset:   m.IsPreset,
	}
}

// NewWidgetList ...
func NewWidgetList() *WidgetList {
	return &WidgetList{
		Widgets: make(map[string][]*Widget),
	}
}

// Add ...
func (wl *WidgetList) Add(widgets ...*Widget) {
	for _, w := range widgets {
		domain := w.Domain
		if _, found := wl.Widgets[domain]; !found {
			wl.Widgets[domain] = []*Widget{}
		}
		wl.Widgets[domain] = append(wl.Widgets[domain], w)
	}
}

// Remove ...
func (wl *WidgetList) Remove(domain string, widgetName string) {
	if _, found := wl.Widgets[domain]; found {
		widgets := wl.Widgets[domain]
		for index, widget := range widgets {
			if widget.Name == widgetName {
				widgets = append(widgets[0:index], widgets[index:0]...)
				break
			}
		}
		wl.Widgets[domain] = widgets
	}
}

// AddUserWidget ...
func (wl *WidgetList) AddUserWidget(widgets ...*Widget) {
	for _, w := range widgets {
		w.Domain = "user"
		wl.Add(w)
	}
}

// RemoveUserWidget ...
func (wl *WidgetList) RemoveUserWidget(widgetName string) {
	wl.Remove("user", widgetName)
}

// Count ...
func (wl *WidgetList) Count() int {
	count := 0
	for _, widgets := range wl.Widgets {
		count += len(widgets)
	}
	return count
}

// CountOfDomain ...
func (wl *WidgetList) CountOfDomain(domain string) int {
	if _, found := wl.Widgets[domain]; found {
		return len(wl.Widgets[domain])
	}
	return -1
}
