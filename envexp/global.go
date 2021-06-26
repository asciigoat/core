package envexp

import (
	"os"
)

var envExpander = &Expander{
	get: os.Getenv,
}

// Expand uses os.GetEnv() to expand a given string
func Expand(s string) string {
	return envExpander.Expand(s)
}
