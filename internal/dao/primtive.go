package dao

import (
	"database/sql"
	"fmt"
	"log"
	"moon-street/config"
	"time"
)

const maxDatabaseConnections = 100

var databaseInstanceSingleton *SqlDatabase = openDatabase(config.ConfSingleton)

type SqlDatabase struct {
	database  *sql.DB
	maxUserId uint64
}

func GetDatabaseInstance() *SqlDatabase {
	return databaseInstanceSingleton
}

func (s *SqlDatabase) Save() {
	stmt := fmt.Sprintf("INSERT into user_tab(name,status, password, email, created_at, updated_at) values ('user%d', 0, 'password', 'user@email.com', current_timestamp, current_timestamp);", s.maxUserId)
	result, err := s.database.Exec(stmt)
	if err != nil {
		log.Printf("Error when save! %v", err) //return the error
	}
	log.Println(result)
}

func openDatabase(configuration config.Config) *SqlDatabase {
	s := &SqlDatabase{}
	var dbUsername string = configuration.Database.Username
	var dbPassword string = configuration.Database.Password
	var dbName string = configuration.Database.Name
	var dbHost string = configuration.Database.Host
	var dbPort int = configuration.Database.Port
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully connect to database: %v", db)
	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxOpenConns(maxDatabaseConnections)
	db.SetMaxIdleConns(maxDatabaseConnections)

	s.database = db
	var maxId uint64 = 0
	stmt := "SELECT max(id) FROM user_tab;" //replace the ping lang
	result := s.database.QueryRow(stmt)
	result.Scan(&maxId)
	log.Printf("Max user id: %d\n", maxId)
	s.maxUserId = maxId
	return s
}
