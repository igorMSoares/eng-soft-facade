package combo

type produto struct {
	// atributos privados
	descricao string
	preco     float64
}

func NewProduto(desc string, preco float64) *produto {
	return &produto{
		descricao: desc,
		preco:     preco,
	}
}

func (p produto) String() string {
	return p.descricao
}
