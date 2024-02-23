package concurrency

import "sync"

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
	go func(c chan result) {
		var wg sync.WaitGroup
		defer close(c) // needs to be closed after wg.Wait()
		defer wg.Wait()
		for _, url := range urls {
			wg.Add(1)
			go func(u string) {
				// sending result value to channel
				resultChannel <- result{u, wc(u)}
				wg.Done()
			}(url)
		}
	}(resultChannel)

	// using range in channel
	// the channel must be closed after sending the last signal (see code above)
	for r := range resultChannel {
		// receiving value from channel
		results[r.string] = r.bool
	}

	return results
}
