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
func UpdateJiraDB(data InjectData, mySQLInfo string, dbTableName string) {
	fmt.Println("MySQL Server Information :", mySQLInfo)

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
		fmt.Println("Method ===> ", method)
	}
	db.Close()
}
