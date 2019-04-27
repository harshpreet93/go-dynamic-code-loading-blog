//go:generate go build -buildmode=plugin

package main

import "fmt"

func Run() {
	fmt.Println("hello, world from pluginC")
}