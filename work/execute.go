package work

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func Execute() {
	text, err := clipboard.ReadAll()
	if err == nil {
		fmt.Println(text)
	} else {
		fmt.Println("Inside error copy")
		fmt.Println(err)
	}
}
