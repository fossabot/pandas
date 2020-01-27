//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package readers

import (
	"fmt"
	"net"

	"github.com/spf13/pflag"
)

type SecureServingOptions struct {
	Address net.IP
	Port    int

	// ServerCert is the TLS cert info for serving secure traffic
	ServerCert GeneratableKeyCert
}

type CertKey struct {
	// CertFile is a file containing a PEM-encoded certificate, and possibly the complete certificate chain
	CertFile string
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string
}

type GeneratableKeyCert struct {
	CertKey CertKey

	// CACertFile is an optional file containing the certificate chain for CertKey.CertFile
	CACertFile string
	// CertDirectory is a directory that will contain the certificates.  If the cert and key aren't specifically set
	// this will be used to derive a match with the "pair-name"
	CertDirectory string
	// PairName is the name which will be used with CertDirectory to make a cert and key names
	// It becomes CertDirector/PairName.crt and CertDirector/PairName.key
	PairName string
}

func NewSecureServingOptions(name string) *SecureServingOptions {
	return &SecureServingOptions{
		Address: net.ParseIP("0.0.0.0"),
		Port:    443,
		ServerCert: GeneratableKeyCert{
			PairName:      name,
			CertDirectory: name + ".local/config/certificates",
		},
	}
}

func (s *SecureServingOptions) Validate() []error {
	errors := []error{}

	if s.Port < 0 || s.Port > 65535 {
		errors = append(errors, fmt.Errorf("--secure-port %v must be between 0 and 65535, inclusive. 0 for turning off secure port.", s.Port))
	}

	return errors
}
func (s *SecureServingOptions) IsTlsEnabled() bool {
	return s.ServerCert.CertKey.CertFile != ""
}

func (s *SecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.IPVar(&s.Address, "reader-address", s.Address, ""+
		"The IP address on which to listen for the --secure-port port. The "+
		"associated interface(s) must be reachable by the rest of the cluster, and by CLI/web "+
		"clients. If blank, all interfaces will be used (0.0.0.0).")

	fs.IntVar(&s.Port, "reader-port", s.Port, ""+
		"The port on which to serve reader with authentication and authorization. If 0, "+
		"don't serve HTTPS at all.")

	fs.StringVar(&s.ServerCert.CertDirectory, "cert-dir", s.ServerCert.CertDirectory, ""+
		"The directory where the TLS certs are located. "+
		"If --tls-cert-file and --tls-private-key-file are provided, this flag will be ignored.")

	fs.StringVar(&s.ServerCert.CertKey.CertFile, "readers-tls-cert-file", s.ServerCert.CertKey.CertFile, ""+
		"File containing the default x509 Certificate for HTTPS. (CA cert, if any, concatenated "+
		"after server cert). If HTTPS serving is enabled, and --tls-cert-file and "+
		"--tls-private-key-file are not provided, a self-signed certificate and key "+
		"are generated for the public address and saved to /var/run/pandas.")

	fs.StringVar(&s.ServerCert.CertKey.KeyFile, "readers-tls-private-key-file", s.ServerCert.CertKey.KeyFile,
		"File containing the default x509 private key matching --tls-cert-file.")

	fs.StringVar(&s.ServerCert.CACertFile, "readers-tls-ca-file", s.ServerCert.CACertFile, "If set, this "+
		"certificate authority will used for secure access from Admission "+
		"Controllers. This must be a valid PEM-encoded CA bundle. Altneratively, the certificate authority "+
		"can be appended to the certificate provided by --tls-cert-file.")

}

func (s *SecureServingOptions) AddDeprecatedFlags(fs *pflag.FlagSet) {
	fs.IPVar(&s.Address, "public-address-override", s.Address,
		"DEPRECATED: see --bind-address instead.")
	fs.MarkDeprecated("public-address-override", "see --bind-address instead.")
}
