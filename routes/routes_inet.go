package routes

import (
	c "go-fiber-test/colltrollers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InetRoutes(app *fiber.App) {
	api := app.Group("/api") // api

	v1 := api.Group("/v1") // v1
	v2 := api.Group("/v2") // v2
	v3 := api.Group("/v3") // v3
	Profile := v1.Group("/profile")
	app.Use(logger.New())
	Profile.Get("/", c.GetProFile)
	// Provide a minimal config
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
			"admin":   "123456",
			"testgo":  "23012023",
		},
	}))

	// //ข้อ5.0****************************************************************
	v1.Get("/BasicAuth", c.BasicAuth)

	// //ข้อ5.1****************************************************************
	v1.Post("/factorial/:number", c.Factorial)

	// //ข้อ5.2****************************************************************
	v3.Post("/First", c.QueryParam)

	// //ข้อ5.3****************************************************************
	v1.Get("/Controller", c.Controller)

	//ข้อ6 ************************************************************************************************

	v1.Post("/register", c.RegisterValidate)

	//ทดลอง
	v1.Post("/register2", c.RegisterValidate2)

	// //********************************************************************************************************************************
	v1.Get("/v1", c.HelloTest)

	v1.Post("/", c.BodyParserTest)

	v1.Get("/user/:name", c.ParamsTest)

	v1.Post("/inet", c.QueryTest)

	v1.Post("/valid", c.ValidateTest)

	v2.Get("/", c.HelloTestV2)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)
	dog.Get("/DogsDelete", c.GetDogsDelete)
	dog.Get("/DogIDGreater", c.GetDogIDGreater)
	dog.Get("/GetDogsColor", c.GetDogsColor)

	v4 := api.Group("/v4") // v4
	company := v4.Group("/company")
	//CRUD Company
	company.Get("/", c.GetCompany)
	company.Post("/", c.AddCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.RemoveCompany)

	//CRUD Profile

	Profile.Post("/", c.AddProFile)
	Profile.Put("/:id", c.UpdateProFile)
	Profile.Delete("/:id", c.RemoveProFile)
	Profile.Get("/ProfileSummary", c.ProfileSummary)
	//search profile
	Profile.Get("/Search", c.SearchProfile)

	// //********************************************************************************************************************************

}
