package main

import (
	"fmt"
)

//Challenge from @davecheney https://twitter.com/davecheney/status/1237209166088790016?s=20

func main() {
	// 65 is the ascii code for "A" -- http://www.asciitable.com/
	s := string(65)
	fmt.Println(s)
	fmt.Println(len(s))
}
