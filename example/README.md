# jiraGetIssues Example
## Get Jira Project Issue Detail Information & Store Data into SQL Server ( Event : Insert/Update/None )

## Usage of this script is below.
* -jiraweb
  Jira Web Site Address, example : -jiraweb=https://inhouse.htcstudio.com/jira
* -password
  Password
* -projectname
  Project Name, example : -projectname=TYGH
* -username
  User Name
* -sqlserverinfo
  Input your sql server information [UserName]:[Password]@tcp([SERVER IP:PORT])/[DatabaseName] , example:  eli:eli@tcp(127.0.0.1:3306)/JiraData
* -sqltablename
  Input your sql talbe name


## Below is the table information for this applicaton
```
 CREATE TABLE `Tickets` (
   `sn` int(10) NOT NULL AUTO_INCREMENT,
   `issuekey` char(20) COLLATE utf8_unicode_ci NOT NULL,
   `issuetype` char(200) COLLATE utf8_unicode_ci NOT NULL,
   `issueid` int(20) NOT NULL,
   `issueself` text COLLATE utf8_unicode_ci NOT NULL,
   `project` char(250) COLLATE utf8_unicode_ci NOT NULL,
   `summary` text COLLATE utf8_unicode_ci NOT NULL,
   `priority` char(200) COLLATE utf8_unicode_ci NOT NULL,
   `resolution` char(100) COLLATE utf8_unicode_ci NOT NULL,
   `status` char(100) COLLATE utf8_unicode_ci NOT NULL,
   `lastchange` datetime NOT NULL,
   `reporter` char(200) COLLATE utf8_unicode_ci NOT NULL,
   `assignee` char(200) COLLATE utf8_unicode_ci NOT NULL,
   `labels` text COLLATE utf8_unicode_ci NOT NULL,
   `fixversions` text COLLATE utf8_unicode_ci NOT NULL,
   `component` varchar(1000) COLLATE utf8_unicode_ci NOT NULL,
   `affectversions` varchar(1000) COLLATE utf8_unicode_ci NOT NULL,
   PRIMARY KEY (`issuekey`),
   UNIQUE KEY `sn` (`sn`),
   UNIQUE KEY `issuekey` (`issuekey`),
   KEY `lastchange` (`lastchange`),
   KEY `status` (`status`),
   KEY `resolution` (`resolution`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=COMPACT;
```
