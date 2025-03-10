// Question 3

package main

import "fmt"

type Employee interface {
	CalculateSalary() int
}

type FullTime struct {
	dailyRate  int
	daysWorked int
}

type Contractor struct {
	dailyRate  int
	daysWorked int
}

type Freelancer struct {
	hourlyRate  int
	hoursWorked int
}

func (f FullTime) CalculateSalary() int {
	return f.dailyRate * f.daysWorked
}

func (c Contractor) CalculateSalary() int {
	return c.dailyRate * c.daysWorked
}

func (fl Freelancer) CalculateSalary() int {
	return fl.hourlyRate * fl.hoursWorked
}

func kh() {
	fullTimeEmployee := FullTime{dailyRate: 500, daysWorked: 30} 
	contractor := Contractor{dailyRate: 100, daysWorked: 30}    
	freelancer := Freelancer{hourlyRate: 100, hoursWorked: 20}   

	employees := []Employee{fullTimeEmployee, contractor, freelancer}

	for _, e := range employees {
		fmt.Println("Salary:", e.CalculateSalary())
	}
}
