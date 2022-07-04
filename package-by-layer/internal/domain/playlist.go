package domain

type Playlist struct {
	ID     string   `json:"id" bson:"_id,omitempty"`
	UserID string   `json:"user_id"`
	Name   string   `json:"name"`
	Songs  []string `json:"songs"`
}
