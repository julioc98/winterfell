package main

import (
	"errors"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkbank/sdk-go/starkbank/invoice"
	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/utils/checks"
)

type projectConfig struct {
	ID         string `conf:"default:6008875367006208"`
	PrivateKey string `conf:"required"`
	Env        string `conf:"default:sandbox"`
}

type config struct {
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

	generateInvoices()
}

func generateInvoices() {
	due := time.Now().Add(time.Hour * 4)
	var invoices []invoice.Invoice

	// random number 8 to 12 (inclusive)
	min := 8
	max := 13 // max is exclusive, so 13 is the limit to include 12
	rn := rand.Intn(max - min)
	num := rn + min

	log.Printf("Creating %d invoices...", num)

	for i := 0; i < num; i++ {
		invoices = append(invoices, invoice.Invoice{
			Amount:   rand.Intn(100000) + 1000, // min R$ 10,00
			Name:     gofakeit.Name(),
			TaxId:    "012.345.678-90",
			Fine:     5,   // 5%
			Interest: 2.5, // 2.5% per month
			Tags:     []string{"imediate"},
			Due:      &due,
		})
	}

	created, err := invoice.Create(invoices, nil)

	if err.Errors != nil {
		for _, e := range err.Errors {
			log.Printf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	log.Printf("Invoices created: %d/%d", len(created), num)
}
