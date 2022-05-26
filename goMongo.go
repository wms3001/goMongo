package goMongo

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goMongo/model"
	"golang.org/x/net/context"
	"os"
	"strconv"
)

var mData *model.MData

//func  OpenConn(coll string) (*mongo.Client,*mongo.Collection) {
func OpenConn(coll string) *model.Conn {
	var conn = &model.Conn{}
	work, _ := os.Getwd()
	viper.SetConfigName("mongo")
	viper.SetConfigType("yml")
	viper.AddConfigPath(work + "/conf")
	viper.ReadInConfig()
	mHost := viper.GetString("mongo.host")
	mPort := viper.GetString("mongo.port")
	mUser := viper.GetString("mongo.user")
	mPass := viper.GetString("mongo.pass")
	mDb := viper.GetString("mongo.dbName")
	// 设置客户端选项
	clientOptions := options.Client().ApplyURI("mongodb://" + mUser + ":" + mPass + "@" + mHost + ":" + mPort + "/" + mDb)
	// 连接 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		conn.Code = -1
		conn.Message = err.Error()
		return conn
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		conn.Code = -1
		conn.Message = err.Error()
		return conn
	}
	collection := client.Database(mDb).Collection(coll)
	//o.Client = client
	conn.Code = 1
	conn.Message = "success"
	conn.Client = client
	conn.Collection = collection
	return conn
}

func SetCollection(client *mongo.Client, database string, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}

func CloseConn(client *mongo.Client) *model.Base {
	var base = &model.Base{}
	base.Code = 1
	base.Message = "关闭成功"
	err := client.Disconnect(context.TODO())
	if err != nil {
		base.Code = -1
		base.Message = err.Error()
	}
	return base
}

func AddOne(collection *mongo.Collection, data interface{}) *model.AddOne {
	var addOne = &model.AddOne{}
	re, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		addOne.Code = -1
		addOne.Message = "添加失败"
	} else {
		addOne.Code = 1
		addOne.Message = "添加成功"
		addOne.InsertId = Strval(re.InsertedID)
	}
	return addOne
}

func AddMany(collection *mongo.Collection, data []interface{}) *model.AddMany {
	var addMany = &model.AddMany{}
	re, err := collection.InsertMany(context.TODO(), data)
	if err != nil {
		addMany.Code = -1
		addMany.Message = "添加失败"
	} else {
		addMany.Code = 1
		addMany.Message = "添加成功"
		//var insertIds []string
		for _, d := range re.InsertedIDs {
			addMany.InsertIds = append(addMany.InsertIds, Strval(d))
		}
	}
	return addMany
}

func GetOneSingl(collection *mongo.Collection, data interface{}) *model.GetOne {
	var getOne = &model.GetOne{}
	var result bson.M
	err := collection.FindOne(context.TODO(), data).Decode(&result)
	if err != nil {
		getOne.Code = -1
		getOne.Message = err.Error()
	} else {
		getOne.Code = 1
		getOne.Message = "查询成功"
		getOne.Data = result
	}
	return getOne
}

func GetOne(collection *mongo.Collection, data *model.MData) *model.GetOne {
	var getOne = &model.GetOne{}
	var result bson.M
	filter, _, _, _, _ := param(data)
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		getOne.Code = -1
		getOne.Message = err.Error()
	} else {
		getOne.Code = 1
		getOne.Message = "查询成功"
		getOne.Data = result
	}
	return getOne
}

func GetMany(collection *mongo.Collection, data *model.MData) *model.GetMany {
	var getMany = &model.GetMany{}
	var result []bson.M
	filter, _, sort, limit, skip := param(data)
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(skip)
	findOptions.SetSort(sort)
	fmt.Println(data)
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		getMany.Code = -1
		getMany.Message = err.Error()
	} else {
		for cur.Next(context.TODO()) {
			var item bson.M
			cur.Decode(&item)
			result = append(result, item)
		}
		getMany.Code = 1
		getMany.Message = "查询成功"
		getMany.Data = result
	}
	return getMany
}

func UpdateOne(collection *mongo.Collection, data *model.MData) *model.UpdateOne {
	var updateOne = &model.UpdateOne{}
	filter, update, _, _, _ := param(data)
	up, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		updateOne.Code = -1
		updateOne.Message = err.Error()
	} else {
		updateOne.Code = 1
		updateOne.Message = "更新成功"
		updateOne.MatchedCount = up.MatchedCount
		updateOne.ModifiedCount = up.ModifiedCount
	}
	return updateOne
}

func UpdateMany(collection *mongo.Collection, data *model.MData) *model.UpdateOne {
	var updateOne = &model.UpdateOne{}
	filter, update, _, _, _ := param(data)
	up, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		updateOne.Code = -1
		updateOne.Message = err.Error()
	} else {
		updateOne.Code = 1
		updateOne.Message = "更新成功"
		updateOne.MatchedCount = up.MatchedCount
		updateOne.ModifiedCount = up.ModifiedCount
	}
	return updateOne
}

func DeleteOne(collection *mongo.Collection, data *model.MData) *model.DeleteOne {
	var deleteOne = &model.DeleteOne{}
	filter, _, _, _, _ := param(data)
	del, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		deleteOne.Code = -1
		deleteOne.Message = err.Error()
	} else {
		deleteOne.Code = 1
		deleteOne.Message = "删除成功"
		deleteOne.DeletedCount = del.DeletedCount
	}
	return deleteOne
}

func DeleteMany(collection *mongo.Collection, data *model.MData) *model.DeleteOne {
	var deleteMany = &model.DeleteOne{}
	filter, _, _, _, _ := param(data)
	del, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		deleteMany.Code = -1
		deleteMany.Message = err.Error()
	} else {
		deleteMany.Code = 1
		deleteMany.Message = "删除成功"
		deleteMany.DeletedCount = del.DeletedCount
	}
	return deleteMany
}

func Count(collection *mongo.Collection, data *model.MData) *model.MCount {
	var mCount = &model.MCount{}
	filter, _, _, _, _ := param(data)
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		mCount.Code = -1
		mCount.Message = err.Error()
	} else {
		mCount.Code = 1
		mCount.Message = "统计成功"
		mCount.DocNum = count
	}
	return mCount
}

func Distinct(collection *mongo.Collection, data *model.MData, field string) *model.MDistinct {
	var mDistinct = &model.MDistinct{}
	filter, _, _, _, _ := param(data)
	dis, err := collection.Distinct(context.TODO(), field, filter)
	if err != nil {
		mDistinct.Code = -1
		mDistinct.Message = err.Error()
	} else {
		mDistinct.Code = 1
		mDistinct.Message = "查询成功"
		mDistinct.Mfilds = dis
	}
	return mDistinct

}

func param(mData *model.MData) (map[string]interface{}, map[string]interface{}, map[string]interface{}, int64, int64) {
	var filter, update, sort map[string]interface{}
	filter = make(map[string]interface{})
	update = make(map[string]interface{})
	sort = make(map[string]interface{})
	switch mData.Type {
	case "and":
		filter = mData.FilterMap
	case "or":
		var fil []bson.M = []bson.M{}
		for index, item := range mData.FilterMap {
			fil = append(fil, bson.M{index: item})
		}
		filter["$or"] = fil
	default:

	}
	sort = mData.SortMap
	update["$set"] = mData.UpMap
	return filter, update, sort, mData.Limit, mData.Skip
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		tt := newValue[1:]
		tt1 := tt[:len(tt)-1]
		key = string(tt1)
	}
	return key
}
