package main

import (
	"github.com/dtas-pm/mailer/pkg/handler"
	pb "github.com/dtas-pm/mailer/proto"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	// go messageLoop()
	//Указываем на каком порту будем слушать запросы
	listener, err := net.Listen("tcp", viper.GetString("port"))
	if err != nil {
		logrus.Fatal("failed to listen", err)
	}
	logrus.Printf("start listening for emails at port %s", viper.GetString("port"))

	//Создаём новый grpc сервер
	rpcserv := grpc.NewServer()

	srv := handler.NewServer()

	//Регистрируем связку сервер + listener
	pb.RegisterMailerServer(rpcserv, srv)
	reflection.Register(rpcserv)

	//Запускаемся и ждём RPC-запросы
	err = rpcserv.Serve(listener)
	if err != nil {
		logrus.Fatal("failed to serve", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
