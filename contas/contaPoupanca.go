package contas

import "estudos/oop/clientes"

type ContaPoupanca struct {
	Titular        clientes.Titular
	NumeroAgencia  int
	NumeroConta    int
	MetaFinanceira float64
	saldo          float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
	}
}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino Conta) {
	podeTransferir := valorDaTransferencia > 0 && c.saldo >= valorDaTransferencia
	if podeTransferir {
		c.Sacar(valorDaTransferencia)
		contaDestino.Depositar(valorDaTransferencia)
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
