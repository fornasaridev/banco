package Contas

import "banco/Clientes"

type ContaPoupanca struct {
	Titular                        Clientes.Titular
	NumAgencia, NumConta, Operacao int
	saldo                          float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "saque realizado com sucesso "
	} else {
		return "saldo insuficiente"
	}

}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito <= 0 {
		return "Valor do deposito menor ou igual a zero", c.saldo
	} else {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo
	}

}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
