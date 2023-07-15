package concurrency

type WebsiteChecker func(string) bool

// structure for channels
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result) // chan of some type, in this case result(struct)

	for _, url := range urls {
		// the only way to start a goroutine is to put go in front of
		// a function call, we often use anonymous functions when we
		// want to start a goroutine
		// passing the url as argument, u is a copy of the value of url
		// and so can't be changed
		go func(u string) {

			// try write into a map without a treatment will generated a bug called race condition
			// Go has builtin race detector. Run the tests with the flag race: $ go test -race
			// to prevent we need use a special structure for that. We can use channel. Channel can
			// communicate between processes(send and receive values)
			// For that, using channel is much better than write directly into map
			// sending value to channel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// using channel instead of sleep
	// time.Sleep(2 * time.Second)

	for i := 0; i < len(urls); i++ {
		// receiving value from channel
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
