/*
Developer | ArkangeL43
Package   | json
Module    | json
File      | modg/json/json.go
Nest      | json

Does:
	Automates the net lookups for CNAME, MX, NS, TXT, etc
*/
package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	v "main/modg/colors"
	"main/modg/constants"
	ec "main/modg/warnings"
	"os"
)

// verified working and tested commands json
type JsonV struct {
	Commands []CmdJV `json:"Completed and Active Commands"`
}

type CmdJV struct {
	Command string `json:"Command"`
}

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

// flags json
type Flagsjson struct {
	Exec  string `json:"Flag"`
	Field string `json:"Category"`
	Desc  string `json:"Description"`
}

type JsonF struct {
	Flaghelp []Flagsjson `json:"Flags"`
}

// json verified
func Open_json_verified_commands(file string) {
	f, err := os.Open(file)
	ec.Ce(err, v.REDHB, "Could not open json file or help file", 1)
	defer f.Close()
	bv, _ := ioutil.ReadAll(f)
	var output JsonV
	json.Unmarshal(bv, &output)
	sum := 1
	for i := 0; i < len(output.Commands); i++ {
		sum++
		fmt.Println(v.REDHB+"\tCommand   | #", sum, " | -> "+v.REDHB+output.Commands[i].Command+v.RED)
	}
}

// json dev help or advanced help menu parsing function
func Open(file string) {
	jsonFile, err := os.Open(file)
	ec.Warning_advanced("<RR6> File Module: Could not open or read JSON file for parsing -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var Commands JsonM
	json.Unmarshal(byteValue, &Commands)
	for i := 0; i < len(Commands.Commands); i++ {
		fmt.Println("____________________________________________________________________")
		fmt.Println(v.RED + "\tCommand Feild    | -> " + v.REDHB + Commands.Commands[i].Field + v.RED)
		fmt.Println(v.RED + "\t                 |")
		fmt.Println(v.RED + "\tDescription      | -> " + v.REDHB + (Commands.Commands[i].Desc) + v.RED)
		fmt.Println(v.RED + "\t                 |")
		fmt.Println(v.RED + "\tCommand Flag     | -> " + v.REDHB + Commands.Commands[i].Exec + v.RED)
		fmt.Println(v.RED + "\t------------------ ")
		fmt.Print(v.RED + "\n")
	}

}

func Read_filepath(filepath string) string {
	newfilepath, err := constants.Parse_filepath(filepath)
	ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
	return newfilepath
}

//call help
func Help_usage_and_menus(isjson bool, filename, outcolor, format string, menu int) {
	if menu == 3 {
		// check json bool if false return false if else return true and output everything
		jsonFile, err := os.Open(filename)
		ec.Warning_advanced("<RR6> File [json] Module: Could not open JSON file for parsing -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var Flaghelp JsonF
		json.Unmarshal(byteValue, &Flaghelp)
		for i := 0; i < len(Flaghelp.Flaghelp); i++ {
			fmt.Print(v.RED + "\n")
			fmt.Println("____________________________________________________________________")
			fmt.Println(v.RED + "\tCommand Feild    | -> " + v.REDHB + Flaghelp.Flaghelp[i].Field + v.RED)
			fmt.Println(v.RED + "\t                 |")
			fmt.Println(v.RED + "\tDescription      | -> " + v.REDHB + (Flaghelp.Flaghelp[i].Desc) + v.RED)
			fmt.Println(v.RED + "\t                 |")
			fmt.Println(v.RED + "\tCommand Flag     | -> " + v.REDHB + Flaghelp.Flaghelp[i].Exec + v.RED)
			fmt.Println(v.RED + "\t------------------ ")
			fmt.Print(v.RED + "\n")
		}
	}
}
