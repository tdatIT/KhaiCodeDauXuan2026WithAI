package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v3"
)

func main() {
	fmt.Println("Happy new year! - 2026")

	app := fiber.New()
	// Simple GET handler (Fiber accepts both func(fiber.Ctx) and func(fiber.Ctx) error)
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Happy new year!")
	})

	app.Get("/fortune", func(c fiber.Ctx) error {
		return c.SendString(getRandomFortune())
	})

	app.Listen(":3000")
}

func getRandomFortune() string {
	fortunes := []string{
		"Chúc mừng năm mới! Năm mới vạn sự như ý, tỷ sự như mơ!",
		"Năm mới phát tài phát lộc, an khang thịnh vượng!",
		"Chúc năm mới sức khỏe dồi dào, tiền vào như nước!",
		"Tân niên hạnh phúc, vạn sự cát tường!",
		"Năm mới thăng quan tiến chức, buôn may bán đắt!",
		"Chúc một năm mới bình an, mọi điều tốt đẹp!",
		"Năm mới tấn tài tấn lộc, sức khỏe bình an!",
		"Chúc xuân mới muôn điều may mắn, gia đình hạnh phúc!",
		"Năm mới làm ăn thịnh vượng, tiền vô như nước!",
		"Chúc năm mới công thành danh toại, mã đáo thành công!",
		"Năm mới rồng bay phượng múa, phú quý vinh hoa!",
		"Chúc tết đến xuân về, hoa nở bốn mùa, phúc lộc đầy nhà!",
		"Năm mới an khang, gia đạo hưng long!",
		"Chúc năm mới vui tươi, hạnh phúc tràn đầy!",
		"Năm mới cầu gì được nấy, ước gì cũng thành!",
	}

	rand.Seed(time.Now().UnixNano())
	return fortunes[rand.Intn(len(fortunes))]
}
