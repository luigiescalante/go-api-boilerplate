package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.api-boilerplate/infrastructure/db"
	"go.api-boilerplate/models"
	"go.api-boilerplate/services"
	"go.api-boilerplate/utils"
)

const PasswordAuthType = "password"

// Login
//
//	@Login			godoc
//	@Description	Auth a user by email and password
//	@Tags			Login
//	@Produce		json
//	@Param			Login	payload	body	string	true	"Fill the payload with the information"	SchemaExample({"email":"","password":""})
//	@Router			/v1/login [post]
func Login(c *fiber.Ctx) error {
	repo, err := userRepo()
	if err != nil {
		return WrapperInternalServerError(c, err)
	}
	user := &models.Users{
		Repo:     repo,
		AuthType: PasswordAuthType,
	}
	err = c.BodyParser(&user)
	if err != nil {
		return WrapperBadRequestError(c, err)
	}
	err = user.ValidateLogin()
	if err != nil {
		return WrapperBadRequestError(c, fmt.Errorf("user or password invalid"))
	}
	token, err := services.Login(user, utils.PasswordAuth{})
	if err != nil {
		return WrapperForbiddenError(c, err)
	}
	response := make(map[string]interface{})
	response["token"] = token
	return WrapperSuccess(c, response, "")
}

// SignUp
//
//	@SignUp			godoc
//	@Description	Signup a new user
//	@Tags			Login
//	@Produce		json
//	@Param			SignUp	payload	body	string	true	"Fill the payload with the information"	SchemaExample({ "first_name":"","last_name":"","email":"", "password":""})
//	@Router			/v1/signup [post]
func SignUp(c *fiber.Ctx) error {
	repo, err := userRepo()
	if err != nil {
		return WrapperInternalServerError(c, err)
	}
	user := &models.Users{
		Repo:     repo,
		AuthType: PasswordAuthType,
	}
	err = c.BodyParser(&user)
	if err != nil {
		return WrapperBadRequestError(c, err)
	}
	err = user.Validate()
	if err != nil {
		return WrapperBadRequestError(c, err)
	}
	token, err := services.UserSignUp(user, utils.PasswordAuth{})
	if err != nil {
		return WrapperInternalServerError(c, err)
	}
	response := make(map[string]interface{})
	response["token"] = token
	return WrapperSuccess(c, response, "")
}

func userRepo() (models.UsersRepo, error) {
	repo, err := db.Db()
	if err != nil {
		return nil, err
	}
	return db.NewUsersRepo(repo), nil
}
