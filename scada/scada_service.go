package scada

import (
	"context"

	pb "github.com/cloustone/pandas/scada/grpc_scada_v1"
)

type ScadaService struct{}

func NewScadaService() *ScadaService {
	return &ScadaService{}
}

func (ss *ScadaService) GetWidgets(context.Context, *pb.GetWidgetsRequest) (*pb.GetWidgetsResponse, error) {
	return nil, nil
}
