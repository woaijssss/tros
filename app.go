// Package tros 友行os
package tros

import (
	"context"
	"errors"
	trlogger "gitee.com/idigpower/tros/logx"
	"gitee.com/idigpower/tros/server/http"
	"gitee.com/idigpower/tros/trkit/mysqlx"
	"gitee.com/idigpower/tros/trkit/redisx"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
)

// App application
type App struct {
	signals []os.Signal
	//logger  trlogger.Adapter

	ctx    context.Context //nolint:contained ctx
	cancel func()

	initializers []Initializer
	servers      []Server
}

// SettingFunc of app
type SettingFunc func(*App)

type (
	// Server interface of server
	Server interface {
		// Start a server
		Start(ctx context.Context) error
		// Stop a server
		Stop() error
	}

	// Initializer interface with Init func
	Initializer interface {
		// Init component
		Init(atx AppContext) error
	}
)

// AppContext app context
type AppContext interface {
	// HTTPRouter http router
	HTTPRouter() http.Router

	// ServiceRegistrar register grpc service
	ServiceRegistrar() grpc.ServiceRegistrar
}

type appContext struct {
	router    http.Router
	registrar grpc.ServiceRegistrar
}

func New(settings ...SettingFunc) *App {
	ctx, cancel := context.WithCancel(context.Background())
	app := &App{
		signals: []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT},
		ctx:     ctx,
		cancel:  cancel,
	}

	for _, f := range settings {
		f(app)
	}

	err := app.Init()
	if err != nil {
		app.exit(err)
	}

	return app
}

func (app *App) Init() error {
	var router http.Router
	var registrar grpc.ServiceRegistrar

	for _, server := range app.servers {
		if r, ok := server.(http.Router); ok {
			router = r
			continue
		}

		if r, ok := server.(grpc.ServiceRegistrar); ok {
			registrar = r
		}
	}

	atx := newAppContext(router, registrar)
	for _, initializer := range app.initializers {
		err := initializer.Init(atx)
		if err != nil {
			return err
		}
	}

	mysqlx.InitMysqlX(app.ctx)
	redisx.Setup(app.ctx)

	return nil
}

func newAppContext(router http.Router, registrar grpc.ServiceRegistrar) *appContext {
	return &appContext{
		router:    router,
		registrar: registrar,
	}
}

// HTTPRouter returns http router
func (atx *appContext) HTTPRouter() http.Router {
	if atx.router == nil {
		panic("http transport not enabled")
	}
	return atx.router
}

// ServiceRegistrar returns gRpc service registrar
func (atx *appContext) ServiceRegistrar() grpc.ServiceRegistrar {
	if atx.registrar == nil {
		panic("gRpc transport not enabled")
	}
	return atx.registrar
}

// Servers register servers to app
func Servers(servers ...Server) SettingFunc {
	return func(app *App) {
		app.servers = servers
	}
}

// WithInitializers register initializers to app
func WithInitializers(initializers ...Initializer) SettingFunc {
	return func(app *App) {
		app.initializers = append(app.initializers, initializers...)
	}
}

func (app *App) exit(err error) {
	trlogger.Infof(app.ctx, "service exit", "error", err)
	//nolint
	os.Exit(1)
}

// Stop application
func (app *App) Stop() error {
	if app.cancel != nil {
		app.cancel()
	}
	return nil
}

// Run application
func (app *App) Run() {
	trlogger.Infof(app.ctx, "start application")
	defer func() {}()

	//go debugServerProcess(logger)
	//go func() {
	//	provider.DefaultProvider()
	//}()

	eg, ctx := errgroup.WithContext(app.ctx)
	for _, server := range app.servers {
		s := server
		eg.Go(func() error {
			// wait for stop signal
			<-ctx.Done()
			return s.Stop()
		})

		eg.Go(func() error {
			err := s.Start(ctx)
			if err != nil {
				trlogger.Infof(ctx, "failed to start server", "error", err)
				return err
			}
			return nil
		})
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, app.signals...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				//nolint
				_ = app.Stop()
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		app.exit(err)
	}

	trlogger.Infof(ctx, "service exit")
}
