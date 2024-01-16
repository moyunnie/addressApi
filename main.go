package main

import (
	"addressApi/areaMap"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressReq struct {
	Data string `json:"data" form:"data"`
}

func main() {
	r := gin.Default()
	r.POST("address", func(c *gin.Context) {
		var request AddressReq
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "请携带data参数 示例张三 13800138000 龙华区龙华街道1980科技文化产业园3栋308 身份证120113196808214821",
			})
			return
		}
		parse := areaMap.Smart(request.Data)
		c.JSON(200, gin.H{
			"code": 200,
			"data": parse,
			"msg":  "解析成功",
		})
		return
	})
	r.GET("ip", func(c *gin.Context) {
		clientIP := c.ClientIP()
		c.String(http.StatusOK, clientIP)
		return
	})
	r.Run(":8001")
}
