package okx_connector

import (
	"os"
	"testing"
)

func requireIntegrationTests(t *testing.T) {
	t.Helper()
	if os.Getenv("OKX_RUN_INTEGRATION_TESTS") == "" {
		t.Skip("set OKX_RUN_INTEGRATION_TESTS=1 to run integration tests")
	}
}

