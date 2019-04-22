// plugin package interfaces are EXPERIMENTAL.

package plugin

import (
	"github.com/qhenkart/gqlgen/codegen"
	"github.com/qhenkart/gqlgen/codegen/config"
)

type Plugin interface {
	Name() string
}

type ConfigMutator interface {
	MutateConfig(cfg *config.Config) error
}

type CodeGenerator interface {
	GenerateCode(cfg *codegen.Data) error
}
