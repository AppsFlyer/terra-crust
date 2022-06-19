package drivers

import "github.com/AppsFlyer/terra-crust/internal/types"

type Parser interface {
	Parse(path string) (map[string]*types.Module, error)
}
