package concurrency

type WebsiteChecker func(url string) bool
type result struct {
	string // anonymous name
	bool   // anonymous name
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// we can solve this data race by coordinating our goroutines using channels.
	// Channels are a Go data structure that can both receive and send values.
	// These operations, along with their details, allow communication between different processes
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// sending result value to channel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// receiving value from channel
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
