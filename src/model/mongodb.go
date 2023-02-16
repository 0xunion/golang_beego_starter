package model

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	conf "github.com/0xunion/exercise_back/src/const/conf"
	routine "github.com/0xunion/exercise_back/src/routine"
	"github.com/0xunion/exercise_back/src/util/meta"
	strings "github.com/0xunion/exercise_back/src/util/strings"
)

type MongoFilter = bson.D
type MongoFilterItem = bson.E
type MongoOptions = options.FindOptions

func IdFilter(id primitive.ObjectID) MongoFilterItem {
	return MongoFilterItem{Key: "_id", Value: id}
}

// check if the value is in the array
func MongoArrayContainsFilter(key string, value ...interface{}) MongoFilterItem {
	return MongoFilterItem{Key: key, Value: bson.M{"$in": value}}
}

// check if the value is not in the array
func MongoArrayNotContainsFilter(key string, value ...interface{}) MongoFilterItem {
	return MongoFilterItem{Key: key, Value: bson.M{"$nin": value}}
}

// normal filter
func MongoKeyFilter(key string, value interface{}) MongoFilterItem {
	return MongoFilterItem{Key: key, Value: value}
}

// bit filter
func MongoNoBitFilter(key string, value int) MongoFilterItem {
	return MongoFilterItem{Key: key, Value: bson.M{"$bitsAllClear": value}}
}

func MongoHasBitFilter(key string, value int) MongoFilterItem {
	return MongoFilterItem{Key: key, Value: bson.M{"$bitsAllSet": value}}
}

// flag filter
func MongoNoFlagFilter(value int) MongoFilterItem {
	return MongoFilterItem{Key: "basictype.flag", Value: bson.M{"$bitsAllClear": value}}
}

func MongoHasFlagFilter(value int) MongoFilterItem {
	return MongoFilterItem{Key: "basictype.flag", Value: bson.M{"$bitsAllSet": value}}
}

// this filter will check if the value is in the range
// eg: MongoValueInRangeFilter("age", 18, 30, true, false) will check if the age is in [18, 30)
//
//			MongoValueInRangeFilter("age", nil, 30, false, true) will check if age is in (-inf, 30]
//	     MongoValueInRangeFilter("age", 18, nil) will check if age is in [18, +inf)
func MongoValueInRangeFilter(key string, min interface{}, max interface{}, equals ...bool) MongoFilterItem {
	order := make(bson.M)

	min_eq := false
	max_eq := false
	if len(equals) > 0 {
		min_eq = equals[0]
	}
	if len(equals) > 1 {
		max_eq = equals[1]
	}

	if min != nil {
		if min_eq {
			order["$gte"] = min
		} else {
			order["$gt"] = min
		}
	}
	if max != nil {
		if max_eq {
			order["$lte"] = max
		} else {
			order["$lt"] = max
		}
	}
	return MongoFilterItem{Key: key, Value: order}
}

func NewMongoFilter(items ...MongoFilterItem) MongoFilter {
	filter := make(MongoFilter, 0)
	for _, item := range items {
		filter = append(filter, item)
	}

	return filter
}

var mongoConnMutex sync.Mutex
var mongoConn *mongo.Client

func init() {
	routine.Info("[MongoDB] Init mongodb monitor")
	mongoConnMutex.Lock()
	routine.Go("mongodb_monitor", func() {
		for {
			// connect to mongodb
			routine.Info("[MongoDB] Start connect mongodb")
			err := mongoConnect()
			if err != nil {
				routine.Error("[MongoDB] Connect error: %v", err)
				time.Sleep(5 * time.Second)
				continue
			}
			routine.Info("[MongoDB] Connect success")
			// release lock
			mongoConnMutex.Unlock()
			// check mongodb connection
			for {
				// check mongodb connection
				err = mongoConn.Ping(context.Background(), nil)
				if err != nil {
					routine.Error("[MongoDB] MongoConnection Lost: %v", err)
					break
				}
				time.Sleep(5 * time.Second)
			}
		}
	})
}

