package drivers

import "gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/types"

type Parser interface {
	Parse(path string) (map[string]*types.Module, error)
}
