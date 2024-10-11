package example

import (
	"fmt"

	"example/pkg/example"
)

/**
* NOTE: Welcome to the /cmd Directory!
*
* This directory contains the main applications for your project.
* Each application directory should match the name of the executable you want to build (e.g., /cmd/myapp).
*
* Try not to place too much code in this directory.
* If the code can be reused in other projects, it belongs in the `/pkg` directory.
* If the code is project-specific and shouldn't be reused, place it in the `/internal` directory.
*
* Be clear about your intentions! Other developers might reuse your code, so make it explicit where reusable code should go.
* It's common practice to have a small main function here that imports and invokes code from `/internal` and `/pkg`, keeping the main logic elsewhere.
**/
func Run() {
	fmt.Println(example.GetHelloWorld())
}
