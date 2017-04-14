package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/luyaotsung/jiraGetIssues/lib"
)

const (
	dbServer    = "127.0.0.1:3306"
	dbUserName  = "eli"
	dbPassword  = "eli"
	dbDBName    = "JiraData"
	dbTableName = "Tickets"
)

// InjectData is the struct for Inject DB Data
type InjectData struct {
	issuekey      string
	issuetype     string
	issueid       string
	issueself     string
	project       string
	summary       string
	priority      string
	resolution    string
	status        string
	lastchange    string
	reporter      string
	assignee      string
	label         string
	fixversion    string
	component     string
	affectversion string
	startdate     string
	duedate       string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// ConfirmMethod is the feature that will confirm sql method is UPDATE or INSERT then return the method.
func ConfirmMethod(TableName string, IssueKey string, LastChange string, mySQLInfo string) (Method string, SQLExtralCMD string) {
	db, err := sql.Open("mysql", mySQLInfo)
	checkErr(err)

	var issuekey string
	var lastchange string
	row := db.QueryRow("SELECT issuekey, lastchange FROM "+TableName+"  WHERE issuekey = ?", IssueKey)
	checkErr(err)
	err = row.Scan(&issuekey, &lastchange)
	db.Close()

	if err == nil {
		fmt.Println(" ------ UPDATE ------ ")

		layout := " 2017-04-14 02:50:33"
		originalTime, _ := time.Parse(layout, lastchange)
		latestTime, _ := time.Parse(layout, LastChange)

		if originalTime.Unix() == latestTime.Unix() {
			fmt.Println("    Same Data    ")
			return "NONE", ""
		}
		return "UPDATE", " WHERE issuekey=?"
	}
	fmt.Println(" ------ INSERT ------ ")
	return "INSERT", ""
}

// UpdateJiraDB is a feature that can add/modify content of one jira ticket.
func UpdateJiraDB(data InjectData) {
	mySQLInfo := dbUserName + ":" + dbPassword + "@tcp(" + dbServer + ")/" + dbDBName
	fmt.Println("MySQL Server Information :", mySQLInfo)

	sqlMethod, sqlExtralCMD := ConfirmMethod(dbTableName, data.issuekey, data.lastchange, mySQLInfo)

	db, err := sql.Open("mysql", mySQLInfo)
	checkErr(err)

	switch method := sqlMethod; method {
	case "UPDATE":
		stmt, err := db.Prepare(sqlMethod + " " + dbTableName + " SET " +
			"issuekey=?," +
			"issuetype=?," +
			"issueid=?," +
			"issueself=?," +
			"project=?," +
			"summary=?," +
			"priority=?," +
			"resolution=?," +
			"status=?," +
			"lastchange=?," +
			"reporter=?," +
			"assignee=?," +
			"labels=?," +
			"fixversions=?," +
			"component=?," +
			"affectversions=?" +
			sqlExtralCMD)
		checkErr(err)
		_, err = stmt.Exec(data.issuekey,
			data.issuetype,
			data.issueid,
			data.issueself,
			data.project,
			data.summary,
			data.priority,
			data.resolution,
			data.status,
			data.lastchange,
			data.reporter,
			data.assignee,
			data.label,
			data.fixversion,
			data.component,
			data.affectversion,
			data.issuekey)
		checkErr(err)

	case "INSERT":
		stmt, err := db.Prepare(sqlMethod + " " + dbTableName + " SET " +
			"issuekey=?," +
			"issuetype=?," +
			"issueid=?," +
			"issueself=?," +
			"project=?," +
			"summary=?," +
			"priority=?," +
			"resolution=?," +
			"status=?," +
			"lastchange=?," +
			"reporter=?," +
			"assignee=?," +
			"labels=?," +
			"fixversions=?," +
			"component=?," +
			"affectversions=?" +
			sqlExtralCMD)
		checkErr(err)
		res, err := stmt.Exec(data.issuekey,
			data.issuetype,
			data.issueid,
			data.issueself,
			data.project,
			data.summary,
			data.priority,
			data.resolution,
			data.status,
			data.lastchange,
			data.reporter,
			data.assignee,
			data.label,
			data.fixversion,
			data.component,
			data.affectversion)
		checkErr(err)
		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println("Last Insert ID : ", id)

	default:
		fmt.Println("Method ===> ", method)
	}
	db.Close()
}

func main() {
	jiraweb := flag.String("jiraweb", "",
		"Jira Web Site Address, example : -jiraweb=https://inhouse.htcstudio.com/jira")
	username := flag.String("username", "", "User Name")
	password := flag.String("password", "", "Password")
	queryproject := flag.String("projects", "",
		"Input your projects to query, example : +Project1+,+Project2,+Project3")

	flag.Parse()

	if flag.NFlag() == 4 {

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

				prepareData := InjectData{
					issuekey:      jiraObject.Issues[x].Key,
					issuetype:     jiraObject.Issues[x].Fields.IssueType.Name,
					issueid:       jiraObject.Issues[x].ID,
					issueself:     jiraObject.Issues[x].Self,
					project:       jiraObject.Issues[x].Fields.Project.Name,
					summary:       jiraObject.Issues[x].Fields.Summary,
					priority:      jiraObject.Issues[x].Fields.Priority.Name,
					resolution:    jiraObject.Issues[x].Fields.Resolution.Name,
					status:        jiraObject.Issues[x].Fields.Status.Name,
					lastchange:    jiraObject.Issues[x].Fields.Updated,
					reporter:      jiraObject.Issues[x].Fields.Reporter.Name,
					assignee:      jiraObject.Issues[x].Fields.Assignee.Name,
					label:         jira.GetLabels(jiraObject.Issues[x].Fields.Labels),
					fixversion:    jira.GetFixVersions(jiraObject.Issues[x].Fields.FixVersion),
					component:     jira.GetComponents(jiraObject.Issues[x].Fields.Components),
					affectversion: jira.GetVersions(jiraObject.Issues[x].Fields.Versions),
					startdate:     jiraObject.Issues[x].Fields.StartDate,
					duedate:       jiraObject.Issues[x].Fields.DueDate,
				}

				UpdateJiraDB(prepareData)

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
