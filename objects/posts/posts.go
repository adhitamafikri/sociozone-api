package objects

import "github.com/adhitamafikri/sociozone-api/objects/photo"

// Post is the actual post object
type Post struct {
	Photos  []objects.Photo
	Caption string
	Likes   int
}
