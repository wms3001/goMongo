package goMongo

import (
	"fmt"
	uuid2 "github.com/google/uuid"
	"testing"
)

func TestGoMongo_Connect(t *testing.T) {
	goMongo := GoMongo{}
	goMongo.Addr = "192.168.4.81"
	goMongo.Port = "27017"
	goMongo.User = "logistics"
	goMongo.Pass = "123456"
	goMongo.Db = "logistics"
	goMongo.Connect()
	goMongo.Ping()
	goMongo.SetCollection("order1")
	goMongo.CloseConn()
}

type OrderTest struct {
	Test   string `bson:"test"`
	Status int    `bson:"status"`
}

func TestGoMongo_AddOne(t *testing.T) {

	goMongo := GoMongo{}
	goMongo.Addr = "192.168.4.81"
	goMongo.Port = "27017"
	goMongo.User = "logistics"
	goMongo.Pass = "123456"
	goMongo.Db = "logistics"
	resp := goMongo.Connect()
	resp1 := goMongo.SetCollection("orderTest")
	ids, _ := uuid2.NewUUID()
	order := OrderTest{ids.String(), 0}
	resp2 := goMongo.AddOne(order)
	resp3 := goMongo.CloseConn()
	fmt.Println(resp.Message)
	fmt.Println(resp1.Message)
	fmt.Println(resp2.Message + Strval(resp2.Data))
	fmt.Println(resp3.Message)
}

func TestGoMongo_AddMany(t *testing.T) {
	goMongo := GoMongo{}
	goMongo.Addr = "192.168.4.81"
	goMongo.Port = "27017"
	goMongo.User = "logistics"
	goMongo.Pass = "123456"
	goMongo.Db = "logistics"
	resp := goMongo.Connect()
	resp1 := goMongo.SetCollection("orderTest")
	id, _ := uuid2.NewUUID()
	o1 := OrderTest{id.String(), 0}
	id1, _ := uuid2.NewUUID()
	o2 := OrderTest{id1.String(), 0}
	id2, _ := uuid2.NewUUID()
	o3 := OrderTest{id2.String(), 0}
	orders := []interface{}{o1, o2, o3}
	resp2 := goMongo.AddMany(orders)
	resp3 := goMongo.CloseConn()
	fmt.Println(resp.Message)
	fmt.Println(resp1.Message)
	fmt.Println(resp2.Message)
	fmt.Println(resp3.Message)
}

func TestGoMongo_GetOne(t *testing.T) {
	goMongo := GoMongo{}
	goMongo.Addr = "192.168.4.81"
	goMongo.Port = "27017"
	goMongo.User = "logistics"
	goMongo.Pass = "123456"
	goMongo.Db = "logistics"
	resp := goMongo.Connect()
	resp1 := goMongo.SetCollection("orderTest")
	var filter = &Filter{}
	filter.Type = "and"
	var fMap = map[string]interface{}{}
	fMap["test"] = "9fb73527-f076-11ec-81a7-2c4d54d02652"
	filter.FMap = fMap
	resp2 := goMongo.GetOne(filter)
	resp3 := goMongo.CloseConn()
	fmt.Println(resp.Message)
	fmt.Println(resp1.Message)
	fmt.Println(resp2.Data)
	fmt.Println(resp3.Message)
}

func TestGoMongo_GetMany(t *testing.T) {
	goMongo := GoMongo{}
	goMongo.Addr = "192.168.4.81"
	goMongo.Port = "27017"
	goMongo.User = "logistics"
	goMongo.Pass = "123456"
	goMongo.Db = "logistics"
	resp := goMongo.Connect()
	resp1 := goMongo.SetCollection("orderTest")
	var filter = &Filter{}
	var option = &Option{}
	filter.Type = "and"
	var fMap = map[string]interface{}{}
	fMap["status"] = 0
	filter.FMap = fMap

	resp2 := goMongo.GetMany(filter, option)
	resp3 := goMongo.CloseConn()
	fmt.Println(resp.Message)
	fmt.Println(resp1.Message)
	fmt.Println(resp2.Data)
	fmt.Println(resp3.Message)
}

func TestGoMongo_UpdateOne(t *testing.T) {
	goMongo := GoMongo{}
	goMongo.Addr = "192.168.4.81"
	goMongo.Port = "27017"
	goMongo.User = "logistics"
	goMongo.Pass = "123456"
	goMongo.Db = "logistics"
	resp := goMongo.Connect()
	resp1 := goMongo.SetCollection("orderTest")
	var filter = &Filter{}
	filter.Type = "and"
	var fMap = map[string]interface{}{}
	fMap["status"] = 0
	filter.FMap = fMap
	var uData = &UData{}
	var uMap = map[string]interface{}{}
	uMap["status"] = 1
	uData.UMap = uMap
	resp2 := goMongo.UpdateOne(filter, uData)
	resp3 := goMongo.CloseConn()
	fmt.Println(resp.Message)
	fmt.Println(resp1.Message)
	fmt.Println(resp2.Data)
	fmt.Println(resp3.Message)
}

