# WebScraper Golang

The project uses Chromedp to build web scrapers and automate web navigation tasks in Go.

## Prerequisites

- Go installed on the system
- Chrome or Chromium browser installed on the system

## Installation

1. Clone the project repository.
   `git clone https://github.com/wenealves10/webscraper-golang.git`
2. Navigate to the project directory.
   `cd webscraper-golang`
3. Install the Chromedp library using the following command:
   `go get github.com/chromedp/chromedp`

## Usage

1. Create a new Go file in the project directory and import the Chromedp library:

```go
    package main

    import (
        "context"
        "fmt"

        "github.com/chromedp/chromedp"
    )
```

2. Create a new context for Chromedp to manage browser actions:

```go
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()
```

3. Add Chromedp commands to navigate to a web page, interact with it, or extract information from it. For example, to navigate to a page and print its title:

```go
    var title string
    err := chromedp.Run(ctx,
        chromedp.Navigate("https://www.example.com"),
        chromedp.Title(&title),
    )
    if err != nil {
        fmt.Println("Error executing Chromedp:", err)
        return
    }
    fmt.Println("Page title:", title)
```

4. Execute the Go program to run the Chromedp actions:
   `go run main.go`

## Advanced Usage

The project can be customized to handle more complex scenarios by leveraging Chromedp's advanced features, such as custom headers, authentication, and proxies. Additionally, the project can be extended to extract more specific information from web pages or perform more complex actions.

## Contributing

Contributions to the project are welcome. To contribute, please create a pull request with your changes.

## License

This project is licensed under the MIT License.
