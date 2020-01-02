package models

import (
	lbsproxy "github.com/cloustone/pandas/lbs/proxy"
)

// AddProjectNotify ...
type AddProjectNotify struct {
	ProjectID string
}

// DeleteProjectNotify ...
type DeleteProjectNotify struct {
	ProjectID string
}

// UpdateProjectNotify ...
type UpdateProjectNotify struct {
	ProjectID string
}

// AddTrackpointNotify ...
type AddTrackpointNotify struct {
	Trackpoint lbsproxy.TrackPoint
	ProjectID  string
	DeviceID   string
	DeviceName string
}
