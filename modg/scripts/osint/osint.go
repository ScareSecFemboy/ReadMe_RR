package osint_utilities

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	simple_errors "main/modg/errors/serr"
	requests "main/modg/requests"
	aws "main/modg/scripts/cloud/aws/aws-constants"
	aws_tools "main/modg/scripts/cloud/aws/aws-json"
	cosint "main/modg/scripts/osint/osintc"
	osint_constants "main/modg/scripts/osint/osintc"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func is_addr_IPv4_o_IPv6(addr string) string {
	for j := 0; j < len(addr); j++ {
		switch addr[j] {
		case '.':
			return "IPv4 Detected"
		case ':':
			return "IPv6 Detected"
		}
	}
	return "[-] Address is neither IPv6 or IPv4???????\n"
}

func Test_if_IP_Is_proton(ip string) {
	ipa := strings.Trim(ip, " ")
	regexcheck, _ := regexp.Compile(cosint.Regex__Ipaddr)
	if !regexcheck.MatchString(ipa) {
		fmt.Println("[-] ERROR: <RR6> | REGEXP | Could not verify IP address must match regex string -> ", cosint.Regex__Ipaddr)
		fmt.Print("Exiting.....")
		os.Exit(0)
	} else {
		response, e := http.Get(cosint.IP_url)
		if e != nil {
			fmt.Println("[ --!-- ] Warning: Could not make a GET request to the URL got error -> \033[39m", e)
		} else {
			if response.StatusCode >= 200 {
				out, err_cr := os.Create("out.txt")
				if err_cr != nil {
					fmt.Println("[ --!-- ] Error: Could not crate the output file of the response body to look for address, got error -> ", err_cr)
					os.Exit(0)
				}
				defer out.Close()
				io.Copy(out, response.Body)
				reader, e := ioutil.ReadFile("out.txt")
				if e != nil {
					fmt.Println("[ --!-- ] Error: Could not open the file to read, this is fatal, script now can not access the IP range to search -> ", e)
					os.Exit(0)
				} else {
					if strings.Contains(string(reader), ip) {
						fmt.Printf("++++>>> RR6: IP [%s] Is apart of the proton mail VPN...\n", ip)
					} else {
						fmt.Printf("---->>> RR6: IP [%s], is NOT apart of the proton mail VPN...\n", ip)
					}
				}
			} else {
				fmt.Println("Response code was not 200......")
				os.Exit(0)
			}
		}
	}
}

func Test_if_email_Is_proton(emails string) {
	mail := strings.Trim(emails, " ")
	valid := regexp.MustCompile(cosint.Regex__Email)
	if !valid.MatchString(mail) {
		fmt.Println("[-] ERROR: <RR6> | REGEXP | Could not verify Email address must match regex string -> ", cosint.Regex__Email)
		fmt.Print("Exiting.....")
		os.Exit(0)
	} else {
		email_string := "https://api.protonmail.ch/pks/lookup?op=index&search=" + emails
		response, e := http.NewRequest("GET", email_string, nil)
		if e != nil {
			fmt.Println("[-] HTTP Error: Couuld not make a new GET request to the url, got error -> ", e)
			os.Exit(0)
		} else {
			response.Header.Set("Accept", "application/json")
			httpresp, e := cosint.Client.Do(response)
			if e != nil {
				fmt.Println("[-] Requests HTTP ERROR: Was not able to make a client request to the http URL, got error -> ", e)
				os.Exit(0)
			} else {
				defer httpresp.Body.Close()
				body, e := io.ReadAll(httpresp.Body)
				if e != nil {
					fmt.Println("[-] READER ERROR: Could not read the http response body from the HTTP client > ", e)
					os.Exit(0)
				} else {
					if strings.Contains(string(body), cosint.Verify_true_proton) {
						fmt.Println("found it!")
					} else {
						fmt.Println("no data.....")
					}
				}

			}
		}
	}

}

