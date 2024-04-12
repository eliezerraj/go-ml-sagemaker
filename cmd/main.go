package main

import(
	"context"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/go-ml-sagemaker/internal/core"
	"github.com/go-ml-sagemaker/internal/handler"
	"github.com/go-ml-sagemaker/internal/service"
)

var(
	logLevel = zerolog.DebugLevel
	infoPod					core.InfoPod
	server					core.Server
	sageMakerEndpoint		core.SageMakerEndpoint

	configOTEL	core.ConfigOTEL

	httpAppServerConfig 	core.HttpAppServer
)

func getEnv() {
	log.Debug().Msg("getEnv")

	if os.Getenv("POD_NAME") !=  "" {
		infoPod.PodName = os.Getenv("POD_NAME")
	}
	if os.Getenv("VERSION") !=  "" {
		infoPod.Version = os.Getenv("VERSION")
	}

	if os.Getenv("PORT") !=  "" {
		intVar, _ := strconv.Atoi(os.Getenv("PORT"))
		server.Port = intVar
	}

	if os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") !=  "" {	
		infoPod.OtelExportEndpoint = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	}

	if os.Getenv("SAGEMAKER_FRAUD_ENDPOINT") !=  "" {	
		sageMakerEndpoint.FraudEndpoint = os.Getenv("SAGEMAKER_FRAUD_ENDPOINT")
	}
}

func init(){
	log.Debug().Msg("init")
	zerolog.SetGlobalLevel(logLevel)

	err := godotenv.Load(".env")
	if err != nil {
		log.Info().Err(err).Msg("No .env file !!!")
	}

	getEnv()
	
	server.ReadTimeout = 60
	server.WriteTimeout = 60
	server.IdleTimeout = 60
	server.CtxTimeout = 60

	configOTEL.TimeInterval = 1
	configOTEL.TimeAliveIncrementer = 1
	configOTEL.TotalHeapSizeUpperBound = 100
	configOTEL.ThreadsActiveUpperBound = 10
	configOTEL.CpuUsageUpperBound = 100
	configOTEL.SampleAppPorts = []string{}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Error().Err(err).Msg("Error to get the POD IP address !!!")
		os.Exit(3)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				infoPod.IPAddress = ipnet.IP.String()
			}
		}
	}

	infoPod.ConfigOTEL 	= &configOTEL
}

func main() {
	log.Debug().Msg("----------------------------------------------------")
	log.Debug().Msg("main")
	log.Debug().Msg("----------------------------------------------------")
	log.Debug().Interface("",server).Msg("")
	log.Debug().Msg("----------------------------------------------------")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration( server.ReadTimeout ) * time.Second)
	defer cancel()

	httpAppServerConfig.Server = &server
	httpAppServerConfig.SageMakerEndpoint = &sageMakerEndpoint
	workerService := service.NewWorkerService(sageMakerEndpoint)
	httpWorkerAdapter := handler.NewHttpWorkerAdapter(workerService)
	httpAppServerConfig.InfoPod = &infoPod
	httpServer := handler.NewHttpAppServer(httpAppServerConfig)

	httpServer.StartHttpAppServer(ctx, httpWorkerAdapter)			
}