package dumbo

import "net/http"


type Resolver struct {
	Path string
	Handler http.HandlerFunc
}

type Schema struct {
	Path string
	Get bool
	Patch bool
	Post bool
	Delete bool
	Head bool
}

type Controller interface{
  Schema() *Schema
  Create() *Resolver
  Update() *Resolver
  Delete()*Resolver
  Read()*Resolver
  Headers() *Resolver
}


func register(c []Controller, mux map[string]*tree) {
  for _,v := range c {
    s := v.Schema()
    if(s.Get) { mux[http.MethodConnect].InsertNode(v.Read().Path, v.Read().Handler) }

	if(s.Post) { mux[http.MethodPost].InsertNode(v.Create().Path, v.Create().Handler) }

	if(s.Patch) { mux[http.MethodPatch].InsertNode(v.Update().Path, v.Update().Handler) }

	if(s.Head) { mux[http.MethodHead].InsertNode(v.Headers().Path, v.Headers().Handler) }

	if(s.Delete) { mux[http.MethodDelete].InsertNode(v.Delete().Path, v.Delete().Handler) }
  }
}