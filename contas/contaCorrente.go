package contas

import "estudos/oop/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino Conta) {
	podeTransferir := valorDaTransferencia > 0 && c.saldo >= valorDaTransferencia
	if podeTransferir {
		c.Sacar(valorDaTransferencia)
		contaDestino.Depositar(valorDaTransferencia)
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
