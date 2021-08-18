package automapper

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"strconv"
)

func Pagination(qpage, qperPage string) (limit, page, offset int) {
	limit = 20
	page = 1
	offset = 0

	page, _ = strconv.Atoi(qpage)
	limit, _ = strconv.Atoi(qperPage)
	if page == 0 && limit == 0 {
		page = 1
		limit = 10
	}
	offset = (page - 1) * limit

	return
}

func GetKeyJsonStruct(value interface{}) []string {

	j, _ := json.Marshal(value)
	// a map container to decode the JSON structure into
	c := make(map[string]json.RawMessage)

	// unmarschal JSON
	e := json.Unmarshal(j, &c)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, _ := range c {
		k[i] = s
		i++
	}

	return k
}

func GetValueStruct(value interface{}) []driver.Value {
	var result []driver.Value
	rv := reflect.ValueOf(value)
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)

		dv := driver.Value(fv.Interface())
		result = append(result, dv)
	}
	return result
}

func GetValueAndColumnStructToDriverValue(value interface{}) ([]driver.Value, []string) {
	var result []driver.Value

	//column
	j, _ := json.Marshal(value)
	// a map container to decode the JSON structure into
	c := make(map[string]json.RawMessage)

	// unmarschal JSON
	e := json.Unmarshal(j, &c)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, e := range c {
		k[i] = s
		v, _ := e.MarshalJSON()
		var val driver.Value
		err := json.Unmarshal(v, &val)
		if err != nil {
			panic(err)
		}

		//dv := driver.Value(fv.Interface())
		result = append(result, val)
		i++
	}

	return result, k

}


