package main

import (
	"fmt"
	"github.com/my-org-org/golang-lib/dependency"
	"github.com/my-org-for-test/golang-local-lib/locallib"
)

func main() {
	fmt.Println("Project 1: DependencyExample")
	dependency.PrintDependencyMessage()

	fmt.Println("Local Deps Demo:")
	locallib.PrintFromLocalLib()
}
