package config

import "os"

// ImgbbAPIKey is API key to ImgBB
var ImgbbAPIKey = os.Getenv("ImgbbAPIKey")

// ImgbbAPIURL is API URL to ImgBB
var ImgbbAPIURL = os.Getenv("ImgbbAPIURL")
