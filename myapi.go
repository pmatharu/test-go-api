package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"database/sql"
	"log"
	"gopkg.in/gorp.v1"
)

type User struct  {
	Id int64 `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname string `db:"lastname" json:"lastname"`
}

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.POST("/users", PostUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
	}

	r.Run(":8088")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "ispy:kjdP9JJnshIUpO@/go_test")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(User{}, "User").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create table failed")
	return dbmap
}


func GetUsers(c *gin.Context) {
	var dbmap = initDb()

	var users []User
	_, err := dbmap.Select(&users, "SELECT * FROM user")
	if err == nil {
		c.JSON(200, users)
	} else {
		c.JSON(404, gin.H{"error": "no user(s) into the table"})
	}
	// curl -i http://localhost:8080/api/v1/users
}

func GetUser(c *gin.Context) {
	var dbmap = initDb()

	id := c.Params.ByName("id")
	var user User
	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)
	if err == nil {
		user_id, _ := strconv.ParseInt(id, 0, 64)
		content := &User{
			Id: user_id,
			Firstname: user.Firstname,
			Lastname: user.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
	// curl -i http://localhost:8080/api/v1/users/1
}

func PostUser(c *gin.Context) {
	// The futur code…
}

func UpdateUser(c *gin.Context) {
	// The futur code…
}
func DeleteUser(c *gin.Context) {
	// The futur code…
}