package http

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/aff-vending-machine/vm-controller/config"
)

func createClientWithCert(cfg config.HTTPConfig) *http.Client {
	timeout := time.Duration(cfg.TimeoutInSec) * time.Second

	// create a new cert for 'foo.com' to be used in the HTTPS server
	certPEM, keyPEM := makeCert(cfg.Host)
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		log.Fatal("tls.X509KeyPair: ", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	// tlsConfig.BuildNameToCertificate()

	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))
	srv.Config.TLSConfig = tlsConfig
	srv.StartTLS()
	defer srv.Close()

	// create a client for the request which has
	// the cert as the only rootCA
	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM([]byte(certPEM))

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            rootCAs,
				InsecureSkipVerify: true,
				// InsecureSkipVerify: true, // this will make it 'work' but then the hostname isn't checked
			},
		},
		Timeout: timeout,
	}
}

// makeCert creates a self-signed certificate for an HTTPS server
// which is valid for 'host' for one minute.
func makeCert(host string) (certPEM, keyPEM []byte) {
	priv, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Fatal("rsa.GenerateKey: ", err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1000),
		Subject: pkix.Name{
			Organization: []string{"At44"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Minute),

		IsCA:                  true,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{host},
	}

	cert, err := x509.CreateCertificate(rand.Reader, &template, &template, priv.Public(), priv)
	if err != nil {
		log.Fatal("x509.CreateCertificate: ", err)
	}

	certPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert})
	keyPEM = pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	return
}
