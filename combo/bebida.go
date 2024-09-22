package combo

import "fmt"

type bebida struct {
	// atributos privados
	ml int
	produto
}

// construtor público
func NewBebida(desc string, preco float64, ml int) *bebida {
	return &bebida{
		produto: produto{
			descricao: desc,
			preco:     preco,
		},
		ml: ml,
	}
}

func (b bebida) String() string {
	if b.ml == 0 {
		return "Nenhuma bebida selecionada"
	}

	return fmt.Sprintf("%s (%dml)", &b.produto, b.ml)
}
