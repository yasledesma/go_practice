package concurrency

type WebsiteChechker func(string) bool
type result struct {
    string
    bool
}

func CheckWebsites(wc WebsiteChechker, urls []string) map[string]bool {
    results := make(map[string]bool)
    resultChannel := make(chan result)

    for _, url := range urls {
        go func(u string) {
            resultChannel <- result{u, wc(u)}
        }(url)
    }

    for range urls {
        r := <- resultChannel
        results[r.string] = r.bool
    }

    return results
}
