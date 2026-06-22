package main

import "fmt"

const bloodType = "A+"

type plansHealth int

const (
	_ plansHealth = iota
	Basico
	Intermediario
	Premium
)

func main() {

	patientName := "Swetony"
	age := 20
	weight := 120.00
	var isActive bool

	fmt.Printf("Nome do paciente: %s\n", patientName)
	fmt.Printf("Idade do paciente: %d\n", age)
	fmt.Printf("Peso do paciente: %.2f kg\n", weight)
	fmt.Printf("Paciente é ativo: %t\n", isActive)
	fmt.Printf("O tipo sanguíneo do paciente: %s\n", bloodType)
	fmt.Printf("Tipo do plano do paciente é: %d\n", Basico)
	fmt.Printf("Tipo da variável weight: %T\n", weight)
	fmt.Printf("Tipo do plano Basico: %T\n", Basico) 

}