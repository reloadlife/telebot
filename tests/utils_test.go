package tests

import (
	"fmt"
	types2 "go/types"
	"golang.org/x/tools/go/packages"
	"reflect"
)

func getAllTypes(packageName string) ([]types2.Type, error) {
	// Load the package
	cfg := &packages.Config{
		Mode: packages.NeedTypes,
	}

	pkgs, err := packages.Load(cfg, packageName)
	if err != nil {
		return nil, fmt.Errorf("error loading package %s: %w", packageName, err)
	}

	// Get all types from the package
	var allTypes []types2.Type
	for _, pkg := range pkgs {
		for _, name := range pkg.Types.Scope().Names() {
			// Construct the fully qualified type name
			typeFullName := name
			// isPrivate ? ignore it
			if !pkg.Types.Scope().Lookup(name).Exported() {
				//_, _ = fmt.Println("Ignoring private type: ", typeFullName)
				continue
			}

			// Resolve the type by name
			typ, err := getTypeByName(typeFullName, pkgs)
			if err != nil {
				return nil, fmt.Errorf("error resolving type %s: %w", typeFullName, err)
			}

			allTypes = append(allTypes, typ)
		}
	}

	return allTypes, nil
}

// getTypeByName returns the reflect.Type for a fully qualified type name
func getTypeByName(typeName string, pkgs []*packages.Package) (types2.Type, error) {
	// Attempt to find the type by name
	for _, pkg := range pkgs {
		if typ := pkg.Types.Scope().Lookup(typeName); typ != nil {
			return typ.Type(), nil
		}
	}

	return nil, fmt.Errorf("type not found: %s", typeName)
}

func removeDuplicates[T any](slice []T) []T {
	keys := make(map[any]bool)
	list := []T{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func isStruct(v interface{}) bool {
	t := reflect.TypeOf(v)

	// Check if the type is a struct and not an interface
	return t.Kind() == reflect.Struct
}
