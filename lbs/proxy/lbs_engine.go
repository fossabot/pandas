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
	"time"

	"github.com/cloustone/pandas/lbs/grpc_lbs_v1"
	serveroptions "github.com/cloustone/pandas/pkg/server/options"
)

type CoordType string

const (
	CoordType_WGS84  CoordType = "wgs84"
	CoordType_GCJ02  CoordType = "gcj02"
	CoordType_BD09LL CoordType = "bd09ll"
)

type Config struct {
	AccessKey string
	ServiceId string
}

type TrackPoint struct {
	EntityName string    `json:"entity_name"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	CoordType  CoordType `json:"coord_type_input"`
	Time       string    `json:"loc_time"`
}

// Location is reported by iot terminal to lbs service
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Time      string  `json:"loc_time"`
}

type CircleGeofence struct {
	Name             string
	MonitoredObjects string
	Longitude        float64
	Latitude         float64
	Radius           float64
	CoordType        CoordType
	Denoise          int
	FenceId          string
}

type PolyGeofence struct {
	Name             string
	MonitoredObjects string
	Vertexes         string
	CoordType        CoordType
	Denoise          int
	FenceId          string
}

type Vertexe struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Geofence struct {
	FenceId         int       `json:"fence_id"`
	FenceName       string    `json:"fence_name"`
	MonitoredObject string    `json:"monitored_person"`
	Shape           string    `json:"shape"`
	Longitude       float64   `json:"longitude"`
	Latitude        float64   `json:"latitude"`
	Radius          float64   `json:"radius"`
	CoordType       CoordType `json:"coord_type"`
	Denoise         int       `json:"denoise"`
	CreateTime      string    `json:"create_time"`
	UpdateTime      string    `json:"modify_time"`
	Vertexes        []Vertexe `json:"vertexes"`
}

type AlarmPoint struct {
	Longitude  float64    `json:"longitude"`
	Latitude   float64    `json:"latitude"`
	Radius     int        `json:"radius"`
	CoordType  string     `json:"coord_type"`
	LocTime    *time.Time `json:"loc_time"`
	CreateTime *time.Time `json:"create_time"`
}

type AlarmPointInfo struct {
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	Radius     int     `json:"radius"`
	CoordType  string  `json:"coord_type"`
	LocTime    int     `json:"loc_time"`
	CreateTime int     `json:"create_time"`
}

type Alarm struct {
	FenceId          string     `json:"fence_id,noempty"`
	FenceName        string     `json:"fence_name,noempty"`
	MonitoredObjects []string   `json:"monitored_objects"`
	Action           string     `json:"action"`
	AlarmPoint       AlarmPoint `json:"alarm_point"`
	PrePoint         AlarmPoint `json:"pre_point"`
}

type AlarmNotification struct {
	Type      int      `json:"type"`
	ServiceId string   `json:"service_id"`
	Alarms    []*Alarm `json:"alarms"`
}

type Point struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	CoordType string  `json:"coord_type"`
	LocTime   int     `json:"loc_time"`
}

type AlarmInfos struct {
	Type      int         `json:"type"`
	ServiceId int         `json:"service_id"`
	Alarms    []AlarmInfo `json:"content"`
}

type AlarmInfo struct {
	FenceId          int            `json:"fence_id,noempty"`
	FenceName        string         `json:"fence_name,noempty"`
	MonitoredObjects string         `json:"monitored_person"`
	Action           string         `json:"action"`
	AlarmPoint       AlarmPointInfo `json:"alarm_point"`
	PrePoint         AlarmPointInfo `json:"pre_point"`
	UserId           string
}

type Engine interface {
	// Add monitored object's track points
	AddTrackPoint(point TrackPoint)
	AddTrackPoints(points []TrackPoint)
	// Create circle geofence and return goefence id if successful
	CreateCircleGeofence(c CircleGeofence) (string, error)
	// Update an existed geofence
	UpdateCircleGeofence(c CircleGeofence) error
	// Delete an existed geofence or monitored objects
	DeleteGeofence(fenceIds []string, objects []string) ([]string, error)
	// List geofences matched with ids or objects
	ListGeofence(fenceIds []string, objects []string) ([]*Geofence, error)
	// Add monitored object for specifed geofence
	AddMonitoredObject(fenceId string, objects []string) error
	// Remove monitored object from specified geofence
	RemoveMonitoredObject(fenceId string, objects []string) error
	// List monitored object in specifed geofence
	ListMonitoredObjects(fenceId string, pageIndex int, pageSize int) (int, []string)
	// Create poly geofence and return goefence id if successful
	CreatePolyGeofence(c PolyGeofence) (string, error)
	// Update an existed poly geofence
	UpdatePolyGeofence(c PolyGeofence) error

	// Alarms
	QueryStatus(monitoredPerson string, fenceIds []string) (BaiduQueryStatusResponse, error)
	GetHistoryAlarms(monitoredPerson string, fenceIds []string) (BaiduGetHistoryAlarmsResponse, error)
	BatchGetHistoryAlarms(input *grpc_lbs_v1.BatchGetHistoryAlarmsRequest) (BaiduBatchHistoryAlarmsResp, error)
	GetStayPoints(input *grpc_lbs_v1.GetStayPointsRequest) (BaiduGetStayPointResp, error)
	UnmarshalAlarmNotification(content []byte) (*AlarmNotification, error)

	//Entity
	AddEntity(EntityName string, EntityDesc string) error
	UpdateEntity(EntityName string, EntityDesc string) error
	DeleteEntity(EntityName string) error
	ListEntity(userId string, collectionId string, CoordTypeOutput string, PageIndex int32, pageSize int32) (int, baiduListEntityStruct)
}

func NewEngine(locationServingOptions *serveroptions.LocationServingOptions) (Engine, error) {
	return newBaiduLbsEngine(locationServingOptions), nil
}
