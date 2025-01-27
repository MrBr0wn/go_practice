package main

import (
	"interfaces/company"
	"interfaces/person"
	"interfaces/robot"
)

func main() {
	pers := person.Person{}
	comp := company.Company{}

	comp.Hire(pers)

	robo := &robot.Robot{}
	comp.Hire(robo)
}
