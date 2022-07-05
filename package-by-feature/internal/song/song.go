package song

type song struct {
	ID         string `json:"id" bson:"_id,omitempty"`
	ArtistName string `json:"artist_name"`
	AlbumName  string `json:"album_name"`
	Name       string `json:"name"`
	Number     int    `json:"number"`
}
