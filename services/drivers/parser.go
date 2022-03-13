package drivers

import "gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/types"

type Parser interface {
	Parse() (*types.Module, error)
}
