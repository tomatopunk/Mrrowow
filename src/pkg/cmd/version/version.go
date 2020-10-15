package version

import (
	"fmt"
	"strings"
)

//todo
func Format() string {

	version := "DEBUG"
	version = strings.TrimPrefix(version, "v")

	return fmt.Sprintf("cat table version %s\n", version)
}
