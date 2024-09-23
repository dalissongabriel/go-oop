package contas

type Conta interface {
	Depositar(valor float64)
	ObterSaldo() float64
}
