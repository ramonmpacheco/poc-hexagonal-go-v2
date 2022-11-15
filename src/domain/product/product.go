package productdomain

type Product struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty" copier:"ID"`
	Name string `json:"name" bson:"name" copier:"Name"`
}

func (p *Product) isValid() bool {
	return len(p.ID) > 0 && len(p.Name) > 0
}

func (ps *Product) FindById(productOutputPort FindProductDataOutputPort) error {
	if err := productOutputPort.ById(ps); err != nil {
		return err
	}
	return nil
}
