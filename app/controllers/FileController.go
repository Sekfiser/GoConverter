package controllers

import (
	"Gonverter/app/service/RabbitMQ"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
	"time"
)

func SaveFileFromBytes(c *fiber.Ctx) error {
	fullFileName := c.Query("FileName")
	fileExt := filepath.Ext(fullFileName)
	fileName := fullFileName[:len(fullFileName)-len(fileExt)]

	newFileName := GenerateFileName(fileName, fileExt)

	path, err := os.Getwd()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Не найден путь"})
	}

	filePath := path + "/FilesToConvert/" + newFileName

	err = os.WriteFile(filePath, c.Body(), 0666)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Ошибка сохранения файла файла"})
	}

	RabbitMQ.MessageToFileQueue(newFileName)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Файл сохранен", "path": "/FilesToConvert/" + newFileName})
}

func SaveFile(c *fiber.Ctx) error {
	file, err := c.FormFile("FileToConvert")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Ошибка чтения файла"})
	}

	fullFileName := file.Filename
	fileExt := filepath.Ext(fullFileName)
	fileName := fullFileName[:len(fullFileName)-len(fileExt)]

	newFileName := GenerateFileName(fileName, fileExt)

	err = c.SaveFile(file, fmt.Sprintf("./%s", "FilesToConvert/"+newFileName))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Ошибка сохранения файла файла"})
	}

	RabbitMQ.MessageToFileQueue(newFileName)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Файл сохранен", "path": "/FilesToConvert/" + newFileName})
}

func SaveFileFromString(c *fiber.Ctx) error {
	fullFileName := c.FormValue("FileName")

	fileExt := filepath.Ext(fullFileName)
	fileName := fullFileName[:len(fullFileName)-len(fileExt)]

	newFileName := GenerateFileName(fileName, fileExt)

	fileString := c.FormValue("FileByteString")

	path, err := os.Getwd()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Ошибка чтения файла"})
	}

	err = os.WriteFile(path+"/FilesToConvert/"+newFileName, []byte(fileString), 0666)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Ошибка сохранения файла файла"})
	}

	RabbitMQ.MessageToFileQueue(newFileName)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Файл сохранен", "path": "/FilesToConvert/" + newFileName})
}

func GenerateFileName(fileName string, fileExt string) string {
	stringToEncode := fileName + time.Now().Format(time.StampNano) + "SaltForFile"
	hash := md5.Sum([]byte(stringToEncode))

	return hex.EncodeToString(hash[:]) + fileExt
}
