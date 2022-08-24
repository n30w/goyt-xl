package main

import "fmt"

// Quick error handling function. I'm lazy to write out that code every single time.
func Enil(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
