// +build integration

package main

import (
	"flag"
	"fmt"
	"testing"
	"net/http"
	"io/ioutil"
	"log"
	"time"
	"math/rand"
)

var addr = flag.String("addr", "http://localhost:8080", "The address to test against")

func makeCertRequest(addr *string, domain string) string {

	res, err := http.Get(fmt.Sprintf("%s/cert/%s", *addr, domain))

	if err != nil {
		log.Printf("Error: %s", err.Error())
		return ""
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}

	return string(body)
}


func TestCertifierServer(t *testing.T) {

	seed := rand.NewSource(time.Now().UnixNano())
	seededRand := rand.New(seed)

	tables := []struct {
		in string
		out string
	}{
		{"amazon.com", "foo-amazon.com"},
		{"google.com", "foo-google.com"},
		{"microsoft.com", "foo-microsoft.com"},
		{"amazon.com", "foo-amazon.com"},
		{"amazon.com", "foo-amazon.com"},
		{"duckduckgo.com", "foo-duckduckgo.com"},
		{".com", "foo-.com"},
		{"reddit.com", "foo-reddit.com"},
		{"certificate.com", "foo-certificate.com"},
		{"haskell.org", "foo-haskell.org"},
		{"emacs.org", "foo-emacs.org"},
		{"amazon.com", "foo-amazon.com"},
	}

	for _, table :=  range tables {
		if cert := makeCertRequest(addr, table.in); cert != table.out {
			t.Errorf("Expected %s, got %s", table.out, cert)

		}
		time.Sleep(time.Duration(seededRand.Intn(6)) * time.Second)
	}
}
