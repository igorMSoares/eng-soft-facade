package combo

import "fmt"

type combo struct {
	burguer   burguer
	bebida    bebida
	sobremesa sobremesa
}

// método público
func NewCombo() *combo {
	return &combo{}
}

// método público
func (c *combo) CriarCombo(tipo int) {
	switch tipo {
	case 1:
		c.criarComboMaster()
	case 2:
		c.criarSuperCombo()
	default:
		fmt.Println("Combo inválido")
	}
}

// método privado
func (c *combo) criarComboMaster() {
	c.burguer = *NewBurguer("Master Burguer", 12.69, 400)
	c.bebida = *NewBebida("Master Tubaína", 8.20, 400)
	c.sobremesa = *NewSobremesa("Master Gateau", 13.39, "M")
}

// método privado
func (c *combo) criarSuperCombo() {
	c.burguer = *NewBurguer("Super Burguer", 18.69, 600)
	c.bebida = *NewBebida("Super Tubaína", 12.20, 500)
	c.sobremesa = *NewSobremesa("Super Gateau", 17.39, "G")
}

func (c *combo) String() string {
	return fmt.Sprintf("-----COMBO-----\n%v\n%v\n%v\n---------------", c.burguer, c.bebida, c.sobremesa)
}
