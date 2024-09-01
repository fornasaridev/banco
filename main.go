package main

import "fmt"

type ContaCorrente struct {
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func main() {
	contaGustavo := ContaCorrente{"Gustavo", 589, 123456, 125.50}
	contaNatalia := ContaCorrente{"Natalia", 589, 123457, 125.50}
	fmt.Println(contaGustavo)
	fmt.Println(contaNatalia)

	var contaMarcelo *ContaCorrente
	contaMarcelo = new(ContaCorrente)
	contaMarcelo.titular = "Marcelo"
	contaMarcelo.numConta = 123458
	contaMarcelo.numAgencia = 589
	contaMarcelo.saldo = 125.50

	fmt.Println(*contaMarcelo)
}
