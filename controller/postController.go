package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/ernestokorpys/gobackend/database"
	"github.com/ernestokorpys/gobackend/models"
	"github.com/ernestokorpys/gobackend/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("invalid Parse")
	}
	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "Invalid Payload 888"})
	}
	return c.JSON(fiber.Map{
		"message": "Post is live 34234",
		"massage": blogpost,
	})
}
func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1")) //se fija el parametro page en la url si no tine asume que es 1
	limit := 2
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	//"User" es el parametro de tabla relacionada en este caso a Blogs que necesito
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	//ceunta el total de publicaciones traidas de la base de datos
	//Esto especifica el modelo sobre el cual se realizará la consulta. En este caso, &models.Blog{} crea una instancia vacía del modelo Blog, que se usa para interactuar con la tabla de base de datos asociada a ese modelo.
	database.DB.Model(&models.Blog{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": getblog,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("Users").First(&blogpost)
	return c.JSON(fiber.Map{
		"data": blogpost,
	})
}

// Actualiza un post en base al id de un usuario
// actualiza el post cuyo id esta en la url de la pagina
func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}

	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse")
	}

	database.DB.Model(&blog).Updates(blog)
	return c.JSON(fiber.Map{
		"message": "post updated",
	})
}
func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.Parsejwt(cookie)
	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id=?", id).Preload("User").Find(&blog)
	return c.JSON(blog)
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}
	deleteQuery := database.DB.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post Deleted",
	})
}
