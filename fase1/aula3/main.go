package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
}

type Person struct {
	Name string
	Age  int
	Address
}

func (p Person) FullInfo() string {
	return fmt.Sprintf("Olá %s! Você é de %s e tem %d anos!", p.Name, p.City, p.Age)
}

type Doctor struct {
	Person
	Specialty string
	CRM string
}

func (d *Doctor) Prescribe(medication string) string {
	return fmt.Sprintf("Dr(a). %s prescreveu %s", d.Name, medication)
}

type Patient struct {
	Person
	BloodType string
	IsActive bool
}

func (p *Patient) Activate() {
	p.IsActive = true
}

func (p Patient) Status() string {
	if p.IsActive {
		return "ativo"
	} else {
		return "inativo"
	}
}

func main() {

	doctor := Doctor{
		Person: Person{Name: "Dr. Ana", Age: 25, Address: 
				Address{State: "PE", City: "Limoeiro", Street: "Pirauira"}},
		Specialty: "Clínico Geral",
		CRM: "PE-12345",
	}

	patient := Patient{
		Person: Person{Name: "Juliana", Age: 22, Address: 
				Address{State: "PE", City: "Limoeiro", Street: "São José"}},
		BloodType: "A+",
		IsActive: false,
	}

	fmt.Println(doctor.FullInfo())
	fmt.Println(patient.FullInfo())
	fmt.Println(doctor.Prescribe("Paracetamol"))
	patient.Activate()
	fmt.Println(patient.Status())

}