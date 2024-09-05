package main

import (
	_ "${Module_Name}/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"${Module_Name}/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
