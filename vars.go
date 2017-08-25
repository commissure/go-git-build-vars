package main

import "fmt"

var (
	buildTime, gitRevision, gitBranch string
)

func main() {
	fmt.Printf("Built at\n\t%s\n", buildTime)
	fmt.Printf("Git revision\n\t%s\n", gitRevision)
	fmt.Printf("Git branch\n\t%s\n", gitBranch)
}
