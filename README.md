# Example of Golang CRUD using MySQL from scratch
In this tutorial, we are going to see an example program to learn how to do database CRUD operations using Golang and MySQL. CRUD is an acronym for Create, Read, Update, and Delete. CRUD operations are basic data manipulation for database.

In this example, we are going to create an interface as database front end to handle these operations. We have Employee table containing Employee information like id, name and city. With this table, we have to perform CRUD using MySQL. 
## Step 1: Prepare and Import MySQL driver into your project
Using Git Bash first install driver for Go's MySQL database package. Run below command and install MySQL driver's
```bash 
go get -u github.com/go-sql-driver/mysql
```
Now create Goblog Database

1. Open PHPMyAdmin/SQLyog or what ever MySQL database management tool that you are using.
2. Create a new database "database_name" 

## Creating the Karyawan Table

Execute the following SQL query to create a table named Employee inside your MySQL database. We will use this table for all of our future operations.

```mysql
DROP TABLE IF EXISTS `karyawan`;
CREATE TABLE `karyawan` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `first_name` varchar(30) NOT NULL,
  `email` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
```

## Creating Struct, Handler and Handler Function

Let's create a file named main.go and put the following code inside it.
We usually import database/sql and use sql to execute database queries on the database.
Function dbConn opens connection with MySQL driver.
We will create Employee struct that has following properties: Id, Name and City.

## Ending
```go
go run main.go
```
