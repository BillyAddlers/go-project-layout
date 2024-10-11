package example

/**
* NOTE: Welcome to the /internal Directory!
*
* This is where you store private application and library code—code you don't want others to import in their projects.
* Unlike other patterns, this one is enforced by the Go compiler, so no one can access these packages from outside your module.
* For more details, check out the Go 1.4 release notes.
*
* You’re not restricted to a single top-level `/internal` directory.
* You can create multiple `/internal` directories at any level in your project tree to organize code as needed.
*
* Optionally, you can add structure to separate shared and non-shared internal code.
* While it's not mandatory (especially in smaller projects), it provides a nice visual indicator of how packages are meant to be used.
* For example, put your main application code in `/internal/app` (e.g., /internal/app/myapp),
* and share common code between apps in `/internal/pkg` (e.g., /internal/pkg/myprivlib).
**/
/**
* NOTE: PrivateFunction() Function
* Basically does nothing.
* returns `void`
**/
func PrivateFunction() {
	// This function is limited only to this project.
	return
}
