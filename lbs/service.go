package lbs

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/cloustone/pandas/lbs/grpc_lbs_v1"
	lbp "github.com/cloustone/pandas/lbs/proxy"
	"github.com/cloustone/pandas/pkg/auth"
	"github.com/cloustone/pandas/pkg/message"
	logr "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var gerrf = status.Errorf

type LbsService struct {
	Proxy *lbp.Proxy
}

// Geofence
func (l *LbsService) CreateCircleGeofence(ctx context.Context, in *pb.CreateCircleGeofenceRequest) (*pb.CreateCircleGeofenceResponse, error) {
	logr.Debugf("CreateCircleGeofence (%s)", in.String())

	name := fmt.Sprintf("%s-%s-%s", in.UserId, in.ProjectId, in.Fence.Name)
	fenceId, err := l.Proxy.CreateCircleGeofence(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		lbp.CircleGeofence{
			Name:             name,
			MonitoredObjects: strings.Join(in.Fence.MonitoredObjects, ","),
			Longitude:        in.Fence.Longitude,
			Latitude:         in.Fence.Latitude,
			Radius:           in.Fence.Radius,
			Denoise:          int(in.Fence.Denoise),
			CoordType:        lbp.CoordType(in.Fence.CoordType),
		})
	if err != nil {
		logr.WithError(err).Errorf("create circle geofence failed")
		return nil, gerrf(codes.Internal, "create circle geofence failed")
	}
	return &pb.CreateCircleGeofenceResponse{FenceId: fenceId}, nil
}

func (l *LbsService) CreatePolyGeofence(ctx context.Context, in *pb.CreatePolyGeofenceRequest) (*pb.CreatePolyGeofenceResponse, error) {
	logr.Debugf("CreatePolyGeofence (%s)", in.String())

	name := fmt.Sprintf("%s-%s-%s", in.UserId, in.ProjectId, in.Fence.Name)
	fenceId, err := l.Proxy.CreatePolyGeofence(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		lbp.PolyGeofence{
			Name:             name,
			MonitoredObjects: strings.Join(in.Fence.MonitoredObjects, ","),
			Vertexes:         in.Fence.Vertexes,
			Denoise:          int(in.Fence.Denoise),
			CoordType:        lbp.CoordType(in.Fence.CoordType),
		})
	if err != nil {
		logr.WithError(err).Errorf("create poly geofence failed")
		return nil, gerrf(codes.Internal, "create poly geofence failed")
	}
	return &pb.CreatePolyGeofenceResponse{FenceId: fenceId}, nil
}

func (l *LbsService) UpdatePolyGeofence(ctx context.Context, in *pb.UpdatePolyGeofenceRequest) (*pb.UpdatePolyGeofenceResponse, error) {
	logr.Debugf("UpdatePolyGeofence (%s)", in.String())

	name := fmt.Sprintf("%s-%s-%s", in.UserId, in.ProjectId, in.Fence.Name)

	err := l.Proxy.UpdatePolyGeofence(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		lbp.PolyGeofence{
			Name:             name,
			MonitoredObjects: strings.Join(in.Fence.MonitoredObjects, ","),
			Vertexes:         in.Fence.Vertexes,
			Denoise:          int(in.Fence.Denoise),
			FenceId:          in.Fence.FenceId,
			CoordType:        lbp.CoordType(in.Fence.CoordType),
		})
	if err != nil {
		logr.WithError(err).Errorf("update poly geofence failed")
		return nil, gerrf(codes.Internal, "update poly geofence failed")
	}
	return &pb.UpdatePolyGeofenceResponse{}, nil
}

