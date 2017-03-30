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

type JiraObject struct {
	Expend     string       `json:"expend"`
	StartAt    int          `json:"startAt"`
	MaxResults int          `json:"maxResults"`
	Total      int          `json:"total"`
	Issues     []JiraIssues `json:"issues"`
}

type JiraIssues struct {
	Expend string     `json:"expend"`
	ID     string     `json:"id"`
	Self   string     `json:"self"`
	Key    string     `json:"key"`
	Fields JiraFields `json:"fields"`
}

type JiraIssuesChangeLog struct {
	Expend    string                  `json:"expend"`
	ID        string                  `json:"id"`
	Self      string                  `json:"self"`
	Key       string                  `json:"key"`
	Fields    JiraFields              `json:"fields"`
	ChangeLog JiraIssuesChangeLogData `json:"changelog"`
}

type JiraFields struct {
	Summary                       string                  `json:"summary"`
	Progress                      interface{}             `json:"progress"`
	IssueType                     JiraFieldsIssueType     `json:"issuetype"`
	Votes                         interface{}             `json:"votes"`
	Resolution                    JiraFieldsResolution    `json:"resolution"`
	FixVersion                    []JiraFieldsFixVersions `json:"fixVersions"`
	ResoluationDate               string                  `json:"resolutiondate"`
	TimeSpent                     int                     `json:"timespent"`
	Reporter                      JiraFieldsReporter      `json:"reporter"`
	AggregateTimeOriginalEstimate int                     `json:"aggregatetimeoriginalestimate"`
	Updated                       string                  `json:"updated"`
	Created                       string                  `json:"created"`
	Description                   string                  `json:"description"`
	Priority                      JiraFieldsPriority      `json:"priority"`
	DueDate                       string                  `json:"duedate"`
	Status                        JiraFieldsStatus        `json:"status"`
	Labels                        []string                `json:"labels"`
	Assignee                      JiraFieldsReporter      `json:"assignee"`
	Project                       JiraFieldsProject       `json:"project"`
	Version                       []JiraFieldsVersions    `json:"versions"`
	Components                    []JiraFieldsResolution  `json:"components"`
}

type JiraFieldsVersions struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Archived    bool   `json:"archived"`
	Released    bool   `json:"released"`
	ReleaseDate string `json:"releaseDate"`
}

type JiraFieldsProject struct {
	Self         string      `json:"self"`
	ID           string      `json:"id"`
	Key          string      `json:"key"`
	Name         string      `json:"name"`
	EmailAddress string      `json:"emailAddress"`
	AvatarUrls   interface{} `json:"avatarUrls"`
}

type JiraFieldsStatus struct {
	Self        string `json:"self"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	ID          string `json:"id"`
}

type JiraFieldsPriority struct {
	Self    string `json:"self"`
	IconURL string `json:"iconUrl"`
	Name    string `json:"name"`
	ID      string `json:"id"`
}

type JiraFieldsReporter struct {
	Self         string      `json:"self"`
	Name         string      `json:"name"`
	EmailAddress string      `json:"emailAddress"`
	AvatarUrls   interface{} `json:"avatarUrls"`
	DisplayName  string      `json:"displayName"`
	Active       bool        `json:"active"`
}

type JiraFieldsResolution struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type JiraFieldsIssueType struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	iconUrl     string `json:"iconUrl"`
	Name        string `json:"name"`
	SubTask     bool   `json:"subtask"`
}

type JiraFieldsFixVersions struct {
	Self           string `json:"self"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	Archievd       bool   `json:"archievd"`
	Released       bool   `json:"released"`
	ResolutionDate string `json:"resolutiondate"`
	TimeSpent      string `json:"timespent"`
}

type JiraIssuesChangeLogData struct {
	StartAt    int                                `json:"startAt"`
	MaxResults int                                `json:"maxResult"`
	Total      int                                `json:"total"`
	Histories  []JiraIssuesChangeLogDataHistories `json:"histories"`
}

