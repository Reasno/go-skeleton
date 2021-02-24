package bootstrap

import (
	"github.com/DoNewsCode/core/config"
	"github.com/DoNewsCode/core/otgorm"
	"github.com/nfangxu/core-skeleton/app"
	"github.com/nfangxu/core-skeleton/cmd"
	"math/rand"
	"time"

	"github.com/DoNewsCode/core"
	"github.com/spf13/cobra"
)

func Bootstrap() (*cobra.Command, *core.C) {
	// setup rand seeder
	rand.Seed(time.Now().UnixNano())

	// init rootCmd
	rootCmd := cmd.NewRootCmd()

	// setup core with config file path
	c := core.Default(core.WithYamlFile(rootCmd.GetCfgPath()))

	// register global dependencies
	providers(c)

	// register global modules
	modules(c)

	// register commands from global modules
	c.ApplyRootCommand(rootCmd.Command)

	return rootCmd.Command, c
}

func providers(c *core.C) {
	c.Provide(otgorm.Provide)
}

func modules(c *core.C) {
	c.AddModuleFunc(app.InjectKernel)
	c.AddModuleFunc(config.New)
	c.AddModuleFunc(core.NewServeModule)
	c.AddModuleFunc(otgorm.New)
}
