// module logs output to a file

package OFI

import (
	"fmt"
	"os"
)

func c_e_file(filename string) bool {
	inf, e := os.Stat(filename)
	if os.IsNotExist(e) {
		return false
	}
	return !inf.IsDir()
}

func Logger(data, filename string) (string, error) {
	p, e := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if e != nil {
		fmt.Println("<RR6> Logging Module: Could not read, write, or append to file file might not exist checking....")
		if !c_e_file(filename) {
			fmt.Println("<RR6> Logging Module: Creating file, filename did not exist")
			created, e := os.Create(filename)
			if e != nil {
				fmt.Println("<RR6> Logging Module: Was not able to create file???")
			}
			created.Close()
		}
	}
	fmt.Fprintln(p, data)
	return "", nil
}
