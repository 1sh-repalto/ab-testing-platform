// utils/env.go
package utils

import "os"

func IsProd() bool {
	return os.Getenv("GIN_MODE") == "release"
}
