package main

import (
	//"log"
	"time"

	//"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var (
	c *cache.Cache
)

func InitCache() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}

func GetFromCache(key string) (interface{}, bool) {
	return c.Get(key)
}

func SetToCache(key string, value interface{}) {
	c.Set(key, value, cache.DefaultExpiration)
}


