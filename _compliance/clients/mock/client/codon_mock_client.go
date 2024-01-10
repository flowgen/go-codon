package client

import (
	"github.com/flowgen/go-codon/testing/clients/mock/client/operations"
	strfmt "github.com/go-openapi/strfmt"
)

func NewHTTPClientWithConfigMap(formats strfmt.Registry, cfgmap map[string]string) *CodonMock {
	return New()
}

func New() *CodonMock {
	cli := new(CodonMock)
	cli.Operations = operations.New()
	return cli
}

type CodonMock struct {
	Operations *operations.Client
}
