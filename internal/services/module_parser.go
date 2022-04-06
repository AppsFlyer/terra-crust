package services

import (
	"fmt"

	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/services/drivers"
	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/types"
)

type ModuleParser struct {
	parsingDriver drivers.Parser
}

func NewParser(driver drivers.Parser) *ModuleParser {
	return &ModuleParser{
		parsingDriver: driver,
	}
}

func (p *ModuleParser) GetModulesList(rootFolder string) (map[string]*types.Module, error) {
	modulesList, err := p.parsingDriver.Parse(rootFolder)
	if err != nil {
		fmt.Println("failed to parse", err.Error())
		return nil, err
	}

	return modulesList, nil
}
