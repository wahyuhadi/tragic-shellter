/*
	author : @wahyuhadi
	Tragic-Shellter :>> Finding hidden gold in human mistake
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wahyuhadi/tragic-shellter/variable"
)

/*
	Main function
*/

func init() {
	_, err := variable.HandlerVar()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	fmt.Println(*variable.Name)
}
