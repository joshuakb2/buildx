package dap

import "github.com/joshuakb2/buildx/dap/common"

type LaunchConfig interface {
	GetConfig() common.Config
}
