package switch_case

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	v "main/modg/colors"
	"main/modg/constants"
	opc "main/modg/copt"
	caller "main/modg/exec"
	execute "main/modg/exec"
	"main/modg/files"
	helpers "main/modg/helpers"
	jsa "main/modg/json"
	requests "main/modg/requests"
	IEEE_Sniff "main/modg/scripts/IEEE-802.11/Network-Capture"
	IEEEParse "main/modg/scripts/IEEE-802.11/pcap_parsing"
	IEEE_Utils "main/modg/scripts/IEEE-802.11/recon"
	IEEE "main/modg/scripts/IEEE-802.11/system-con"
	brutes "main/modg/scripts/brute-forcing"
	cloudfront_utils "main/modg/scripts/cloud/cloudfront/cloudfront-json"
	mcafe_utils "main/modg/scripts/cloud/mcafe/mcafe-json"
	forensics "main/modg/scripts/file-forensics"
	cr "main/modg/scripts/hashing/hashatk"
	generator "main/modg/scripts/hashing/hashgen"
	watchers "main/modg/scripts/httpr"
	"main/modg/scripts/lookup"
	stalk "main/modg/scripts/lookup"
	osintutils "main/modg/scripts/osint"
	phone_utils "main/modg/scripts/osint/phone"
	stego "main/modg/scripts/steg-osint"
	vuln "main/modg/scripts/vulnscans"
	win "main/modg/scripts/windows/pe"
	str "main/modg/sub"
	cv "main/modg/switch/casev"
	emod "main/modg/switch/errc"
	ios "main/modg/switch/scanners"
	"main/modg/warnings"
	ec "main/modg/warnings"
	"os"
	"os/exec"
	"time"

	"main/modg/system-runscript"
)

func Banner(file, color string) {
	content, err := ioutil.ReadFile(file)
	ec.Ce(err, v.RED, "Could not read file", 1)
	fmt.Println(constants.Clear_hex, color, string(content))
}

func Help(file, color string) {
	content, err := ioutil.ReadFile(file)
	ec.Ce(err, v.RED, "Could not read file", 1)
	fmt.Println(color, string(content))
}

func SSB(Screen_rotation string) {
	if Screen_rotation == "landscape" || Screen_rotation == "Landscape" {
		newfilepath, err := constants.Parse_filepath(constants.Team_logo_rr)
		ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		Banner(newfilepath, v.RED)
	}
	if Screen_rotation == "verticle" {
		newfilepath, err := constants.Parse_filepath(constants.Verticle_banner)
		ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		Banner(newfilepath, v.RED)
	}
	if Screen_rotation == "" {
		newfilepath, err := constants.Parse_filepath(constants.Verticle_banner)
		ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		Banner(newfilepath, v.RED)
	} else if Screen_rotation == "shark" {
		newfilepath, err := constants.Parse_filepath(constants.Shark)
		ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, err, 1, 233, "")
		Banner(newfilepath, v.RED)
	} else if Screen_rotation == "none" {
		fmt.Println(constants.Clear_hex)
		fmt.Println(v.BLKHB, "[+] <<< Starting red rabbit console >>>")
		fmt.Println(v.BLKHB, "[!] <<< Setting: Screen resolution set to `none` no output banner or format id specified >>>")
	}
}

// simple io
func sio(msg, color string) string {
	var s string
	fmt.Print(color, msg)
	fmt.Scanf("%s", &s)
	return s
}

func Parse_options_for_netcap(layer_to_capture string, flag bool, sniff string) {
	if flag {
		var filteryn string
		a := sio("Enter an interface to use >>>> ", v.GRN)
		fmt.Print("[+] Init Setting: {Interface} -> ", a, "\n")
		fmt.Print("[-] Would you like to use a BPF (Berkly Packet Filter) (y/n)>>>> ")
		fmt.Scanf("%s", &filteryn)
		if filteryn == "y" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter a filter ex (tcp and port 80):>>>> ")
			filt, _ := reader.ReadString('\n')
			IEEE_Sniff.Live_Run(a, 1024, false, true, filt, 30*time.Second, layer_to_capture)
		} else {
			IEEE_Sniff.Live_Run(a, 1024, false, false, "tcp and port 80", 30*time.Second, layer_to_capture)
		}
	} else {
		var filteryn string
		fmt.Print("[-] Would you like to use a BPF (Berkly Packet Filter) (y/n)>>>> ")
		fmt.Scanf("%s", &filteryn)
		if filteryn == "yes" {
			filt := sio("Enter a filter ex(tcp and port 80) >>> ", v.RED)
			IEEE_Sniff.Live_Run(sniff, 1024, false, true, filt, 30*time.Second, layer_to_capture)
		} else {
			IEEE_Sniff.Live_Run(sniff, 1024, false, false, "tcp and port 80", 30*time.Second, layer_to_capture)
		}
	}
}

