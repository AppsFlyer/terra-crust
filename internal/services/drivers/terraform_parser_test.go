package drivers_test

import (
	"testing"

	"gitlab.appsflyer.com/real-time-platform/terra-crust/internal/services/drivers"
)

func TestParse(t *testing.T) {
	t.Parallel()
	parser := drivers.NewTerraformParser()

	m, err := parser.Parse("../../../mock/modules")
	if err != nil {
		t.Errorf("failed to parse , reason: %s", err.Error())
	}

	if len(m) != 2 {
		t.Errorf("Failed to parse, expected 2 modules received : %d", len(m))
	}

	if _, ok := m["consul_sync"]; !ok {
		t.Errorf("Expected consul_sync module to exist on the parsing , and not existing")
	}

	if _, ok := m["zookeeper"]; !ok {
		t.Errorf("Expected zookeeper module to exist on the parsing , and not existing")
	}
}

func TestParseBadPath(t *testing.T) {
	t.Parallel()
	parser := drivers.NewTerraformParser()

	m, err := parser.Parse("../../../internal")
	if err != nil {
		t.Errorf("failed to parse , reason: %s", err.Error())
	}

	if len(m) != 0 {
		t.Errorf("Failed to parse, expected 0 modules received : %d", len(m))
	}
}

func TestParseNotExistingPath(t *testing.T) {
	t.Parallel()
	parser := drivers.NewTerraformParser()

	_, err := parser.Parse("../../internal")
	if err == nil {
		t.Errorf("Expected error for bad route")
	}
}
