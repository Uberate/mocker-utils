package gins

import "github.com/gin-gonic/gin"

func GinStart(engine *gin.Engine, webConfig WebConfig) error {
	gin.SetMode(webConfig.Mod)
	return engine.Run(webConfig.Addr...)
}
