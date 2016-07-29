package main

import "fmt"

type Fetcher interface {
    Fectch(url string) (string, []string, error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
    if depth <= 0 {
        return
    }

    body, urls, err := fetcher.Fectch(url)

    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    for _, u := range urls {
        Crawl(u, depth-1, fetcher)
    }
    return
}

func main() {
    Crawl("http://golang.org/", 4, fetcher)
}

type fakeFetcher map[string]*fetchResult

type fetchResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fectch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }

    return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
    "http://golang.org/": &fetchResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
        },
    },
    "http://golang.org/pkg/": &fetchResult{
        "packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt": &fetchResult{
        "Package fmt",
        []string{
            "http://golang.org",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fetchResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}


