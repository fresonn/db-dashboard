package app

import (
	"context"
	"dashboard/api/gen/openapi"
	"dashboard/api/internal/config"
	"dashboard/api/internal/delivery/rest"
	"dashboard/api/internal/postgres"
	"dashboard/api/internal/scopes/cluster"
	"dashboard/api/internal/scopes/cluster/repo"
	"dashboard/api/internal/utils"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
)

type App struct {
	config config.AppConfig
	router http.Handler
	log    *slog.Logger
}

func New(cfg config.AppConfig, logger *slog.Logger) *App {

	pgManager := postgres.New(cfg, logger)

	clusterStorage := repo.New(cfg, logger, pgManager)

	cluserScope := cluster.New(cfg, logger, clusterStorage, pgManager)

	r := chi.NewRouter()
	r.Use(requestIDMiddleware)

	restHandler := rest.New(cluserScope)

	strictHandler := openapi.NewStrictHandler(restHandler, nil)

	handler := openapi.HandlerFromMuxWithBaseURL(strictHandler, r, "/api")

	return &App{
		config: cfg,
		log:    logger,
		router: handler,
	}
}

func (a *App) Run() {

	ctx, cancel := context.WithCancel(context.Background())

	server := &http.Server{
		Addr:    ":" + utils.IntToString(a.config.Env.Port),
		Handler: a.router,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if a.config.Env.IsDev {
		a.log.Info(fmt.Sprintf("http server is running http://localhost%s/api", server.Addr))
	}

	defer cancel()

	go func() {
		<-done
		cancel()
	}()

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	g.Go(func() error {
		<-gCtx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return server.Shutdown(ctx)
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		fmt.Printf("shutdown with error: %v", err)
	} else {
		fmt.Println("✅ server shutdown gracefully")
	}
}
