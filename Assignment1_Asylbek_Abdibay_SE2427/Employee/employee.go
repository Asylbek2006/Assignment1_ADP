package Employee

type Employee interface {
	CalculateSalary() float64
	GetWorkHours() int
}

type FullTime struct {
	MonthlySalary float64
	BonusRate     float64
}

func (f FullTime) CalculateSalary() float64 {
	return f.MonthlySalary + (f.MonthlySalary * f.BonusRate)
}
func (f FullTime) GetWorkHours() int { return 160 }

type PartTime struct {
	HourlyRate  float64
	HoursWorked int
}

func (p PartTime) CalculateSalary() float64 {
	return p.HourlyRate * float64(p.HoursWorked)
}
func (p PartTime) GetWorkHours() int { return p.HoursWorked }

type Contractor struct {
	ProjectRate       float64
	ProjectsCompleted int
}

func (c Contractor) CalculateSalary() float64 {
	return c.ProjectRate * float64(c.ProjectsCompleted)
}
func (c Contractor) GetWorkHours() int { return c.ProjectsCompleted * 40 }

type Intern struct {
	DailyRate  float64
	DaysWorked int
}

func (i Intern) CalculateSalary() float64 {
	return i.DailyRate * float64(i.DaysWorked)
}
func (i Intern) GetWorkHours() int { return i.DaysWorked * 8 }
