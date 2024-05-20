package helper

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func IsValidItNIP(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		var nipRegex = regexp.MustCompile(`^615[12](200[0-9]|20[1-2][0-9]|202[0-4])(0[1-9]|1[0-2])([0-9]{3,5})$`)
		return nipRegex.MatchString(value)
	}
	return false
}

func IsValidNurseNIP(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		var nipRegex = regexp.MustCompile(`^303[12](200[0-9]|20[1-2][0-9]|202[0-4])(0[1-9]|1[0-2])([0-9]{3,5})$`)
		return nipRegex.MatchString(value)
	}
	return false
}

func IsItNiP(nip string) bool {
	var nipRegex = regexp.MustCompile(`^615`)
	return nipRegex.MatchString(nip)
}

func IsNurseNiP(nip string) bool {
	var nipRegex = regexp.MustCompile(`^303`)
	return nipRegex.MatchString(nip)
}

func IsValidUrl(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		var urlRegex = regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(/[a-zA-Z0-9-._~:?#@!$&'()*+,;=]*)*$`)
		return urlRegex.MatchString(value)
	}
	return false
}

func IsGender(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value == "male" {
			return true
		}
		if value == "female" {
			return true
		}
	}
	return false
}

func IdentityNumber(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(int64)
	if ok {
		strValue := strconv.FormatInt(value, 10)
		return len(strValue) == 16
	}
	return false

}

func CustomMessageValidation(valErr validator.FieldError) string {
	switch valErr.ActualTag() {
	case "nip_it":
		return "nip it invalid"
	case "nip_nurse":
		return "nip nurse invalid"
	case "required":
		return fmt.Sprintf("%s required", valErr.Field())
	case "min":
		minParam := valErr.Param()
		return fmt.Sprintf("%s cannot be less than %s", valErr.Field(), minParam)
	case "max":
		maxParam := valErr.Param()
		return fmt.Sprintf("%s can't be more than %s", valErr.Field(), maxParam)
	case "category":
		return "The product does not match the existing type"
	case "url_image":
		return "image not url"
	case "numeric":
		return "numeric required"
	case "gender":
		return "gender invalid"
	case "identity_number":
		return "identity number invalid required 16 digit"
	case "startswith":
		return "phone number required startswith +62"
	case "datetime":
		return "datetime invalid value format example 2006-01-02"
	default:
		return "Tag not implement custom message error"
	}
}
