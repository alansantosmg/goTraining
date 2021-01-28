// Teste de unidade

// O nome do arquivo tem que ser sempre algumacoisa_test.go
// O nome do pacote tem que ser o mesmo do pacote  testado

package enderecos

import (
	"testing"
)

// a funcao de teste tem comecar com Test
// seguido do nome da função com Inicio em letra maiuscula
// O parametro é sempre um t do tipo ponteiro de T do pacote testing

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioTeste{
		{"Rua ABC", "rua"},
		{"Avenida ABC", "avenida"},
		{"Estrada ABC", "estrada"},
	}

	for _, cenario := range cenariosDeTeste {

		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}

}