func (l *LbsService) UpdateCircleGeofence(ctx context.Context, in *pb.UpdateCircleGeofenceRequest) (*pb.UpdateCircleGeofenceResponse, error) {
	logr.Debugf("UpdateCircleGeofence (%s)", in.String())

	name := fmt.Sprintf("%s-%s-%s", in.UserId, in.ProjectId, in.Fence.Name)

	err := l.Proxy.UpdateCircleGeofence(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		lbp.CircleGeofence{
			Name:             name,
			MonitoredObjects: strings.Join(in.Fence.MonitoredObjects, ","),
			Longitude:        in.Fence.Longitude,
			Latitude:         in.Fence.Latitude,
			Radius:           in.Fence.Radius,
			Denoise:          int(in.Fence.Denoise),
			FenceId:          in.Fence.FenceId,
			CoordType:        lbp.CoordType(in.Fence.CoordType),
		})
	if err != nil {
		logr.WithError(err).Errorf("update circle geofence failed")
		return nil, gerrf(codes.Internal, "update circle geofence failed")
	}
	return &pb.UpdateCircleGeofenceResponse{}, nil
}

func (l *LbsService) DeleteGeofence(ctx context.Context, in *pb.DeleteGeofenceRequest) (*pb.DeleteGeofenceResponse, error) {
	logr.Debugf("DeleteGeofence (%s)", in.String())

	fenceIds, err := l.Proxy.DeleteGeofence(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in.FenceIds, in.Objects)
	if err != nil {
		logr.WithError(err).Errorf("delete geofence failed")
		return nil, gerrf(codes.Internal, "delete geofence failed")
	}
	return &pb.DeleteGeofenceResponse{FenceIds: fenceIds}, nil
}

func (l *LbsService) ListGeofences(ctx context.Context, in *pb.ListGeofencesRequest) (*pb.ListGeofencesResponse, error) {
	logr.Debugf("ListGeofence (%s)", in.String())

	fences, err := l.Proxy.ListGeofence(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in.FenceIds, in.Objects)
	if err != nil {
		logr.WithError(err).Errorf("list geofence failed")
		return nil, gerrf(codes.Internal, "list geofence failed")
	}
	fenceList := []*pb.Geofence{}
	for _, f := range fences {
		fence := &pb.Geofence{
			FenceId:         fmt.Sprint(f.FenceId),
			FenceName:       f.FenceName,
			MonitoredObject: strings.Split(f.MonitoredObject, ","),
			Shape:           f.Shape,
			Longitude:       f.Longitude,
			Latitude:        f.Latitude,
			Radius:          f.Radius,
			CoordType:       string(f.CoordType),
			Denoise:         int32(f.Denoise),
			CreateTime:      f.CreateTime,
			UpdateTime:      f.UpdateTime,
		}
		for _, vtx := range f.Vertexes {
			vertexe := &pb.Vertexe{
				Latitude:  vtx.Latitude,
				Longitude: vtx.Longitude,
			}
			fence.Vertexes = append(fence.Vertexes, vertexe)
		}
		fenceList = append(fenceList, fence)
	}
	return &pb.ListGeofencesResponse{Fences: fenceList}, nil
}

func (l *LbsService) AddMonitoredObject(ctx context.Context, in *pb.AddMonitoredObjectRequest) (*pb.AddMonitoredObjectResponse, error) {
	logr.Debugf("AddMonitoredObject (%s)", in.String())

	err := l.Proxy.AddMonitoredObject(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in.FenceId, in.Objects)
	if err != nil {
		logr.WithError(err).Errorf("add monitored object failed")
		return nil, gerrf(codes.Internal, "add monitored object failed")
	}
	return &pb.AddMonitoredObjectResponse{}, nil
}

func (l *LbsService) RemoveMonitoredObject(ctx context.Context, in *pb.RemoveMonitoredObjectRequest) (*pb.RemoveMonitoredObjectResponse, error) {
	logr.Debugf("RemoveMonitoredObject(%s)", in.String())

	err := l.Proxy.RemoveMonitoredObject(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in.FenceId, in.Objects)
	if err != nil {
		logr.WithError(err).Errorf("remove monitored object failed")
		return nil, gerrf(codes.Internal, "remove monitored object failed")
	}
	return &pb.RemoveMonitoredObjectResponse{}, nil
}

