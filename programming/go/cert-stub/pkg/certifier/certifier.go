package certifier

import (
	"sync"
	"fmt"
	"time"
	"log"
)

// State represents certificate storage by way of map
// Example:
// {
//    "amazon.com": 2019/06/18 00:16:17,
//    "microsoft.com": 2019/06/19 01:16:05,
//    "google.com": 2019/06/20 02:18:00,
// }
type State struct {
	Certificates map[string]time.Time
}

// Certifier represents a certifier service
type Certifier struct {
	mu     *sync.Mutex
	state  State
	WgJobs *sync.WaitGroup
}

// NewCertifier initializes and returns a new Certifier instance.
func NewCertifier() *Certifier {
	return &Certifier{
		mu: &sync.Mutex{},
		state: State{
			Certificates: map[string]time.Time{},
		},
		WgJobs: &sync.WaitGroup{},
	}
}

// GenerateCertificate generates a certificate and adds it to the Certifier state
func (c *Certifier) GenerateCertificate(domain string, expiration time.Time) string {
	go c.certifyDomain(domain, expiration)
	log.Printf("%s: Requested certificate.", domain)
	// Always return the certificate even though there may be work to do in certifyDomain
	return fmt.Sprintf("foo-%s", domain)
}

// certifyDomain checks for a certificate before storing a certificate and setting the cleanup for expired domain certificates
func (c *Certifier) certifyDomain(domain string, expiration time.Time) {
	c.WgJobs.Add(1)
	c.mu.Lock()
	if _, ok := c.state.Certificates[domain]; !ok {
		log.Printf("%s: Storing certificate. Will expire on %s", domain, expiration.Format(time.Stamp))
		c.state.Certificates[domain] = expiration
		go c.expireCertificate(domain, expiration)
		log.Printf("%s: Sleeping for 10 seconds", domain)
		time.Sleep(10 * time.Second)
	}
	c.mu.Unlock()
	c.WgJobs.Done()
}

// expireCertificate removes the certificate when it is expired
func (c *Certifier) expireCertificate(domain string, expiration time.Time) {
	c.WgJobs.Add(1)
	duration := time.Until(expiration)
	time.Sleep(duration)
	c.mu.Lock()
	log.Printf("%s: certificate expired", domain)
	delete(c.state.Certificates, domain)
	c.mu.Unlock()
	c.WgJobs.Done()
}
