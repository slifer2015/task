//go:generate swag init --parseDependency -d ../internal/api -g ../../cmd/main.go
package main

import (
	golog "log"

	"project/cmd/docs"
	"project/internal/api"
	"project/internal/config"
	"project/pkg/logger"
)

// @title Task API
// @version 1.0
// @description Task micro-service
// @termsOfService http://swagger.io/terms/

// @contact.name CONTACT NAME
// @contact.url http://www.contact.url
// @contact.email contact@email.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9090
// @BasePath /api/v1
func main() {
	cfg, err := config.Read()
	if err != nil {
		golog.Fatal(err)
	}

	log, err := logger.NewLogger(cfg.ServiceName, cfg.LogLevel, []string{})
	if err != nil {
		golog.Fatal(err)
	}

	docs.SwaggerInfo.Host = cfg.HostWithoutProtocol()

	server, err := api.NewServer(cfg, log)
	if err != nil {
		golog.Fatalf("can't start server: %s", err)
	}

	defer func() {
		_ = server.Close()
	}()

	log.Info("server started", logger.String("port", cfg.Port))

	golog.Fatal(server.Start(cfg.Port))
}
