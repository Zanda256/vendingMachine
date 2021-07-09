package roles

type Buyer struct {
	ID int
}

type Seller struct {
	ID int
}

func (s *Seller) addProduct() {}

func (s *Seller) updateProducts() {}

func (s *Seller) removeProduct() {}

func (b *Buyer) deposit() {}

func (b *Buyer) buy() {}
