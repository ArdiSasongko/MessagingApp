package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kooroshh/fiber-boostrap/app/repository"
	"github.com/kooroshh/fiber-boostrap/pkg/response"
)

func GetHistory(ctx *fiber.Ctx) error {
	resp, err := repository.GetAllMessage(ctx.Context())
	if err != nil {
		errResponse := fmt.Errorf("failed get all message : %s", err)
		fmt.Println(errResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errResponse.Error(), nil)
	}

	return response.SendSuccessResponse(ctx, resp)
}
