/*
Copyright © 2024 @pol-cova
*/
package main

import (
	"github.com/pol-cova/GoGinit/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
