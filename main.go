package main

import (
	"banco/Contas"
	"fmt"
)

func main() {
	contaGustavo := Contas.ContaPoupanca{}

	contaGustavo.Depositar(100)
	fmt.Println(contaGustavo.ObterSaldo())
	PagarBoleto(&contaGustavo, 60)
	fmt.Println(contaGustavo.ObterSaldo())

	contaNatalia := Contas.ContaCorrente{}
	contaNatalia.Depositar(500)
	PagarBoleto(&contaNatalia, 400)
	fmt.Println(contaNatalia.ObterSaldo())

}
func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)

}

type verificarConta interface {
	Sacar(valor float64) string
}