func Test_if_IP_Is_CLOUDFLARE(ip string) {
	if is_addr_IPv4_o_IPv6(ip) == "IPv4 Detected" {
		for _, c := range osint_constants.CloudFlareCIDR_IP4 {
			_, sa, _ := net.ParseCIDR(c)
			ipa, _, _ := net.ParseCIDR(ip + "/0")
			if sa.Contains(ipa) {
				fmt.Printf("\033[31m<RR6> Address Range | %s | Does match to list  | %s | > | %s | Is a possible Cloudflare IP range \n", c, ip, c)
			} else {
				fmt.Printf("\033[32m<RR6> Address Range | %s | Does not match with | %s | > | %s | Is not a possible cloudflare IP range\n", c, ip, c)
			}
		}
	} else {
		for _, a := range osint_constants.CloudFlareCIDR_IP6 {
			_, sc, _ := net.ParseCIDR(a)
			ipa, _, _ := net.ParseCIDR(ip + "/0")
			if sc.Contains(ipa) {
				fmt.Printf("\033[31m<RR6> Address Range | %s | Does match to list  | %s | > | %s | Is a possible Cloudflare IP range \n", a, ip, a)
			} else {
				fmt.Printf("\033[32m<RR6> Address Range | %s | Does not match with | %s | > | %s | Is not a possible cloudflare IP range\n", a, ip, a)
			}
		}
	}
}

func Test_if_IP_Is_AWS(ip string) {
	if is_addr_IPv4_o_IPv6(ip) == "IPv4 Detected" || is_addr_IPv4_o_IPv6(ip) == "IPv6 Detected" {
		response_aws, e := http.Get(aws.IP_URL)
		simple_errors.See_errorbased(e, "\033[39m", " Was not able to make a GET request to the given URL for AWS matching...", false)
		if response_aws.StatusCode == 200 {
			defer response_aws.Body.Close()
			prefix, e := ioutil.ReadAll(response_aws.Body)
			simple_errors.See_errorbased(e, "\033[39m", "Could not read the entire response body from the url", false)
			u, err := url.Parse(aws.IP_URL)
			simple_errors.See_errorbased(err, "\033[39m", "Could not properly parse the AWS url, error -> ", false)
			domain := strings.Split(u.Hostname(), ".")
			d := domain[len(domain)-2] + "." + domain[len(domain)-1]
			filename := d + ".json"
			requests.Write(filename, string(prefix))
		} else {
			fmt.Println("<RR6> Requests module: Could not make a good decent 200 status code request to the url")
		}
	}
	aws_tools.Output_test(ip)
}

func Test_Public_addr() {
	uli := "https://api.ipify.org?format=text"
	response, e := http.Get(uli)
	simple_errors.See_errorbased(e, "\033[31m", "Could not make a get request to the API url", false)
	defer response.Body.Close()
	ip, e := ioutil.ReadAll(response.Body)
	simple_errors.See_errorbased(e, "\033[31m", "Could not read the entire response body from the url", false)
	fmt.Printf("[+] Address ~>  | %s | \n", ip)
}

func Test_Username_Site(username string, wordlist string) string {
	contents, err := os.Open(wordlist)
	if err != nil {
		log.Fatal("<RR6> OSINT Module: Could not open the wordlist for the website gathering.. -> ", err)
	} else {
		scanner := bufio.NewScanner(contents)
		for scanner.Scan() {
			parser := scanner.Text() + username
			response, err := http.Get(parser)
			if err != nil {
				log.Fatal("<RR6> OSINT Module: Could not make a proper request to the URL and host got error, -> ", err)
			} else {
				if response.StatusCode == 200 {
					fmt.Printf("[+] Found username < %s > IS  on website < %s > | Stat code < 200 > \n", username, scanner.Text())
				} else {
					fmt.Printf("[-] Error username < %s > NOT on website < %s > | Stat code < %v > \n", username, scanner.Text(), response.StatusCode)
				}
			}
		}
	}
	return "Error?"
}

func Test_IP_Address_Location(ip string) {
	parser := fmt.Sprintf("https://ipapi.co/%s/json/", ip)
	return_value, _, e := requests.Create_GET_Body(parser)
	if e != nil {
		fmt.Println("<RR6> OSINT Module talking to <- Requests module Caused -> error ->  ", e)
	} else {
		fmt.Println(return_value)
	}
}
