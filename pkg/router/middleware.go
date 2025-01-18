package router

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kooroshh/fiber-boostrap/app/repository"
	jwttoken "github.com/kooroshh/fiber-boostrap/pkg/jwt_token"
	"github.com/kooroshh/fiber-boostrap/pkg/response"
)

func MiddlewareAuth(ctx *fiber.Ctx) error {
	auth := ctx.Get("authorization")
	if auth == "" {
		errResponse := fmt.Errorf("authorized empty")
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusUnauthorized, errResponse.Error(), nil)
	}

	_, err := repository.GetUserSessionByToken(ctx.Context(), auth)
	if err != nil {
		errResponse := fmt.Errorf("failed get user session by token : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusUnauthorized, errResponse.Error(), nil)
	}

	claims, err := jwttoken.ValidateToken(ctx.Context(), auth)
	if err != nil {
		errResponse := fmt.Errorf("failed validate token : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusUnauthorized, errResponse.Error(), nil)
	}

	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		errResponse := fmt.Errorf("token expired time : %v", claims.ExpiresAt)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusUnauthorized, errResponse.Error(), nil)
	}

	ctx.Set("username", claims.Username)
	ctx.Set("fullname", claims.Fullname)
	return ctx.Next()
}

func MiddlewareRefreshToken(ctx *fiber.Ctx) error {
	auth := ctx.Get("authorization")
	if auth == "" {
		errResponse := fmt.Errorf("authorized empty")
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusUnauthorized, errResponse.Error(), nil)
	}

	claims, err := jwttoken.ValidateToken(ctx.Context(), auth)
	if err != nil {
		errResponse := fmt.Errorf("failed validate token : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusUnauthorized, errResponse.Error(), nil)
	}

	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		errResponse := fmt.Errorf("token expired time : %v", claims.ExpiresAt)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusUnauthorized, errResponse.Error(), nil)
	}

	ctx.Locals("username", claims.Username)
	ctx.Locals("fullname", claims.Fullname)
	return ctx.Next()
}
