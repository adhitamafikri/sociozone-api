package objects

// PhotosObject stores URL of a photo in a post
type PhotosObject struct {
	URL string
}

// Post is the actual post object
type Post struct {
	Photos  []PhotosObject
	Caption string
	Likes   int
}
