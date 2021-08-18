package cache

import (
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
)

var Cache = cache.New(24*time.Hour, 24*time.Hour)

type Emp []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func SetCache(key string, emp interface{},exp time.Duration) bool {
	Cache.Set(key, emp, exp)
	return true
}

func GetCacheToken(key string) (string, bool) {
	var emp *time.Time
	var found bool
	data, found := Cache.Get(key)
	if found {
		emp = data.(*time.Time)
		return emp.String(), found
	}

	return "", found
}
func GetCache(key string) (string, bool) {
	var emp int
	var found bool
	data, found := Cache.Get(key)
	if found {
		emp = data.(int)
	}

	return strconv.Itoa(emp) , found
}
