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

	"github.com/luyaotsung/jiraGetIssues/lib"
)

// this feature can get labes merge strings
func getLabels(d []string) string {
	name := ""
	for z := range d {
		name = name + "," + d[z]
	}
	if name == "" {
		return ""
	}
	return name[1:len(name)]
}

// this feature can get fix versions merge strings
func getFixVersions(d []jira.JiraFieldsFixVersions) string {
	name := ""
	for z := range d {
		name = name + "," + d[z].Name
	}
	if name == "" {
		return ""
	}
	return name[1:len(name)]
}

// this feature can get versions merge strings
func getVersions(d []jira.JiraFieldsVersions) string {
	name := ""
	for z := range d {
		name = name + "," + d[z].Name
	}
	if name == "" {
		return ""
	}
	return name[1:len(name)]
}

// this feature can get component merge strings
func getComponents(d []jira.JiraFieldsResolution) string {
	name := ""
	for z := range d {
		name = name + "," + d[z].Name
	}
	if name == "" {
		return ""
	}
	return name[1:len(name)]
}

func getReturnJSON(webaddress string, username string, password string, startcount int, totalcount int) jira.JiraObject {

	client := &http.Client{}

	// Number is from 0 maximum is 1000 at one times.
	getProjectIssueList := fmt.Sprintf("%s/rest/api/2/search?jql=project+in+(+GMM+,+TES+)+order+by+id&startAt=%d&maxResults=%d", webaddress, startcount, totalcount)
	req, err := http.NewRequest("GET", getProjectIssueList, nil)
	req.SetBasicAuth(username, password)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	// Clone http response
	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
	}
	respError := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	xSeraphLoginReason := resp.Header.Get("X-Seraph-LoginReason")

	fmt.Println(" Login Feedback is ==>", xSeraphLoginReason)

	var jiraObject jira.JiraObject
	errorMessage := json.NewDecoder(resp.Body).Decode(&jiraObject)
	if errorMessage != nil {
		log.Fatal(errorMessage)
	}

	// If value of Total it seems Project name is error.
	if jiraObject.Total == 0 {
		var errorMessageObject jira.ObjectErrorMessage
		json.NewDecoder(respError).Decode(&errorMessageObject)
		for x := range errorMessageObject.ErrorMessages {
			fmt.Println(errorMessageObject.ErrorMessages[x])
		}
		log.Fatal("Issue of Project is empty")
	}
	return jiraObject
}

func main() {

	jiraweb := flag.String("jiraweb", "", "Jira Web Site Address, example : -jiraweb=https://inhouse.htcstudio.com/jira")
	username := flag.String("username", "", "User Name")
	password := flag.String("password", "", "Password")

	flag.Parse()

	if flag.NFlag() == 3 {

		// func getReturnJSON( webaddress string, username string, password string, startcount int, totalcount int) jira.JiraObject {
		jiraObject := getReturnJSON(*jiraweb, *username, *password, 0, 1)

		fmt.Println("Total : ", jiraObject.Total)
		fmt.Println("MaxResult : ", jiraObject.MaxResults)

		// Start to do our job to query all issues.
		rangeCount := (jiraObject.Total / 1000) + 1
		fmt.Println("Range Count => ", rangeCount)

		for i := 0; i < rangeCount; i++ {

			startCount := 1000 * i
			totalCount := 1

			fmt.Println("Strat Count --> ", startCount)

			// func getReturnJSON( webaddress string, username string, password string, startcount int, totalcount int) jira.JiraObject {
			jiraObject = getReturnJSON(*jiraweb, *username, *password, startCount, totalCount)

			for x := range jiraObject.Issues {
				Key := jiraObject.Issues[x].Key

				fmt.Println("Key:", Key, "Issue TYPE:", jiraObject.Issues[x].Fields.IssueType.Name)

				componentMerged := getComponents(jiraObject.Issues[x].Fields.Components)
				labelsMerged := getLabels(jiraObject.Issues[x].Fields.Labels)

				fmt.Println("Components -->", componentMerged, "Labels -->", labelsMerged)
			}
		}
	} else {
		CommandName := strings.Split(os.Args[0], "/")
		fmt.Printf("Please chekc help file by \" %s -h \" \n", CommandName[len(CommandName)-1])
	}
}
