package main

import (
	"fmt"
)

// Employee interface defines a method to calculate salary
type Employee interface {
	CalculateSalary() float64
	GetType() string
}

// FullTime struct represents a full-time employee
type FullTime struct {
	DailyRate  float64
	DaysWorked float64
}

// Contractor struct represents a contractor
type Contractor struct {
	DailyRate  float64
	DaysWorked float64
}

// Freelancer struct represents a freelancer
type Freelancer struct {
	HourlyRate  float64
	HoursWorked float64
}

// CalculateSalary method for FullTime employees
func (f FullTime) CalculateSalary() float64 {
	return f.DailyRate * f.DaysWorked
}

// GetType method for FullTime
func (f FullTime) GetType() string {
	return "Full-Time Employee"
}

// CalculateSalary method for Contractors
func (c Contractor) CalculateSalary() float64 {
	return c.DailyRate * c.DaysWorked
}

// GetType method for Contractor
func (c Contractor) GetType() string {
	return "Contractor"
}

// CalculateSalary method for Freelancers
func (fr Freelancer) CalculateSalary() float64 {
	return fr.HourlyRate * fr.HoursWorked
}

// GetType method for Freelancer
func (fr Freelancer) GetType() string {
	return "Freelancer"
}

func main() {
	fullTimeEmp := FullTime{DailyRate: 500.75, DaysWorked: 22.5}
	contractorEmp := Contractor{DailyRate: 300.50, DaysWorked: 20}
	freelancerEmp := Freelancer{HourlyRate: 150.25, HoursWorked: 30}

	employees := []Employee{fullTimeEmp, contractorEmp, freelancerEmp}

	for _, emp := range employees {
		fmt.Printf("%s Salary: %.2f\n", emp.GetType(), emp.CalculateSalary())
	}
}
