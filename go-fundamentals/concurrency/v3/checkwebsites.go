package concurrency

import "time"

type WebsiteChecker func(url string) bool

func CheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {
	checked := make(map[string]bool)

	for _, url := range urls {

		// to run tests to check 'race condition' issue of concurrency process: go test -race .
		// in this case, two or more different goroutines performing writes on a map.

		// the only way to start a goroutine is to put 'go' in front of a function call
		//// we often use anonymous functions
		// to avoid reuse the same reference of url, we need pass the value to function
		// to goroutines have their own independent copy
		go func(u string) {
			checked[u] = checker(u)
		}(url)

		// try to fix the concurrency issue if a delay, but it will be too slow
		// Tests passed but this is not a correct approach
		time.Sleep(2 * time.Second)
	}
	return checked
}
