package handler

import "github.com/gofiber/fiber/v2"

// Hello godoc
// @Summary Greet User
// @Description Say Hello
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

// Mello godoc
// @Summary Greet User
// @Description Say Mello
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /Mello [get]
func Mello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Mello i'm ok!", "data": nil})
}
