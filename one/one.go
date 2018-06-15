package one

import (
	"InitFlow/two"
	"fmt"
)

var Number = 1

func init() {

	fmt.Println(`No.`, Number, `is ready`)
}

func i() int {
	return two.Number
}
