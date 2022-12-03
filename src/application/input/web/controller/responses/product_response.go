package responses

type ProductResponse struct {
	ID   string `json:"id,omitempty" bson:"id,omitempty" copier:"ID"`
	Name string `json:"name" bson:"name" copier:"Name"`
}
