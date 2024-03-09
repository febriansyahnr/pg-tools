package main

import (
	"embed"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/febrianpaper/pg-tools/config"
	http_snapCore "github.com/febrianpaper/pg-tools/internal/adapter/http/snapCore"
	mysql_trxHistory "github.com/febrianpaper/pg-tools/internal/adapter/mysql/trxHistory"
	"github.com/febrianpaper/pg-tools/internal/handler"
	handler_virtualAccount "github.com/febrianpaper/pg-tools/internal/handler/virtualAccount"
	"github.com/febrianpaper/pg-tools/internal/service"
	service_trxHistory "github.com/febrianpaper/pg-tools/internal/service/trxHistory"
	service_virtualAccount "github.com/febrianpaper/pg-tools/internal/service/virtualAccount"
	"github.com/febrianpaper/pg-tools/pkg/httpRequestExt"
	"github.com/febrianpaper/pg-tools/pkg/mySqlExt"
	"github.com/febrianpaper/pg-tools/pkg/redisExt"
	"github.com/febrianpaper/pg-tools/port"
	"github.com/go-chi/chi/v5"
)

//go:embed public
var FS embed.FS

type Module struct {
	Config         *config.Config
	Secret         *config.Secret
	HttpClient     httpRequestExt.IHTTPRequest
	DatabaseClient mySqlExt.IMySqlExt
}

func (m *Module) setupAdapters() port.Adapters {
	cacheClient, err := redisExt.New(
		m.Config.RedisConfig.Host,
		m.Config.RedisConfig.Port,
		m.Secret.RedisSecret.Password,
		m.Config.RedisConfig.CacheDB,
	)
	if err != nil {
		log.Fatal(err)
	}

	trxLogRepo := mysql_trxHistory.NewTrxLogRepo(m.DatabaseClient)

	return port.Adapters{
		SnapCore: http_snapCore.New(m.HttpClient, m.Config, m.Secret, trxLogRepo),
		TrxQueue: mysql_trxHistory.NewTrxQueueRepo(m.DatabaseClient),
		Cache:    cacheClient,
		TrxLog:   trxLogRepo,
	}
}

func (m *Module) setupServices(adapters *port.Adapters) service.Services {
	return service.Services{
		VirtualAccountService: service_virtualAccount.New(adapters.SnapCore, adapters.Cache, adapters.TrxQueue),
		TrxHistoryService:     service_trxHistory.New(adapters.TrxLog),
	}
}

func (m *Module) Run() {
	router := chi.NewMux()

	adapters := m.setupAdapters()
	services := m.setupServices(&adapters)

	vaHandler := handler_virtualAccount.New(services.VirtualAccountService, services.TrxHistoryService)

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/virtual-account", handler.MakeHandler(vaHandler.Index))
	router.Post("/virtual-account/inquiry", handler.MakeHandler(vaHandler.Inquiry))
	router.Post("/virtual-account/pay", handler.MakeHandler(vaHandler.Payment))
	router.Get("/virtual-account/list-processed-va", handler.MakeHandler(vaHandler.ListProcessedVA))
	router.Get("/virtual-account/log-page", handler.MakeHandler(vaHandler.LogPage))
	router.Get("/virtual-account/log/detail", handler.MakeHandler(vaHandler.LogDetailPage))

	slog.Info("Application running", "port", ":3033")

	log.Fatal(http.ListenAndServe(":3033", router))
}

func main() {
	config, secret, err := config.LoadConfig("./.config.yaml", "./.secret.yaml")
	if err != nil {
		log.Fatalf("Unable to load configuration and secret: %v", err)
	}
	mysqlExtConfig := mySqlExt.Config{
		Host:         config.MySQLConfig.Host,
		Port:         config.MySQLConfig.Port,
		Username:     secret.MySQLSecret.Username,
		Password:     secret.MySQLSecret.Password,
		DBName:       secret.MySQLSecret.Database,
		MaxIdleConns: config.MySQLConfig.MaxIdleConns,
		MaxIdleTime:  config.MySQLConfig.MaxOpenConns,
		MaxLifeTime:  config.MySQLConfig.MaxLifeTime,
		MaxOpenConns: config.MySQLConfig.MaxOpenConns,
	}
	dbClient, err := mySqlExt.New(mysqlExtConfig)
	if err != nil {
		fmt.Printf("Unable to init mysql, %v", err)
		panic(err)
	}
	defer dbClient.Close()

	module := &Module{
		Config:         config,
		Secret:         secret,
		HttpClient:     httpRequestExt.New(),
		DatabaseClient: dbClient,
	}

	module.Run()
}
