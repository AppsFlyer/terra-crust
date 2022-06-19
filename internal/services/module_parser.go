package services

import (
	logger "github.com/AppsFlyer/go-logger"
	"github.com/AppsFlyer/terra-crust/internal/services/drivers"
	"github.com/AppsFlyer/terra-crust/internal/types"
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
		p.logger.Error("Failed parsing root folder", err.Error())

		return nil, err
	}

	return modulesList, nil
}
