package example

import (
	"os"
)

// IsDocker reports whether the application is running inside Docker,
// detected by the presence of the `/.dockerenv` file.
//
// returns `boolean`
func IsDocker() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}

// GetHelloWorld returns a greeting that varies depending on whether
// the instance is running inside Docker.
//
// returns `string`
func GetHelloWorld() string {
	if IsDocker() {
		return `
			こんにちわ from Go and Docker!
			This message shows that your Go application is running smoothly within a Docker container,
			demonstrating the seamless collaboration between the two.
			Edit /cmd/example/main.go to get started.
			`
	}

	return `
		こんにちわ from Go!
		Your Go application is successfully running, delivering efficient
		and powerful code execution straight to your terminal.
		Edit /cmd/example/main.go to get started.
		`
}
