package example

import (
	"fmt"
	"os"
)

/**
* NOTE: IsDocker() Function
* Used to detect whether application is running inside Docker or not.
* returns `boolean`
**/
func IsDocker() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}

/**
* NOTE: GetHelloWorld() Function
* Used to return a Hello World string based on whether
* instance is running inside Docker or not.
**/
func GetHelloWorld() string {
	if IsDocker() {
		return fmt.Sprintf(`
			こんにちわ from Go and Docker!
			This message shows that your Go application is running smoothly within a Docker container,
			demonstrating the seamless collaboration between the two.
			Edit /cmd/example/main.go to get started.
			`)
	} else {
		return fmt.Sprintf(`
			こんにちわ from Go!
			Your Go application is successfully running, delivering efficient
			and powerful code execution straight to your terminal.
			Edit /cmd/example/main.go to get started.
			`)
	}
}
