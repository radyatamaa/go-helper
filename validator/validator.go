package validator

import (
	"go-helper/converter"
	"go-helper/models"
	"strconv"
)

func ValidationQueryParamsDataTypeNumberFloat(request []models.DataTypeNumberFloatValidation) (bool, string) {

	for _, v := range request {
		if _, err := strconv.ParseFloat(v.Value, 64); err != nil {
			return false, v.Key + " " + models.ErrInvalidDataType.Error()
		}
	}
	return true, ""
}
func ValidationQueryParamsDataTypeNumberInt(request []models.DataTypeNumberIntValidation) (bool, string) {

	for _, v := range request {
		if _, err := strconv.Atoi(v.Value); err != nil {
			return false, v.Key + " " + models.ErrInvalidDataType.Error()
		}
	}
	return true, ""
}

func ValidationQueryParamsRequired(request []models.RequiredValidation) (bool, string) {

	for _, v := range request {
		if v.Value == "" {
			return false, v.Key + " " + models.ErrIsRequired.Error()
		}
	}
	return true, ""
}

func ValidationQueryParamsMaxMinLonglat(request []models.MaxMinLonglatValidation) (bool, string) {

	for _, v := range request {
		if v.Key == "latitude_now" {
			lat := converter.StringToFloat(v.Value)
			if lat < -90 {
				return false, v.Key + " " + models.ErrInvalidValue.Error()
			}

			if lat > 90 {
				return false, v.Key + " " + models.ErrInvalidValue.Error()
			}
		}

		if v.Key == "longitude_now" {
			lat := converter.StringToFloat(v.Value)
			if lat < -180 {
				return false, v.Key + " " + models.ErrInvalidValue.Error()
			}

			if lat > 180 {
				return false, v.Key + " " + models.ErrInvalidValue.Error()
			}
		}
	}

	return true, ""
}
func ValidationQueryMaxMinNumber(validation []models.MaxMinNumberValidation) (bool, string) {

	for _, v := range validation {
		val := converter.StringToFloat(v.Value)
		if v.ValueMinNumber == -1 {
			continue
		}
		if val < v.ValueMinNumber {
			return false, v.Key + " " + models.ErrInvalidValue.Error() + " min " + converter.FloatToString(v.ValueMinNumber)
		}

		if val > v.ValueMaxNumber {
			return false, v.Key + " " + models.ErrInvalidValue.Error() + " max " + converter.FloatToString(v.ValueMaxNumber)
		}
	}

	return true, ""
}

func ValidationQueryParamsPageLimit(request []models.PageLimitValidation) (bool, string) {

	for _, v := range request {
		if v.Key == "page" {
			p := converter.StringToInt(v.Value)
			if p < 1 {
				return false, v.Key + " " + models.ErrInvalidValue.Error()
			}
		}

		if v.Key == "limit" {
			l := converter.StringToInt(v.Value)
			if l < 1 {
				return false, v.Key + " " + models.ErrInvalidValue.Error()
			}

			if l > 100 {
				return false, v.Key + " " + models.ErrInvalidValue.Error()
			}
		}
	}

	return true, ""
}

func GlobalValidationQueryParams(request models.GlobalValidation) (bool, string) {
	checkQueryparams, message := ValidationQueryParamsRequired(request.RequiredValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryParamsDataTypeNumberInt(request.DataTypeNumberIntValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryParamsDataTypeNumberFloat(request.DataTypeNumberFloatValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryParamsMaxMinLonglat(request.MaxMinLonglatValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryParamsPageLimit(request.PageLimitValidation)
	if checkQueryparams == false {
		return false, message
	}

	checkQueryparams, message = ValidationQueryMaxMinNumber(request.MaxMinNumberValidation)
	if checkQueryparams == false {
		return false, message
	}

	return true, ""
}
