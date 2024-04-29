package main

import (
	"github.com/banggibima/go-fiber-jwt/config"
	"github.com/banggibima/go-fiber-jwt/internal/interface/http"
	"github.com/banggibima/go-fiber-jwt/pkg/fiber"
	"github.com/banggibima/go-fiber-jwt/pkg/gorm"
	"github.com/banggibima/go-fiber-jwt/pkg/postgres"
	"github.com/banggibima/go-fiber-jwt/pkg/viper"
)

func main() {
	v, err := viper.New()
	if err != nil {
		panic(err)
	}

	c, err := config.Init(v)
	if err != nil {
		panic(err)
	}

	p, err := postgres.New(c)
	if err != nil {
		panic(err)
	}

	g, err := gorm.New(p)
	if err != nil {
		panic(err)
	}

	f, err := fiber.New(c)
	if err != nil {
		panic(err)
	}

	if err := http.NewServer(c, g, f).Start(); err != nil {
		panic(err)
	}
}
