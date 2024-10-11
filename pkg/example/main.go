package example

/**
* NOTE: Welcome to the /pkg Directory!
*
* This is where you store library code that's safe for external applications to use (e.g., /pkg/mypubliclib).
* Other projects might import these libraries, so be mindful about what you place here.
* For better encapsulation, consider using the `/internal` directory, since Go enforces access restrictions there.
*
* The `/pkg` directory is an explicit way to communicate that the code here is available for others to use.
* Travis Jeffery's blog post, "I'll take pkg over internal," offers a great overview of how to choose between these directories.
* This directory can also help you organize Go code when your root directory has many non-Go components,
* making it easier to use Go tools effectively.
*
* Keep in mind, this is not a universal pattern—many repos use it, but many others don't.
* It's up to you! Just remember that more people will understand this layout than not.
* It's not a complex pattern, and once new Go developers grasp it, they'll find it simple to work with.
* For small projects, you can skip using `/pkg`, as the extra nesting might not add much value.
* But when your project grows and the root directory becomes crowded, especially with non-Go components, `/pkg` might come in handy.
**/

/**
* NOTE: Add(a, b int) Function
* Returns the sum of two integers.
* returns `int`
**/
func Add(a, b int) int {
	return a + b
}
