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
    
	routes map[string]*tree

	controllers []Controller

	handler http.Handler
}

type EngineParams struct {
	Addr         string
	IdleTimeout  int
	WriteTimeout int
	Middlewares  []func(http.Handler) http.Handler
	Controllers []Controller
	ErrorHandler func(error)
}

func Engine(e EngineParams) *dumbo {
	d := &dumbo{
		http:            &http.Server{Addr: e.Addr, IdleTimeout: time.Duration(e.IdleTimeout), WriteTimeout: time.Duration(e.WriteTimeout)},
		middlewares:     e.Middlewares,
		notFoundHandler: func(w http.ResponseWriter, r *http.Request) {},
		errorHandler:    e.ErrorHandler,
		pool:            &sync.Pool{},
		state:           false,
		handler: nil,
		routes:  map[string]*tree{
			"GET":     NewTree(),
			"POST":    NewTree(),
			"PUT":     NewTree(),
			"DELETE":  NewTree(),
			"PATCH":   NewTree(),
			"OPTIONS": NewTree(),
			"HEAD":    NewTree(),
		},
		controllers: e.Controllers,
	}
	d.Register(d.controllers...)

	return d;
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


func (d*dumbo) Register(c ...Controller) {
	if d.handler == nil {
		d.buildHandler()
	}
	for _,v := range c {
	  s := v.Schema()
	  if(s.Get) { d.routes[http.MethodConnect].InsertNode(v.Read().Path, v.Read().Handler) }
  
	  if(s.Post) { d.routes[http.MethodPost].InsertNode(v.Create().Path, v.Create().Handler) }
  
	  if(s.Patch) { d.routes[http.MethodPatch].InsertNode(v.Update().Path, v.Update().Handler) }
  
	  if(s.Head) { d.routes[http.MethodHead].InsertNode(v.Headers().Path, v.Headers().Handler) }
  
	  if(s.Delete) { d.routes[http.MethodDelete].InsertNode(v.Delete().Path, v.Delete().Handler) }
	}
  }	

  func (d *dumbo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, params := d.routes[r.Method].GetNode(r.URL.Path);
	print(params)
	if f != nil {
		if err := r.ParseForm(); err != nil {
			log.Printf("Error parsing form: %s", err)
			return
		}
		if d.handler != nil {
			d.handler.ServeHTTP(w, r)
		}
		f := f.handler
		f.ServeHTTP(w, r);

	} else {
		if d.notFoundHandler != nil {
			d.notFoundHandler.ServeHTTP(w,r)
		}else {
			http.NotFoundHandler().ServeHTTP(w,r)
		}
	}



  }
 


func (d *dumbo) buildHandler() {
	d.handler = chain(d.middlewares, http.HandlerFunc(middlewareHTTP))
}