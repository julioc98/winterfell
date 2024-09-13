package main

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/ardanlabs/conf/v3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/julioc98/winterfell/internal/app"
	"github.com/julioc98/winterfell/internal/infra/api"
	"github.com/julioc98/winterfell/internal/infra/gateway"
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/utils/checks"
)

type projectConfig struct {
	ID         string `conf:"default:6008875367006208"`
	PrivateKey string `conf:"required"`
	Env        string `conf:"default:sandbox"`
}

type config struct {
	Port     string `conf:"default:3000"`
	Language string `conf:"default:pt-BR"`
	Project  projectConfig
}

func main() {

	var cfg config
	help, err := conf.Parse("", &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			log.Println(help)
			return
		}
		log.Printf("parsing config: %s", err.Error())
		return
	}

	starkbank.Language = cfg.Language

	// the private key is being read from environment variable and the line breaks are being replaced
	privatekey := strings.ReplaceAll(cfg.Project.PrivateKey, "\\n", "\n")

	user := project.Project{
		Id:          cfg.Project.ID,
		PrivateKey:  checks.CheckPrivateKey(privatekey),
		Environment: checks.CheckEnvironment(cfg.Project.Env),
	}

	starkbank.User = user

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/health"))
	r.Use(middleware.AllowContentType("application/json", "text/xml"))

	gtw := gateway.NewTransferGateway()
	uc := app.NewUseCase(gtw)
	h := api.NewRestHandler(r, uc)

	h.RegisterHandlers()

	http.Handle("/", r)

	// Start server.
	log.Printf("Starting server on port %s", cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
