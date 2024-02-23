package concurrency

type WebsiteChecker func(url string) bool

func CheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {
	checked := make(map[string]bool)

	for _, url := range urls {
		checked[url] = checker(url)
	}

	return checked
}
