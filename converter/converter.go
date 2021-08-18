package converter

import (
	"encoding/json"
	"strconv"
	"time"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	if input_num != 0{
		return strconv.FormatFloat(input_num, 'f', 0, 64)
	}else {
		return "0"
	}
}
func StringNullableToFloat(value *string) float64 {
	if value != nil{
		res,_:= strconv.ParseFloat(*value,64)
		return res
	}
	return 0
}
func StringToFloat(value string) float64 {
	if value != "" {
		res, _ := strconv.ParseFloat(value, 64)
		return res
	}
	return 0
}
func FloatNUllableToString(input_num *float64) string {
	// to convert a float number to a string
	if input_num != nil{
		return strconv.FormatFloat(*input_num, 'f', 0, 64)
	}else {
		return ""
	}
}
func FloatNUllableToFloat(value *float64) float64 {
	if value != nil{
		return *value
	}
	return 0
}
func DateTimeToDateTimeNUllable(value time.Time) *time.Time {
	return &value
}
func DateTimeNullableToDateTime(value *time.Time) time.Time {
	if value == nil{
		return time.Time{}
	}
	return *value
}
func IntToIntNullable(value int) *int {
	return &value
}
func IntNullableToInt(value *int) int {
	if value == nil{
		return 0
	}
	return *value
}
func StringToStringNullable(value string) *string {
	return &value
}
func ObjectToString(value interface{}) string {
	result ,_:= json.Marshal(value)
	return string(result)
}
func StringNullableToString(value *string) string {
	if value != nil{
		return *value
	}
	return ""
}
func IntNullableToStringNullable(value *int) *string {

	if value != nil{
		result := strconv.Itoa(*value)
		return &result
	}
	return nil
}
func IntNullableToString(value *int) string {

	if value != nil{
		result := strconv.Itoa(*value)
		return result
	}
	return "0"
}
func StringToIntNullable(value string) *int {

	if value != ""{
		result ,_:= strconv.Atoi(value)
		return &result
	}
	return nil
}
func StringToInt(value string) int {

	if value != ""{
		result ,_:= strconv.Atoi(value)
		return result
	}
	return 0
}
func StringNullableToInt(value *string) int {

	if value != nil{
		result ,_:= strconv.Atoi(*value)
		return result
	}
	return 0
}
func StringNullableToDateTimeNullable(value *string) *time.Time {
	if value != nil{
		var layoutFormat string
		var date time.Time

		layoutFormat = "2006-01-02 15:04:05"
		date, _ = time.Parse(layoutFormat, *value)
		return &date
	}

	return nil
}
func DateTimeNullableToStringNullable(value *time.Time) *string {
	if value != nil{
		layoutFormat := "2006-01-02 15:04:05"
		date := value.Format(layoutFormat)
		return &date
	}

	return nil
}
func DateTimeNullableToStringNullableWithFormat(value *time.Time,format string) *string {
	if value != nil{
		layoutFormat := format
		date := value.Format(layoutFormat)
		return &date
	}

	return nil
}
func StringNullableToStringDefaultFormatDate(value *string) *string {
	if value != nil{
		var layoutFormat string
		var date time.Time

		layoutFormat = "2006-01-02T15:04:05Z"
		date, _ = time.Parse(layoutFormat, *value)
		dateString := date.Format("2006-01-02 15:04:05")
		return &dateString
	}

	return nil
}
func StringToDateTimeNullable(value string) *time.Time {
	if value != ""{
		var layoutFormat string
		var date time.Time
		layoutFormat = "2006-01-02T15:04:05.000Z"
		date, _ = time.Parse(layoutFormat, value)
		return &date
	}

	return nil
}
func StringToDate(value string) time.Time {
	if value != ""{
		var layoutFormat string
		var date time.Time

		layoutFormat = "2006-01-02"
		date, _ = time.Parse(layoutFormat, value)
		return date
	}

	return time.Time{}
}
func StringNullableToDateNullable(value *string) *string {
	if value != nil{
		var layoutFormat string
		var date time.Time

		layoutFormat = "20060102"
		date, _ = time.Parse(layoutFormat, *value)
		dateString := date.Format("20060102")
		return &dateString
	}

	return nil
}
func IntNullableToBool(value *int) bool  {
	if value != nil{
		if *value == 1{
			return true
		}
	}
	return false
}
func MetersToMils(meters float64) float64 {
	return meters * 0.000621371192
}
func MilsToMeters(mils float64) float64 {
	return mils * 1609.344
}

func MetersToKilometers(meters float64) float64 {
	return meters / 1000
}
func NowYmd() string {
	t := time.Now()
	timeFormated := t.Format("2006-01-02 15:04:05")
	return timeFormated
}

func NowAddDay() string {
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = NowYmd()
	date, _ = time.Parse(layoutFormat, value)

	return date.AddDate(0, 0, 1).Format(layoutFormat)
}

