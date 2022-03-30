package contas

import (
	"fmt"

	"github.com/garcia-paulo/bank-go/clientes"
)

type ContaPoupanca struct {
	Titular    clientes.Titular
	AgenciaNum int
	ContaNum   int
	Operacao   int
	saldo      float64
}

func (c *ContaPoupanca) Sacar(valor float64) {
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

func (c *ContaPoupanca) Depositar(valor float64) {
	if valor < 0 {
		fmt.Println("Valor inválido para depósito")
		return
	}

	c.saldo += valor
	fmt.Println("Depósito realizado com sucesso")
}

func (c *ContaPoupanca) Transferir(valor float64, ContaCorrenteDestino *ContaCorrente) {
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

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
