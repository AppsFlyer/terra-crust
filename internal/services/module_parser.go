package services

import (
	"gitlab.appsflyer.com/real-time-platform/terra-crust/internal/services/drivers"
	"gitlab.appsflyer.com/real-time-platform/terra-crust/internal/types"
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
		return nil, err
	}

	return modulesList, nil
}
