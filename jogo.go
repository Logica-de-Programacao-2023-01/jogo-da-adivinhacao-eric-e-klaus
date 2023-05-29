package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var JogarNovamente string
	TotaldeTentativas := 0
	var NumDigitado int
	fmt.Println("Bem vindo ao jogo da adivinhação!")

	for {
		RandomNum := rand.Intn(99) + 1
		NumTentativas := 0
		fmt.Println("Tente adivinhar o número aleatório.")
		fmt.Print("Digite um número inteiro entre 1 e 100: ")
		fmt.Scan(&NumDigitado)
		for {
			if RandomNum > NumDigitado {
				fmt.Println("O número é maior que", NumDigitado)
				fmt.Print("Tente novamente: ")
				fmt.Scan(&NumDigitado)
				NumTentativas++
				TotaldeTentativas++
			} else if RandomNum < NumDigitado {
				fmt.Println("O número é menor que", NumDigitado)
				fmt.Print("Tente novamente: ")
				fmt.Scan(&NumDigitado)
				NumTentativas++
				TotaldeTentativas++
			} else if RandomNum == NumDigitado {
				fmt.Println("Você acertou!")
				NumTentativas++
				TotaldeTentativas++
				break
			}
		}
		fmt.Printf("Número de tentativas: %d\n", NumTentativas)
		fmt.Println("Deseja jogar novamente? (s/n)")
		fmt.Scan(&JogarNovamente)
		if JogarNovamente == "n" {
			break
		}
	}
}
