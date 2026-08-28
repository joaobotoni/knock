package khttp

import "strings"

func join(prefix, path string) string {
	return strings.TrimRight(prefix, "/") + "/" + strings.TrimLeft(path, "/")
}
