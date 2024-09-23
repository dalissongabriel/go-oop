package main

import (
	"fmt"

	"estudos/oop/clientes"
	"estudos/oop/contas"
)

func main() {
	contaAlisson := contas.ContaCorrente{Titular: clientes.Titular{Documento: "123.345.678-9", Nome: "Alisson", Profissao: "Programador"}, NumeroAgencia: 123, NumeroConta: 9987}
	contaAna := contas.ContaPoupanca{Titular: clientes.Titular{Documento: "443.345.548-9", Nome: "Ana", Profissao: "UX Designer"}, NumeroAgencia: 441, NumeroConta: 1112}

	contaAlisson.Depositar(100)
	contaAna.Depositar(500)

	contaAlisson.Transferir(50, &contaAna)

	fmt.Println(contaAlisson.ObterSaldo(), contaAna.ObterSaldo())
}