// case test function parser
func M_TTY(command string, flags_rr6 *opc.RR6_options) {
	switch command {
	case "settings":
		fmt.Println(flags_rr6)
	case "commands":
		SSB(flags_rr6.Screen_rotation)
		if flags_rr6.Screen_rotation == "verticle" {
			nfp, e := constants.Parse_filepath("/text/menus/verticle_menu.txt")
			ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, e, 1, 233, "")
			Help(nfp, "\033[38m")
		} else {
			nfp, e := constants.Parse_filepath("/text/menus/landscape_menu.txt")
			ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, e, 1, 233, "")
			Help(nfp, "\033[38m")
		}
	case "h":
		SSB(flags_rr6.Screen_rotation)
		if flags_rr6.Screen_rotation == "verticle" {
			nfp, e := constants.Parse_filepath("/text/menus/verticle_menu.txt")
			ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, e, 1, 233, "")
			Help(nfp, "\033[38m")
		} else {
			nfp, e := constants.Parse_filepath("/text/menus/landscape_menu.txt")
			ec.Warning_advanced("<RR6> File Module: Could not open file for parsing -> ", v.REDHB, 1, false, false, true, e, 1, 233, "")
			Help(nfp, "\033[38m")
		}
	case "clear":
		SSB(flags_rr6.Screen_rotation)
	case "cls":
		SSB(flags_rr6.Screen_rotation)
	case "exit":
		fmt.Println("[+] :D")
		os.Exit(0)
	case "help flags":
		fp := constants.Flags
		newfp := jsa.Read_filepath(fp)
		fmt.Println(newfp)
		jsa.Help_usage_and_menus(true, newfp, v.RED, "\n", 3)
	case "help commands":
		fp := constants.Sets
		newfp := jsa.Read_filepath(fp)
		jsa.Open(newfp)
	case "help verified":
		fp := constants.Verified_commands
		newfp := jsa.Read_filepath(fp)
		jsa.Open_json_verified_commands(newfp)
	case "help advanced":
		helpers.Usage("text/help/advance.txt", v.RED)
	case "help all":
		helpers.Readall()
	case "load examples":
		files.Getfp()
	case "time":
		system.Time()
	//*************************************************************************************
	//*          This break concludes the file and help menu cases and inputs             *
	//*                                                                                   *
	//*************************************************************************************
	case "stalk mac":
		s, err := ios.Scanv(1, "Enter a MAC> ", v.REDHB, "e", cv.Scan_variable)
		emod.Return_error(err, "Could not scan line or input", 1, 2, false)
		system.Sep("\n\033[31m")
		stalk.Mac_trace(s)
	case "stalk oui":
		fmt.Println("- EX: macs.txt ")
		s, err := ios.Scanv(1, "Enter a MAC file> ", v.RET_RED, "e", cv.Filename)
		emod.Return_error(err, "Could not scan line or input", 1, 2, false)
		system.Sep("\n\033[31m")
		stalk.Oui_filename(s)
	case "stalk hosts":
		fmt.Println(" - EX: 10.0.0")
		s, err := ios.Scanv(1, "Enter 3 network octets> ", v.RET_RED, "e", cv.Subnet)
		emod.Return_error(err, "Could not scan line or input", 1, 2, false)
		system.Sep("\n\033[31m")
		stalk.Sub_main(s)
	//*************************************************************************************
	//*          This break concludes the stalk modules for MACS                          *
	//*                                                                                   *
	//*************************************************************************************
	case "search head":
		s, err := ios.Scanv(1, "Enter a URL> ", v.RET_RED, "e", cv.Url)
		emod.Return_error(err, "Could not scan line or input", 1, 2, false)
		system.Sep("\n\033[31m")
		stalk.LookHead(s)
	case "search ns":
		s, err := ios.Scanv(1, "Enter a domain> ", v.RET_RED, "e", cv.Domain)
		emod.Return_error(err, "Could not scan line or input", 1, 2, false)
		system.Sep("\n\033[31m")
		stalk.LookNS(s)
	case "search mx":
		s, err := ios.Scanv(1, "Enter a domain> ", v.RET_RED, "e", cv.Domain)
		emod.Return_error(err, "Could not scan line or input", 1, 2, false)
		system.Sep("\n\033[31m")
		stalk.LookMX(s)
	case "search endpoint":
		if flags_rr6.Filepath_general == "" {
			s, err := ios.Scanv(1, "Enter a filepath to a host file> ", v.RET_RED, "e", cv.Filename)
			emod.Return_error(err, "Could not scan line or input", 1, 2, false)
			input := stalk.Scan_target_file(s)
			results := stalk.RetrieveContents(str.Rdv(input))
			for _, elem := range results {
				fmt.Println(elem[1 : len(elem)-1])
			}
		} else {
			s, err := ios.Scanv(1, "Enter a filepath to a host file> ", v.RET_RED, "e", cv.Filename)
			emod.Return_error(err, "Could not scan line or input", 1, 2, false)
			input := stalk.Scan_target_file(s)
			results := stalk.RetrieveContents(str.Rdv(input))
			for _, elem := range results {
				fmt.Println(elem[1 : len(elem)-1])
			}
		}
	case "search ip":
		s, err := ios.Scanv(1, "Enter a host> ", v.RET_RED, "e", cv.Ipa)
		emod.Return_error(err, "Could not scan line or input", 1, 2, false)
		system.Sep("\n")
		stalk.LookIP(s)
		system.Sep("\n")
	case "search txt":
		s, err := ios.Scanv(1, "Enter a domain > ", v.RET_RED, "e", cv.Domain)
		emod.Return_error(err, "Could not scan line or input", 1, 2, false)
		system.Sep("\n\033[31m")
		stalk.Looktxt(s)
	case "search hostname":
		s, err := ios.Scanv(1, "Enter a IPA or host > ", v.RET_RED, "e", cv.Ipa)
		emod.Return_error(err, "Could not scan line or input", 1, 2, false)
		system.Sep("\n\033[31m")
		stalk.LookHSIP(s)
	case "search cname":
		s, err := ios.Scanv(1, "Enter a domain> ", v.RET_RED, "e", cv.Domain)
		emod.Return_error(err, "Could not scan line or input", 1, 2, false)
		system.Sep("\n\033[31m")
		cname, e := stalk.LookCNAME(string(s))
		if e != nil {
			fmt.Println("<RR6> Lookup Module: Can not lookup CNAME's - ", e)
		}
		fmt.Println("CNAME -> ", cname)
	case "search srv":
		fmt.Println("[+] Value1 should be -> domain")
		fmt.Println("[+] Value2 should be -> xmpp-server or server type")
		fmt.Println("[!] When done hit enter with no extra values, and example will be the following input\n[!] example.com xmpp-server")
		s, s2, err := ios.ScanV2(1, "", v.RET_RED, "", cv.Domain, cv.Server_lookup, "")
		if err != 0x00 {
			fmt.Println(v.REDHB, "<RR6> IO MODULE: Could not make a new scanner -> ", err)
		}
		stalk.LookSRV("tcp", s, s2)
	case "search robots":
		if flags_rr6.Url == "" {
			url, e := ios.Scanv(1, "Enter a URL> ", v.RET_RED, "e", cv.Domain)
			emod.Return_error(e, "Could not scan line or input", 1, 2, false)
			fmt.Println("[+] Setting: Initilized variable - ", url, " As a url to scan")
			stalk.LookRobot(url, "/robots.txt")
		} else {
			stalk.LookRobot(flags_rr6.Url, "/robots.txt")
		}
	case "search urls":
		url, e := ios.Scanv(1, "Enter a URL> ", v.RET_RED, "e", cv.Url)
		emod.Return_error(e, "Could not scan line or input", 1, 2, false)
		domain, e := ios.Scanv(1, "Enter a domain> ", v.RET_RED, "e", cv.Domain)
		emod.Return_error(e, "Could not scan line or input", 1, 2, false)
		httpurl, e := ios.Scanv(1, "Enter a HTTP url> ", v.RET_RED, "e", cv.Httpurl)
		emod.Return_error(e, "Could not scan line or input", 1, 2, false)
		lookup.Search_urls(url, domain, httpurl)
	case "search banner":
		ip, e := ios.Scanv(1, "Enter a host/ip to dial> ", v.RET_RED, "e", cv.Ipa)
		emod.Return_error(e, "Could not scan line or input", 1, 2, false)
		stalk.Lookup_banner_main(ip)
	case "search ipa":
		if flags_rr6.Url == "" {
			url, e := ios.Scanv(1, "Enter a URL> ", v.RET_RED, "e", cv.Url)
			emod.Return_error(e, "Could not scan line or input", 1, 2, false)
			lookup.LookIP(url)
		} else {
			lookup.LookIP(flags_rr6.Url)
		}
	case "search server":
		if flags_rr6.Url == "" {
			url, e := ios.Scanv(1, "Enter a URL>", v.RET_RED, "e", cv.Url)
			emod.Return_error(e, "Could not scan line or input", 1, 2, false)
			a, e := requests.GET_val(url, "GET", "server")
			emod.Return_error(e, "Could make a proper response", 1, 2, false)
			fmt.Println("\n[+] Server -> ", a)
		} else {
			a, e := requests.GET_val(flags_rr6.Url, "GET", "server")
			emod.Return_error(e, "Could make a proper response", 1, 2, false)
			fmt.Println("\n[+] Server -> ", a)
		}
	case "search route":
		caller.Traceroute()
	case "search admin":
		fmt.Println(flags_rr6.Url)
		fmt.Println(flags_rr6.PayloadList)
		// if ios.scanv does not work, then use sio, this works better there must be something wrong with the way the valuesd are being read and sent back in ios
		if flags_rr6.Url == "" || flags_rr6.PayloadList == "" {
			a := sio("Enter a URL> ", v.RET_RED)
			b := sio("Enter path to admin panel payload list or a payload list>", v.RET_RED)
			vuln.AdminFinder(a, b)
		} else {
			vuln.AdminFinder(flags_rr6.Url, flags_rr6.PayloadList)
		}
	case "search sqli":
		if flags_rr6.Url == "" || flags_rr6.PayloadList == "" {
			a := sio("Enter a URL> ", v.RET_RED)
			b := sio("Enter path to admin panel payload list or a payload list>", v.RET_RED)
			vuln.SQLIfinderMAIN(a, b)
		} else {
			vuln.SQLIfinderMAIN(flags_rr6.Url, flags_rr6.PayloadList)
		}
	case "search lfi":
		if flags_rr6.Url == "" {
			a := sio("Enter a URL> ", v.RET_RED)
			system.Call_perl_lfiscan(a)
		} else {
			system.Call_perl_lfiscan(flags_rr6.Url)

		}
	case "search sig":
		if flags_rr6.Input == "" {
			a := sio("Enter a Filepath or filename to be scanned> ", v.RET_RED)
			stego.File_sig(a)
		} else {
			stego.File_sig(flags_rr6.Input)
		}
	case "search filepath":
		if flags_rr6.Filepath_general == "" {
			a := sio("Enter a filepath to be searched> ", v.RET_RED)
			system.Walker_caller(a)
		} else {
			system.Walker_caller(flags_rr6.Filepath_general)
		}
	case "search archive":
		if flags_rr6.Input == "" {
			a := sio("Enter a JPG Image to be looked at> ", v.RET_RED)
			caller.Call_perl_s(a)
			forensics.Extract_ZIP(a)
		} else {
			caller.Call_perl_s(flags_rr6.Input)
			forensics.Extract_ZIP(flags_rr6.Input)
		}
	case "search geo":
		if flags_rr6.Input == "" {
			fmt.Println("WARN: FILE MUST BE JPG/JPEG")
			a := sio("Enter a JPG Image to be looked at> ", v.RET_RED)
			stego.Geo_loc(a)
		} else {
			stego.Geo_loc(flags_rr6.Input)
		}
		//*************************************************************************************
		//*          This break concludes the search modules for OSINT and web recon          *
		//*                                                                                   *
		//*************************************************************************************
	case "fetch links":
		if flags_rr6.Url == "" {
			a := sio("Enter a URL> \033[39m", v.BLKHB)
			r, e := watchers.Fetch_Links(a)
			warnings.Warning_simple("<RR6> UIO Module: Could not call watchers - ", v.REDHB, e)
			fmt.Println(r)
		} else {
			r, e := watchers.Fetch_Links(flags_rr6.Url)
			warnings.Warning_simple("<RR6> UIO Module: Could not call watchers - ", v.REDHB, e)
			fmt.Println(r)
		}
	case "fetch title":
		if flags_rr6.Url == "" {
			a := sio("Enter a URL> \033[39m", v.BLKHB)
			r, e := watchers.Fetch_Title(a)
			warnings.Warning_simple("<RR6> UIO Module: Could not call watchers - ", v.REDHB, e)
			fmt.Println(r)
		} else {
			r, e := watchers.Fetch_Title(flags_rr6.Url)
			warnings.Warning_simple("<RR6> UIO Module: Could not call watchers - ", v.REDHB, e)
			fmt.Println(r)
		}
	case "fetch lists":
		if flags_rr6.Url == "" {
			a := sio("Enter a URL> \033[39m", v.BLKHB)
			r, e := watchers.Fetch_LI(a)
			warnings.Warning_simple("<RR6> UIO Module: Could not call watchers - ", v.REDHB, e)
			fmt.Println(r)
		} else {
			r, e := watchers.Fetch_LI(flags_rr6.Url)
			warnings.Warning_simple("<RR6> UIO Module: Could not call watchers - ", v.REDHB, e)
			fmt.Println(r)
		}
	case "fetch documents":
		if flags_rr6.Url == "" {
			a := sio("Enter a URL> \033[39m", v.BLKHB)
			r, e := watchers.Fetch_doc(a)
			warnings.Warning_simple("<RR6> UIO Module: Could not call watchers - ", v.REDHB, e)
			fmt.Println(r)
		} else {
			r, e := watchers.Fetch_doc(flags_rr6.Url)
			warnings.Warning_simple("<RR6> UIO Module: Could not call watchers - ", v.REDHB, e)
			fmt.Println(r)
		}

		//*************************************************************************************
		//*          This break concludes the Watchers for HTML recon / parsing               *
		//*                                                                                   *
		//*************************************************************************************
	case "inject zip":
		if flags_rr6.Input == "" || flags_rr6.Filepath_general == "" || flags_rr6.Output == "" {
			a := sio("Enter a JPG Image to be injected> ", v.RET_RED)
			system.Sep("\n")
			b := sio("Enter a ZIP file to be injected into the image> ", v.RET_RED)
			system.Sep("\n")
			c := sio("Enter a output file> ", v.RET_RED)
			caller.Call_perl_s(a)
			forensics.Inject_ZIP(a, b, c)
		} else {
			caller.Call_perl_s(flags_rr6.Input)
			forensics.Inject_ZIP(flags_rr6.Input, flags_rr6.Filepath_general, flags_rr6.Output)
		}
	case "dump bootsec info":
		if flags_rr6.Filepath_general == "" {
			a := sio("Enter a Boot sec filepath> ", v.RET_RED)
			forensics.Dump_boots(a)
		} else {
			forensics.Dump_boots(flags_rr6.Filepath_general)
		}
	case "dump file info":
		if flags_rr6.Filepath_general == "" {
			a := sio("Enter a filepath or file to be dumped> ", v.RET_RED)
			system.File_inf(a)
		} else {
			system.File_inf(flags_rr6.Filepath_general)
		}
	case "dump file binary":
		if flags_rr6.Filepath_general == "" {
			a := sio("Enter a file to be dumped into hex> ", v.RET_RED)
			stego.Hex_dump(a)
		} else {
			stego.Hex_dump(flags_rr6.Filepath_general)
		}
	case "dump pe info":
		if flags_rr6.Filepath_general == "" {
			a := sio("Enter a PE file to be dumped> ", v.RET_RED)
			inf, e := win.Parser(a)
			if e != nil {
				fmt.Println("<RR6> Windows Module: Could not get information on PE file, parse file, or open file -> ", e)
			} else {
				fmt.Println(inf)
			}
		} else {
			inf, e := win.Parser(flags_rr6.Filepath_general)
			if e != nil {
				fmt.Println("<RR6> Windows Module: Could not get information on PE file, parse file, or open file -> ", e)
			} else {
				fmt.Println(inf)
			}
		}
	case "dump image metadata":
		if flags_rr6.Input == "" {
			a := sio("Enter a filepath to an image to dump> ", v.RET_RED)
			caller.Call_perl_s(a)
		} else {
			caller.Call_perl_s(flags_rr6.Input)
		}
	case "run RR6 GUI":
		caller.Call_perl_RR6_GUI()
	case "run RR6 scan gui":
		if flags_rr6.Iprange == "" {
			a := sio("Enter a CIDR> ", v.RET_RED)
			caller.Call_perl_GUI_host_scanner(a)
		} else {
			caller.Call_perl_GUI_host_scanner(flags_rr6.Iprange)
		}
		//*************************************************************************************
		//*          This break concludes the dumping and file information/sys calls          *
		//*                                                                                   *
		//*************************************************************************************
	case "inject bmp":
		if flags_rr6.Input == "" || flags_rr6.Payload == "" {
			a := sio("Enter a image to   inject> ", v.RET_RED)
			system.Sep("\n")
			b := sio("Enter a payload to inject> ", v.RET_RED)
			system.Sep("\n")
			caller.Run_bmp(b, a)
			caller.Call_perl_s(a)
		} else {
			caller.Run_bmp(flags_rr6.Payload, flags_rr6.Input)
			caller.Call_perl_s(flags_rr6.Input)
		}
	case "inject jpg":
		if flags_rr6.Input == "" || flags_rr6.Payload == "" || flags_rr6.Type == "" {
			a := sio("Enter a image to   inject> ", v.RET_RED)
			system.Sep("\n")
			b := sio("Enter a payload to inject> ", v.RET_RED)
			system.Sep("\n")
			fmt.Println("\033[32mChunks -> {'COM' or 'DQT'}")
			c := sio("\033[31mEnter a chunk to   inject> ", v.RET_RED)
			caller.Run_JPG(c, b, a)
			caller.Call_perl_s(a)
		} else {
			caller.Run_JPG(flags_rr6.Jpgchunk, flags_rr6.Payload, flags_rr6.Input)
			caller.Call_perl_s(flags_rr6.Input)
		}
	case "inject gif":
		if flags_rr6.Input == "" || flags_rr6.Payload == "" || flags_rr6.Pheight == "" || flags_rr6.Pwidth == "" {
			a := sio("Enter a image to   inject> ", v.RET_RED)
			system.Sep("\n")
			b := sio("Enter a payload to inject> ", v.RET_RED)
			system.Sep("\n")
			c := sio("Enter a pixel height> ", v.RET_RED)
			system.Sep("\n")
			d := sio("Enter a pixel width> ", v.RET_RED)
			caller.Run_GIF(a, b, c, d)
			caller.Call_perl_s(a)
		} else {
			caller.Run_GIF(flags_rr6.Input, flags_rr6.Payload, flags_rr6.Pheight, flags_rr6.Pwidth)
			caller.Call_perl_s(flags_rr6.Input)
		}
	case "inject webp":
		if flags_rr6.Input == "" || flags_rr6.Payload == "" || flags_rr6.Pheight == "" || flags_rr6.Pwidth == "" {
			a := sio("Enter a image to   inject> ", v.RET_RED)
			system.Sep("\n")
			b := sio("Enter a payload to inject> ", v.RET_RED)
			system.Sep("\n")
			caller.Run_Webp(a, b)
			caller.Call_perl_s(a)
		} else {
			caller.Run_Webp(flags_rr6.Input, flags_rr6.Payload)
			caller.Call_perl_s(flags_rr6.Input)
		}
	//*************************************************************************************
	//*          This break concludes the injection / stegonography section               *
	//*                                                                                   *
	//*************************************************************************************
	case "ping udp":
		if flags_rr6.Iprange == "" {
			a := sio("Enter a CIDR> ", v.RET_RED)
			execute.Host_discover_r6("discover_udp", a)
		} else {
			execute.Host_discover_r6("discover_udp", flags_rr6.Iprange)
		}
	case "ping tcp":
		if flags_rr6.Iprange == "" {
			a := sio("Enter a CIDR> ", v.RET_RED)
			execute.Host_discover_r6("discover_tcp", a)
		} else {
			execute.Host_discover_r6("discover_tcp", flags_rr6.Iprange)
		}
	case "ping syn":
		if flags_rr6.Iprange == "" {
			a := sio("Enter a CIDR> ", v.RET_RED)
			execute.Host_discover_r6("discover_syn", a)
		} else {
			execute.Host_discover_r6("discover_syn", flags_rr6.Iprange)
		}
	case "ping icmp":
		if flags_rr6.Iprange == "" {
			a := sio("Enter a CIDR> ", v.RET_RED)
			execute.Host_discover_r6("discover_icmp", a)
		} else {
			execute.Host_discover_r6("discover_icmp", flags_rr6.Iprange)
		}
	case "ping arp":
		IEEE_Utils.Return_Values_andCall()

	//*************************************************************************************
	//*          This break concludes the network recon / scanning sectuion               *
	//*                                                                                   *
	//*************************************************************************************
	case "crack sha1 list":
		if flags_rr6.Hashlist == "" || flags_rr6.Brute_list == "" {
			a := sio("Enter a Hash list> ", v.RET_RED)
			b := sio("Enter a Wordlist > ", v.RET_RED)
			cr.Brute_SHA1_wordlist(b, a)
		} else {
			cr.Brute_SHA1_wordlist(flags_rr6.Brute_list, flags_rr6.Hashlist)
		}
	case "crack sha1 single":
		a := sio("Enter a SHA1 hash> ", v.RET_RED)
		b := sio("Enter a Wordlist > ", v.RET_RED)
		cr.Brute_SHA1_single(b, a)
	case "crack md5 list":
		if flags_rr6.Hashlist == "" || flags_rr6.Brute_list == "" {
			a := sio("Enter a Hash list> ", v.RET_RED)
			b := sio("Enter a Wordlist > ", v.RET_RED)
			cr.Brute_MD5_wordlist(b, a)
		} else {
			cr.Brute_MD5_wordlist(flags_rr6.Brute_list, flags_rr6.Hashlist)
		}
	case "crack md5 single":
		a := sio("Enter a MD5 hash> ", v.RET_RED)
		b := sio("Enter a Wordlist > ", v.RET_RED)
		cr.Brute_MD5_Single(b, a)
	case "crack sha256 list":
		if flags_rr6.Hashlist == "" || flags_rr6.Brute_list == "" {
			a := sio("Enter a Hash list> ", v.RET_RED)
			b := sio("Enter a Wordlist > ", v.RET_RED)
			cr.Brute_SHA256_main(a, b)
		} else {
			cr.Brute_SHA256_main(flags_rr6.Hashlist, flags_rr6.Brute_list)
		}
	case "crack mysql vb":
		if flags_rr6.Brute_list == "" || flags_rr6.Filepath_general == "" {
			a := sio("Enter a path to the wordlist> ", v.RET_RED)
			b := sio("Enter a path to the SQL File> ", v.RET_RED)
			helpers.Call_sql(a, b)
		} else {
			helpers.Call_sql(flags_rr6.Brute_list, flags_rr6.Filepath_general)
		}
	case "crack sha256 single":
		a := sio("Enter a SHA256 hash> ", v.RET_RED)
		b := sio("Enter a Wordlist > ", v.RET_RED)
		cr.Brute_SHA256_main_Single(b, a)
	case "encode md5":
		a := sio("Enter a string> ", v.RET_RED)
		generator.Call_all("md5", "", a)
	case "encode sha1":
		a := sio("Enter a string> ", v.RET_RED)
		generator.Call_all("sha1", "", a)
	case "encode sha256":
		a := sio("Enter a string> ", v.RET_RED)
		generator.Call_all("sha256", "", a)
	case "encode sha512":
		a := sio("Enter a string> ", v.RET_RED)
		generator.Call_all("sha512", "", a)
	case "encode base64":
		a := sio("Enter a string> ", v.RET_RED)
		generator.Call_all("base64", "", a)
	case "encode base32":
		a := sio("Enter a string> ", v.RET_RED)
		generator.Call_all("base32", "", a)
	case "encode rot13":
		a := sio("Enter a string> ", v.RET_RED)
		generator.Call_all("rot13", "", a)
	case "encode HMAC":
		a := sio("Enter a string> ", v.RET_RED)
		b := sio("Enter a key   > ", v.RET_RED)
		generator.Call_all("HMAC", a, b)
	// file type generation, file.txt all hashes or words in that file will be hashed
	case "encode md5 list":
		if flags_rr6.Hashlist == "" || flags_rr6.Output == "" {
			a := sio("Enter a wordlist to hash> ", v.RET_RED)
			b := sio("Enter a output list to out data to> ", v.RET_RED)
			fmt.Println("--------------------------------------------------------")
			fmt.Println("[+] Converting all words in the list to a MD5 hash ---")
			fmt.Print("\n")
			generator.Listed_Generation(a, b, 1)
		} else {
			fmt.Println("--------------------------------------------------------")
			fmt.Println("[+] Converting all words in the list to a MD5 hash ---")
			fmt.Print("\n")
			generator.Listed_Generation(flags_rr6.Hashlist, flags_rr6.Output, 1)
		}
	case "encode sha1 list":
		if flags_rr6.Hashlist == "" || flags_rr6.Output == "" {
			a := sio("Enter a wordlist to hash> ", v.RET_RED)
			b := sio("Enter a output list to out data to> ", v.RET_RED)
			fmt.Println("--------------------------------------------------------")
			fmt.Println("[+] Converting all words in the list to a SHA1 hash ---")
			fmt.Print("\n")
			generator.Listed_Generation(a, b, 2)
		} else {
			fmt.Println("--------------------------------------------------------")
			fmt.Println("[+] Converting all words in the list to a SHA1 hash ---")
			fmt.Print("\n")
			generator.Listed_Generation(flags_rr6.Hashlist, flags_rr6.Output, 2)
		}
	case "encode sha256 list":
		if flags_rr6.Hashlist == "" || flags_rr6.Output == "" {
			a := sio("Enter a wordlist to hash> ", v.RET_RED)
			b := sio("Enter a output list to out data to> ", v.RET_RED)
			fmt.Println("--------------------------------------------------------")
			fmt.Println("[+] Converting all words in the list to a SHA256 hash ---")
			fmt.Print("\n")
			generator.Listed_Generation(a, b, 3)
		} else {
			fmt.Println("--------------------------------------------------------")
			fmt.Println("[+] Converting all words in the list to a SHA256 hash ---")
			fmt.Print("\n")
			generator.Listed_Generation(flags_rr6.Hashlist, flags_rr6.Output, 3)
		}
	case "encode sha512 list":
		if flags_rr6.Hashlist == "" || flags_rr6.Output == "" {
			a := sio("Enter a wordlist to hash> ", v.RET_RED)
			b := sio("Enter a output list to out data to> ", v.RET_RED)
			fmt.Println("--------------------------------------------------------")
			fmt.Println("[+] Converting all words in the list to a SHA512 hash ---")
			fmt.Print("\n")
			generator.Listed_Generation(a, b, 7)
		} else {
			fmt.Println("--------------------------------------------------------")
			fmt.Println("[+] Converting all words in the list to a SHA512 hash ---")
			fmt.Print("\n")
			generator.Listed_Generation(flags_rr6.Hashlist, flags_rr6.Output, 7)

		}
	case "encode base32 list":
		if flags_rr6.Hashlist == "" || flags_rr6.Output == "" {
			a := sio("Enter a wordlist to hash> ", v.RET_RED)
			b := sio("Enter a output list to out data to> ", v.RET_RED)
			fmt.Println("--------------------------------------------------------")
			fmt.Println("[+] Converting all words in the list to a Base32 hash ---")
			fmt.Print("\n")
			generator.Listed_Generation(a, b, 4)
		} else {
			fmt.Println("--------------------------------------------------------")
			fmt.Println("[+] Converting all words in the list to a Base32 hash ---")
			fmt.Print("\n")
			generator.Listed_Generation(flags_rr6.Hashlist, flags_rr6.Output, 4)
		}
	case "encode base64 list":
		if flags_rr6.Hashlist == "" || flags_rr6.Output == "" {
			a := sio("Enter a wordlist to hash> ", v.RET_RED)
			b := sio("Enter a output list to out data to> ", v.RET_RED)
			fmt.Println("--------------------------------------------------------")
			fmt.Println("[+] Converting all words in the list to a Base64 hash ---")
			fmt.Print("\n")
			generator.Listed_Generation(a, b, 5)
		} else {
			fmt.Println("--------------------------------------------------------")
			fmt.Println("[+] Converting all words in the list to a Base64 hash ---")
			fmt.Print("\n")
			generator.Listed_Generation(flags_rr6.Hashlist, flags_rr6.Output, 5)

		}

	//*************************************************************************************
	//*          This break concludes the hashing, crypto, encoding modules and commands  *
	//*                                                                                   *
	//*************************************************************************************
	case "dump pcap dot11":
		if flags_rr6.Filepath_general == "" {
			a := sio("Path to 802.11 Pcap file> ", v.RET_RED)
			IEEEParse.Open_parse_BSSID(a)
		} else {
			IEEEParse.Open_parse_BSSID(flags_rr6.Filepath_general)
		}
	case "dump pcap ftp":
		if flags_rr6.Filepath_general == "" {
			a := sio("Path to FTP/net Pcap file> ", v.RET_RED)
			IEEEParse.Ftp_sniff_OFFLINE_PCAP(a, true)
		} else {
			IEEEParse.Ftp_sniff_OFFLINE_PCAP(flags_rr6.Filepath_general, true)
		}
	case "dump pcap ospf":
		if flags_rr6.Filepath_general == "" {
			a := sio("Path to OSPF/net Pcap file> ", v.RET_RED)
			IEEEParse.OSPF_OFFLINE_Parsing(a)
		} else {
			IEEEParse.OSPF_OFFLINE_Parsing(flags_rr6.Filepath_general)
		}
	case "dump pcap smtppa":
		if flags_rr6.Filepath_general == "" {
			a := sio("Path to SMTP/net Pcap file> ", v.RET_RED)
			IEEEParse.Pcap_parser_OFFLINE(a, "AUTH PLAIN")
		} else {
			IEEEParse.Pcap_parser_OFFLINE(flags_rr6.Filepath_general, "AUTH PLAIN")
		}
	case "dump pcap smtpe":
		if flags_rr6.Filepath_general == "" {
			a := sio("Path to SMTP/net Pcap file> ", v.RET_RED)
			IEEEParse.Pcap_parser_OFFLINE(a, "RCPT TO")
			IEEEParse.Pcap_parser_OFFLINE(a, "MAIL FROM")
		} else {
			IEEEParse.Pcap_parser_OFFLINE(flags_rr6.Filepath_general, "RCPT TO")
			IEEEParse.Pcap_parser_OFFLINE(flags_rr6.Filepath_general, "MAIL FROM")
		}
	case "dump pcap sipa":
		if flags_rr6.Filepath_general == "" {
			a := sio("Path to SIP/net Pcap File> ", v.RET_RED)
			fmt.Println("[ >>>>> ] Username ")
			IEEEParse.Pcap_parser_OFFLINE_byte(a, "USER")
			fmt.Println("[ >>>>> ] Password ")
			IEEEParse.Pcap_parser_OFFLINE_byte(a, "PASS")
		} else {
			fmt.Println("[ >>>>> ] Username ")
			IEEEParse.Pcap_parser_OFFLINE_byte(flags_rr6.Filepath_general, "USER")
			fmt.Println("[ >>>>> ] Password ")
			IEEEParse.Pcap_parser_OFFLINE_byte(flags_rr6.Filepath_general, "PASS")
		}
	case "dump pcap sipok":
		if flags_rr6.Filepath_general == "" {
			a := sio("Path to SIP/net Pcap File> ", v.RET_RED)
			IEEEParse.Pcap_parser_OFFLINE_byte(a, "+OK")
		} else {
			IEEEParse.Pcap_parser_OFFLINE_byte(flags_rr6.Filepath_general, "+OK")
		}
	case "dump pcap sipinv":
		if flags_rr6.Filepath_general == "" {
			a := sio("Path to SIP/Net Pcap File> ", v.RET_RED)
			IEEEParse.Pcap_parser_OFFLINE_byte(a, "INVITE")
		} else {
			IEEEParse.Pcap_parser_OFFLINE_byte(flags_rr6.Filepath_general, "INVITE")
		}
	case "dump pcap sipreg":
		if flags_rr6.Filepath_general == "" {
			a := sio("Path to SIP/Net Pcap File> ", v.RET_RED)
			IEEEParse.Pcap_parser_OFFLINE_byte(a, "REGISTERS")
		} else {
			IEEEParse.Pcap_parser_OFFLINE_byte(flags_rr6.Filepath_general, "REGISTERS")
		}
	case "dump pcap sippg":
		if flags_rr6.Filepath_general == "" {
			a := sio("Path to SIP/Net Pcap File> ", v.RET_RED)
			IEEEParse.Pcap_parser_OFFLINE(a, "POST")
			IEEEParse.Pcap_parser_OFFLINE(a, "GET")
		} else {
			IEEEParse.Pcap_parser_OFFLINE(flags_rr6.Filepath_general, "POST")
			IEEEParse.Pcap_parser_OFFLINE(flags_rr6.Filepath_general, "GET")
		}
	case "dump pcap imaplogn":
		if flags_rr6.Filepath_general == "" {
			a := sio("Path to IMAP/Net Pcap File> ", v.RET_RED)
			IEEEParse.Pcap_parser_OFFLINE(a, "LOGIN")
		} else {
			IEEEParse.Pcap_parser_OFFLINE(flags_rr6.Filepath_general, "LOGIN")
		}
	case "dump pcap custom":
		fmt.Println("SMTP DEFUALT PARAMS")
		fmt.Println("\033[49m")
		fmt.Println(constants.Paramaters_SMTP)
		fmt.Println("--------------------------")
		fmt.Println("SMTP PLAIN AUTH DEFUALT PARAMS")
		fmt.Println("\033[49m")
		fmt.Println(constants.Paramaters_SMTP_plain)
		fmt.Println("--------------------------")
		fmt.Println("HTTP PARAMATERS")
		fmt.Println("\033[49m")
		fmt.Println(constants.Paramaters_HTTP)
		fmt.Println("--------------------------")
		fmt.Println("HTTP IMAGE DEFUALT PARAMS")
		fmt.Println("\033[49m")
		fmt.Println(constants.Paramaters_HTTP_img)
		fmt.Println("--------------------------")
		fmt.Println("IMAP PLAIN TEXT DEFUALT PARAMS")
		fmt.Println("\033[49m")
		fmt.Println(constants.Paramaters_IMAP_PLAIN_TEXT)
		fmt.Println("--------------------------")
		fmt.Println("\033[49m")
		fmt.Println(v.RED, "SIP Paramaters")
		fmt.Println(constants.Paramaters_SIP)
		fmt.Print("\n\n\n Note: You can add your own values, however this is just a simple help menu to guide you in the right direction :> ")
		opt := sio("Enter a value or paramater to parse      (1) allowed >>> ", v.BLKHB)
		f := sio("Enter the path to the network pcap file  (1) allowed >>> ", v.BLKHB)
		IEEEParse.Pcap_parser_OFFLINE_byte(f, opt)
	case "dump pcap":
		if flags_rr6.Filepath_general == "" {
			a := sio("Enter the filepath to the PCAP file (1) allowed >>>> ", v.BLKHB)
			IEEEParse.Parser(a)
		} else {
			IEEEParse.Parser(flags_rr6.Filepath_general)
		}
	//*************************************************************************************
	//*          This break concludes the Offline PCAP parsing and file parsers           *
	//*                                                                                   *
	//*************************************************************************************
	case "Brute SMTP":
		if flags_rr6.Brute_list == "" {
			a := sio("Enter a filepath to the wordlist you want to use> \033[39m", v.BLKHB)
			e := sio("Enter a email service to brute force> \033[39m", v.BLKHB)
			i := sio("Enter a email to brute force> \033[39m", v.BLKHB)
			brutes.Brute_SMTP(a, i, e)
		} else {
			e := sio("Enter a email service to brute force> \033[39m", v.BLKHB)
			i := sio("Enter a email to brute force> \033[39m", v.BLKHB)
			brutes.Brute_SMTP(flags_rr6.Brute_list, i, e)
		}
	case "Brute SSH":
		if flags_rr6.Brute_list == "" {
			a := sio("Enter an SSH Host> \033[39m", v.BLKHB)
			b := sio("Enter an SSH User> \033[39m", v.BLKHB)
			c := sio("Enter a Wordlist > \033[39m", v.BLKHB)
			brutes.Brute_SSH_(b, c, a)
		} else {
			a := sio("Enter an SSH Host> \033[39m", v.BLKHB)
			b := sio("Enter an SSH User> \033[39m", v.BLKHB)
			brutes.Brute_SSH_(b, flags_rr6.Brute_list, a)
		}
	case "Brute FTP":
		if flags_rr6.Brute_list == "" {
			a := sio("Enter a FTP Host> \033[39m", v.BLKHB)
			B := sio("Enter a FTP User> \033[39m", v.BLKHB)
			C := sio("Enter a Wordlist> \033[39m", v.BLKHB)
			P := sio("Enter FTP Port  > \033[39m", v.BLKHB)
			brutes.Brute_FTP(B, C, a, P)
		} else {
			a := sio("Enter a FTP Host> \033[39m", v.BLKHB)
			B := sio("Enter a FTP User> \033[39m", v.BLKHB)
			P := sio("Enter a FTP Port> \033[39m", v.BLKHB)
			brutes.Brute_FTP(B, flags_rr6.Brute_list, a, P)
		}
	case "Brute HTTPA":
		if flags_rr6.Brute_list == "" {
			a := sio("Enter a URL> \033[39m", v.BLKHB)
			b := sio("Enter a Wordlist> \033[39m", v.BLKHB)
			c := sio("Enter a Username> \033[39m", v.BLKHB)
			brutes.Brute_BASIC_HTTP_AUTH(a, b, c, 2)
		} else {
			a := sio("Enter a URL> \033[39m", v.BLKHB)
			c := sio("Enter a Username> \033[39m", v.BLKHB)
			brutes.Brute_BASIC_HTTP_AUTH(a, flags_rr6.Brute_list, c, 2)
		}
	case "Brute Telnet":
		a := sio("Enter a Hostname> \033[39m\033[49m", v.BLKHB)
		b := sio("Enter a Wordlist> \033[39m\033[49m", v.BLKHB)
		c := sio("Enter a Username> \033[39m\033[49m", v.BLKHB)
		p := "perl"
		p1 := "r6.pl"
		p00 := "-o"
		arg00 := "brute_telnet"
		p2 := "-f"
		arg := b
		p3 := "-u"
		arg1 := c
		p4 := "-t"
		arg2 := a
		exe := exec.Command(p, p1, p00, arg00, p2, arg, p3, arg1, p4, arg2)
		fmt.Println("[*] Executed....\nSilencing Output......")
		fmt.Println("[!] IF this takes too long CTRL+C is key to SIGINT ( Signal Interrupt )")
		stdout, e := exe.Output()
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println("Racing....")
		fmt.Print(string(stdout))
	case "Brute Cpan":
		a := sio("Enter a host IPA> ", v.BLKHB)
		b := sio("Enter a Username> ", v.BLKHB)
		c := sio("Enter a CPANPort> ", v.BLKHB)
		d := sio("Enter a Wordlist> ", v.BLKHB)
		p := "perl"
		p1 := constants.Perl_CPAN_Brute
		p00 := "-h"
		arg00 := a
		p2 := "-u"
		arg := b
		p3 := "-p"
		arg1 := c
		p4 := "-l"
		arg2 := d
		exe := exec.Command(p, p1, p00, arg00, p2, arg, p3, arg1, p4, arg2)
		fmt.Println(v.REDHB, "[*] Executed....\nSilencing Output......", v.RET_RED)
		fmt.Println("[!] IF this takes too long CTRL+C is key to SIGINT ( Signal Interrupt )")
		stdout, e := exe.Output()
		if e != nil {
			log.Fatal(e)
		}
		fmt.Println("Racing....")
		fmt.Print(string(stdout))
	//*************************************************************************************
	//*          This break concludes the Online brute forcing attack modules             *
	//*                                                                                   *
	//*************************************************************************************
	case "sniff interfaces":
		IEEE.HOME_Interfaces()
	case "sniff application":
		if flags_rr6.Sniffc == "" {
			Parse_options_for_netcap("app", true, "")
		} else {
			Parse_options_for_netcap("app", false, flags_rr6.Sniffc)
		}
	case "sniff tcp":
		if flags_rr6.Sniffc == "" {
			Parse_options_for_netcap("tcp", true, "")
		} else {
			Parse_options_for_netcap("tcp", false, flags_rr6.Sniffc)
		}
	case "sniff icmp":
		if flags_rr6.Sniffc == "" {
			Parse_options_for_netcap("icmp", true, "")
		} else {
			Parse_options_for_netcap("icmp", false, flags_rr6.Sniffc)
		}
	case "sniff ip":
		if flags_rr6.Sniffc == "" {
			Parse_options_for_netcap("ip", true, "")
		} else {
			Parse_options_for_netcap("ip", false, flags_rr6.Sniffc)
		}
	case "sniff dhcp":
		if flags_rr6.Sniffc == "" {
			Parse_options_for_netcap("dhcmp", true, "")
		} else {
			Parse_options_for_netcap("dhcmp", false, flags_rr6.Sniffc)
		}
	case "sniff ethernet":
		if flags_rr6.Sniffc == "" {
			Parse_options_for_netcap("eth", true, "")
		} else {
			Parse_options_for_netcap("eth", false, flags_rr6.Sniffc)
		}
	//*********************************************************************************************************
	//*          This break concludes the Online network attacks/sniffing and credential sniffers             *
	//*                                                                                                       *
	//*********************************************************************************************************
	case "check proton ip":
		ipa := sio("Enter a IP>> ", v.GRN)
		osintutils.Test_if_IP_Is_proton(ipa)
	case "check proton email":
		email := sio("Enter an email>> ", v.GRN)
		osintutils.Test_if_email_Is_proton(email)
	case "check cloudflare ip":
		ip := sio("Enter a IP>> ", v.BBLU)
		osintutils.Test_if_IP_Is_CLOUDFLARE(ip)
	case "check aws ip":
		ip := sio("Enter a IP>> ", "\033[39m")
		osintutils.Test_if_IP_Is_AWS(ip)
	case "check cloudfront ip":
		ip := sio("Enter a IP>> ", "\033[39m")
		cloudfront_utils.Compare(ip)
	case "check mcafe ip":
		ip := sio("Enter a IP>> ", "\033[39m")
		mcafe_utils.Output_test(ip)
	case "check myip":
		osintutils.Test_Public_addr()
	case "check number":
		a := sio("Enter a Number hash>> ", "\033[39m")
		phone_utils.Parser(a)
	case "check number us":
		a := sio("Enter a Number hash>> ", "\033[39m")
		phone_utils.Parser_US(false, a)
	case "trace number us":
		a := sio("Enter the first 3 numbers >> ", "\033[39m")
		b := sio("Enter the second 3 numbers >> ", "\033[39m")
		c := sio("Enter the last 4 numbers   >> ", "\033[39m")
		system.Run_number(a, b, c)
	case "check number be":
		a := sio("Enter a number hash>> ", "\033[39m")
		phone_utils.Parser_BE(a)
	case "check username":
		a := sio("Enter a username>> ", "\033[39m")
		osintutils.Test_Username_Site(a, "config/urllist_user.txt")
	case "trace ip":
		a := sio("Enter a IP>> ", "\033[39m")
		osintutils.Test_IP_Address_Location(a)
	default:
		fmt.Println("command does not exist")
	}
}
