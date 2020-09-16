package api

import (
	"net/url"

	hydra "github.com/ory/hydra-client-go/client"
)

func getHydraAdminClient(hydraURL string) *hydra.OryHydra {
	adminURL, _ := url.Parse(hydraURL)
	return hydra.NewHTTPClientWithConfig(
		nil,
		&hydra.TransportConfig{
			Schemes:  []string{adminURL.Scheme},
			Host:     adminURL.Host,
			BasePath: adminURL.Path,
		},
	)
}
