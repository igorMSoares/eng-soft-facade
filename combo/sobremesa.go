package combo

import "fmt"

type sobremesa struct {
	// atributos privados
	tamanho string
	produto
}

// construtor público
func NewSobremesa(desc string, preco float64, tamanho string) *sobremesa {
	return &sobremesa{
		produto: produto{
			descricao: desc,
			preco:     preco,
		},
		tamanho: tamanho,
	}
}

func (s sobremesa) String() string {
	if s.tamanho == "" {
		return "Nenhuma sobremesa selecionada"
	}

	return fmt.Sprintf("%s (tamanho: %s)", &s.produto, s.tamanho)
}
