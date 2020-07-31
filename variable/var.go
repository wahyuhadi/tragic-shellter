/*
	author : @wahyuhadi
*/

package variable

import (
	"errors"
	"flag"
)

/*
	Define const
*/
const (
	CodeName = "Shellter Framework"
)

/*
	Dynamic variable. If your change the variable, constant not use
	Defult value for this variable is constants
*/
var (
	Name = flag.String("name", CodeName, "codename")
)

/*
	Handle Variable .
	execute in init() function in package main
*/
func HandlerVar() (string, error) {

	/*
		Example handle variable
	*/
	if *Name == "" {
		return "", errors.New("[!] Variable Name undifined")
	}

	return "", nil
}
