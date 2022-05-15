package drivers

import "gitlab.appsflyer.com/real-time-platform/terra-crust/internal/types"

type Parser interface {
	Parse(path string) (map[string]*types.Module, error)
}
