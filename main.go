package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zommage/leisureRoom/conf"
	"github.com/zommage/leisureRoom/controllers/room"
	. "github.com/zommage/leisureRoom/logs"
	models "github.com/zommage/leisureRoom/models"
	pb "github.com/zommage/leisureRoom/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var (
	confInfo *conf.Config
	confPath = flag.String("config", "./conf/app.dev.ini", "platform draw console profilePath")
)

func Init() error {
	flag.Parse()
	var err error
	confInfo, err = conf.InitConfig(confPath)
	if err != nil {
		return fmt.Errorf("init config err: %v", err)
	}

	err = InitLog(confInfo.LogConf.LogPath, confInfo.LogConf.LogLevel)
	if err != nil {
		return fmt.Errorf("Init log is err: %v", err)
	}

	err = models.InitDb()
	if err != nil {
		return fmt.Errorf("Init nsq is err: %v", err)
	}
	return nil
}

func main() {
	//catch global panic
	defer func() {
		if err := recover(); err != nil {
			Log.Errorf("panic err: %v", err)
			fmt.Println("panic err: ", err)
		}
	}()

	err := Init()
	if err != nil {
		fmt.Println("main init err: ", err)
		return
	}

	grpcLis, err := net.Listen("tcp", confInfo.GrpcConf.GrpcHost+":"+confInfo.GrpcConf.GrpcPort)
	if err != nil {
		Log.Errorf("net listen failed: %v", err)
		return
	}
	grpcServer := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(confInfo.GrpcConf.GrpcMaxConnectIdleSec) * time.Second,
		MaxConnectionAge:      time.Duration(confInfo.GrpcConf.GrpcMaxConnectAgeSec) * time.Second,
		MaxConnectionAgeGrace: time.Duration(confInfo.GrpcConf.GrpcMaxConnectAgeGraceSec) * time.Second,
		Time:    time.Duration(confInfo.GrpcConf.GrpcTimeSec) * time.Second,
		Timeout: time.Duration(confInfo.GrpcConf.GrpcTimeTimeoutSec) * time.Second,
	}))

	// register grpc service
	pb.RegisterLeisureRoomServiceServer(grpcServer, &room.RoomService{})

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = grpcServer.Serve(grpcLis); err != nil {
			Log.Errorf("grpc error:%v", err)
			panic(err)
		}
	}()

	Log.Info("leisureRoom server is start")
	fmt.Println("leisureRoom server is start")

	<-done
	grpcServer.GracefulStop()
	if err = models.Close(); err != nil {
		Log.Errorf("db close failed:%v", err)
	}

	fmt.Println("leisureRoom server is down")
	Log.Info("leisureRoom server is down")
}
