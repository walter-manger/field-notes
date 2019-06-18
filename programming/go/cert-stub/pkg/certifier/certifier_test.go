package certifier

import (
	"testing"
	"time"
)


func TestCertifierInitializes(t *testing.T) {
	c := NewCertifier()
	if c == nil {
		t.Errorf("Certifier instance is nil!")
	}
}

func TestGenerateCertificateReturnsCorrectValue(t *testing.T) {
	c := NewCertifier()
	cert := c.GenerateCertificate("some-domain", time.Now().Add(10 * time.Minute))

	if len(cert) == 0 {
		t.Errorf("Generated Certificate is empty")
	}

	if cert != "foo-some-domain" {
		t.Errorf("Generated certificate is not correct. Expected foo-some-domain but got %s", cert)
	}
}

func TestGenerateCertificateHandlesCertsForDifferentDomains(t *testing.T) {
	c := NewCertifier()

	tables := []struct {
		in string
		out string
	}{
		{"amazon.com", "foo-amazon.com"},
		{"google.com", "foo-google.com"},
		{"microsoft.com", "foo-microsoft.com"},
		{"microsoft.com", "foo-microsoft.com"},
	}

	for _, table :=  range tables {
		cert := c.GenerateCertificate(table.in, time.Now().Add(1 * time.Minute))
		if cert != table.out {
			t.Errorf("Expected %s got %s", table.out, cert)
		}
	}
}

func TestCertifyDomainStoresCertificate(t *testing.T) {
	c := NewCertifier()

	c.certifyDomain("amazon.com", time.Now().Add(1 * time.Minute))

	_, ok := c.state.Certificates["amazon.com"]
	if !ok {
		t.Errorf("amazon.com is not a certified domain")
	}
}
