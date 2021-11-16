package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vodolaz095/gin-cache"
	"github.com/vodolaz095/gin-cache/memory"
	"github.com/vodolaz095/ticker/config"
	"github.com/vodolaz095/ticker/models"
	"github.com/vodolaz095/ticker/public"
)

func main() {
	var err error
	err = models.Connect(config.Token)
	if err != nil {
		log.Fatalf("%s : while connecting to Tinkoff Investement API", err)
	}
	r := gin.Default()
	memoryCache := memory.New(config.CacheTTL)
	r.StaticFS("/public", public.AssetFile())
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.FileFromFS("favicon.ico", public.AssetFile())
	})
	r.GET("/robots.txt", func(c *gin.Context) {
		c.FileFromFS("robots.txt", public.AssetFile())
	})
	r.GET("/", func(c *gin.Context) {
		c.FileFromFS("main.html", public.AssetFile()) // https://github.com/gin-gonic/gin/issues/2654
	})
	r.Use(gincache.New(memoryCache, func(c *gin.Context) (key string, ttl time.Duration, err error) {
		return "ticker", config.CacheTTL, nil
	}))
	r.GET("/portfolio.json", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
		defer cancel()
		positions, err1 := models.GetPositions(ctx, config.AccountID)
		if err1 != nil {
			panic(err1)
		}
		currencies, err1 := models.GetCurrencies(ctx, config.AccountID)
		if err1 != nil {
			panic(err1)
		}
		now := time.Now()
		c.JSON(http.StatusOK, gin.H{
			"now":        now.Format(time.Stamp),
			"timestamp":  now.Unix(),
			"positions":  positions,
			"currencies": currencies,
		})
	})
	err = r.Run(fmt.Sprintf("%s:%s", config.Address, config.Port))
	if err != nil {
		log.Fatalf("%s : while starting app on %s:%s", err, config.Address, config.Port)
	}
}
