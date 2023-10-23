package dumbo

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

type dumbo struct {
	http *http.Server

	pool *sync.Pool

	notFoundHandler http.HandlerFunc

	middlewares []func(http.Handler) http.Handler

	errorHandler func(error)

	state bool
    
	routes     map[string]*tree
}

type EngineParams struct {
	Addr         string
	IdleTimeout  int
	WriteTimeout int
	middlewares  []func(http.Handler) http.Handler
	errorHandler func(error)
}

func Engine(e EngineParams) *dumbo {
	return &dumbo{
		http:            &http.Server{Addr: e.Addr, IdleTimeout: time.Duration(e.IdleTimeout), WriteTimeout: time.Duration(e.WriteTimeout)},
		middlewares:     e.middlewares,
		notFoundHandler: func(w http.ResponseWriter, r *http.Request) {},
		errorHandler:    e.errorHandler,
		pool:            &sync.Pool{},
		state:           false,
	}
}

func (d *dumbo) Deploy(cb callback) *dumbo {
	err := d.http.ListenAndServe()
	if err != nil {
		d.errorHandler(err)
		return &dumbo{}
	}
	d.state = true
	return &dumbo{}
}

func (d *dumbo) Http() *http.Server {
	return d.http
}

func (d *dumbo) Shutdown(ctx context.Context) error {
	log.Println("Stopping the server")
	return d.http.Shutdown(ctx)
}

