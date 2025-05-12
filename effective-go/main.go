package main

import (
	"fmt"
)

type Owner struct {
	Name string
}

type Object struct {
	owner Owner
}

func (o *Object) Owner() Owner {
	return o.owner
}

func main() {
obj := Object {
	owner: Owner{Name: "Peter"},
}

owner := obj.Owner()
if owner.Name == "Tom" {
	fmt.Println("It's an owner")
}
}

