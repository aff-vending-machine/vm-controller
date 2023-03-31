package lugentpay

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

type apiImpl struct {
	client *http.Client
}

func New() *apiImpl {
	timeout := 5 * time.Minute

	// create a new cert for 'foo.com' to be used in the HTTPS server
	certPEM, keyPEM := makeCert("localhost")
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		log.Fatal("tls.X509KeyPair: ", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tlsConfig.BuildNameToCertificate()

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

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            rootCAs,
				InsecureSkipVerify: true,
				// InsecureSkipVerify: true, // this will make it 'work' but then the hostname isn't checked
			},
		},
		Timeout: timeout,
	}

	return &apiImpl{
		client: &client,
	}
}
