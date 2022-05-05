package main

import (
	"fmt"
	"log"
	emsg "main/modules/go-main/web-server/errors"
	okmsg "main/modules/go-main/web-server/ok++"
	"net"
)

type Letters struct {
	Crypt string
}

const (
	//_____________________________________________________________________________________//
	Server_port_primary = 5501             // Primary port for the webserver and fileserver
	Server_port_backup  = 5502             // Backup port in case 5501 is not ready to be used
	Server_pass_primary = "RR6__ADMIN"     // Password for server authentication
	Server_host_primary = "localhost"      // Host for the webserver
	Server_Port_DBPrime = 5432             // Common postgreSQL port in the case we run a DB
	Server_WebU_primary = "127.0.0.1:5501" // This is the primary URL for the primary port
	Server_WebU_backup  = "127.0.0.1:5502" // If the primary port can not be used, this url will be set
	Server_port_backup2 = "127.0.0.1:8080" // standard HTTP
	//--------------------------------------------------------------------------------------------//
)

func T_port(portnum int) (bool, string, uint) {
	port := fmt.Sprint(rune(portnum))
	listener, e := net.Listen("tcp", ":"+port)
	if e != nil {
		fmt.Println(emsg.Server_TCP_Listener_FAIL, portnum)
		return false, port, 0x00
	} else {
		e := listener.Close()
		if e != nil {
			log.Fatal(okmsg.Server_Stat_Failed_To_Listen, e)
		} else {
			fmt.Println("<RR6> Server: Stat TCPL | Closed | TCP port listener closed")
		}
		return true, port, 0x00
	}
}

func Call_port() string {
	a, b, e := T_port(Server_port_primary)
	if a {
		if e != 0x00 {
			fmt.Println(emsg.Internal_Server_Error, e)
		} else {
			fmt.Println("<RR6> Server: Stat Port | ", b, " | Is good to use as a port")
			return b
		}
	} else {
		fmt.Println(okmsg.Server_Stat_Primary_Port_Good)
		a, b, e := T_port(Server_port_backup)
		if a {
			if e != 0x00 {
				fmt.Println(okmsg.Server_Stat_Backup_Port_Good)
				fmt.Println(b)
				return b
			} else {
				fmt.Println(emsg.Internal_Server_Error, e)
			}
		}
	}
	return b
}

func main() {
	a := Call_port()
	fmt.Println("<RR6> Server: Stat Port | ", a, " | Using current port for server ")

}
