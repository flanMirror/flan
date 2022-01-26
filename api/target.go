package openapi

import (
	"fmt"
	"os"
)

var _ = func() interface{} {
	if len(os.Args) == 2 && os.Args[1] == "target" {
		fmt.Print(Target)
		os.Exit(0)
	}
	return nil
}()

const Target = "12.101.1"
