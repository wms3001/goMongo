package goMongo

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"strconv"
)

type GoMongo struct {
	Addr     string
	Port     string
	Password string
	User     string
	Pass     string
	Db       string
	Client   *mongo.Client
	Collect  *mongo.Collection
}

func (goMongo *GoMongo) Connect() *Resp {
	var resp = &Resp{}
	clientOptions := options.Client().ApplyURI("mongodb://" + goMongo.User + ":" + goMongo.Pass + "@" + goMongo.Addr +
		":" + goMongo.Port + "/" + goMongo.Db)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "connected"
		goMongo.Client = client
	}
	return resp
}

func (goMongo *GoMongo) Ping() *Resp {
	var resp = &Resp{}
	err := goMongo.Client.Ping(context.TODO(), nil)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "connected"
	}
	return resp
}

func (goMongo *GoMongo) SetCollection(coll string) *Resp {
	var resp = &Resp{}
	collection := goMongo.Client.Database(goMongo.Db).Collection(coll)
	if collection != nil {
		resp.Code = 1
		resp.Message = "success"
		goMongo.Collect = collection
	} else {
		resp.Code = -1
		resp.Message = "connect error"
	}
	return resp
}

func (goMongo *GoMongo) CloseConn() *Resp {
	var resp = &Resp{}
	resp.Code = 1
	resp.Message = "关闭成功"
	err := goMongo.Client.Disconnect(context.TODO())
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	}
	return resp
}

func (goMongo *GoMongo) AddOne(data interface{}) *Resp {
	var resp = &Resp{}
	re, err := goMongo.Collect.InsertOne(context.TODO(), data)
	if err != nil {
		resp.Code = -1
		resp.Message = "添加失败"
	} else {
		resp.Code = 1
		resp.Message = "添加成功"
		resp.Data = Strval(re.InsertedID)
	}
	return resp
}

func (goMongo *GoMongo) AddMany(data []interface{}) *Resp {
	var resp = &Resp{}
	re, err := goMongo.Collect.InsertMany(context.TODO(), data)
	if err != nil {
		resp.Code = -1
		resp.Message = "添加失败"
	} else {
		resp.Code = 1
		resp.Message = "添加成功"
		resp.Data = re.InsertedIDs
	}
	return resp
}

func (goMongo *GoMongo) GetOne(data *Filter) *Resp {
	var resp = &Resp{}
	var result bson.M
	condition := param(data)
	err := goMongo.Collect.FindOne(context.TODO(), condition).Decode(&result)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "查询成功"
		resp.Data = result
	}
	return resp
}

func (goMongo *GoMongo) GetMany(data *Filter, option *Option) *Resp {
	var resp = &Resp{}
	var result []bson.M
	condition := param(data)
	findOptions := options.Find()
	findOptions.SetLimit(option.Limit)
	findOptions.SetSkip(option.Skip)
	findOptions.SetSort(option.SMap)
	fmt.Println(data)
	cur, err := goMongo.Collect.Find(context.TODO(), condition, findOptions)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		for cur.Next(context.TODO()) {
			var item bson.M
			cur.Decode(&item)
			result = append(result, item)
		}
		resp.Code = 1
		resp.Message = "查询成功"
		resp.Data = result
	}
	return resp
}

func (goMongo *GoMongo) UpdateOne(data *Filter, uData *UData) *Resp {
	var resp = &Resp{}
	condition := param(data)
	updata := paramData(uData)
	up, err := goMongo.Collect.UpdateOne(context.TODO(), condition, updata)
	fmt.Println(up)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "更新成功"
		resp.Data = up
	}
	return resp
}

func (goMongo *GoMongo) UpdateMany(data *Filter, uData *UData) *Resp {
	var resp = &Resp{}
	condition := param(data)
	updata := paramData(uData)
	up, err := goMongo.Collect.UpdateMany(context.TODO(), condition, updata)
	fmt.Println(up)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "更新成功"
		resp.Data = up
	}
	return resp
}

func (goMongo *GoMongo) DeleteOne(data *Filter) *Resp {
	var resp = &Resp{}
	condition := param(data)
	del, err := goMongo.Collect.DeleteOne(context.TODO(), condition)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "删除成功"
		resp.Data = del.DeletedCount
	}
	return resp
}

func (goMongo *GoMongo) DeleteMany(data *Filter) *Resp {
	var resp = &Resp{}
	condition := param(data)
	del, err := goMongo.Collect.DeleteMany(context.TODO(), condition)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "删除成功"
		resp.Data = del.DeletedCount
	}
	return resp
}

func (goMongo *GoMongo) Count(data *Filter) *Resp {
	var resp = &Resp{}
	condition := param(data)
	count, err := goMongo.Collect.CountDocuments(context.TODO(), condition)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "统计成功"
		resp.Data = count
	}
	return resp
}

func (goMongo *GoMongo) Distinct(data *Filter, field string) *Resp {
	var resp = &Resp{}
	condition := param(data)
	dis, err := goMongo.Collect.Distinct(context.TODO(), field, condition)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "查询成功"
		resp.Data = dis
	}
	return resp

}

func param(condition *Filter) map[string]interface{} {
	var filter map[string]interface{}
	filter = make(map[string]interface{})
	switch condition.Type {
	case "and":
		filter = condition.FMap
	case "or":
		var fil []bson.M = []bson.M{}
		for index, item := range condition.FMap {
			fil = append(fil, bson.M{index: item})
		}
		filter["$or"] = fil
	default:

	}
	return filter
}

func paramData(uData *UData) map[string]interface{} {
	var updata map[string]interface{}
	updata = make(map[string]interface{})
	updata["$set"] = uData.UMap
	return updata
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
