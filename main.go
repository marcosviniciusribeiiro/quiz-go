package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type GameState struct{
	Name string
	Points int
	Quiz string
	Type string
	Questions []Question
}

type Question struct{
	Text string
	Options []string
	Answer int
}

func (g *GameState) Choose(){
	var num int
	for{
		fmt.Println("Tipos de Quiz:\n 1 - Quiz de Conhecimentos Gerais\n 2 - Quiz de História\n 3 - Quiz de Inglês")
		fmt.Println("\nEscolha um Tipo de Quiz:")
		reader := bufio.NewReader(os.Stdin)
		read,_ := reader.ReadString('\n')
		num,_ = toInt(read[:len(read)-2])
		
		switch num{
		case 1:
			g.Quiz = "quiz-conhecimentos-gerais.csv"
			g.Type = "Quiz de Conhecimentos Gerais"
		case 2:
			g.Quiz = "quiz-historia.csv"
			g.Type = "Quiz de História do Brasil"

		case 3:
			g.Quiz = "quiz-ingles.csv"
			g.Type = "Quiz de Inglês"
		default:
			fmt.Println("Por favor digite um número, de 1 à 3")
		continue
		}
		break
	}
}

func (g *GameState) ProcessCSV(s string){
	f, err := os.Open(s)

	if err != nil{
		panic("Erro ao ler o arquivo csv.")
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil{
		panic("Erro ao ler csv.")
	}
	
	for index, record := range records{
		if index > 0{
			correctAnswer, _ := toInt(record[5]) 
			question := Question{
				Text: record[0],
				Options: record[1:5],
				Answer: correctAnswer,
			}
			g.Questions = append(g.Questions, question)
		}
	}
}

func(g *GameState) Init(){
	fmt.Printf("\nSeja bem vindo(a) ao %s!\n", g.Type)
	fmt.Println("Escreva o seu nome:")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil{
		panic("Erro ao ler a string.")
	}
	g.Name = name
	fmt.Printf("Vamos jogar %s\n", g.Name)
}

func (g *GameState) Run(){
	for index, question := range g.Questions{
		fmt.Printf("\033[33m%d. %s\033[0m\n", index+1, question.Text)
				
		for j, option := range question.Options{
			fmt.Printf(" [%d] %s\n", j+1, option)
		}

		fmt.Println("Digite uma alternativa:")

		var answer int
		var err error
		for{
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')
			answer, err = toInt(read[:len(read)-2])

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break			
		}

		if answer == question.Answer{
			fmt.Println("Parabéns, você acertou a resposta!")
			g.Points += 10
			fmt.Println("----------------------------------")
		} else {
			fmt.Println("Ops, você errou...")
			fmt.Println("--------------")
		}
	}
}

func (g *GameState) Grade(n int){
	if n > 80 {
		fmt.Printf("Nota S")
	} else if n > 60 {
		fmt.Printf("Nota A")
	} else if n > 40 {
		fmt.Printf("Nota B")
	} else if n > 20 {
		fmt.Printf("Nota C")
	} else {
		fmt.Printf("Nota D")
	}
}

func toInt(s string) (int, error){
	i, err := strconv.Atoi(s)
	if err != nil{
		return 0, errors.New("por favor digite um número para responder o quiz")
	}
	return i, nil
}

func main() {
	game := &GameState{Points: 0}
	game.Choose()
	go game.ProcessCSV(game.Quiz)
	game.Init()
	game.Run()
	fmt.Printf("\nFim de Jogo!\nVocê fez %d pontos: ", game.Points)
	game.Grade(game.Points)
}