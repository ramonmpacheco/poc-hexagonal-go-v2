package model

type ProductModel struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty" copier:"ID"`
	Name string `json:"name" bson:"name" copier:"Name"`
}
