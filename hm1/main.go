package main

import (
	"fmt"
	"hm1/builder"
	"hm1/cook"
	"hm1/driver"
	"hm1/employee"
	"hm1/president"
	"hm1/teacher"
)

func main() {
	teacherEmp := teacher.NewTeacher("teacher", 1000, "Almaty")
	builderEmp := builder.NewBuilder("builder", 500, "Shymkent")
	driverEmp := driver.NewDriver("driver", 1500, "Aktobe")
	cookEmp := cook.NewCook("cook", 3000, "Almaty")
	presidentEmp := president.NewPresident("president", 99999999999, "Astana")

	printEmployeeInfo(teacherEmp)
	printEmployeeInfo(builderEmp)
	printEmployeeInfo(driverEmp)
	printEmployeeInfo(cookEmp)
	printEmployeeInfo(presidentEmp)

}

func printEmployeeInfo(e employee.Employee) {
	fmt.Printf("Position: %s\n", e.GetPosition())
	fmt.Printf("Salary: $%.2f\n", e.GetSalary())
	fmt.Printf("Address: %s\n\n", e.GetAddress())
}
