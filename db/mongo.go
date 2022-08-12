package db

import (
	"fmt"
	"crypto/tls"
	// "net"
	"strings"

	"go-echo-restful-api/config"
	"github.com/labstack/gommon/log"
	mgo "gopkg.in/mgo.v2"
)

var (
	session      *mgo.Session
	databaseName = "example"
)

func init() {
	stgConnection()
}

func stgConnection() {
	hosts := strings.Split(config.Config.Database.Address, ",")
	tlsConfig := &tls.Config{}
	fmt.Println("-------tlsConfig------", tlsConfig)

	dialInfo := &mgo.DialInfo{
		Addrs:          hosts,
		Database:       "admin",
		Username:       config.Config.Database.Username,
		Password:       config.Config.Database.Password,
		ReplicaSetName: "ClusterSVC-shard-0",
	}

	fmt.Println("-------------", dialInfo)
	// dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
	// 	fmt.Println("----------Inner---")

	// 	// conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
	// 	conn, err := mgo.Dial("localhost")
	// 	fmt.Println("-------------", conn)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return conn, err

	// }
	mgoSession, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	session = mgoSession
}

func pullSession() *mgo.Session {
	return session.Copy()
}

// Ping connection
func Ping() error {
	sessionCopy := pullSession()
	defer sessionCopy.Close()
	return sessionCopy.Ping()
}
