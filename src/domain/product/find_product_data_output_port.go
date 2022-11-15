package productdomain

type FindProductDataOutputPort interface {
	ById(product *Product) error
}
