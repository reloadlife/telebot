package di

import (
	"fmt"
	"github.com/sarulabs/di"
)

var Container di.Container

func setupServices(services ...*di.Def) {
	builder, _ := di.NewBuilder()

	for _, service := range services {
		err := builder.Add(*service)
		if err != nil {
			fmt.Printf("Failed to register service: %v, Err: %v\n", service.Name, err.Error())
		}
	}
	Container = builder.Build()
	for _, service := range services {
		_, err := Container.SafeGet(service.Name)
		if err != nil {
			fmt.Printf("Failed to load service: %v, Err: %v\n", service.Name, err.Error())
		}
	}

	fmt.Printf("Loaded %v services\n", len(Container.Definitions()))
}

func getService[T interface{}](name string) T {
	service, err := getServiceSafe[T](name)
	if err != nil {
		fmt.Printf("Failed to load service: %v, Err: %v", name, err.Error())
		return service
	}
	return service
}

func getServiceSafe[T interface{}](name string) (T, error) {
	service, err := Container.SafeGet(name)
	return service.(T), err
}

func removeService() {
	clean()
	err := Container.DeleteWithSubContainers()
	if err != nil {
		fmt.Printf("Failed to remove servicees, Err: %v", err.Error())
		return
	}
}

func clean() {
	err := Container.Clean()
	if err != nil {
		fmt.Printf("Failed to clean, Err: %v", err.Error())
		return
	}
}
