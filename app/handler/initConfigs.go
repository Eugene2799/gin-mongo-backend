package handler

import (
	"fmt"
	"gin-mongo-backend/models"
	"gin-mongo-backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 读取配置
var config utils.Config

// GetAppConfigs 获取app初始化数据
func GetAppConfigs(c *gin.Context) {
	ctx, over := models.GetCtx()
	defer over()

	fullName := c.PostForm("fullName")
	configs := models.MongoDB.Configs
	var configsSchema models.ConfigsSchema

	// 查找原来的文档
	query := bson.M{"full_name": fullName}
	dbError := configs.Collection.FindOne(ctx, query).Decode(&configsSchema)
	if dbError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    dbError.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   configsSchema,
	})
}

// UpdateAppConfigs 更新app初始化数据
func UpdateAppConfigs(c *gin.Context) {
	ctx, over := models.GetCtx()
	defer over()

	configs := models.MongoDB.Configs
	var configsSchema models.ConfigsSchema

	err := c.BindJSON(&configsSchema)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"full_name", c.Param("full_name")}}
	data := bson.D{
		{"force_update", configsSchema.ForceUpdate},
		{"force_route", configsSchema.ForceRoute},
		{"title", configsSchema.Title},
		{"message", configsSchema.Message},
		{"version", configsSchema.Version},
		{"user_id", configsSchema.UserID},
		{"created_at", time.Now().Unix()},
		{"users_url", config.APIs.UsersAPI},
		{"tools_url", config.APIs.ToolsAPI},
	}
	update := bson.D{{"$set", data}}

	// 更新
	var updatedDocument bson.M
	dbError := configs.Collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedDocument)
	if dbError != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if dbError == mongo.ErrNoDocuments {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"msg":    "no documents",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"msg":    dbError.Error(),
			})
		}
		return
	}
	fmt.Printf("updated document %v", updatedDocument)

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
	})
}
