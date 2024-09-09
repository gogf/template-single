package main

import (
	_ "${MODULE_NAME}/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"${MODULE_NAME}/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
