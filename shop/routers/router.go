package routers

import (
	"net/http"
	"os"
	v1 "shop/routers/api/v1"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "shop/docs"

	"shop/pkg/export"
	"shop/pkg/qrcode"
	"shop/pkg/upload"
	"shop/routers/api"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	// logger中间件将日志写到gin.DefaultWriter,即使GIN_MODE=release
	// 默认gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// recovery中间件从任何panic恢复，如果出现panic,它会写一个500错误
	r.Use(gin.Recovery())

	temPath := ""
	staticPath := ""
	if os.Getenv("mysql_addr") == "" {
		temPath = "D:/k8s/gocode/src/shop/templates/**/*"
		staticPath = "D:/k8s/gocode/src/shop/statics/assets"
	} else {
		temPath = "templates/**/*"
		staticPath = "statics/assets"
	}
	r.LoadHTMLGlob(temPath)
	r.Static("/statics", staticPath)

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	apiv1.Use()
	{
		apiv1.GET("/goods/:id", v1.GetGoods)
		apiv1.GET("/order/:id", v1.GetOrder)
		apiv1.Any("/order", v1.GetOrderAllList)
		apiv1.GET("/user/:userId/order", v1.GetOrderList)

		//获取标签列表
		//apiv1.GET("/tags", v1.GetTags)
		////新建标签
		//apiv1.POST("/tags", v1.AddTag)
		////更新指定标签
		//apiv1.PUT("/tags/:id", v1.EditTag)
		////删除指定标签
		//apiv1.DELETE("/tags/:id", v1.DeleteTag)
		////导出标签
		//r.POST("/tags/export", v1.ExportTag)
		////导入标签
		//r.POST("/tags/import", v1.ImportTag)
		//
		////获取文章列表
		//apiv1.GET("/articles", v1.GetArticles)
		////获取指定文章
		//apiv1.GET("/articles/:id", v1.GetArticle)
		////新建文章
		//apiv1.POST("/articles", v1.AddArticle)
		////更新指定文章
		//apiv1.PUT("/articles/:id", v1.EditArticle)
		////删除指定文章
		//apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		////生成文章海报
		//apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
