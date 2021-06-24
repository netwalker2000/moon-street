package dao

import (
	"database/sql"
	"fmt"
	"moon-street/config"
	"time"

	"github.com/spf13/viper"
)

const maxDatabaseConnections = 100

var dbUsername string
var dbPassword string
var dbName string
var dbHost string
var dbPort string

type SqlDatabase struct {
	database  *sql.DB
	maxUserId uint64
}

func init() {
	//to check: race-condition?
	config.InitConfig()
	summary := viper.GetViper().AllKeys()
	fmt.Println(summary)
	dbUsername = viper.GetString("database.username")
	dbPassword = viper.GetString("database.password")
	dbHost = viper.GetString("database.host")
	dbPort = viper.GetString("database.port")
	dbName = viper.GetString("database.name")
}

func NewDatabaseInstance() *SqlDatabase {
	s := &SqlDatabase{}
	s.openDatabase()
	return s
}

func (s *SqlDatabase) Save() {
	stmt := fmt.Sprintf("INSERT into user_tab(name,status, password, email, created_at, updated_at) values ('user%d', 0, 'password', 'user@email.com', current_timestamp, current_timestamp);", s.maxUserId)
	result, err := s.database.Exec(stmt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func (s *SqlDatabase) openDatabase() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxOpenConns(maxDatabaseConnections)
	db.SetMaxIdleConns(maxDatabaseConnections)

	s.database = db
	var maxId uint64 = 0
	stmt := "SELECT max(id) FROM user_tab;"
	result := s.database.QueryRow(stmt)
	result.Scan(&maxId)
	fmt.Printf("Max user id: %d\n", maxId)
	s.maxUserId = maxId
}
