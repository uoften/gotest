package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TheGroup struct{
	Id bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Totalamount int
	Dayofyear int
	Actualyear string
	Transactiondate string
	Count int
}

var results []TheGroup

func main() {
	// 连接数据库
	session, err := mgo.Dial("192.168.10.198:27018")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// 获取数据库,获取集合
	c := session.DB(DeviceOrderDbName).C(DeviceOrderCollectionName)

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{ "transactiontype": transactiontype }
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"year": bson.M{ "$year": "$transactiondate" },
					"dayOfYear": bson.M{ "$dayOfYear": "$transactiondate" }
				},
				"totalamount": bson.M{ "$sum": "$qty" },
				"count": bson.M{ "$sum": 1 },
				"date": bson.M{ "$first": "$transactiondate"}
			}
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
				"totalamount": 1,
				"dayOfYear": "$_id.dayOfYear",
				"actualyear": bson.M{ "$substr": []interface{}{ "$_id.year", 0, 4 } },
				"transactiondate": bson.M{
					"$dateToString": bson.M{ "format": "%Y-%m-%d %H:%M", "date": "$date" }
				},
				"count": 1
			}
		}
	}

	pipe := collection.Pipe(pipeline)

	fmt.Println(pipe);
}
