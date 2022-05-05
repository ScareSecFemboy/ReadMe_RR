package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonM struct {
	Commands []Command `json:"Commands"`
}

// User struct which contains a name
// a type and a list of social links
type Command struct {
	Exec  string `json:"command"`
	Field string `json:"Category"`
	Desc  string `json:"desc"`
}

func open(file string) {
	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var Commands JsonM

	json.Unmarshal(byteValue, &Commands)

	for i := 0; i < len(Commands.Commands); i++ {
		fmt.Println("\tCommand Feild    | -> " + Commands.Commands[i].Field)
		fmt.Println("\tDescription      | -> " + (Commands.Commands[i].Desc))
		fmt.Println("\tCommand          | -> " + Commands.Commands[i].Exec)
		fmt.Println("\t------------------ ")
	}

}

func main() {
	open("sets.json")
}
