package config

import "os"

// APIURL is a base URL of the API
var APIURL = os.Getenv("API_BASE_URL")
