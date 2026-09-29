package example

import (
	"os"
)

// isContainer reports whether the application is running inside a
// container runtime, detected via the markers each runtime leaves behind.
//
// Docker creates `/.dockerenv`. Podman creates `/run/.containerenv`.
// Neither marker is set when running directly on the host.
//
// returns `boolean`
func isContainer() bool {
	markers := []string{
		"/.dockerenv",        // Docker
		"/run/.containerenv", // Podman
	}

	for _, marker := range markers {
		if _, err := os.Stat(marker); err == nil {
			return true
		}
	}

	// Fall back to environment variables for runtimes started from a
	// service manager, which may not create the marker files above.
	return os.Getenv("CONTAINER_RUNTIME") == "podman" ||
		os.Getenv("RUNNING_IN_PODMAN") == "true"
}

// GetHelloWorld returns a greeting that varies depending on whether
// the instance is running inside a container.
//
// returns `string`
func GetHelloWorld() string {
	if isContainer() {
		return `
			こんにちわ from Go inside a container!
			This message shows that your Go application is running smoothly within a container,
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
