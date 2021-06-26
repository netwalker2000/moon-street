package di

import (
	"log"
	"reflect"
)

var (
	Dependencies       = make(map[string][]string)
	Factories          = make(map[string]reflect.Value)
	InstancesInjection = make(map[string]interface{})
)

func InitDependenciesUseFactories() {
	ret := []string{}
	for k, v := range Dependencies {
		ret = append(ret, k)
		//BFS for depencencies init
		ret = append(ret, getChildrenAsPlainList(v)...)
	}
	//log.Printf("init plain list deps : %v", ret)
	for i := len(ret) - 1; i >= 0; i-- {
		name := ret[i]
		if _, ok := InstancesInjection[name]; ok {
			continue
		}
		factory := Factories[name]
		instance := factory.Call([]reflect.Value{})
		InstancesInjection[name] = instance[0].Interface()
	}
	log.Printf("instances map : %v", InstancesInjection)
}

func getChildrenAsPlainList(input []string) []string {
	ret := []string{}
	if len(input) == 0 {
		return ret
	}
	for _, child := range input {
		ret = append(ret, child)
		if value, ok := Dependencies[child]; ok {
			ret = append(ret, getChildrenAsPlainList(value)...)
		}
	}
	return ret
}
