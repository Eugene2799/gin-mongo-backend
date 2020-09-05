package models

import (
	"gin-mongo-backend/utils"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoDB *Model

// ErrNotExist 数据不存在
var ErrNotExist = errors.New("not_exist")

type Model struct {
	client  *mongo.Client
	db      *mongo.Database
	Configs *ConfigsModel
}

// GetCtx 获取并发上下文(默认10秒超时)
func GetCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 15*time.Second)
}

// createIndexes 检查并创建索引
func createIndexes(ctx context.Context, name string, indexes []bson.M) error {
	collection := MongoDB.db.Collection(name)
	collectionIndexes := collection.Indexes()
	cur, err := collectionIndexes.List(ctx)
	if err != nil { // 读取索引发生错误
		return err
	}
	if cur == nil { // 指针不存在
		return errors.New("can't read collection")
	}
	if !cur.Next(ctx) { // 索引不存在，创建索引
		log.Info().Msg("Init index for " + name)
		for i := range indexes { // 创建唯一索引
			if _, err := collectionIndexes.CreateOne(ctx, mongo.IndexModel{
				Keys:    indexes[i],
				Options: options.Index().SetUnique(false),
			}); err != nil {
				return err
			}
		}
	}
	return cur.Close(ctx) // 关闭指针
}

// initCollection 初始化集合
func initCollection() error {
	ctx, cancel := GetCtx()
	defer cancel()
	// 初始化索引
	DBIndexes := []struct {
		name    string
		indexes []bson.M
	}{
		{name: "configs", indexes: []bson.M{{"full_name": 1}}},
	}
	for _, i := range DBIndexes {
		if err := createIndexes(ctx, i.name, i.indexes); err != nil {
			return err
		}
	}
	return nil
}

//InitDB 初始化数据库
func InitDB(config *utils.DBConfig) error {
	MongoDB = &Model{}
	err := connectDB(config)
	if err != nil {
		return err
	}
	MongoDB.db = MongoDB.client.Database(config.DBName)

	// 初始化集合
	if err := initCollection(); err != nil {
		return err
	}

	// 初始化 Model
	// 初始化数据库
	MongoDB.Configs = &ConfigsModel{
		Collection: MongoDB.db.Collection("configs"),
	}

	return nil
}

// Connect to MongoDB
func connectDB(config *utils.DBConfig) error {
	ctx, cancel := GetCtx()
	defer cancel()

	uri := os.Getenv("MONGODB_URL")

	if len(uri) == 0 {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", config.User, config.Password, config.Host, config.Port, config.DBName)
	}
	var err error
	if MongoDB.client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri)); err != nil {
		return err
	}
	// 测试连接
	if err := MongoDB.client.Ping(ctx, readpref.Primary()); err != nil {
		log.Error().Err(err).Msg("Failure to connect MongoDB!!!")
		return err
	}

	log.Info().Msg("Successful connection to MongoDB.")
	return nil
}

// Disconnect to MongoDB
func DisconnectDB() error {
	if MongoDB == nil {
		return nil
	}
	ctx, cancel := GetCtx()
	defer cancel()
	err := MongoDB.client.Disconnect(ctx)
	MongoDB = nil
	return err
}
