package block

import (
	"crypto/x509"
	"encoding/pem"
	"strings"

	"github.com/pkg/errors"
)

// Cert cert
type Cert struct {
	Pem string `json:"pem,omitempty" db:"pem"`
	CN  string `json:"cn,omitempty" db:"cn"`
	OU  string `json:"ou,omitempty" db:"ou"`
	Org string `json:"org,omitempty" db:"org"`
}

// NewCert return new Cert
func NewCert(certRaw []byte) (*Cert, error) {
	c := &Cert{
		Pem: string(certRaw),
	}
	b, _ := pem.Decode([]byte(certRaw))
	if b == nil {
		return nil, errors.New("decode cert pem error")
	}
	cert, err := x509.ParseCertificate(b.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "parse cert error")
	}
	c.CN = cert.Subject.CommonName
	c.OU = strings.Join(cert.Subject.OrganizationalUnit, ",")
	c.Org = strings.Join(cert.Issuer.Organization, ",")
	return c, nil
}
