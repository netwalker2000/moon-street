package dao

import (
	"log"
	"moon-street/internal/di"
	"testing"
)

func TestMax(t *testing.T) {
	di.InitDependenciesUseFactories()
	instance := di.InstancesInjection[ComponentName].(UserRepo)
	log.Println(instance)
}
