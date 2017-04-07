package jira

// JiraObject is the first layer of jql search
type JiraObject struct { //xxxx
	Expend     string       `json:"expend"`
	StartAt    int          `json:"startAt"`
	MaxResults int          `json:"maxResults"`
	Total      int          `json:"total"`
	Issues     []JiraIssues `json:"issues"`
}

// JiraIssues is for JiraObject Issues
type JiraIssues struct {
	Expend string     `json:"expend"`
	ID     string     `json:"id"`
	Self   string     `json:"self"`
	Key    string     `json:"key"`
	Fields JiraFields `json:"fields"`
}

// JiraIssuesChangeLog is the feedback json of specific jira ticket
type JiraIssuesChangeLog struct {
	Expend    string                  `json:"expend"`
	ID        string                  `json:"id"`
	Self      string                  `json:"self"`
	Key       string                  `json:"key"`
	Fields    JiraFields              `json:"fields"`
	ChangeLog JiraIssuesChangeLogData `json:"changelog"`
}

// JiraFields is the field structure for JIRA issues
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
	Versions                      []JiraFieldsVersions    `json:"versions"`
	Components                    []JiraFieldsResolution  `json:"components"`
}

// JiraFieldsVersions is the fields of jira versions
type JiraFieldsVersions struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Archived    bool   `json:"archived"`
	Released    bool   `json:"released"`
	ReleaseDate string `json:"releaseDate"`
}

// JiraFieldsProject is the filds of Jira Project
type JiraFieldsProject struct {
	Self         string      `json:"self"`
	ID           string      `json:"id"`
	Key          string      `json:"key"`
	Name         string      `json:"name"`
	EmailAddress string      `json:"emailAddress"`
	AvatarUrls   interface{} `json:"avatarUrls"`
}

// JiraFieldsStatus is the fiedls of Jira Status
type JiraFieldsStatus struct {
	Self        string `json:"self"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	ID          string `json:"id"`
}

// JiraFieldsPriority is the fields of priority
type JiraFieldsPriority struct {
	Self    string `json:"self"`
	IconURL string `json:"iconUrl"`
	Name    string `json:"name"`
	ID      string `json:"id"`
}

//JiraFieldsReporter is the fields of Jira Reporter
type JiraFieldsReporter struct {
	Self         string      `json:"self"`
	Name         string      `json:"name"`
	EmailAddress string      `json:"emailAddress"`
	AvatarUrls   interface{} `json:"avatarUrls"`
	DisplayName  string      `json:"displayName"`
	Active       bool        `json:"active"`
}

// JiraFieldsResolution is the fields of Jira Resolution
type JiraFieldsResolution struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// JiraFieldsIssueType is the fields of Issue Type
type JiraFieldsIssueType struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	SubTask     bool   `json:"subtask"`
}

// JiraFieldsFixVersions is the fields of Fix Versions
type JiraFieldsFixVersions struct {
	Self           string `json:"self"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	Archievd       bool   `json:"archievd"`
	Released       bool   `json:"released"`
	ResolutionDate string `json:"resolutiondate"`
	TimeSpent      string `json:"timespent"`
}

// JiraIssuesChangeLogData is the fiedls of change log data
type JiraIssuesChangeLogData struct {
	StartAt    int                                `json:"startAt"`
	MaxResults int                                `json:"maxResult"`
	Total      int                                `json:"total"`
	Histories  []JiraIssuesChangeLogDataHistories `json:"histories"`
}

// JiraIssuesChangeLogDataHistories is the fields of data history
type JiraIssuesChangeLogDataHistories struct {
	ID      string                                  `json:"id"`
	Author  JiraIssuesChangeLogDataHistoriesAuthor  `json:"author"`
	Created string                                  `json:"created"`
	Items   []JiraIssuesChangeLogDataHistoriesItems `json:"items"`
}

// JiraIssuesChangeLogDataHistoriesAuthor is the fields of Author
type JiraIssuesChangeLogDataHistoriesAuthor struct {
	Self         string      `json:"self"`
	Name         string      `json:"name"`
	EmailAddress string      `json:"emailAddress"`
	AvatarUrls   interface{} `json:"avatarUrls"`
	DisplayName  string      `json:"displayName"`
	Active       bool        `json:"active"`
}

// JiraIssuesChangeLogDataHistoriesItems is the fields of Items
type JiraIssuesChangeLogDataHistoriesItems struct {
	Field      string `json:"field"`
	FieldType  string `json:"fieldtype"`
	From       string `json:"from"`
	FromString string `json:"fromString"`
	To         string `json:"to"`
	ToString   string `json:"toString"`
}

// ObjectErrorMessage is for error message use
type ObjectErrorMessage struct {
	ErrorMessages []string    `json:"errorMessages"`
	Errors        interface{} `json:"errors"`
}

// GetMergetName is the get merge name struct used
type GetMergeName struct {
	Name string `json:"name"`
}
