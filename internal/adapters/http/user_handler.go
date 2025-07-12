package http

import (
	"encoding/json"
	"github.com/aknEvrnky/notification-system/internal/application/core/domain"
	"github.com/gofiber/fiber/v2"
)

func (a *Adapter) CreateUserHandler(ctx *fiber.Ctx) error {
	var user domain.User
	if err := json.Unmarshal(ctx.Body(), &user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid payload")
	}
	if err := a.api.CreateUser(ctx.Context(), user); err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).Send([]byte("user created"))
}

func (a *Adapter) UpdateUserHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user domain.User
	if err := json.Unmarshal(ctx.Body(), &user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid payload")
	}
	user.Id = id
	if err := a.api.UpdateUser(ctx.Context(), user); err != nil {
		return err
	}
	return ctx.Send([]byte("user updated"))
}

func (a *Adapter) DeleteUserHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := a.api.DeleteUser(ctx.Context(), id); err != nil {
		return err
	}
	return ctx.Send([]byte("user deleted"))
}
