package util

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"strings"
	"time"

	"github.com/cloustone/pandas/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

const (
	sslSuffix = "?ssl=true"
)

func ConnectMongoWithViper() (*mgo.Session, error) {
	mongoURI := config.GetMongoAddr()
	username := viper.GetString("mongo.username")
	password := viper.GetString("mongo.password")
	cert := viper.GetString("mongo.cert")

	return ConnectMongo(mongoURI, username, password, cert)
}

// Connect connects to a mongo database collection, using the provided username, password, and certificate file
// It returns a pointer to the session and collection objects, or an error if the connection attempt fails.
func ConnectMongo(mongoURI string, username string, password string, cert string) (*mgo.Session, error) {
	uri := strings.TrimSuffix(mongoURI, sslSuffix)
	dialInfo, err := mgo.ParseURL(uri)
	if err != nil {
		logrus.WithError(err).Errorf("Cannot parse Mongo Connection URI")
		return nil, err
	}
	dialInfo.FailFast = true
	dialInfo.Timeout = 10 * time.Second

	// only do ssl if we have the suffix
	if strings.HasSuffix(mongoURI, sslSuffix) {
		logrus.Debugf("Using TLS for mongo ...")
		tlsConfig := &tls.Config{}
		roots := x509.NewCertPool()
		if ca, err := ioutil.ReadFile(cert); err == nil {
			roots.AppendCertsFromPEM(ca)
		}
		tlsConfig.RootCAs = roots
		tlsConfig.InsecureSkipVerify = false
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
	}

	// in case the username/password are not part of the URL string
	if username != "" && password != "" {
		dialInfo.Username = username
		dialInfo.Password = password
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func EnsureIndex(s *mgo.Session, dbname string, collection string, keys ...string) {
	session := s.Clone()
	defer session.Close()

	c := session.DB(dbname).C(collection)
	index := mgo.Index{
		Key:        keys,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		logrus.Errorf("Create dev INDEX error! %s\n", err)
	}
}
