package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	opcoes := []string{"pedra", "papel", "tesoura"}
	vence := map[string]string{
		"pedra":   "tesoura",
		"papel":   "pedra",
		"tesoura": "papel",
	}

	const pontosParaVencer = 3
	meusPontos := 0
	pontosMaquina := 0

	rand.Seed(time.Now().UnixNano())
	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("=== PEDRA, PAPEL E TESOURA ===")
	fmt.Printf("Quem fizer %d pontos primeiro ganha!\n\n", pontosParaVencer)

	for meusPontos < pontosParaVencer && pontosMaquina < pontosParaVencer {
		var escolha string

		// Só aceita as três opções válidas
		for {
			fmt.Print("Sua jogada (pedra, papel ou tesoura): ")
			entrada, _ := leitor.ReadString('\n')
			escolha = strings.ToLower(strings.TrimSpace(entrada))

			if escolha == "pedra" || escolha == "papel" || escolha == "tesoura" {
				break
			}
			fmt.Println("Jogada inválida! Digite apenas: pedra, papel ou tesoura.")
		}

		maquina := opcoes[rand.Intn(len(opcoes))]
		fmt.Println("A máquina escolheu:", maquina)

		if escolha == maquina {
			fmt.Println("Empate!")
		} else if vence[escolha] == maquina {
			meusPontos++
			fmt.Println("Você ganhou essa rodada!")
		} else {
			pontosMaquina++
			fmt.Println("A máquina ganhou essa rodada!")
		}

		fmt.Printf("Placar -> Você: %d | Máquina: %d\n\n", meusPontos, pontosMaquina)
	}

	fmt.Println("=== FIM DE JOGO ===")
	if meusPontos > pontosMaquina {
		fmt.Println("🏆 VOCÊ GANHOU O JOGO!")
	} else {
		fmt.Println("🤖 A MÁQUINA GANHOU O JOGO!")
	}
}
