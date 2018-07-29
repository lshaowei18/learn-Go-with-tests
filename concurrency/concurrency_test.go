package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsite(t *testing.T) {
	t.Run("Multiple url no concurrent", func(t *testing.T) {
		urls := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}
		want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}
		got := CheckWebsite1(mockWebsiteChecker, urls)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Multiple url", func(t *testing.T) {
		urls := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}
		want := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}
		got := CheckWebsite(mockWebsiteChecker, urls)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func benchmarkCheckWebsite(cf CheckFunction, b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "waat://furhurterwe.geds"
	}
	for i := 0; i < b.N; i++ {
		cf(slowStubWebsiteChecker, urls)
	}
}

func BenchmarkCheckWebsite(b *testing.B)           { benchmarkCheckWebsite(CheckWebsite1, b) }
func BenchmarkCheckWebsiteConcurrent(b *testing.B) { benchmarkCheckWebsite(CheckWebsite, b) }
