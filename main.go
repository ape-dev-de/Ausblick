package main

import "fmt"

const (
	AppName    = "Ausblick"
	AppVersion = "1.0.0"
)

func getGreeting() string {
	return fmt.Sprintf("%s v%s", AppName, AppVersion)
}

func main() {
	fmt.Println(getGreeting())
}
