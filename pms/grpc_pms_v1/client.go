package grpc_pms_v1

import (
	"fmt"

	serveroptions "github.com/cloustone/pandas/pkg/server/options"
	logr "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	DefaultPort = 30011
	ServiceName = "pms"
)

var (
	// LbsV1LocalAddress exposes the local address that is used if we run with DNS disabled
	V1LocalAddress = fmt.Sprintf(":%d", DefaultPort)
)

// Client is a client interface for interacting with the tuning service.
type Client interface {
	ProjectManager() ProjectManagementClient
	Close() error
}

type client struct {
	pm   ProjectManagementClient
	conn *grpc.ClientConn
}

// NewClient create a new load-balanced client to talk to the Tuning
// service. If the dns_server config option is set to 'disabled', it will
// default to the pre-defined LocalPort of the service.
func NewClient() (Client, error) {
	return NewWithAddress(V1LocalAddress)
}

// NewWithAddress create a new load-balanced client to talk to the Tuning
// service. If the dns_server config option is set to 'disabled', it will
// default to the pre-defined LocalPort of the service.
func NewWithAddress(addr string) (Client, error) {
	address := serveroptions.GetServiceAddress(addr, ServiceName)
	caKey := serveroptions.GetCAKey()

	logr.Infof("NewWithAddress: address: %s, ca: %s", address, caKey)

	var dialOpts []grpc.DialOption

	if serveroptions.IsTLSEnabled() {
		creds, err := credentials.NewClientTLSFromFile(caKey, ServiceName)
		if err != nil {
			return nil, err
		}
		dialOpts = []grpc.DialOption{grpc.WithTransportCredentials(creds), grpc.WithBlock()}
	} else {
		dialOpts = []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	}

	conn, err := grpc.Dial(address, dialOpts...)
	if err != nil {
		return nil, err
	}
	logr.Infof("Dial pms service ok")

	return &client{
		conn: conn,
		pm:   NewProjectManagementClient(conn),
	}, nil
}

func (c *client) ProjectManager() ProjectManagementClient {
	return c.pm
}

func (c *client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
