package main

import "github.com/tkitsunai/edinet-go/api/rest"

func main() {
	if err := rest.Run(); err != nil {
		panic(err)
	}
}
