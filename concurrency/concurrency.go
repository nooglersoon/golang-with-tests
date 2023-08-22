package concurrency

import "fmt"

type WebsiteChecker func(string) bool
type result struct {
	// Key
	string
	// Value
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {

	results := make(map[string]bool)
	// Create a result channel
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Instead of access directly to results map, send a result struct to results channel
			// Send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	/*
		The next for loop iterates once for each of the urls. Inside we're using a receive expression, which assigns a value received from a channel to a variable.
		This also uses the <- operator, but with the two operands now reversed: the channel is now on the right and the variable that we're assigning to is on the left:
	*/

	// Receive expression to update the maps
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		fmt.Println(r)
		results[r.string] = r.bool
	}

	return results

}
