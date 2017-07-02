package httpreq

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

func CertTransport(certFile string, keyFile string, caFile string) (transport *http.Transport, err error) {
	// Load client cert
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	var caCertPool *(x509.CertPool)
	if caFile != "" {
		// Load CA cert
		caCert, err := ioutil.ReadFile(caFile)
		if err != nil {
			return nil, err
		}
		caCertPool = x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	return &http.Transport{TLSClientConfig: tlsConfig}, nil
}