type JiraIssuesChangeLogDataHistories struct {
	ID      string                                  `json:"id"`
	Author  JiraIssuesChangeLogDataHistoriesAuthor  `json:"author"`
	Created string                                  `json:"created"`
	Items   []JiraIssuesChangeLogDataHistoriesItems `json:"items"`
}

type JiraIssuesChangeLogDataHistoriesAuthor struct {
	Self         string      `json:"self"`
	Name         string      `json:"name"`
	EmailAddress string      `json:"emailAddress"`
	AvatarUrls   interface{} `json:"avatarUrls"`
	DisplayName  string      `json:"displayName"`
	Active       bool        `json:"active"`
}

type JiraIssuesChangeLogDataHistoriesItems struct {
	Field      string `json:"field"`
	FieldType  string `json:"fieldtype"`
	From       string `json:"from"`
	FromString string `json:"fromString"`
	To         string `json:"to"`
	ToString   string `json:"toString"`
}

type objectErrorMessage struct {
	ErrorMessages []string    `json:"errorMessages"`
	Errors        interface{} `json:"errors"`
}

func main() {

	jiraweb := flag.String("jiraweb", "", "Jira Web Site Address, example : -jiraweb=http://jira.sw.studio.htc.com")
	projectname := flag.String("projectname", "", "Project Name, example : -projectname=TYGH")
	username := flag.String("username", "", "User Name")
	password := flag.String("password", "", "Password")

	flag.Parse()

	if flag.NFlag() == 4 {

		client := &http.Client{}
		getProjectIssueList := fmt.Sprintf("%s/rest/api/2/search?jql=project=%s&maxResults=-1", *jiraweb, *projectname)
		req, err := http.NewRequest("GET", getProjectIssueList, nil)
		req.SetBasicAuth(*username, *password)
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

		var jiraObject JiraObject
		errorMessage := json.NewDecoder(resp.Body).Decode(&jiraObject)
		if errorMessage != nil {
			log.Fatal(errorMessage)
		}

		// If value of Total it seems Project name is error.
		if jiraObject.Total == 0 {
			var errorMessageObject objectErrorMessage
			json.NewDecoder(respError).Decode(&errorMessageObject)
			for x := range errorMessageObject.ErrorMessages {
				fmt.Println(errorMessageObject.ErrorMessages[x])
			}
			log.Fatal("Issue of Project is empty")
		}

		fmt.Println("Total : ", jiraObject.Total)
		fmt.Println("MaxResult : ", jiraObject.MaxResults)

		for x := range jiraObject.Issues {
			Key := jiraObject.Issues[x].Key
			fmt.Println("Issue TYPE", jiraObject.Issues[x].Fields.IssueType.Name)
			fmt.Println(Key)

			changeLogCommand := fmt.Sprintf("%s/rest/api/2/issue/%s?expand=changelog", *jiraweb, Key)

			reqChangeLog, err := http.NewRequest("GET", changeLogCommand, nil)
			reqChangeLog.SetBasicAuth(*username, *password)
			reqChangeLog.Header.Add("Content-Type", "application/json")
			respChangeLog, err := client.Do(reqChangeLog)

			if err != nil {
				log.Fatal(err)
			}

			var jiraChangeLog JiraIssuesChangeLog
			json.NewDecoder(respChangeLog.Body).Decode(&jiraChangeLog)

			for y := range jiraChangeLog.ChangeLog.Histories {
				ChangeLogID := jiraChangeLog.ChangeLog.Histories[y].ID
				ChangeLogCreated := jiraChangeLog.ChangeLog.Histories[y].Created

				fmt.Printf("   ==>  ID -> %s , Created -> %s  \n", ChangeLogID, ChangeLogCreated)
			}
		}
	} else {
		CommandName := strings.Split(os.Args[0], "/")
		fmt.Printf("Please chekc help file by \" %s -h \" \n", CommandName[len(CommandName)-1])
	}
}
