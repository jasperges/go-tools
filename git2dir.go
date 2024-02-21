// git2dir is a simple CLI tool to convert a git url (https or ssh) to a directory name
package main

import (
	"fmt"
	"os"
	"regexp"
)

// main passes the first CLI argument to urlToDir and prints the result
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gclone <url>")
		os.Exit(1)
	}
	url := os.Args[1]
	fmt.Print(urlToDir(url))
}

// urlToDir converts a git url to a directory name, it accepts both http and ssh urls
func urlToDir(url string) string {
	re := regexp.MustCompile(`(?:https://|git@)(.*?)(?:/|:)(.*)\.git$`)
	match := re.FindStringSubmatch(url)
	if match == nil {
		return ""
	}
	return fmt.Sprintf("%s/%s", match[1], match[2])
}
