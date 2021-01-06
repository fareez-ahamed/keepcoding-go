package main

import "fmt"

// Employee is a struct
type Employee struct {
	Name string
}

// Manager is a struct
type Manager struct {
	title string
}

// Title returns manager's title
func (m Manager) Title() string {
	return m.title
}

// HasTitle is an interface
type HasTitle interface {
	Title() string
}

func main() {
	x := 235
	print(x)
	print(&Employee{Name: "Fareez"})
	print(Manager{title: "Associate Manager"})
}

func print(data interface{}) {
	switch val := data.(type) {
	case string, int32:
		fmt.Printf("%s is a string\n", val)
	case int:
		fmt.Printf("%d is an integer\n", val)
	case *Employee:
		fmt.Printf("%s is an Employee\n", val.Name)
	case HasTitle:
		fmt.Printf("%s is title of given type\n", val.Title())
	default:
		fmt.Printf("%v Unsupported type\n", val)
	}
}
