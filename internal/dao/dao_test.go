package dao

import (
	"log"
	"testing"
)

func TestMax(t *testing.T) {
	instance := NewDatabaseInstance() //todo: mock
	log.Println(instance.maxUserId)
}

func TestSave(t *testing.T) {
	instance := NewDatabaseInstance()
	instance.Save()
}
