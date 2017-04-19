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

// Below is the table information for this applicaton
// CREATE TABLE `Tickets` (
//   `sn` int(10) NOT NULL AUTO_INCREMENT,
//   `issuekey` char(20) COLLATE utf8_unicode_ci NOT NULL,
//   `issuetype` char(200) COLLATE utf8_unicode_ci NOT NULL,
//   `issueid` int(20) NOT NULL,
//   `issueself` text COLLATE utf8_unicode_ci NOT NULL,
//   `project` char(250) COLLATE utf8_unicode_ci NOT NULL,
//   `summary` text COLLATE utf8_unicode_ci NOT NULL,
//   `priority` char(200) COLLATE utf8_unicode_ci NOT NULL,
//   `resolution` char(100) COLLATE utf8_unicode_ci NOT NULL,
//   `status` char(100) COLLATE utf8_unicode_ci NOT NULL,
//   `lastchange` datetime NOT NULL,
//   `reporter` char(200) COLLATE utf8_unicode_ci NOT NULL,
//   `assignee` char(200) COLLATE utf8_unicode_ci NOT NULL,
//   `labels` text COLLATE utf8_unicode_ci NOT NULL,
//   `fixversions` text COLLATE utf8_unicode_ci NOT NULL,
//   `component` varchar(1000) COLLATE utf8_unicode_ci NOT NULL,
//   `affectversions` varchar(1000) COLLATE utf8_unicode_ci NOT NULL,
//   PRIMARY KEY (`issuekey`),
//   UNIQUE KEY `sn` (`sn`),
//   UNIQUE KEY `issuekey` (`issuekey`),
//   KEY `lastchange` (`lastchange`),
//   KEY `status` (`status`),
//   KEY `resolution` (`resolution`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=COMPACT;

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
			}
		}
	} else {
		CommandName := strings.Split(os.Args[0], "/")
		fmt.Printf("Please chekc help file by \" %s -h \" \n", CommandName[len(CommandName)-1])
	}
}
