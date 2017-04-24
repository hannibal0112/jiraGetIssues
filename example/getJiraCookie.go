package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type CookieReturn struct {
	Self      string          `json:"self"`
	Name      string          `json:"name"`
	LoginInof CookieLoginInfo `json:"loginInfo"`
}

type CookieLoginInfo struct {
	FailedLoginCount  int    `json:"failLoginCount"`
	LoginCount        int    `json:"loginCount"`
	LastFailLoginTime string `json:"lastFailedLoginTime"`
	PreviousLoginTime string `json:"previousLoginTime"`
}

type ObjectErrorMessage struct {
	ErrorMessages []string    `json:"errorMessages"`
	Errors        interface{} `json:"errors"`
}

func main() {
	username := flag.String("username", "", "User Name")
	password := flag.String("password", "", "Password")
	jiraweb := flag.String("jiraweb", "",
		"Jira Web Site Address, example : -jiraweb=https://inhouse.htcstudio.com/jira")

	flag.Parse()

	if flag.NFlag() == 3 {

		client := &http.Client{}

		getProjectIssueList := fmt.Sprintf("%s/rest/auth/1/session", *jiraweb)
		req, err := http.NewRequest("GET", getProjectIssueList, nil)
		req.SetBasicAuth(*username, *password)
		req.Header.Add("Content-Type", "application/json")
		resp, err := client.Do(req)

		if err != nil {
			log.Fatal(err)
		}

		// JSESSIONID=3181977F1AC7392A1C43EE87FC8BCB36; Path=/jira/; Secure; HttpOnly
		jiraCookie := strings.Split(resp.Header.Get("Set-Cookie"), ";")[0]

		fmt.Println("Jira Cookie ==> ", jiraCookie)

		// Clone http response
		var bodyBytes []byte
		if resp.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(resp.Body)
		}
		respError := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		var jiraObject CookieReturn
		errorMessage := json.NewDecoder(resp.Body).Decode(&jiraObject)
		if errorMessage != nil {
			log.Fatal(errorMessage)
		}

		if jiraObject.LoginInof.LoginCount == 0 {
			var errorMessageObject ObjectErrorMessage
			json.NewDecoder(respError).Decode(&errorMessageObject)
			for x := range errorMessageObject.ErrorMessages {
				fmt.Println(errorMessageObject.ErrorMessages[x])
			}
			log.Fatal("You are not login ever")
		}

	} else {
		CommandName := strings.Split(os.Args[0], "/")
		fmt.Printf("Please chekc help file by \" %s -h \" \n", CommandName[len(CommandName)-1])
	}
}
