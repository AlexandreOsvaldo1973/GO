package main

import (
	"fmt"
	"strings"
)

/*func main(){
	fmt.Println("Hello Golanguinho!")
}

# In $GOPATH

/go
	/src
		/project1
		/project2
	/bin
		compiled-binary-1
		compiled-binary2
	/pkg
		objects-to-be-linked */

/*func main(){
	var n1 = 0
	var n2 = 0
	var n3 = 2
	var n4 = 5
	fmt.Println(n1 + n2 + n3 + n4)
}*/

func main(){
	m1 := "nome"
	m2 := "meu nome"
	fmt.Println(strings.Contains(m1, m2))
	fmt.Println(strings.ReplaceAll(m1, "m", "xi"))
}