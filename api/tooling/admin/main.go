package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aborilov/hippo/api/tooling/admin/commands"
	"github.com/aborilov/hippo/business/sdk/sqldb"
	"github.com/ardanlabs/conf/v3"
)

var build = "develop"

type config struct {
	conf.Version
	Args conf.Args
	DB   struct {
		User         string `conf:"default:postgres"`
		Password     string `conf:"default:postgres,mask"`
		Host         string `conf:"default:database-service"`
		Name         string `conf:"default:postgres"`
		MaxIdleConns int    `conf:"default:0"`
		MaxOpenConns int    `conf:"default:0"`
		DisableTLS   bool   `conf:"default:true"`
	}
}

func main() {
	if err := run(); err != nil {
		if !errors.Is(err, commands.ErrHelp) {
			fmt.Println("msg", err)
		}
		os.Exit(1)
	}
}

func run() error {
	cfg := config{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
	}

	const prefix = "HIPPO"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}

		out, err := conf.String(&cfg)
		if err != nil {
			return fmt.Errorf("generating config for output: %w", err)
		}
		fmt.Println(out)

		return fmt.Errorf("parsing config: %w", err)
	}

	return processCommands(cfg.Args, cfg)
}

// processCommands handles the execution of the commands specified on
// the command line.
func processCommands(args conf.Args, cfg config) error {
	dbConfig := sqldb.Config{
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		Host:         cfg.DB.Host,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	}

	switch args.Num(0) {
	case "migrate":
		if err := commands.Migrate(dbConfig); err != nil {
			return fmt.Errorf("migrating database: %w", err)
		}

	case "seed":
		if err := commands.Seed(dbConfig); err != nil {
			return fmt.Errorf("seeding database: %w", err)
		}

	case "migrate-seed":
		if err := commands.Migrate(dbConfig); err != nil {
			return fmt.Errorf("migrating database: %w", err)
		}
		if err := commands.Seed(dbConfig); err != nil {
			return fmt.Errorf("seeding database: %w", err)
		}

	default:
		fmt.Println("migrate:    create the schema in the database")
		fmt.Println("seed:       add data to the database")
		fmt.Println("provide a command to get more help.")
		return commands.ErrHelp
	}

	return nil
}
