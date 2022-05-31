package services

import (
	logger "gitlab.appsflyer.com/go/af-go-logger/v1"
	"gitlab.appsflyer.com/real-time-platform/terra-crust/internal/services/drivers"
	"gitlab.appsflyer.com/real-time-platform/terra-crust/internal/types"
)

type ModuleParser struct {
	parsingDriver drivers.Parser
	logger        logger.Logger
}

func NewParser(logger logger.Logger, driver drivers.Parser) *ModuleParser {
	return &ModuleParser{
		parsingDriver: driver,
		logger:        logger,
	}
}

func (p *ModuleParser) GetModulesList(rootFolder string) (map[string]*types.Module, error) {
	modulesList, err := p.parsingDriver.Parse(rootFolder)
	if err != nil {
		p.logger.ErrorWithError("Failed parsing root folder", err)

		return nil, err
	}

	return modulesList, nil
}
