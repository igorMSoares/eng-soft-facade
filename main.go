package main

import (
	"fmt"
	"igorMSoares/eng-soft-facade/combo"
)

func main() {
	for true {
		fmt.Println("Selecione seu combo:")
		fmt.Println(">> 1. Combo Master")
		fmt.Println(">> 2. Super Combo")
		fmt.Println("\nSua escolha (1 ou 2):")

		tipo := 0

		_, err := fmt.Scanln(&tipo)
		if err != nil {
			fmt.Println(">> Opção inválida. Digite 1 ou 2.")
			fmt.Println()

			var flushStdin string
			fmt.Scanln(&flushStdin)

			continue
		}

		combo := combo.NewCombo()
		combo.CriarCombo(tipo)

		fmt.Println()
		fmt.Println(combo)

		break
	}
}
