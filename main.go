package main

import (
	"bufio"
	"clitut/file"
	"clitut/network"
	"clitut/work"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	if _,err := file.ReadFile(); err != nil {

		// Ask input from user
		fmt.Println("Enter your Email ID:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		email := scanner.Text()

		fmt.Println("Enter your password:")
		scanner.Scan()
		password := scanner.Text()

		resp, loginerr := network.Login(email, password)

		if resp != nil {
			m := new(file.Message)

			if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {

				fmt.Println(err)
			} else {

				// Check if this is a valid login
				if (file.Message{} != *m) {
					file.SaveFile(m)
					work.Execute()
				} else {
					fmt.Println("Not a Valid id pwd combination, Run the program again")
				}

			}
		} else {
			fmt.Println("inside error callback")
			fmt.Println(loginerr)
		}

	} else {
		work.Execute()
	}

}
