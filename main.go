package main

import (
	"github.com/nfangxu/core-skeleton/bootstrap"
	"github.com/nfangxu/core-skeleton/cmd"
)

func main() {
	// bootstrap
	root, c := bootstrap.Bootstrap()

	// add custom command
	root.AddCommand(cmd.NewVersionCmd(c))

	// 运行
	(func() {
		defer c.Shutdown()
		_ = root.Execute()
	})()
}
