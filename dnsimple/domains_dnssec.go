package dnsimple

import (
	"context"
	"fmt"
)

// Dnssec represents the current DNSSEC settings for a domain in DNSimple.
type Dnssec struct {
	Enabled bool `json:"enabled"`
}

func dnssecPath(accountID string, domainIdentifier string) (path string) {
	path = fmt.Sprintf("%v/dnssec", domainPath(accountID, domainIdentifier))
	return
}

// dnssecResponse represents a response from an API method that returns a Dnssec struct.
type dnssecResponse struct {
	Response
	Data *Dnssec `json:"data"`
}

// EnableDnssec enables DNSSEC on the domain.
//
// See https://developer.dnsimple.com/v2/domains/dnssec/#enable

func (s *DomainsService) EnableDnssec(ctx context.Context, accountID string, domainIdentifier string) (*dnssecResponse, error) {
	path := versioned(dnssecPath(accountID, domainIdentifier))
	dnssecResponse := &dnssecResponse{}

	resp, err := s.client.post(ctx, path, dnssecResponse, nil)
	if err != nil {
		return nil, err
	}

	dnssecResponse.HttpResponse = resp
	return dnssecResponse, nil
}

// DisableDnssec disables DNSSEC on the domain.
//
// See https://developer.dnsimple.com/v2/domains/dnssec/#disable
func (s *DomainsService) DisableDnssec(ctx context.Context, accountID string, domainIdentifier string) (*dnssecResponse, error) {
	path := versioned(dnssecPath(accountID, domainIdentifier))
	dnssecResponse := &dnssecResponse{}

	resp, err := s.client.delete(ctx, path, dnssecResponse, nil)
	if err != nil {
		return nil, err
	}

	dnssecResponse.HttpResponse = resp
	return dnssecResponse, nil
}

// GetDnssec retrieves the current status of DNSSEC on the domain.
//
// See https://developer.dnsimple.com/v2/domains/dnssec/#get
func (s *DomainsService) GetDnssec(ctx context.Context, accountID string, domainIdentifier string) (*dnssecResponse, error) {
	path := versioned(dnssecPath(accountID, domainIdentifier))
	dnssecResponse := &dnssecResponse{}

	resp, err := s.client.get(ctx, path, dnssecResponse)
	if err != nil {
		return nil, err
	}

	dnssecResponse.HttpResponse = resp
	return dnssecResponse, nil
}
