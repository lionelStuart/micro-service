package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	proto "github.com/micro/go-plugins/config/source/grpc/proto"
	grpc2 "google.golang.org/grpc"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	mux        sync.RWMutex
	configMaps = make(map[string]*proto.ChangeSet)
	apps       = []string{"micro"}
)

type Service struct {
}

func (s Service) Read(ctx context.Context, req *proto.ReadRequest) (rsp *proto.ReadResponse, err error) {
	appName := parsePath(req.Path)

	rsp = &proto.ReadResponse{
		ChangeSet: getConfig(appName)}
	return
}

func (s Service) Watch(req *proto.WatchRequest, server proto.Source_WatchServer) (err error) {
	appName := parsePath(req.Path)

	rsp := &proto.WatchResponse{
		ChangeSet: getConfig(appName)}
	if err = server.Send(rsp); err != nil {
		log.Logf("[Watch] watch handle err, %s", err)
		return err
	}

	return
}

func loadAndWatchConfigFile() (err error) {
	for _, app := range apps {
		if err := config.Load(file.NewSource(
			file.WithPath("./conf/" + app + ".yml"),
		)); err != nil {
			log.Fatalf("[loadAndWatchConfigFile] load file path err, %s", err)
		}
	}

	watcher, err := config.Watch()
	if err != nil {
		log.Fatalf("[loadAndWatchConfigFile] load file changed err, %s", err)
		return err
	}

	go func() {
		for {
			v, err := watcher.Next()
			if err != nil {
				log.Fatalf("[loadAndWatchConfigFile] watch file changed err, %s", err)
				return
			}

			//handle changes
			log.Logf("[loadAndWatchConfigFile] file changed, %s", string(v.Bytes()))
		}
	}()

	return
}

func getConfig(appName string) *proto.ChangeSet {
	bytes := config.Get(appName).Bytes()

	log.Logf("[getConfig] appName: %s", appName)
	return &proto.ChangeSet{
		Data:      bytes,
		Checksum:  fmt.Sprintf("%x", md5.Sum(bytes)),
		Format:    "yml",
		Source:    "file",
		Timestamp: time.Now().Unix(),
	}
}

func parsePath(path string) (appName string) {
	paths := strings.Split(path, "/")

	if paths[0] == "" && len(paths) > 1 {
		return paths[1]
	}

	return paths[0]
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Logf("[main] Recoverd in f%v", r)
		}
	}()

	err := loadAndWatchConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	service := grpc2.NewServer()
	proto.RegisterSourceServer(service, new(Service))
	ts, err := net.Listen("tcp", ":9600")
	if err != nil {
		log.Fatal(err)
	}
	log.Log("configServer started")

	err = service.Serve(ts)
	if err != nil {
		log.Fatalf("err %v", err)
	}
}
