package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/luyaotsung/jiraGetIssues/lib/jira"
	"github.com/luyaotsung/jiraGetIssues/lib/jirasql"
)

func main() {
	jiraweb := flag.String("jiraweb", "",
		"Jira Web Site Address, example : -jiraweb=https://inhouse.htcstudio.com/jira")
	username := flag.String("username", "", "User Name")
	password := flag.String("password", "", "Password")
	queryproject := flag.String("projects", "",
		"Input your projects to query, example : +Project1+,+Project2,+Project3")
	sqlserverinfo := flag.String("sqlserverinfo", "",
		"Input your sql server information [UserName]:[Password]@tcp([SERVER IP:PORT])/[DatabaseName] , example:  eli:eli@tcp(127.0.0.1:3306)/JiraData ")
	sqltablename := flag.String("sqltablename", "",
		"Input your sql name")

	flag.Parse()

	if flag.NFlag() == 6 {

		jiraObject := jira.GetReturnJSON(*jiraweb, *username, *password, *queryproject, 0, 1)

		fmt.Println("Total : ", jiraObject.Total)
		fmt.Println("MaxResult : ", jiraObject.MaxResults)

		// Start to do our job to query all issues.
		rangeCount := (jiraObject.Total / 1000) + 1
		fmt.Println("Range Count => ", rangeCount)

		for i := 0; i < rangeCount; i++ {
			startCount := 1000 * i
			totalCount := 1000
			fmt.Println("Strat Count --> ", startCount)

			jiraObject = jira.GetReturnJSON(*jiraweb, *username, *password, *queryproject, startCount, totalCount)

			for x := range jiraObject.Issues {

				prepareData := jirasql.InjectData{
					Issuekey:      jiraObject.Issues[x].Key,
					Issuetype:     jiraObject.Issues[x].Fields.IssueType.Name,
					Issueid:       jiraObject.Issues[x].ID,
					Issueself:     jiraObject.Issues[x].Self,
					Project:       jiraObject.Issues[x].Fields.Project.Name,
					Summary:       jiraObject.Issues[x].Fields.Summary,
					Priority:      jiraObject.Issues[x].Fields.Priority.Name,
					Resolution:    jiraObject.Issues[x].Fields.Resolution.Name,
					Status:        jiraObject.Issues[x].Fields.Status.Name,
					Lastchange:    jiraObject.Issues[x].Fields.Updated,
					Reporter:      jiraObject.Issues[x].Fields.Reporter.Name,
					Assignee:      jiraObject.Issues[x].Fields.Assignee.Name,
					Label:         jira.GetLabels(jiraObject.Issues[x].Fields.Labels),
					Fixversion:    jira.GetFixVersions(jiraObject.Issues[x].Fields.FixVersion),
					Component:     jira.GetComponents(jiraObject.Issues[x].Fields.Components),
					Affectversion: jira.GetVersions(jiraObject.Issues[x].Fields.Versions),
					Startdate:     jiraObject.Issues[x].Fields.StartDate,
					Duedate:       jiraObject.Issues[x].Fields.DueDate,
				}

				jirasql.UpdateJiraDB(prepareData, *sqlserverinfo, *sqltablename)

				issuekey := jiraObject.Issues[x].Key
				issuetype := jiraObject.Issues[x].Fields.IssueType.Name
				issueid := jiraObject.Issues[x].ID
				issueself := jiraObject.Issues[x].Self
				project := jiraObject.Issues[x].Fields.Project.Name
				summary := jiraObject.Issues[x].Fields.Summary
				priority := jiraObject.Issues[x].Fields.Priority.Name
				resolution := jiraObject.Issues[x].Fields.Resolution.Name
				status := jiraObject.Issues[x].Fields.Status.Name
				lastchange := jiraObject.Issues[x].Fields.Updated
				reporter := jiraObject.Issues[x].Fields.Reporter.Name
				assignee := jiraObject.Issues[x].Fields.Assignee.Name
				label := jira.GetLabels(jiraObject.Issues[x].Fields.Labels)
				fixversion := jira.GetFixVersions(jiraObject.Issues[x].Fields.FixVersion)
				component := jira.GetComponents(jiraObject.Issues[x].Fields.Components)
				affectversion := jira.GetVersions(jiraObject.Issues[x].Fields.Versions)
				startdate := jiraObject.Issues[x].Fields.StartDate
				duedate := jiraObject.Issues[x].Fields.DueDate

				fmt.Println("=============================================================")
				fmt.Println("Issue Key : ", issuekey)
				fmt.Println("Issue Type : ", issuetype)
				fmt.Println("Issue ID : ", issueid)
				fmt.Println("Issue Self : ", issueself)
				fmt.Println("Project : ", project)
				fmt.Println("Summary : ", summary)
				fmt.Println("Priority : ", priority)
				fmt.Println("Resolution : ", resolution)
				fmt.Println("Status : ", status)
				fmt.Println("LastChange : ", lastchange)
				fmt.Println("Reporter : ", reporter)
				fmt.Println("Assignee : ", assignee)
				fmt.Println("Label : ", label)
				fmt.Println("Fix Version : ", fixversion)
				fmt.Println("Component : ", component)
				fmt.Println("Affect Version : ", affectversion)
				fmt.Println("Start Date : ", startdate)
				fmt.Println("DueDate : ", duedate)

			}
		}
	} else {
		CommandName := strings.Split(os.Args[0], "/")
		fmt.Printf("Please chekc help file by \" %s -h \" \n", CommandName[len(CommandName)-1])
	}
}
