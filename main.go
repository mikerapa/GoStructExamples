package main

import "fmt"

type employee struct {
	name string
	id   int
}

func (e employee) ChangeIDByVal(newID int) {
	e.id = newID
}

func (e *employee) ChangeIDByPointer(newID int) {
	e.id = newID
}

// passed in a pointer
func (e *employee) EmployeeToString() string {
	formattedString := fmt.Sprintf("Employee: %s (%d)", e.name, e.id)
	return formattedString
}

// passed in a value
func (e employee) EmployeeToString2() string {
	formattedString := fmt.Sprintf("Employee: %s (%d)", e.name, e.id)
	return formattedString
}

func main() {
	emp1 := employee{name: "Mike", id: 22}

	fmt.Println("employee 1: ", emp1)

	fmt.Println(emp1.EmployeeToString())
	fmt.Println(emp1.EmployeeToString2())
	// change id by value
	emp1.ChangeIDByVal(99) // Does not really change the value
	fmt.Println(emp1.EmployeeToString2())
	emp1.ChangeIDByPointer(101) // changes the value
	fmt.Println(emp1.EmployeeToString2())
}
