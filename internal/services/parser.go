package services

import "gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/services/drivers"

type Parser struct {
	parsingDriver drivers.Parser
}

func NewParser(driver drivers.Parser) *Parser {
	return &Parser{
		parsingDriver: driver,
	}
}
