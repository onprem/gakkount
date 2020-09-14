package main

import (
	"net/url"

	hydra "github.com/ory/hydra-client-go/client"
)

func getHydraAdminClient() *hydra.OryHydra {
	adminURL, _ := url.Parse("http://localhost:4445")
	return hydra.NewHTTPClientWithConfig(
		nil,
		&hydra.TransportConfig{
			Schemes:  []string{adminURL.Scheme},
			Host:     adminURL.Host,
			BasePath: adminURL.Path,
		},
	)
}
