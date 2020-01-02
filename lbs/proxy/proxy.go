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
package proxy

import (
	"github.com/cloustone/pandas/lbs/grpc_lbs_v1"
	"github.com/cloustone/pandas/pkg/auth"
	genericoptions "github.com/cloustone/pandas/pkg/server/options"
)

type Proxy struct {
	engineName string
	repo       Repository
	engine     Engine
}

func NewProxy(locationServingOptions *genericoptions.LocationServingOptions) *Proxy {
	return &Proxy{
		engine:     newBaiduLbsEngine(locationServingOptions),
		engineName: locationServingOptions.Provider,
		repo:       NewRepository(),
	}
}

func (p *Proxy) AddTrackPoint(principal auth.Principal, point TrackPoint) {
	p.engine.AddTrackPoint(point)
}

func (p *Proxy) AddTrackPoints(principal auth.Principal, points []TrackPoint) {
	p.engine.AddTrackPoints(points)
}

func (p *Proxy) CreateCircleGeofence(principal auth.Principal, c CircleGeofence) (string, error) {
	return p.engine.CreateCircleGeofence(c)
}

func (p *Proxy) UpdateCircleGeofence(principal auth.Principal, c CircleGeofence) error {
	return p.engine.UpdateCircleGeofence(c)
}

func (p *Proxy) DeleteGeofence(principal auth.Principal, fenceIds []string, objects []string) ([]string, error) {
	return p.engine.DeleteGeofence(fenceIds, objects)
}

func (p *Proxy) ListGeofence(principal auth.Principal, fenceIds []string, objects []string) ([]*Geofence, error) {
	return p.engine.ListGeofence(fenceIds, objects)
}

func (p *Proxy) AddMonitoredObject(principal auth.Principal, fenceId string, objects []string) error {
	return p.engine.AddMonitoredObject(fenceId, objects)
}

func (p *Proxy) RemoveMonitoredObject(principal auth.Principal, fenceId string, objects []string) error {
	return p.engine.RemoveMonitoredObject(fenceId, objects)
}

func (p *Proxy) ListMonitoredObjects(principal auth.Principal, fenceId string, pageIndex int, pageSize int) (int, []string) {
	return p.engine.ListMonitoredObjects(fenceId, pageIndex, pageSize)
}

func (p *Proxy) CreatePolyGeofence(principal auth.Principal, c PolyGeofence) (string, error) {
	return p.engine.CreatePolyGeofence(c)
}

func (p *Proxy) UpdatePolyGeofence(principal auth.Principal, c PolyGeofence) error {
	return p.engine.UpdatePolyGeofence(c)
}

// Alarms
func (p *Proxy) QueryStatus(principal auth.Principal, monitoredPerson string, fenceIds []string) (BaiduQueryStatusResponse, error) {
	return p.engine.QueryStatus(monitoredPerson, fenceIds)
}

func (p *Proxy) GetHistoryAlarms(principal auth.Principal, monitoredPerson string, fenceIds []string) (BaiduGetHistoryAlarmsResponse, error) {
	return p.engine.GetHistoryAlarms(monitoredPerson, fenceIds)
}

func (p *Proxy) BatchGetHistoryAlarms(principal auth.Principal, input *grpc_lbs_v1.BatchGetHistoryAlarmsRequest) (BaiduBatchHistoryAlarmsResp, error) {
	return p.engine.BatchGetHistoryAlarms(input)
}

func (p *Proxy) GetStayPoints(principal auth.Principal, input *grpc_lbs_v1.GetStayPointsRequest) (BaiduGetStayPointResp, error) {
	return p.engine.GetStayPoints(input)
}

func (p *Proxy) UnmarshalAlarmNotification(principal auth.Principal, content []byte) (*AlarmNotification, error) {
	return p.engine.UnmarshalAlarmNotification(content)
}

//Entity
func (p *Proxy) AddEntity(principal auth.Principal, entityName string, entityDesc string) error {
	return p.engine.AddEntity(entityName, entityDesc)
}

func (p *Proxy) UpdateEntity(principal auth.Principal, entityName string, entityDesc string) error {
	return p.engine.UpdateEntity(entityName, entityDesc)
}

func (p *Proxy) DeleteEntity(principal auth.Principal, entityName string) error {
	return p.engine.DeleteEntity(entityName)
}

func (p *Proxy) ListEntity(principal auth.Principal, collectionId string, coordTypeOutput string, pageIndex int32, pageSize int32) (int, baiduListEntityStruct) {
	return p.engine.ListEntity(principal.UserId(), collectionId, coordTypeOutput, pageIndex, pageSize)
}
