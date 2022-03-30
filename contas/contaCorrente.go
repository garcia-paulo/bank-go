package contas

import (
	"fmt"

	"github.com/garcia-paulo/bank-go/clientes"
)

type ContaCorrente struct {
	Titular    clientes.Titular
	AgenciaNum int
	ContaNum   int
	saldo      float64
}

func (c *ContaCorrente) Sacar(valor float64) {
	if valor < 0 {
		fmt.Println("Valor inválido para saque")
		return
	}

	if valor > c.saldo {
		fmt.Println("saldo insuficiente")
		return
	}

	c.saldo -= valor
	fmt.Println("Saque realizado com sucesso")
}

func (c *ContaCorrente) Depositar(valor float64) {
	if valor < 0 {
		fmt.Println("Valor inválido para depósito")
		return
	}

	c.saldo += valor
	fmt.Println("Depósito realizado com sucesso")
}

func (c *ContaCorrente) Transferir(valor float64, ContaCorrenteDestino *ContaCorrente) {
	if valor < 0 {
		fmt.Println("Valor inválido para transferência")
		return
	}

	if valor > c.saldo {
		fmt.Println("saldo insuficiente")
		return
	}

	c.saldo -= valor
	ContaCorrenteDestino.saldo += valor
	fmt.Println("Transferência realizada com sucesso")
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