func TestGoMongo_UpdateMany(t *testing.T) {
	goMongo := GoMongo{}
	goMongo.Addr = "192.168.4.81"
	goMongo.Port = "27017"
	goMongo.User = "logistics"
	goMongo.Pass = "123456"
	goMongo.Db = "logistics"
	resp := goMongo.Connect()
	resp1 := goMongo.SetCollection("orderTest")
	var filter = &Filter{}
	filter.Type = "and"
	var fMap = map[string]interface{}{}
	fMap["status"] = 0
	filter.FMap = fMap
	var uData = &UData{}
	var uMap = map[string]interface{}{}
	uMap["status"] = 1
	uData.UMap = uMap
	resp2 := goMongo.UpdateMany(filter, uData)
	resp3 := goMongo.CloseConn()
	fmt.Println(resp.Message)
	fmt.Println(resp1.Message)
	fmt.Println(resp2.Data)
	fmt.Println(resp3.Message)
}

func TestGoMongo_Count(t *testing.T) {
	goMongo := GoMongo{}
	goMongo.Addr = "192.168.4.81"
	goMongo.Port = "27017"
	goMongo.User = "logistics"
	goMongo.Pass = "123456"
	goMongo.Db = "logistics"
	resp := goMongo.Connect()
	resp1 := goMongo.SetCollection("orderTest")
	var filter = &Filter{}
	filter.Type = "and"
	var fMap = map[string]interface{}{}
	fMap["status"] = 1
	filter.FMap = fMap
	resp2 := goMongo.Count(filter)
	resp3 := goMongo.CloseConn()
	fmt.Println(resp.Message)
	fmt.Println(resp1.Message)
	fmt.Println(resp2.Data)
	fmt.Println(resp3.Message)
}

func TestGoMongo_Distinct(t *testing.T) {
	goMongo := GoMongo{}
	goMongo.Addr = "192.168.4.81"
	goMongo.Port = "27017"
	goMongo.User = "logistics"
	goMongo.Pass = "123456"
	goMongo.Db = "logistics"
	resp := goMongo.Connect()
	resp1 := goMongo.SetCollection("orderTest")
	var filter = &Filter{}
	filter.Type = "and"
	var fMap = map[string]interface{}{}
	fMap["status"] = 1
	filter.FMap = fMap
	resp2 := goMongo.Distinct(filter, "test")
	resp3 := goMongo.CloseConn()
	fmt.Println(resp.Message)
	fmt.Println(resp1.Message)
	fmt.Println(resp2.Data)
	fmt.Println(resp3.Message)
}

func TestGoMongo_DeleteOne(t *testing.T) {
	goMongo := GoMongo{}
	goMongo.Addr = "192.168.4.81"
	goMongo.Port = "27017"
	goMongo.User = "logistics"
	goMongo.Pass = "123456"
	goMongo.Db = "logistics"
	resp := goMongo.Connect()
	resp1 := goMongo.SetCollection("orderTest")
	var filter = &Filter{}
	filter.Type = "and"
	var fMap = map[string]interface{}{}
	fMap["status"] = 1
	filter.FMap = fMap
	resp2 := goMongo.DeleteOne(filter)
	resp3 := goMongo.CloseConn()
	fmt.Println(resp.Message)
	fmt.Println(resp1.Message)
	fmt.Println(resp2.Data)
	fmt.Println(resp3.Message)
}

func TestGoMongo_DeleteMany(t *testing.T) {
	goMongo := GoMongo{}
	goMongo.Addr = "192.168.4.81"
	goMongo.Port = "27017"
	goMongo.User = "logistics"
	goMongo.Pass = "123456"
	goMongo.Db = "logistics"
	resp := goMongo.Connect()
	resp1 := goMongo.SetCollection("orderTest")
	var filter = &Filter{}
	filter.Type = "and"
	var fMap = map[string]interface{}{}
	fMap["status"] = 1
	filter.FMap = fMap
	resp2 := goMongo.DeleteMany(filter)
	resp3 := goMongo.CloseConn()
	fmt.Println(resp.Message)
	fmt.Println(resp1.Message)
	fmt.Println(resp2.Data)
	fmt.Println(resp3.Message)
}
