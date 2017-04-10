package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/luyaotsung/jiraGetIssues/lib"
)

func main() {
	jiraweb := flag.String("jiraweb", "", "Jira Web Site Address, example : -jiraweb=https://inhouse.htcstudio.com/jira")
	username := flag.String("username", "", "User Name")
	password := flag.String("password", "", "Password")

	flag.Parse()

	if flag.NFlag() == 3 {

		// func getReturnJSON( webaddress string, username string, password string, startcount int, totalcount int) jira.JiraObject {
		jiraObject := jira.GetReturnJSON(*jiraweb, *username, *password, 0, 1)

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
			jiraObject = jira.GetReturnJSON(*jiraweb, *username, *password, startCount, totalCount)

			for x := range jiraObject.Issues {
				Key := jiraObject.Issues[x].Key

				fmt.Println("Key:", Key, "Issue TYPE:", jiraObject.Issues[x].Fields.IssueType.Name)

				componentMerged := jira.GetComponents(jiraObject.Issues[x].Fields.Components)
				labelsMerged := jira.GetLabels(jiraObject.Issues[x].Fields.Labels)

				fmt.Println("Components -->", componentMerged, "Labels -->", labelsMerged)
			}
		}
	} else {
		CommandName := strings.Split(os.Args[0], "/")
		fmt.Printf("Please chekc help file by \" %s -h \" \n", CommandName[len(CommandName)-1])
	}
}
