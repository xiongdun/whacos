package ccc

import (
	"reflect"
	"testing"
)

//
//func main() {
//
//}

func TestCheckWebsties(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	actualResults := Checkwebsites(mockWebsiteChecker, websites)
	want := len(websites)
	got := len(actualResults)

	if want != got {
		t.Fatalf("wnanted %v , got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("wanted %v , got %v", expectedResults, actualResults)
	}
}

type WebsiteChecker func(string) bool

func Checkwebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = wc(url)
	}
	return results
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}
