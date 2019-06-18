package main

import (
	"log"
	"flag"
	"os"
	"os/signal"

	"github.com/walter-manger/cert-stub/pkg/api"
)

var signalChan chan os.Signal

func init() {
	signalChan = make(chan os.Signal)
}

func shutdown(capi *api.CertifierAPI) {
	if capi == nil {
		return
	}
	<-signalChan
	close(capi.RequestChannel)
	log.Println("Cert-Stub Service Shutting Down Gracefully")
	capi.CertifierJobs().Wait()
}

func main() {
	flag.Usage = func() {
		flag.PrintDefaults()
	}

	log.Printf("Starting Cert-Stub Server...\n\n")

	port := flag.String("port", "8080", "The port to listen on")
	expirationInMinutes := flag.Int("expirationInMinutes", 1, "The expiration duration for all certificates")
	flag.Parse()

	capi := api.NewCertifierAPI(*port, *expirationInMinutes)
	capi.Start()

	signal.Notify(signalChan, os.Interrupt)
	shutdown(capi)
}
