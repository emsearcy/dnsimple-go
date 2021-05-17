package dnsimple

import (
	"context"
	"fmt"
)

// EnableDomainAutoRenewal enables auto-renewal for the domain.
//
// See https://developer.dnsimple.com/v2/registrar/auto-renewal/#enable
func (s *RegistrarService) EnableDomainAutoRenewal(ctx context.Context, accountID string, domainName string) (*domainResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/auto_renewal", accountID, domainName))
	domainResponse := &domainResponse{}

	resp, err := s.client.put(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	domainResponse.HttpResponse = resp
	return domainResponse, nil
}

// DisableDomainAutoRenewal disables auto-renewal for the domain.
//
// See https://developer.dnsimple.com/v2/registrar/auto-renewal/#enable
func (s *RegistrarService) DisableDomainAutoRenewal(ctx context.Context, accountID string, domainName string) (*domainResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/domains/%v/auto_renewal", accountID, domainName))
	domainResponse := &domainResponse{}

	resp, err := s.client.delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	domainResponse.HttpResponse = resp
	return domainResponse, nil
}