func (l *LbsService) ListMonitoredObjects(ctx context.Context, in *pb.ListMonitoredObjectsRequest) (*pb.ListMonitoredObjectsResponse, error) {
	logr.Debugf("ListMonitoredObjects(%s)", in.String())

	total, objects := l.Proxy.ListMonitoredObjects(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in.FenceId, int(in.PageIndex), int(in.PageSize))

	return &pb.ListMonitoredObjectsResponse{TotalFences: int32(total), Objects: objects}, nil
}

func (l *LbsService) QueryStatus(ctx context.Context, in *pb.QueryStatusRequest) (*pb.QueryStatusResponse, error) {
	logr.Debugf("QueryStatus(%s)", in.String())

	fenceStatus, err := l.Proxy.QueryStatus(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in.MonitoredPerson, in.FenceIds)
	if err != nil {
		logr.Errorln("QueryStatus failed:", err)
		return nil, gerrf(codes.NotFound, "not found")
	}

	rsp := getMonitoredStatus(fenceStatus)

	return rsp, nil
}

func getMonitoredStatus(fenceStatus lbp.BaiduQueryStatusResponse) *pb.QueryStatusResponse {
	rsp := &pb.QueryStatusResponse{
		Status:  int32(fenceStatus.Status),
		Message: fenceStatus.Message,
		Size:    int32(fenceStatus.Size),
	}
	for _, mpVal := range fenceStatus.MonitoredStatuses {
		monitoredStatus := &pb.MonitoredStatus{
			FenceId:         int32(mpVal.FenceId),
			MonitoredStatus: mpVal.MonitoredStatus,
		}
		rsp.MonitoredStatuses = append(rsp.MonitoredStatuses, monitoredStatus)
	}
	logr.Debugln("get monitered status:", rsp)
	return rsp
}

func (l *LbsService) GetHistoryAlarms(ctx context.Context, in *pb.GetHistoryAlarmsRequest) (*pb.GetHistoryAlarmsResponse, error) {
	logr.Debugf("GetHistoryAlarms (%s)", in.String())

	alarmPoint, err := l.Proxy.GetHistoryAlarms(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in.MonitoredPerson, in.FenceIds)
	if err != nil {
		logr.Errorln("GetHistoryAlarms failed:", err)
		return nil, gerrf(codes.NotFound, "not found")
	}

	rsp := getHistoryAlarmPoint(alarmPoint)

	return rsp, nil
}

func getHistoryAlarmPoint(alarmPoint lbp.BaiduGetHistoryAlarmsResponse) *pb.GetHistoryAlarmsResponse {
	rsp := &pb.GetHistoryAlarmsResponse{
		Status:  int32(alarmPoint.Status),
		Message: alarmPoint.Message,
		Size:    int32(alarmPoint.Size),
	}
	for _, haVal := range alarmPoint.Alarms {

		alarm := &pb.Alarm{
			FenceId:         int32(haVal.FenceId),
			FenceName:       haVal.FenceName,
			MonitoredPerson: haVal.MonitoredPerson,
			Action:          haVal.Action,
			AlarmPoint: &pb.AlarmPoint{
				Longitude:  haVal.AlarmPoint.Longitude,
				Latitude:   haVal.AlarmPoint.Latitude,
				Radius:     int32(haVal.AlarmPoint.Radius),
				CoordType:  haVal.AlarmPoint.CoordType,
				LocTime:    haVal.AlarmPoint.LocTime,
				CreateTime: haVal.AlarmPoint.CreateTime,
			},
			PrePoint: &pb.PrePoint{
				Longitude:  haVal.AlarmPoint.Longitude,
				Latitude:   haVal.AlarmPoint.Latitude,
				Radius:     int32(haVal.AlarmPoint.Radius),
				CoordType:  haVal.AlarmPoint.CoordType,
				LocTime:    haVal.AlarmPoint.LocTime,
				CreateTime: haVal.AlarmPoint.CreateTime,
			},
		}

		rsp.Alarms = append(rsp.Alarms, alarm)
	}
	logr.Debugln("getHistoryAlarmPoint:", rsp)
	return rsp
}

