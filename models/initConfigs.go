package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ConfigsModel 初始化数据库
type ConfigsModel struct {
	Collection *mongo.Collection
}

// ConfigsSchema 基本数据结构
type ConfigsSchema struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`         // ID [索引]
	FullName    string             `bson:"full_name"`             // 软件全称
	ForceUpdate bool               `bson:"force_update"`          // 强制更新？
	ForceRoute  string             `bson:"force_route,omitempty"` // 强制跳转
	UpdateUrl   string             `bson:"update_url"`            // 更新地址
	Title       string             `bson:"title"`                 // 信息标题
	Message     string             `bson:"message"`               // 版本信息
	Version     string             `bson:"version"`               // 版本号
	UserID      primitive.ObjectID `bson:"user_id"`               // 用户ID
	UsersUrl    string             `bson:"users_url"`             // 用户相关接口url
	ToolsUrl    string             `bson:"tools_url"`             // 工具接口url
	CreatedAt   int64              `bson:"created_at"`            // 创建时间
}
