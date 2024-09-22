package combo

import "fmt"

type burguer struct {
	// atributos privados
	gramas int
	produto
}

// construtor público
func NewBurguer(desc string, preco float64, gramas int) *burguer {
	return &burguer{
		produto: produto{
			descricao: desc,
			preco:     preco,
		},
		gramas: gramas,
	}
}

func (b burguer) String() string {
	if b.gramas == 0 {
		return "Nenhum burguer selecionado"
	}

	return fmt.Sprintf("%s (%dg)", &b.produto, b.gramas)
}
