package main

import (
	"context"
	"github.com/RyanTrue/GophKeeper/internal/config"
	"github.com/RyanTrue/GophKeeper/internal/repository"
	"github.com/RyanTrue/GophKeeper/internal/repository/postgres"
	"github.com/RyanTrue/GophKeeper/internal/server"
	servicesPkg "github.com/RyanTrue/GophKeeper/internal/services"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "03:04:05PM"})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	cfg := config.NewConfig("./config")
	log.Debug().Interface("config", cfg).Send()

	db, err := postgres.NewPostgres(ctx, cfg.ReposConfig.Postgres)
	if err != nil {
		log.Fatal().Err(err).Msg("Connecting to the Postgres database")
	}

	factory := postgres.NewFactory(db)
	repos := repository.NewRepository(factory)

	services := servicesPkg.NewServices(repos, cfg.ServerConfig.JWTSecret)

	coreServer := server.NewServer(
		cfg.ServerConfig.Address,
		services,
		cfg.ServerConfig.SSLCertPath,
		cfg.ServerConfig.SSLKeyPath,
	)

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		log.Info().Msg("The server has just started!")
		return coreServer.Run()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return coreServer.Shutdown()
	})

	if err = g.Wait(); err != nil {
		log.Info().Err(err).Msg("Reason for graceful shutdown")
	}

	log.Info().Msg("The application is shutdown")
}
