package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/need/go-backend/utils/jwt"
)

func IsAdmin(c *fiber.Ctx) error {
	return checkRole(c, "admin")
}

func IsUser(c *fiber.Ctx) error {
	return checkRole(c, "user")
}

func IsMod(c *fiber.Ctx) error {
	return checkRole(c, "mod")
}

func checkRole(c *fiber.Ctx, requiredRole string) error {
	// ดึง token จาก header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Missing or invalid token"})
	}

	// แยก Bearer ออกจาก token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid token format"})
	}

	// ตรวจสอบและแยกข้อมูลจาก token
	claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid token", "data": err.Error()})
	}

	// ดึง roles จาก claims
	roles := claims.Roles

	// ตรวจสอบว่าผู้ใช้มีบทบาทที่ต้องการหรือไม่
	for _, role := range roles {
		if role == requiredRole {
			return c.Next()
		}
	}

	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "error", "message": "Forbidden: insufficient permissions"})
}
