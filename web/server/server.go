package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/piotrkira/microservices-calc/gateway/client"

	"github.com/gorilla/mux"
)

// Server struct
type Server struct {
	r *mux.Router
}

// NewRouter assign rotings to paths
func NewRouter(s *Server) {
	s.r.HandleFunc("/", indexHandler())
}

func indexHandler() http.HandlerFunc {
	gatewayCli := client.New("gateway")
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		log.Println("Index request")
		add, err := gatewayCli.Add(2, 4)
		if err != nil {
			add = ""
		}
		sub, err := gatewayCli.Sub(2, 4)
		if err != nil {
			add = ""
		}
		mul, err := gatewayCli.Mul(2, 4)
		if err != nil {
			add = ""
		}
		div, err := gatewayCli.Div(2, 4)
		if err != nil {
			add = ""
		}
		fmt.Fprintf(w, "This is a calculator\n%v\n%v\n%v\n%v\n", add, sub, mul, div)
	}
}

// New returns new server
func New() *Server {
	s := Server{r: mux.NewRouter()}
	s.r.Host("127.0.0.1")
	NewRouter(&s)
	return &s
}

// Start starts server
func (s *Server) Start() {
	log.Println("Web server started")
	if err := http.ListenAndServe(":8080", s.r); err != nil {
		log.Println(err)
	}
}
