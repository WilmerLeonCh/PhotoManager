package config

import (
	"github.com/Netflix/go-env"
	_ "github.com/joho/godotenv/autoload"

	"github.com/PhotoManager/utils"
)

var Parsed Config

func init() {
	utils.Must(env.UnmarshalFromEnviron(&Parsed))
}
