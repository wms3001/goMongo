# goMongo

### 1. install
```go
go get github.com/wms3001/goMongo
```
### 2.Add folder in project directory and add mongo.yml file
```yml
mongo:
  host: 192.168.1.1
  port: 27017
  user: test
  pass: 123456
  dbName: test
```
### 3. example
```go
conn := goMongo.OpenConn("order")
defer goMongo.CloseConn(conn.Client)
```
order is Collection
add
```go
addOne = goMongo.AddOne(conn.Collection, &order)
addMany = goMongo.AddMany(conn.Collection,orders)
```
getone
```go
ser := bson.D{{"_id",id}}
one := goMongo.GetOneSingl(conn.Collection,ser)
```
mData :
{
"filterMap":{
"status":0   
},
"type":"and"
,
"upMap":{
"status":1
},
"limit":5,
"skip":0,
"sortMap":{}
}

get
```go
one := goMongo.GetOne(conn.Collection,mData)
many := goMongo.GetMany(conn.Collection,mData)
```
update
```go
up := goMongo.UpdateOne(conn.Collection,mData)
up := goMongo.UpdateMany(conn.Collection,mData)
```
delete
```go
del := goMongo.DeleteOne(conn.Collection,mData)
del := goMongo.DeleteMany(conn.Collection,mData)
```
count
```go
cou := goMongo.Count(conn.Collection,mData)
```
distinct
```go
dis := goMongo.Distinct(conn.Collection,mData,"carrier")
```
carrier is field
