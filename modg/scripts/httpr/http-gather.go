/*(
	Dev -> ArkAngeL43
	Pkg -> gather
	fp  -> RR6/modg/scripts/httpr

TODO:
	This package is made to grab data off of a HTML page or file such as finding
	emails, titles, code notes, links in the HTML,
)
*/
package gather

import (
	"fmt"
	"io"
	"io/ioutil"
	c "main/modg/colors"
	ex "main/modg/warnings"
	"net/http"
	"os"
	"strings"
	"time"

	httpvars "main/modg/http-const"
	httpconst "main/modg/http-const/constants"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// fetch all links in the HREF section
func Fetch_Links(url string) (string, error) {
	request_to := url
	response, e := http.Get(request_to)
	ex.Warning_simple("<RR6> Requests Module: Could not make a new HTTP request to the target URL =>> ", c.REDHB, e)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		fmt.Println("<RR6> Requests Module: Status got code =>> < 200 > ")
	} else {
		fmt.Println("<RR6> Requests Module: Got a unwanted status code =>> !! ", response.StatusCode, " !! This might be a fatal STAT code")
	}
	document, e := goquery.NewDocumentFromReader(response.Body)
	ex.Warning_simple("<RR6> HTTP Document: Could not create a new reader with the response body ", c.REDHB, e)
	function := func(integer int, section *goquery.Selection) bool {
		links, _ := section.Attr(httpconst.Href) // looks for any HREF links in the response body
		return strings.HasPrefix(links, "https")
	}
	document.Find("body a").FilterFunction(function).Each(func(_ int, tagged *goquery.Selection) {
		link, _ := tagged.Attr("href")
		linkt := tagged.Text()
		fmt.Printf("[+] Found link < %s >  | < %s > \n", linkt, link)
	})
	return "", nil
}

// fetch all titles
func Fetch_Title(url string) (string, error) {
	response, e := http.Get(url)
	ex.Warning_simple("<RR6> Requests Module: Could not make a new HTTP request to the target URL =>> ", c.REDHB, e)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		fmt.Println("<RR6> Requests Module: Status got code =>> < 200 > ")
	} else {
		fmt.Println("<RR6> Requests Module: Got a unwanted status code =>> !! ", response.StatusCode, " !! This might be a fatal STAT code")
	}
	document, e := goquery.NewDocumentFromReader(response.Body)
	ex.Warning_simple("<RR6> HTTP Document: Could not create a new reader with the response body ", c.REDHB, e)
	searchfor := document.Find("title").Text()
	fmt.Printf("[+] Url < %s > Has a title of {<- %s ->} ", url, searchfor)
	return "", nil
}

// fetch all list tags
func Fetch_LI(url string) ([]string, error) {
	response, e := http.Get(url)
	ex.Warning_simple("<RR6> Requests Module: Could not make a new HTTP request to the target URL =>> ", c.REDHB, e)
	if response.StatusCode == 200 {
		fmt.Println("<RR6> Requests Module: Status got code =>> < 200 > ")
	} else {
		fmt.Println("<RR6> Requests Module: Got a unwanted status code =>> !! ", response.StatusCode, " !! This might be a fatal STAT code")
	}
	defer response.Body.Close()
	f, e := os.Create("response.html")
	if e != nil {
		// this needs to be customized
		fmt.Println("<RR6> Requests Module: Could not make a new filename, got error =>> ", e)
	}
	defer f.Close()
	_, e = io.Copy(f, response.Body)
	if e != nil {
		fmt.Println("<RR6> OS/System Module: Could not copy the data into the new file; got error =>> ", e)
	}
	fmt.Println("[*] Sleeping to read the file......5(s)")
	time.Sleep(5 * time.Second)
	reader, e := ioutil.ReadFile("response.html")
	ex.Warning_simple("<RR6> Requests Module: Could not read the response HTML file ==>> (sometimes this is not a fatal error, and not a glitch, user has been warned) ", c.REDHB, e)
	txt := string(reader)
	documentate, e := html.Parse(strings.NewReader(txt))
	ex.Warning_simple("<RR6> Requests Module: Could not make a new reader for the html parser ==>>  ", c.REDHB, e)
	var output []string
	Watch_dog_1(documentate, &output, httpconst.L)
	fmt.Println("--------------------------------------------------")
	fmt.Print("[*] Found new data - \n")
	return output, nil
}

func Watch_dog_1(document *html.Node, str *[]string, html_tag string) (string, error) {
	httpvars.Finder_func = func(n *html.Node, tagged string) *html.Node {
		for J := n.FirstChild; J != nil; J = J.NextSibling {
			if J.Type == html.TextNode && J.Parent.Data == tagged {
				*str = append(*str, J.Data)
			}
			watchdog := httpvars.Finder_func(J, tagged)
			if watchdog != nil {
				return watchdog
			}
		}
		return nil
	}
	httpvars.Finder_func(document, html_tag)
	return "", nil
}

// fetch all documents on the page
func Fetch_doc(url string) (string, error) {
	response, e := http.Get(url)
	ex.Warning_simple("<RR6> Requests Module: Could not make a proper GET request to the target URL ==>>  ", c.REDHB, e)
	document, e := goquery.NewDocumentFromReader(response.Body)
	ex.Warning_simple("<RR6> Requests Module: Could not make a new reader for the html parser ==>>  ", c.REDHB, e)
	document.Find(httpconst.Atag).Each(func(j int, sections *goquery.Selection) {
		find_by_tag, ex := sections.Attr(httpconst.Href)
		if ex && Watch_dog_2(find_by_tag) {
			fmt.Println("\033[49m<RR6> Found document          | ", find_by_tag, "|\033[39m\033[49m")
		}
	})
	return "", nil
}

func Watch_dog_2(parsing_what string) bool {
	p := strings.Split(parsing_what, ".")
	if len(p) < 2 {
		return true
	} else {
		for _, ext := range httpvars.DOCEXE {
			if p[len(p)-1] == ext {
				return true
			}
		}
	}
	return false
}
