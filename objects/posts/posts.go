package objects

import "github.com/adhitamafikri/sociozone-api/objects/photos"

// Post is the actual post object
type Post struct {
	Photos  []objects.Photos
	Caption string
	Likes   int
}
