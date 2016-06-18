package file

import (
	"encoding/gob"
	"fmt"
	"os"
)

func SaveFile(m *Message) {

	// create a file
	dataFile, err := os.Create("integerdata.gob")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// serialize the data
	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(m)

	dataFile.Close()

}
