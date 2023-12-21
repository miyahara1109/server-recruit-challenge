package model

type ExtendAlbum struct {
	ID       AlbumID  `json:"id"`
	Title    string   `json:"title"`
	Singer   *Singer   `json:"singer"`
}
