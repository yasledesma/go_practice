package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://alquilerargentina.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://alquilerargentina.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(fakeWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func fakeWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
    urls := make([]string, 100)

    for i := 0; i < len(urls); i++ {
        urls[i] = "a url"
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        CheckWebsites(SlowStubWebsiteChecker, urls)
    }
}

func SlowStubWebsiteChecker(_ string) bool {
    time.Sleep(20 * time.Millisecond)
    return true
}

