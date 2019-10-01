package apiserver

import (
	"github.com/kanatmg/http-rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
    "github.com/gorilla/mux"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

//создание сервера
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Запуск сервера
func (server *APIServer) Start() error {
	if err := server.configureLogger(); err != nil {
		return err
	}

	server.configureRouter()
	server.logger.Info("starting api server")
	if err := server.configureStore(); err != nil{
		return nil
	}

	return http.ListenAndServe(server.config.BindAddr, server.router)
}

func (server *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}
	server.logger.SetLevel(level)
	return nil
}

func (server *APIServer) configureStore() error {
	st := store.New(server.config.Store)
	if err := st.Open(); err != nil {
		return nil
	}
	server.store = st
	return nil
}

func(server *APIServer) configureRouter() {
	server.router.HandleFunc("/",server.handleHello())
}

func (server *APIServer) handleHello() http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request) {
		io.WriteString(w,"Hello, world or こんにちは世界\n")
	}
}



