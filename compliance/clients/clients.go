package clients

import (
	mock_client "github.com/flowgen/go-codon/testing/clients/mock/client"
)

var Mock = mock_client.NewHTTPClientWithConfigMap(nil, nil).Operations
