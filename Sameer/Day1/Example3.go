package main

import "fmt"

type Employee interface {
	CalculateSalary() float64
}

type FTE struct {
	monthlySalary float64
	Days          int
}

func (f FTE) CalculateSalary() float64 {
	return f.monthlySalary / 30 * float64(f.Days)
}

type Contractor struct {
	monthlyPay float64
	Days       int
}

func (c Contractor) CalculateSalary() float64 {
	return c.monthlyPay / 30 * float64(c.Days)
}

type Freelancer struct {
	hourlyRate float64
	Hour       int
}

func (f Freelancer) CalculateSalary() float64 {
	return f.hourlyRate * float64(f.Hour)
}

func main() {

	fullTime := FTE{monthlySalary: 15000, Days: 20}
	contractor := Contractor{monthlyPay: 3000, Days: 20}
	freelancer := Freelancer{hourlyRate: 100, Hour: 60}

	fmt.Print("FullTime Pay:", fullTime)
	fmt.Print("Contractor Pay:", contractor)
	fmt.Print("Freelance Pay:", freelancer)

}
