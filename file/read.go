package file

import (
	"encoding/gob"
	"fmt"
	"os"
)

func ReadFile() (Message,error) {

	// open data file
	dataFile2, err := os.Open("integerdata.gob")

	if err != nil {
		// fmt.Println(err)
		return Message{},err
	}

	var mew Message
	dataDecoder := gob.NewDecoder(dataFile2)
	err = dataDecoder.Decode(&mew)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataFile2.Close()

	// fmt.Println(mew)

	return mew,nil

}
