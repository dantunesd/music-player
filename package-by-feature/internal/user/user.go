package user

type User struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name"`
}
