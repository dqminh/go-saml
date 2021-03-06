package saml

import (
	"errors"

	"github.com/RobotsAndPencils/go-saml/util"
)

// ServiceProviderSettings provides settings to configure server acting as a SAML Service Provider.
// Expect only one IDP per SP in this configuration. If you need to configure multipe IDPs for an SP
// then configure multiple instances of this module
type ServiceProviderSettings struct {
	PublicCertPath              string
	PrivateKeyPath              string
	IDPSSOURL                   string
	IDPSSODescriptorURL         string
	IDPPublicCertContent        string
	IDPPublicCertPath           string
	AssertionConsumerServiceURL string
	SPSignRequest               bool

	hasInit       bool
	publicCert    string
	privateKey    string
	iDPPublicCert string
}

type IdentityProviderSettings struct {
}

func (s *ServiceProviderSettings) Init() (err error) {
	if s.hasInit {
		return nil
	}

	if s.SPSignRequest {
		s.publicCert, err = util.LoadCertificate(s.PublicCertPath)
		if err != nil {
			return err
		}

		s.privateKey, err = util.LoadCertificate(s.PrivateKeyPath)
		if err != nil {
			return err
		}
	}

	if s.IDPPublicCertPath != "" {
		s.iDPPublicCert, err = util.LoadCertificate(s.IDPPublicCertPath)
		if err != nil {
			return err
		}
	} else if s.IDPPublicCertContent != "" {
		s.iDPPublicCert = util.SanitizeCertificate(s.IDPPublicCertContent)
	} else {
		return errors.New("missing idp public cert or its path")
	}

	s.hasInit = true
	return nil
}

func (s *ServiceProviderSettings) PublicCert() string {
	if !s.hasInit {
		panic("Must call ServiceProviderSettings.Init() first")
	}
	return s.publicCert
}

func (s *ServiceProviderSettings) PrivateKey() string {
	if !s.hasInit {
		panic("Must call ServiceProviderSettings.Init() first")
	}
	return s.privateKey
}

func (s *ServiceProviderSettings) IDPPublicCert() string {
	if !s.hasInit {
		panic("Must call ServiceProviderSettings.Init() first")
	}
	return s.iDPPublicCert
}
