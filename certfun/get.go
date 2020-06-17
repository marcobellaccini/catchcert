package certfun

import (
	"bytes"
	"crypto/tls"
	"encoding/pem"
)

// GetCertsPEM gets PEM certificate given an address string.
// Address string should be in the form "name:port".
// It returns cert and err.
// A TLS handshake will be started, in order to fetch server certificate.
// If no errors occurred, cert will contain
// the PEM-encoded string of the certificate
// and err will be nil.
func GetCertsPEM(address string) (string, error) {
	conn, err := tls.Dial("tcp", address, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return "", err
	}
	defer conn.Close()
	var b bytes.Buffer
	for _, cert := range conn.ConnectionState().PeerCertificates {
		err := pem.Encode(&b, &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: cert.Raw,
		})
		if err != nil {
			return "", err
		}
	}
	return b.String(), nil
}
