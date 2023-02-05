// Copyright 2022 AppsFlyer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parser_test

import (
	"github.com/AppsFlyer/terra-crust/internal/services/drivers/parser"
	"testing"

	logger "github.com/AppsFlyer/go-logger"
)

func TestParse(t *testing.T) {
	t.Parallel()
	log := logger.NewSimple()
	parserDriver := parser.NewTerraformParser(log)

	m, err := parserDriver.Parse("../../../../mock/modules")
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
	log := logger.NewSimple()
	parserDriver := parser.NewTerraformParser(log)

	m, err := parserDriver.Parse("../../../../internal")
	if err != nil {
		t.Errorf("failed to parse , reason: %s", err.Error())
	}

	if len(m) != 0 {
		t.Errorf("Failed to parse, expected 0 modules received : %d", len(m))
	}
}

func TestParseNotExistingPath(t *testing.T) {
	t.Parallel()
	log := logger.NewSimple()
	parserDriver := parser.NewTerraformParser(log)

	_, err := parserDriver.Parse("../../../internal")
	if err == nil {
		t.Errorf("Expected error for bad route")
	}
}
