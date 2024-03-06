package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ernestokorpys/gobackend/database"
	"github.com/ernestokorpys/gobackend/models"
	"github.com/ernestokorpys/gobackend/util"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]interface{} //map diccionario clave valor [string] las claves son strings si o si; interface de cualquier tipo de valor

	var userData models.User
	if err := c.BodyParser(&data); //se llama al metodo y se almacena la posición del cuerpo de la solicitud en data
	err != nil {
		fmt.Println("Unable to parse body")
	}
	//Check password
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"messaje": "Password must be at least 7 characters long.",
		})
	}
	//Revisa si el email ya esta en la base de datos se le pasa una conversion del Json como argumento
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid email format.",
		})
	}
	//Revisa si el email ya esta en la base de datos se le pasa una conversion del Json como argumento
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "This email is already registered.",
		})
	}

	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Acount create succesfully.",
	})
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z\d\-]+)*\.[a-z]+\z`)
	return Re.MatchString(email)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse Body")
	}
	var user models.User
	database.DB.Where("email = ? ", data["email"]).First(&user)
	//Busaca el primer registro en la DB en que aparezca el email y lo guarda en user es con
	//&user debido a que first debe cargar directamente el valor en una posicion especifica
	//si no hay nada carga user.id=0 y quedan el resto de espacios vacios"
	fmt.Println(user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Invalid Email or Password",
			// "user":    user,
		})
	}

	//user. pasa los datos guardados luego del where en &user que se buscaron y los compara con los enviados en el ultimo json que son solo correo y pass
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	//generación del token y confverison del id a entero de string ya que es un tipo unit
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	//guarda y mantiene la sesion iniciada en base al token provisto, si este caduca cierra la sesion
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"massage": "Succesful login",
		"user":    user,
	})
}

type Claims struct {
	jwt.StandardClaims
}