func (l *LbsService) BatchGetHistoryAlarms(ct context.Context, in *pb.BatchGetHistoryAlarmsRequest) (*pb.BatchGetHistoryAlarmsResponse, error) {
	logr.Debugf("RemoveCollection (%s)", in.String())

	historyAlarms, err := l.Proxy.BatchGetHistoryAlarms(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in)
	if err != nil {
		logr.Errorln("BatchGetHistoryAlarms failed:", err)
		return nil, gerrf(codes.NotFound, "not found")
	}

	rsp := getBatchHistoryAlarmPoint(historyAlarms)

	return rsp, nil

}

func getBatchHistoryAlarmPoint(historyAlarms lbp.BaiduBatchHistoryAlarmsResp) *pb.BatchGetHistoryAlarmsResponse {
	rsp := &pb.BatchGetHistoryAlarmsResponse{
		Status:  int32(historyAlarms.Status),
		Message: historyAlarms.Message,
		Size:    int32(historyAlarms.Size),
		Total:   int32(historyAlarms.Total),
	}
	for _, haVal := range historyAlarms.Alarms {

		alarm := &pb.Alarm{
			FenceId:         int32(haVal.FenceId),
			FenceName:       haVal.FenceName,
			MonitoredPerson: haVal.MonitoredPerson,
			Action:          haVal.Action,
			AlarmPoint: &pb.AlarmPoint{
				Longitude:  haVal.AlarmPoint.Longitude,
				Latitude:   haVal.AlarmPoint.Latitude,
				Radius:     int32(haVal.AlarmPoint.Radius),
				CoordType:  haVal.AlarmPoint.CoordType,
				LocTime:    haVal.AlarmPoint.LocTime,
				CreateTime: haVal.AlarmPoint.CreateTime,
			},
			PrePoint: &pb.PrePoint{
				Longitude:  haVal.AlarmPoint.Longitude,
				Latitude:   haVal.AlarmPoint.Latitude,
				Radius:     int32(haVal.AlarmPoint.Radius),
				CoordType:  haVal.AlarmPoint.CoordType,
				LocTime:    haVal.AlarmPoint.LocTime,
				CreateTime: haVal.AlarmPoint.CreateTime,
			},
		}

		rsp.Alarms = append(rsp.Alarms, alarm)
	}
	logr.Debugln("getHistoryAlarmPoint:", rsp)
	return rsp
}

func (l *LbsService) GetStayPoints(ctx context.Context, in *pb.GetStayPointsRequest) (*pb.GetStayPointsResponse, error) {
	logr.Debugf("GetStayPoints (%s)", in.String())

	stayPoints, err := l.Proxy.GetStayPoints(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in)
	if err != nil {
		logr.Errorln("BatchGetHistoryAlarms failed:", err)
		return nil, gerrf(codes.NotFound, "not found")
	}

	rsp := getGrpcStayPoints(stayPoints)

	return rsp, nil
}

