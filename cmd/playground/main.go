package main

import (
	"fmt"

	"github.com/1024casts/1024casts/util"
)

func main() {

	// test markdown to html
	var md = "This is some sample code.\n\n```go\n" +
		`func main() {
	fmt.Println("Hi")
}
` + "```"

	testHtml := util.MarkdownToHtml(md)
	fmt.Printf("markdown:%s to html: %s", md, testHtml)
}
