package api

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/walter-manger/cert-stub/pkg/certifier"
)

// CertifierAPI represents the API for interacting with the Certifier package
type CertifierAPI struct {
	ExpirationInMinutes int
	RequestChannel      chan bool
	certifier           *certifier.Certifier
	port                string
}

// CertifierJobs returns Certifier's WgJobs
// (Because the consumer of this package will not have access -- A friend of a friend is not my friend)
func (capi *CertifierAPI) CertifierJobs() *sync.WaitGroup {
	return capi.certifier.WgJobs
}

// NewCertifierAPI returns a pointer to a new instance of CertifierAPI
func NewCertifierAPI(port string, expirationInMinutes int) *CertifierAPI {
	certifier := certifier.NewCertifier()
	return &CertifierAPI{
		ExpirationInMinutes: expirationInMinutes,
		RequestChannel: make(chan bool),
		certifier:      certifier,
		port:           port,
	}
}

// Start creates all the endpoints and listens on localhost:capi.port
func (capi *CertifierAPI) Start() {

	p := capi.port
	go func() {
		log.Printf("Certifier API accepting connections on port %s\n", p)
		log.Printf("Default certificate expiration time is %d minutes\n", capi.ExpirationInMinutes)
		http.HandleFunc("/cert/", capi.createHandler(capi.getCertificateHandler, "GET"))
		log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", p), nil))
	}()
}

// createHandler creates a generic handler in a closure to enforce verbs and respects the API Channel
func (capi *CertifierAPI) createHandler(handler http.HandlerFunc, method string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Gracefully shut down and stop all requests until the jobs have finished if
		// a shutdown is requested
		select {
		case <-capi.RequestChannel:
			http.Error(w, "Not accepting new connections", http.StatusLocked)
			return
		default:
		}

		if r.Method != method {
			http.Error(w, fmt.Sprintf("This endpoint only supports %s requests", method), http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	}
}

// getCertificateHandler returns a handler for the /cert/:domain path
func (capi *CertifierAPI) getCertificateHandler(w http.ResponseWriter, r *http.Request) {

	domain := r.URL.Path[len("/cert/"):]
	if len(domain) == 0 {
		http.Error(w, "domain is a required parameter", http.StatusUnprocessableEntity)
		return
	}

	cert := capi.certifier.GenerateCertificate(domain, time.Now().Add(time.Duration(capi.ExpirationInMinutes)*time.Minute))
	fmt.Fprintf(w, fmt.Sprintf("%s", cert))
}
