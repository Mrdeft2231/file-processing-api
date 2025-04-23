package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h2non/filetype"
	"io"
)

// Анализируем сигнатуру файла для определения формата
func AnalysisFormat() fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Файл не найден или произошла ошибка при загрузке",
			})
		}
		src, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Ошибка открытия файла",
			})
		}
		defer src.Close()

		head := make([]byte, 261)
		_, err = src.Read(head)
		if err != nil && err != io.EOF {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "ошибка при чтении файла",
			})
		}
		king, _ := filetype.Match(head)
		if king == filetype.Unknown {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Файл не поддерживается, прикрепите другой файл",
			})
		} else {
			c.Locals("fileType", king.MIME.Value)
			c.Locals("fileExtension", king.Extension)
		}
		return c.Next()
	}
}
