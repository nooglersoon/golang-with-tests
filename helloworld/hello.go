// Packages are ways of grouping up related Go code together.
package hello

// With import "fmt" we are importing a package which contains the Println function that we use to print.
import "fmt"

/*
The main function will be executed after all the initialization tasks have been performed.
In this function, you usually call other packages and implement your program logic.
*/
func init() {
	fmt.Println("launch initialization")
	fmt.Println(createHello("Aji", "spanish"))
}

// Is good to seperate domain code from the outside world (side-effects)
// Separate these conserns so its easier to test them
// Followed string is to tell the compiler that the followed function return a string data type
func createHello(name string, language string) string {

	if name == "" {
		name = "World"
	}
	// Declare prefix
	prefix := greetingPrefix(language)
	return prefix + name + "!"
}

// In our function signature we have made a named return value (prefix string).
// This will create a variable called prefix in your function.
// You can return whatever it's set to by just calling return rather than return prefix.
func greetingPrefix(name string) (prefix string) {
	switch name {
	case "english":
		prefix = englishHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	case "french":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// Defining a constants
// const englishHelloPrefix = "Hello "
// const spanishHelloPrefix = "Hola "
// const frenchHelloPrefix = "Bonjour "

// We can group constants in a block instead of declaring them each on their own line
const (
	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Bonjour "
)
