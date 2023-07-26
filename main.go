package main

import (
	"fmt"
	"strconv"
)

var idade int = 39

//Struct
type Pessoa struct {
	Nome  string
	Idade int
}

//método dizer nome
func (p Pessoa) DizerNome() string {
	return p.Nome
}

func (p Pessoa) dizerNome() {
	fmt.Println(p.Nome)
}

func main() {

	fmt.Println("Hello World!")
	var nome string = "Evandro"
	nome = "Evandro"

	fmt.Println("Meu nome é " + nome + " minha idade" + strconv.Itoa(idade))

	for i := 0; i < 10; i++ {
		fmt.Println(" Iteração " + strconv.Itoa(i))
	}

	i := 0

	for i < 10 {
		fmt.Println(" Iteração " + strconv.Itoa(i))
		i++
	}

	meuArray := [5]int{1, 2, 3, 4, 5}

	primeiroIndice := meuArray[0]

	fmt.Println("Primeiro indice " + strconv.Itoa(primeiroIndice))

	fmt.Println("Indice 0 " + strconv.Itoa(meuArray[0]))
	tamanho := len(meuArray)
	fmt.Println("Tamanho do array " + strconv.Itoa(tamanho))

	//slices
	minhaFatia := []int{10, 20, 30, 40, 50}
	fmt.Println(minhaFatia)

	subFatia := minhaFatia[2:]

	fmt.Println("Subfatia")
	fmt.Println(subFatia)

	fmt.Println(meuArray)

	var idade int
	fmt.Println("Digite sua idade " + strconv.Itoa(idade))

	//mapas
	var meuMapa map[string]string = map[string]string{
		"nome":  "Evandro",
		"idade": "39",
	}

	fmt.Println("Meu nome " + meuMapa["nome"])

	var pessoa1 Pessoa = Pessoa{Nome: "John", Idade: 30}

	fmt.Println(pessoa1.DizerNome())

	idades := []int{10, 20, 30, 40, 50}

	for _, idade := range idades {
		fmt.Println("Idade Dentro do Loop Range " + strconv.Itoa(idade))
	}

	pessoa1.dizerNome()

}
