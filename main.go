package main

import(
     "database/sql"
     "log"
     "net/http"
     "text/template"
     _ "github.com/go-sql-driver/mysql"
)

type User struct{
	Id int
	Name string
	Email string
}

func dbConn() (db *sql.DB){
	dbDriver := "mysql"
	dbUser := "anovan"
	dbPass := "K@W!anaconda098"
	dbName := "core_auth"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("template/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    sql := "SELECT id,first_name,email FROM tbl_user ORDER BY id DESC"
    log.Println("Query: "+sql)
    selDB, err := db.Query(sql)
    if err != nil {
        panic(err.Error())
    }
    emp := User{}
    res := []User{}
    for selDB.Next() {
        var id int
        var first_name, email string
        err = selDB.Scan(&id, &first_name, &email)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = first_name
        emp.Email = email
        res = append(res, emp)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    log.Println("Server started on:"+nId)
    selDB, err := db.Query("SELECT * FROM tbl_user WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := User{}
    for selDB.Next() {
        var id int
        var first_name, email string
        err = selDB.Scan(&id, &first_name, &email)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = first_name
        emp.Email = email
    }
    tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM tbl_user WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := User{}
    for selDB.Next() {
        var id int
        var name, email string
        err = selDB.Scan(&id, &name, &email)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.Email = email
    }
    tmpl.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        insForm, err := db.Prepare("INSERT INTO tbl_user(first_name, email) VALUES(?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, city)
        log.Println("INSERT: Name: " + name + " | City: " + city)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        email := r.FormValue("city")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE tbl_user SET first_name=?, email=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, email, id)
        log.Println("UPDATE: Name: " + name + " | City: " + email)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM tbl_user WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func main() {
    
    log.Println("Server started on: http://localhost:8080")
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":8080", nil)
}