func getGrpcStayPoints(stayPoints lbp.BaiduGetStayPointResp) *pb.GetStayPointsResponse {
	rsp := &pb.GetStayPointsResponse{
		Status:  int32(stayPoints.Status),
		Message: stayPoints.Message,
		Size:    int32(stayPoints.Size),
		Total:   int32(stayPoints.Total),
		StartPoint: &pb.Point{
			Latitude:  stayPoints.StartPoint.Latitude,
			Longitude: stayPoints.StartPoint.Longitude,
			CoordType: stayPoints.StartPoint.CoordType,
			LocTime:   fmt.Sprint(stayPoints.StartPoint.LocTime),
		},
		EndPoint: &pb.Point{
			Latitude:  stayPoints.EndPoint.Latitude,
			Longitude: stayPoints.EndPoint.Longitude,
			CoordType: stayPoints.EndPoint.CoordType,
			LocTime:   fmt.Sprint(stayPoints.EndPoint.LocTime),
		},
	}
	for _, val := range stayPoints.Points {
		point := &pb.Point{
			Latitude:  val.Latitude,
			Longitude: val.Longitude,
			CoordType: val.CoordType,
			LocTime:   fmt.Sprint(val.LocTime),
		}
		rsp.Points = append(rsp.Points, point)
	}
	logr.Debugln("getGrpcStayPoints rsp:", rsp)
	return rsp
}

func (l *LbsService) NotifyAlarms(ctx context.Context, in *pb.NotifyAlarmsRequest) (*pb.NotifyAlarmsResponse, error) {
	logr.Debugf("NotifyAlarms(%s)", in.String())

	alarm, err := l.Proxy.UnmarshalAlarmNotification(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in.Content)
	if err != nil {
		logr.WithError(err).Errorf("unmarshal alarm failed")
		return nil, gerrf(codes.Internal, "unmarshal alarm failed")
	}
	config := message.NewConfigWithViper()
	producer, err := message.NewProducer(config, false)
	if err != nil {
		logr.WithError(err).Errorf("create message producer failed")
		return nil, gerrf(codes.Internal, "create message producer failed")
	}

	if err := producer.SendMessage(&lbp.AlarmTopic{Alarm: alarm}); err != nil {
		logr.WithError(err).Errorf("send alarm failed")
		return nil, gerrf(codes.Internal, "send alarm failed")
	}
	return &pb.NotifyAlarmsResponse{}, nil
}

func (l *LbsService) ListCollections(ctx context.Context, in *pb.ListCollectionsRequest) (*pb.ListCollectionsResponse, error) {
	logr.Debugln("ListCollections")

	//	productList := getProductList(products)
	//	logr.Debugln("productList:", productList)

	//	return productList, nil
	return nil, nil
}

/*
func getProductList(products []*Collection) *pb.ListCollectionsResponse {
	productList := &pb.ListCollectionsResponse{}
	for _, val := range products {
		productList.ProjectIds = append(productList.ProjectIds, val.ProjectId)
	}
	return productList
}
*/

func (l *LbsService) AddEntity(ctx context.Context, in *pb.AddEntityRequest) (*pb.AddEntityResponse, error) {
	logr.Debugf("AddEntity (%s)", in.String())

	err := l.Proxy.AddEntity(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in.EntityName, in.EntityDesc)
	if err != nil {
		logr.WithError(err).Errorf("AddEntity failed")
		return nil, gerrf(codes.Internal, "AddEntity failed")
	}
	return &pb.AddEntityResponse{}, nil
}

func (l *LbsService) DeleteEntity(ctx context.Context, in *pb.DeleteEntityRequest) (*pb.DeleteEntityResponse, error) {
	logr.Debugf("DeleteEntity (%s)", in.String())

	err := l.Proxy.DeleteEntity(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in.EntityName)
	if err != nil {
		logr.WithError(err).Errorf("DeleteEntity failed")
		return nil, gerrf(codes.Internal, "DeleteEntity failed")
	}
	return &pb.DeleteEntityResponse{}, nil
}

func (l *LbsService) UpdateEntity(ctx context.Context, in *pb.UpdateEntityRequest) (*pb.UpdateEntityResponse, error) {
	logr.Debugf("UpdateEntity (%s)", in.String())

	err := l.Proxy.UpdateEntity(
		auth.NewPrincipal(in.UserId, in.ProjectId),
		in.EntityName, in.EntityDesc)
	if err != nil {
		logr.WithError(err).Errorf("UpdateEntity failed")
		return nil, gerrf(codes.Internal, "UpdateEntity failed")
	}
	/*
		entity := EntityRecord{
			//	UserId:        in.UserId,
			//	ProjectId:     in.ProjectId,
			EntityName:    in.EntityName,
			LastUpdatedAt: time.Now(),
		}
	*/

	return &pb.UpdateEntityResponse{}, nil
}

