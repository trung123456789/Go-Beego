package implement

import "gopkg.in/go-playground/validator.v10"

// use a single instance of Validate, it caches struct info
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}
