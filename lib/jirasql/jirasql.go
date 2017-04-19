package jirasql

import (
	"database/sql"
	"fmt"
	"time"
)

// InjectData is the struct for Inject DB Data
type InjectData struct {
	Issuekey      string
	Issuetype     string
	Issueid       string
	Issueself     string
	Project       string
	Summary       string
	Priority      string
	Resolution    string
	Status        string
	Lastchange    string
	Reporter      string
	Assignee      string
	Label         string
	Fixversion    string
	Component     string
	Affectversion string
	Startdate     string
	Duedate       string
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
		layout := " 2017-04-14 02:50:33"
		originalTime, _ := time.Parse(layout, lastchange)
		latestTime, _ := time.Parse(layout, LastChange)

		if originalTime.Unix() == latestTime.Unix() {
			fmt.Printf("MySQL Method is %s for %s \n", "NONE", IssueKey)
			return "NONE", ""
		}
		fmt.Printf("MySQL Method is %s for %s \n", "UPDATE", IssueKey)
		return "UPDATE", " WHERE issuekey=?"
	}
	fmt.Printf("MySQL Method is %s for %s \n", "INSERT", IssueKey)
	return "INSERT", ""
}

// UpdateJiraDB is a feature that can add/modify content of one jira ticket.
func UpdateJiraDB(data InjectData, mySQLInfo string, dbTableName string) {
	sqlMethod, sqlExtralCMD := ConfirmMethod(dbTableName, data.Issuekey, data.Lastchange, mySQLInfo)

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
		_, err = stmt.Exec(data.Issuekey,
			data.Issuetype,
			data.Issueid,
			data.Issueself,
			data.Project,
			data.Summary,
			data.Priority,
			data.Resolution,
			data.Status,
			data.Lastchange,
			data.Reporter,
			data.Assignee,
			data.Label,
			data.Fixversion,
			data.Component,
			data.Affectversion,
			data.Issuekey)
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
		res, err := stmt.Exec(data.Issuekey,
			data.Issuetype,
			data.Issueid,
			data.Issueself,
			data.Project,
			data.Summary,
			data.Priority,
			data.Resolution,
			data.Status,
			data.Lastchange,
			data.Reporter,
			data.Assignee,
			data.Label,
			data.Fixversion,
			data.Component,
			data.Affectversion)
		checkErr(err)
		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println("Last Insert ID : ", id)

	default:
	}
	db.Close()
}

// ShowInejctData is the feature that will display inject data by fmt.println
func ShowINjecData(data InjectData) {
	fmt.Println("=============================================================")
	fmt.Println("Issue Key : ", data.Issuekey)
	fmt.Println("Issue Type : ", data.Issuetype)
	fmt.Println("Issue ID : ", data.Issueid)
	fmt.Println("Issue Self : ", data.Issueself)
	fmt.Println("Project : ", data.Project)
	fmt.Println("Summary : ", data.Summary)
	fmt.Println("Priority : ", data.Priority)
	fmt.Println("Resolution : ", data.Resolution)
	fmt.Println("Status : ", data.Status)
	fmt.Println("LastChange : ", data.Lastchange)
	fmt.Println("Reporter : ", data.Reporter)
	fmt.Println("Assignee : ", data.Assignee)
	fmt.Println("Label : ", data.Label)
	fmt.Println("Fix Version : ", data.Fixversion)
	fmt.Println("Component : ", data.Component)
	fmt.Println("Affect Version : ", data.Affectversion)
	fmt.Println("Start Date : ", data.Startdate)
	fmt.Println("DueDate : ", data.Duedate)
}