func (l *LbsService) ListEntity(ctx context.Context, in *pb.ListEntityRequest) (*pb.ListEntityResponse, error) {
	logr.Debugf("ListEntity (%s)", in.String())
	/*
		entitiesInfo, err := l.Proxy.GetEntity(
			auth.NewPrincipal(in.UserId, in.ProjectId))
		if err != nil {
			logr.WithError(err).Errorf("mongo get entities info failed")
			return nil, gerrf(codes.Internal, "mongo get entities info failed")
		}
		entitiesName := getEntitiesName(entitiesInfo)

		total, entityInfo := l.Proxy.ListEntity(in.UserId, in.ProjectId, in.CoordTypeOutput, in.PageIndex, in.PageSize)
		if total == -1 {
			logr.Errorf("ListEntity failed")
			return nil, gerrf(codes.Internal, "ListEntity failed")
		}

		entitys := &pb.ListEntityResponse{}
		for _, val := range entityInfo.Entities {
			entityInfo := &pb.EntityInfo{
				EntityName: val.EntityName,
				Longitude:  val.LastLocation.Longitude,
				Latitude:   val.LastLocation.Latitude,
			}
			if isEntityInCollection(val.EntityName, entitiesName) == false {
				continue
			}
			entitys.EntityInfo = append(entitys.EntityInfo, entityInfo)
		}
		entitys.Total = int32(len(entitys.EntityInfo))

		logr.Debugln("EntityInfo:", entityInfo, "total:", total)

		return entitys, nil
	*/
	return nil, nil
}

func getEntitiesName(entitiesInfo []*lbp.EntityRecord) []string {
	entitiesName := make([]string, 0)
	for _, val := range entitiesInfo {
		entitiesName = append(entitiesName, val.EntityName)
	}
	return entitiesName
}
func isEntityInCollection(entityName string, entitiesName []string) bool {
	var isEntityExist bool = false
	for _, val := range entitiesName {
		if val == entityName {
			isEntityExist = true
		}
	}
	return isEntityExist
}

func (l *LbsService) GetFenceIds(ctx context.Context, in *pb.GetFenceIdsRequest) (*pb.GetFenceIdsResponse, error) {
	logr.Debugf("GetFenceIds (%s)", in.String())
	/*

		fences, err := l.Proxy.GetFenceIds(
			auth.NewPrincipal(in.UserId, in.ProjectId))
		if err != nil {
			logr.WithError(err).Errorln("ListCollections err:", err)
			return nil, err
		}

		fenceIds := getFenceIds(fences)

		return fenceIds, nil
	*/
	return nil, nil
}

/*
func getFenceIds(fences []*GeofenceRecord) *pb.GetFenceIdsResponse {
	fenceIdsResp := &pb.GetFenceIdsResponse{}
	for _, val := range fences {
		fenceIdsResp.FenceIds = append(fenceIdsResp.FenceIds, val.FenceId)
	}
	logr.Debugln("fenceIds:", fenceIdsResp.FenceIds)
	return fenceIdsResp
}
*/
func (l *LbsService) GetFenceUserId(ctx context.Context, in *pb.GetFenceUserIdRequest) (*pb.GetFenceUserIdResponse, error) {
	/*
		userId, err := l.Proxy.GetFenceUserId(nil, in.FenceId)
		if err != nil {
			logr.WithError(err).Errorf("add collection '%s' failed", in.FenceId)
			return nil, err
		}
		rsp := &pb.GetFenceUserIdResponse{
			UserId: userId,
		}
		return rsp, nil
	*/
	return nil, nil
}
