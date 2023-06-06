
package main

import (
	_ "net/http/pprof"
	"os"

	"github.com/ixbaseANT/gord/app"
)

func main() {
	if err := app.StartApp(); err != nil {
		os.Exit(1)
	}
}
