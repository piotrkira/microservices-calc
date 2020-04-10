package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

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
	s.r.HandleFunc("/add", addHandler())
	s.r.HandleFunc("/sub", subHandler())
	s.r.HandleFunc("/mul", mulHandler())
	s.r.HandleFunc("/div", divHandler())
}

func indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		log.Println("Index request")
		tmp, _ := template.ParseFiles("index.html")
		tmp.Execute(w, "")
	}
}

func addHandler() http.HandlerFunc {
	gatewayCli := client.New("gateway")
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		defer r.Body.Close()
		log.Println("Addition requested")
		params := mux.Vars(r)
		a, _ := strconv.ParseInt(params["a"], 10, 64)
		b, _ := strconv.ParseInt(params["b"], 10, 64)
		result, err := gatewayCli.Add(a, b)
		if err != nil || result == "" {
			fmt.Fprintf(w, "Addition is currently unaviable")
		}
		fmt.Fprintf(w, "%v", result)
	}
}

func subHandler() http.HandlerFunc {
	gatewayCli := client.New("gateway")
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		log.Println("Addition requested")
		params := mux.Vars(r)
		a, _ := strconv.ParseInt(params["a"], 10, 64)
		b, _ := strconv.ParseInt(params["b"], 10, 64)
		result, err := gatewayCli.Sub(a, b)
		if err != nil {
			fmt.Fprintf(w, "Subtraction is currently unaviable")
		}
		fmt.Fprintf(w, "%v", result)
	}
}

func mulHandler() http.HandlerFunc {
	gatewayCli := client.New("gateway")
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		log.Println("Addition requested")
		params := mux.Vars(r)
		a, _ := strconv.ParseInt(params["a"], 10, 64)
		b, _ := strconv.ParseInt(params["b"], 10, 64)
		result, err := gatewayCli.Mul(a, b)
		if err != nil {
			fmt.Fprintf(w, "Multiplication is currently unaviable")
		}
		fmt.Fprintf(w, "%v", result)
	}
}

func divHandler() http.HandlerFunc {
	gatewayCli := client.New("gateway")
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		log.Println("Addition requested")
		params := mux.Vars(r)
		a, _ := strconv.ParseInt(params["a"], 10, 64)
		b, _ := strconv.ParseInt(params["b"], 10, 64)
		result, err := gatewayCli.Div(a, b)
		if err != nil {
			fmt.Fprintf(w, "Division is currently unaviable")
		}
		fmt.Fprintf(w, "%v", result)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
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
