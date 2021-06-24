package dao

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	instance := NewDatabaseInstance()
	fmt.Println(instance.maxUserId)
}

func TestSave(t *testing.T) {
	instance := NewDatabaseInstance()
	instance.Save()
}
