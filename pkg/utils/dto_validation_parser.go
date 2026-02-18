package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var fieldNameDict = map[string]string {
	"Name":     		"Имя",
	"Sign":     		"Обозначение",
	"Code":     		"Код",
	"Login": 			"Логин",
	"Password": 		"Пароль",
	"Surname":        	"Фамилия",
	"Patronymic":     	"Отчество",
	"RoleID":         	"Должность",
	"RepeatPassword": 	"Повторение пароля",
}

func ValidationErrors(err error) string {
	errInfo := ""

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			tag := fieldErr.Tag()
			field := fieldErr.Field()
			if name, ok := fieldNameDict[field]; ok {
				field = name
			}

			var errMsg string
			errInfo += field + ": "

			switch tag {
			case "required":
				errMsg = "обязательное поле"
			case "email":
				errMsg = "некорректный формат"
			case "min":
				errMsg = fmt.Sprintf("минимальная длина %s символов", fieldErr.Param())
			case "max":
				errMsg = fmt.Sprintf("максимальная длина %s символов", fieldErr.Param())
			case "len":
				errMsg = fmt.Sprintf("длина должна быть %s символов", fieldErr.Param())
			case "eqfield":
				errMsg = "значения не совпадают"
			default:
				errMsg = "некорректное значение"
			}

			errInfo += errMsg + "; "
		}
	}

	if errInfo == "" {
		return "Некорректные данные."
	} else {
		return "Некорректные данные: " + errInfo[:len(errInfo) - 1]
	}
}