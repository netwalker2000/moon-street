package dao

import (
	"database/sql"
	"fmt"
	"log"
	"moon-street/config"
	"moon-street/internal/di"
	"moon-street/internal/model"
	"reflect"
	"time"
)

const maxDatabaseConnections = 100

type UserDataAccessPrimitiveImpl struct {
	database  *sql.DB
	maxUserId uint64
}

const ComponentName = "daoComponent"

func init() {
	di.Dependencies[ComponentName] = []string{}
	di.Factories[ComponentName] = reflect.ValueOf(openDatabase)
}

func (s *UserDataAccessPrimitiveImpl) GetByName(name string) (model.User, error) {
	stmt := fmt.Sprintf("select id, password, email from user_tab where name = '%s'; ", name)
	result, err := s.database.Query(stmt)
	user := &model.User{}
	if err != nil {
		log.Printf("Error when query by name! %v", err) //return the error
		return *user, err
	}
	//log.Println(result)
	var (
		id       int
		password string
		email    string
	)
	for result.Next() {
		err := result.Scan(&id, &password, &email)
		if err != nil {
			log.Printf("error when parse sql result: %v", err)
		}
		//log.Println(id, name, email)
		user.Password = password
		user.Email = email
	}
	return *user, nil
}

func (s *UserDataAccessPrimitiveImpl) Save(user model.User) (int64, error) {
	stmt := fmt.Sprintf("INSERT into user_tab(name, status, password, email, created_at, updated_at) values ('%s', 0, '%s', '%s', current_timestamp, current_timestamp);",
		user.Name, user.Password, user.Email)
	result, err := s.database.Exec(stmt)
	if err != nil {
		log.Printf("Error when save! %v", err) //return the error
		return 0, err
	}
	log.Println(result)
	return 1, nil
}

func openDatabase() *UserDataAccessPrimitiveImpl {
	configuration := config.ConfigSingleton
	s := &UserDataAccessPrimitiveImpl{}
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
