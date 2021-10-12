package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	pb "go.kolesa-team.org/gl/go-course/rpc"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"sync"
)

// Config структура конфигов
type Config struct {
	Listen   string // Адрес и порт
	LogLevel string // Уровень логирования
}

// main точка входа
func main() {
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	cfg := new(Config)
	if _, err := toml.DecodeFile(*configPath, cfg); err != nil {
		logrus.
			WithField("config", *configPath).
			WithError(err).
			Fatalln("Невозможно загрузить конфигурацию сервиса")
	}

	lvl, _ := logrus.ParseLevel(cfg.LogLevel)

	logrus.SetLevel(lvl)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	s := NewServer(Options{
		GrpcPort: 8080,
		RESTPort: 8081,
	})

	s.Start()
	s.WaitStop()
}

type Handler struct {
	pb.UnimplementedServiceServer
}

func (h *Handler) Health(ctx context.Context, empty *pb.Empty) (*pb.StringMessage, error) {
	return &pb.StringMessage{Value: "ok"}, nil
}

// Сервер
type Server struct {
	wg      sync.WaitGroup
	server  *Handler
	options Options
}

// Конфиги сервера
type Options struct {
	GrpcPort  int
	RESTPort  int
	PprofPort int
}

// Новый сервер
func NewServer(options Options) *Server {
	return &Server{options: options}
}

// Старт сервера
func (s *Server) Start() {
	var err error
	s.wg.Add(1)

	go func() {
		err = s.startGRPC()
		s.wg.Done()

		if err != nil {
			logrus.WithError(err).Error("Ошибка при запуске gRPC сервера")
		}
	}()

	s.wg.Add(1)

	go func() {
		err = s.startREST()
		s.wg.Done()

		if err != nil {
			logrus.WithError(err).Error("Ошибка при запуске REST сервера")
		}
	}()

	s.wg.Add(1)
}

// Запускает gRPC сервис
func (s *Server) startGRPC() error {
	port := s.options.GrpcPort
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))

	if err != nil {
		logrus.WithError(err).Errorf("Невозможно запустить gRPC сервис на порту %d", port)
		return err
	}

	grpcServer := grpc.NewServer()

	pb.RegisterServiceServer(grpcServer, s.server)

	logrus.Infof("gRPC сервис запущен на порту %d", port)

	return grpcServer.Serve(lis)
}

// Запускает REST сервис
func (s *Server) startREST() error {
	port := s.options.RESTPort
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterServiceHandlerFromEndpoint(ctx, gwMux, fmt.Sprintf(":%d", s.options.GrpcPort), opts)

	if err != nil {
		logrus.WithError(err).Errorf("Невозможно запустить REST сервис на порту %d", port)
		return err
	}

	logrus.Infof("REST сервис запущен на порту %d", port)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), gwMux)
}

// Ожидаю завершения всех процессов
func (s *Server) WaitStop() {
	s.wg.Wait()
}

// health используется для проверки состояния микросервиса
func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("ok"))
}
