package productdomain

type Product struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty" copier:"ID"`
	Name string `json:"name" bson:"name" copier:"Name"`
}
