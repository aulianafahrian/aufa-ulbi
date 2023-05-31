package main

import (
	"log"

	"github.com/aulianafahrian/aufa-ulbi/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/aulianafahrian/aufa-ulbi/url"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/swagger"
	_ "github.com/aulianafahrian/aufa-ulbi/docs"
)

// @title Anganaufea TES SWAG
// @version 1.0
// @description this is a simple server

// @contact.name API Support
// @contact.url https://github.com/aulianafahrian
// @contact.email 1214049@std.ulbi.ac.id

// @host aufa-ulbi.herokuapp.com
// @BasePath /
// @scheme https http

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