func mongoConnect() error {
	uri := strings.StringJoin(
		"mongodb://",
		conf.MongoDBHost(),
		":",
		conf.MongoDBPort(),
	)

	client_options := options.Client().ApplyURI(uri)
	if conf.MongoDBUser() != "" {
		client_options = client_options.SetAuth(options.Credential{
			Username: conf.MongoDBUser(),
			Password: conf.MongoDBPass(),
		})
	}

	client, err := mongo.NewClient(client_options)
	if err != nil {
		return err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return err
	}

	mongoConn = client
	return nil
}

func checkAlive() bool {
	if mongoConn == nil {
		mongoConnMutex.Lock()
		defer mongoConnMutex.Unlock()
	}

	// the code below will cause serval performance problems
	// every query will need 2 request to mongodb, one for checkAlive, one for query
	// so, if it cause serious problem, we delete this function
	// check mongodb connection
	err := mongoConn.Ping(context.Background(), nil)
	if err != nil {
		// reconnect mongodb
		routine.Info("[MongoDB] MongoConnection Lost: %v", err)
		err = mongoConnect()
		if err != nil {
			routine.Error("[MongoDB] Connect error: %v", err)
			return false
		}
	}
	return true
}

// NativeQuery return a mongo cursor for query
func NativeQuery(name string) (*mongo.Cursor, error) {
	if !checkAlive() {
		return nil, nil
	}
	collection := mongoConn.Database(conf.MongoDBName()).Collection(name)
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

// ModelGet return the model we want
func ModelGet[T any](filter MongoFilter) (*T, error) {
	var result T
	collection := meta.GetTypeName(result)
	if !checkAlive() {
		return &result, errors.New("mongodb connection lost")
	}
	coll := mongoConn.Database(conf.MongoDBName()).Collection(collection)
	err := coll.FindOne(context.Background(), filter).Decode(&result)
	return &result, err
}

// ModelGetAll return all the model we want
func ModelGetAll[T any](filter MongoFilter, options ...*MongoOptions) ([]T, error) {
	var result []T
	collection := meta.GetTypeName(result)
	if !checkAlive() {
		return result, errors.New("mongodb connection lost")
	}
	coll := mongoConn.Database(conf.MongoDBName()).Collection(collection)
	cursor, err := coll.Find(context.Background(), filter, options...)
	if err != nil {
		return result, err
	}
	err = cursor.All(context.Background(), &result)
	return result, err
}

// ModelInsert insert a model to mongodb
func ModelInsert[T any](model *T, id *primitive.ObjectID) error {
	collection := meta.GetTypeName(model)
	if !checkAlive() {
		return errors.New("mongodb connection lost")
	}
	coll := mongoConn.Database(conf.MongoDBName()).Collection(collection)
	result, err := coll.InsertOne(context.Background(), model)
	if err != nil {
		return err
	}

	// set id
	if id != nil {
		*id = result.InsertedID.(primitive.ObjectID)
	}
	return err
}

// ModelUpdate update a model to mongodb
// warning: the model must has the complete data structure, or it will cause field empty
func ModelUpdate[T any](filter MongoFilter, model *T) error {
	collection := meta.GetTypeName(model)
	if !checkAlive() {
		return errors.New("mongodb connection lost")
	}
	coll := mongoConn.Database(conf.MongoDBName()).Collection(collection)
	_, err := coll.UpdateOne(context.Background(), filter, bson.D{{Key: "$set", Value: model}})
	return err
}

// ModelDelete delete a model to mongodb
func ModelDelete[T any](filter MongoFilter) error {
	var data T
	collection := meta.GetTypeName(data)
	if !checkAlive() {
		return errors.New("mongodb connection lost")
	}
	coll := mongoConn.Database(conf.MongoDBName()).Collection(collection)
	_, err := coll.DeleteOne(context.Background(), filter)
	return err
}
