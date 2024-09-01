package Contas

import "banco/Clientes"

type ContaCorrente struct {
	Titular              Clientes.Titular
	NumAgencia, NumConta int
	saldo                float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "saque realizado com sucesso "
	} else {
		return "saldo insuficiente"
	}

}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito <= 0 {
		return "Valor do deposito menor ou igual a zero", c.saldo
	} else {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo
	}

}
func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > 0 && valorTransferencia <= c.saldo {
		c.saldo -= valorTransferencia
		contaDestino.saldo += valorTransferencia
		return true
	} else {
		return false
	}

}
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
