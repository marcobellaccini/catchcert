package certfun

import (
	"strings"
	"testing"
)

func TestGetCertsPEM(t *testing.T) {
	cases := []struct {
		in, expPrefix, expSuffix string
		returnsErr               bool
	}{

		{"github.com:443", "-----BEGIN CERTIFICATE-----\n", "----END CERTIFICATE-----\n", false},
		{"127.0.0.1:3210", "", "", true},
	}
	for _, c := range cases {
		cert, err := GetCertsPEM(c.in)

		if (err == nil && c.returnsErr == true) || (err != nil && c.returnsErr == false) {
			t.Errorf("GetCertsPEM unexpectedly returned err %q", err)
		} else if c.returnsErr == false {
			if !strings.HasPrefix(cert, c.expPrefix) || !strings.HasSuffix(cert, c.expSuffix) {
				t.Errorf("GetCertsPEM returned bad certificate %q", cert)
			}
		}
	}
}
