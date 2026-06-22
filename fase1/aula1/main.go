package main

import (
	"bufio"   // leitura de input com buffer (mais robusto que Scanln)
	"fmt"     // formatação de output (como System.out em Java)
	"os"      // acesso ao sistema operacional (stdin, stdout, args)
	"strings" // manipulação de strings
	"time"    // data e hora
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o seu nome: ")

	inputName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler o nome: ", err)
		os.Exit(1)
	}

	fmt.Print("Digite a sua cidade: ")

	inputCity, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler a cidade: ", err)
		os.Exit(1)
	}

	name := strings.TrimSpace(inputName)
	city := strings.TrimSpace(inputCity)

	if name == ""{
		fmt.Println("O nome é obrigatório")
		os.Exit(1)
	}

	if city == "" {
		fmt.Println("A cidade é obrigatória")
		os.Exit(1)
	}

	greeting := getGreeting()

	message := fmt.Sprintf("%s, %s! Você é de %s. Bem-vindo ao Go.", greeting, name, city)

	fmt.Println(message)
}

func getGreeting() string {
	hour := time.Now().Hour()

	switch {
	case hour >= 5 && hour < 12:
		return "Bom dia"
	case hour >= 12 && hour < 18:
		return "Boa tarde"
	default:
		return "Boa noite"
	}
}
