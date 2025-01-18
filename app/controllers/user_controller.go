package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kooroshh/fiber-boostrap/app/models"
	"github.com/kooroshh/fiber-boostrap/app/repository"
	jwttoken "github.com/kooroshh/fiber-boostrap/pkg/jwt_token"
	"github.com/kooroshh/fiber-boostrap/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	err := ctx.BodyParser(user)
	if err != nil {
		errResponse := fmt.Errorf("failed to parse request : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	err = user.Validate()
	if err != nil {
		errResponse := fmt.Errorf("failed to validate request : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		errResponse := fmt.Errorf("failed to hash password : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}
	user.Password = string(hash)

	err = repository.InsertUser(ctx.Context(), user)
	if err != nil {
		errResponse := fmt.Errorf("failed to parse request : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, errResponse.Error(), nil)
	}

	resp := user
	resp.Password = ""
	return response.SendSuccessResponse(ctx, resp)
}

func Login(ctx *fiber.Ctx) error {
	loginReq := new(models.LoginRequest)

	err := ctx.BodyParser(loginReq)
	if err != nil {
		errResponse := fmt.Errorf("failed to parse request : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	err = loginReq.Validate()
	if err != nil {
		errResponse := fmt.Errorf("failed to validate request : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	user, err := repository.GetUser(ctx.Context(), loginReq.Username)
	if err != nil {
		errResponse := fmt.Errorf("invalid credentials : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusNotFound, errResponse.Error(), nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		errResponse := fmt.Errorf("invalid credentials : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	token, err := jwttoken.GeneratedToken(ctx.Context(), user.Username, user.FullName, "token")
	if err != nil {
		errResponse := fmt.Errorf("failed generated token : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	refreshToken, err := jwttoken.GeneratedToken(ctx.Context(), user.Username, user.FullName, "refresh_token")
	if err != nil {
		errResponse := fmt.Errorf("failed generated token : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	session := models.UserSession{
		UserID:              int(user.ID),
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        time.Now().Add(jwttoken.MapToken["token"]),
		RefreshTokenExpired: time.Now().Add(jwttoken.MapToken["refresh_token"]),
	}

	err = repository.InsertSession(ctx.Context(), &session)
	if err != nil {
		errResponse := fmt.Errorf("failed inserrt token : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	resp := models.LoginResponse{
		Username:     user.Username,
		Fullname:     user.FullName,
		Token:        token,
		RefreshToken: refreshToken,
	}
	return response.SendSuccessResponse(ctx, resp)
}

func Logout(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	err := repository.DeleteSession(ctx.Context(), token)
	if err != nil {
		errResponse := fmt.Errorf("failed delete token : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	return response.SendSuccessResponse(ctx, nil)
}

func RefreshToken(ctx *fiber.Ctx) error {
	refreshToken := ctx.Get("Authorization")
	username := ctx.Locals("username").(string)
	fullName := ctx.Locals("fullname").(string)

	token, err := jwttoken.GeneratedToken(ctx.Context(), username, fullName, "token")
	if err != nil {
		errResponse := fmt.Errorf("failed generated token : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	err = repository.UpdateSession(ctx.Context(), token, refreshToken, time.Now().Add(jwttoken.MapToken["token"]))
	if err != nil {
		errResponse := fmt.Errorf("failed update token : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	return response.SendSuccessResponse(ctx, fiber.Map{
		"token": token,
	})
}
