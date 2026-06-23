package main

import "fmt"

type HealthProvider interface {
	GetName() string
	GetRole() string
}

type Doctor struct {
	Name      string
	CRM       string
	Specialty string
}

type Nurse struct {
	Name  string
	COREN string
	Unit  string
}

func (d Doctor) GetName() string {
	return fmt.Sprintf("Dr(a). %s", d.Name)
}

func (n Nurse) GetName() string {
	return fmt.Sprintf("Enfermeiro(a): %s", n.Name)
}

func (d Doctor) GetRole() string {
	return "Médico(a)"
}

func (n Nurse) GetRole() string {
	return "Enfermeiro(a)"
}

func RegisterProvider(p HealthProvider) {
	fmt.Printf("[%s] %s - Registrado(a) no sistema\n", p.GetRole(), p.GetName())
}

type Notifier interface {
	Notify(message string) error
}

type EmailNotifier struct {
	Email string
}

type SMSNotifier struct {
	Phone string
}

func (e EmailNotifier) Notify(message string) error {
	fmt.Printf("Email para %s: %s\n", e.Email, message)
	return nil
}

func (s SMSNotifier) Notify(message string) error {
	fmt.Printf("SMS para %s: %s\n", s.Phone, message)
	return nil
}

type NotificationService struct {
	notifier Notifier
}

func (ns *NotificationService) Send(msg string) {
	if err := ns.notifier.Notify(msg); err != nil {
		fmt.Println("Erro ao notificar: ", err)
	}
}

func main() {

	doctor := Doctor{Name: "Swetony", CRM: "PE-12345", Specialty: "Clínico Geral"}
	nurse := Nurse{Name: "Ana", COREN: "COREN-PE 123456", Unit: "SUS"}

	RegisterProvider(doctor)
	RegisterProvider(nurse)

	email := NotificationService{notifier: EmailNotifier{Email: "luiz@gmail.com"}}
	sms := NotificationService{notifier: SMSNotifier{Phone: "11999999999"}}

	email.Send("Olá Luiz")
	sms.Send("Olá José")

}