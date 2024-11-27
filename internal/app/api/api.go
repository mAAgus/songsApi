package api

import (
	"fmt"
	"go/scr/hhruxongs/storage"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Base API server instance description
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//Добавление поля для работы с хранилищем
	storage *storage.Storage
}

// API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server/configure loggers, router, database connection and etc....
func (api *API) Start() error {
	//Проверяем, что логгер правильно сконфигурирован
	if err := api.configureLoggerField(); err != nil {
		return fmt.Errorf("failed to configure logger: %w", err)
	}

	//Логируем начало запуска сервера
	api.logger.Infof("Starting API server at port: %s", api.config.BindAddr)

	//Конфигурируем маршрутизатор
	api.configureRouterField()

	//Конфигурируем хранилище
	if err := api.configuresStorageFild(); err != nil {
		return fmt.Errorf("failed to configure storage: %w", err)
	}

	//Логируем успешное завершение конфигурации
	api.logger.Info("API server configured successfully")

	//Стартуем HTTP-сервер
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
