package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	c "main/modg/colors"
	e "main/modg/warnings"
)

const (
	content_type_PUT = "Content-Type"
	application_PUT  = "application/json"
	CHARSET_PUT      = "charset=utf-8"
)

// defualt for all requests, this is a package to preform HEAD, PUT, DELETE, METHOD, POST, GET and other HTTP/HTTPS methods

//GetRequest performs a GET request
func Create_GET(target string) (string, int, error) {
	resp, err := http.Get(target)
	e.Warning_advanced("<RR6> Requests Module: Could not make a GET request to the target (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the body of the response? (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	a := string(rbody)
	return resp.Status, len(a), nil
}

//POST
func Create_POST(target string) (string, int, error) {
	pb, _ := json.Marshal("{data}")
	rb := bytes.NewBuffer(pb)
	response, err := http.Post(target, "application/json", rb)
	e.Warning_advanced("<RR6> Requests Module: Could not make a POST methodized request to the target? (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the body of the response? (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	a := string(body)
	return response.Status, len(a), nil
}

//HEAD
func Create_HEAD(target string) (string, int, error) {
	response, ec := http.Head(target)
	e.Warning_advanced("<RR6> Requests Module: Could not make a HEAD methodized request to the target (>>>) ", c.REDHB, 1, false, false, true, ec, 1, 255, "")
	defer response.Body.Close()
	b, ea := ioutil.ReadAll(response.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the body of the response? (>>>) ", c.REDHB, 1, false, false, true, ea, 1, 255, "")
	sb := string(b)
	return response.Status, len(sb), nil
}

// Create GET and read body with the HTTP client and return body
func Create_GET_Body(url string) (string, uint, error) {
	client := &http.Client{}
	request, e := http.NewRequest("GET", url, nil)
	if e != nil {
		return "<RR6> Requests Module: Could not create a client GET request to the server -> ", 0x00, e
	} else {
		response, e := client.Do(request)
		if e != nil {
			return "<RR6> Requests Module: Could not make a GET request to the server -> ", 0x00, e
		} else {
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)
			return_B := string(body)
			return return_B, 0x00, nil
		}
	}
}

//PUT
func Create_PUT(target string) (string, int, error) {
	client := &http.Client{}
	json, _ := json.Marshal("{data}")
	req, err := http.NewRequest(http.MethodPut, target, bytes.NewBuffer(json))
	e.Warning_advanced("<RR6> Requests Module: Could not make a new PUT methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	req.Header.Set(content_type_PUT, application_PUT)
	resp, err := client.Do(req)
	e.Warning_advanced("<RR6> Requests Module: Could not send the PUT methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	body, err := ioutil.ReadAll(resp.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the response body (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	b := string(body)
	return resp.Status, len(b), nil
}

//send the methodized request
func Request(target string, method string) (string, int, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, target, nil)
	e.Warning_advanced("<RR6> Requests Module: Could not make a new PUT methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	resp, err := client.Do(req)
	e.Warning_advanced("<RR6> Requests Module: Could not send the methodized request (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	e.Warning_advanced("<RR6> Requests Module: Could not read the response body (>>>) ", c.REDHB, 1, false, false, true, err, 1, 255, "")
	sb := string(body)
	return resp.Status, len(sb), nil
}

func GET_val(target, method, key string) (string, uint) {
	cl := &http.Client{}
	req, e := http.NewRequest(method, target, nil)
	if e != nil {
		fmt.Println("<RR6> Request module: Could not make a new methodized request (>>>) ", c.REDHB, e)
	}
	resp, e := cl.Do(req)
	if e != nil {
		fmt.Println("<RR6> Request module: Could not execute the new methodized request (>>>) ", c.REDHB, e)
	}
	defer resp.Body.Close()
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("<RR6> Request module: Could not read the response body (>>>) ", c.REDHB, e)
	}
	return resp.Header.Get("server"), 0x00
}

// function will remove all duplicate values in the URL that is found in the slice
func Remove_URL_vals(s []string) []string {
	l := []string{}
	k := make(map[string]bool)
	for _, e := range s {
		if _, v := k[e]; !v {
			k[e] = true
			l = append(l, e)
		}
	}
	return l
}

func ce(err error, msg string, typer string, exit_code int) bool {
	if err != nil {
		if typer == "fmt" {
			fmt.Println(err, msg, exit_code)
			os.Exit(exit_code)
		}
		if typer == "log" {
			log.Fatal(err, msg)
			os.Exit(exit_code)
		}
	} else {
		return true
	}
	return true
}

func Isexisting(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Write(Filename_of_data, data string) {
	if !Isexisting(Filename_of_data) {
		d1 := []byte(data)
		err := os.WriteFile(Filename_of_data, d1, 0644)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("\033[31m<RR6> Logging: File and data have been saved...")
		}
	}
}
