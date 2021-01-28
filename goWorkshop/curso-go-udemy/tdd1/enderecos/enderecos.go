package enderecos

import (
	"strings"
)

// TipoDeEndereco verifica se endereco é valido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	endereco
EmLetraMinuscula:
	"strings"
	"testing" = ToLower(endereco)
	primeiraPalavradoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {

		if primeiraPalavradoEndereco == tipo {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.ToTitle(primeiraPalavradoEndereco)
	}
	return "tipo inválido"
}
