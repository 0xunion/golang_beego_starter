package model

import (
	"context"
	"errors"
	"reflect"
	gostrings "strings"
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
func MongoArrayContainsFilter(key string, value interface{}) MongoFilterItem {
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

func MongoSort(key string, order int) bson.D {
	return bson.D{{Key: key, Value: order}}
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

var last_check_time int64

func checkAlive() bool {
	if mongoConn == nil {
		mongoConnMutex.Lock()
		defer mongoConnMutex.Unlock()
	}

	// check if there is a checkAlive request in 30 seconds
	if time.Now().Unix()-last_check_time < 30 {
		return true
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
	last_check_time = time.Now().Unix()
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

type updateMongoField struct {
	UpdateType string      // $set, $inc, $push, $pull
	Field      string      // field name
	Value      interface{} // if UpdateType is $inc, Value must be int
}

func MongoIncField(field string, value interface{}) updateMongoField {
	return updateMongoField{
		UpdateType: "$inc",
		Field:      field,
		Value:      value,
	}
}

func MongoDecField(field string, value interface{}) updateMongoField {
	return updateMongoField{
		UpdateType: "$inc",
		Field:      field,
		Value:      -value.(int),
	}
}

func MongoSetField(field string, value interface{}) updateMongoField {
	return updateMongoField{
		UpdateType: "$set",
		Field:      field,
		Value:      value,
	}
}

func MongoPushField(field string, value interface{}) updateMongoField {
	return updateMongoField{
		UpdateType: "$push",
		Field:      field,
		Value:      value,
	}
}

func MongoPullField(field string, value interface{}) updateMongoField {
	return updateMongoField{
		UpdateType: "$pull",
		Field:      field,
		Value:      value,
	}
}

// ModelUpdateField update a model field to mongodb
func ModelUpdateField[T any](filter MongoFilter, fields ...updateMongoField) error {
	var data T
	collection := meta.GetTypeName(data)
	if !checkAlive() {
		return errors.New("mongodb connection lost")
	}

	coll := mongoConn.Database(conf.MongoDBName()).Collection(collection)
	update := bson.D{}
	for _, field := range fields {
		update = append(update, bson.E{Key: field.UpdateType, Value: bson.D{{Key: field.Field, Value: field.Value}}})
	}

	result, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("no matched data")
	}

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

// ModelGetJoin return the model join with other model
// R is the result model with join data, T is the model we want
// foreign Model info must be set in lookup
type Lookup struct {
	From         string
	LocalField   string
	ForeignField string
	As           string
}

var lookup_cache = make(map[string][]Lookup)
var lookup_cache_mutex = sync.RWMutex{}

func getLookupCache(key string) ([]Lookup, bool) {
	lookup_cache_mutex.RLock()
	defer lookup_cache_mutex.RUnlock()
	lookup, ok := lookup_cache[key]
	return lookup, ok
}

func setLookupCache(key string, lookup []Lookup) {
	lookup_cache_mutex.Lock()
	defer lookup_cache_mutex.Unlock()
	lookup_cache[key] = lookup
}

func getLookup[T any, R any]() ([]Lookup, error) {
	var origin T
	var result R
	collection := reflect.TypeOf(origin).Name()
	// get lookup info from cache
	lookup, ok := getLookupCache(collection)
	if !ok {
		// scan R to get lookup info
		reflectR := reflect.ValueOf(result)
		lookup = make([]Lookup, 0)
		for i := 0; i < reflectR.NumField(); i++ {
			field := reflectR.Type().Field(i)
			if field.Type.Kind() == reflect.Struct {
				// get tag
				tag := field.Tag.Get("join")
				if tag != "" {
					if tag == "-" {
						continue
					}
					// parse tag "uid=id"
					tagList := gostrings.Split(tag, "=")
					if len(tagList) != 2 {
						return nil, errors.New("join tag error")
					}
					// get field type
					fieldType := reflectR.Field(i).Type()
					// get foreign collection name
					foreignCollection := meta.GetTypeName(reflect.New(fieldType).Interface())
					localField := ""
					foreignField := ""
					// use reflect to get bson tag
					// walk through all fields in struct to find the bson tag
					for j := 0; j < reflectR.NumField(); j++ {
						field := reflectR.Type().Field(j)
						if field.Type.Kind() == reflect.Struct {
							continue
						}
						bsonTag := field.Tag.Get("bson")
						// split bson tag to get the real tag
						bsonTagList := gostrings.Split(bsonTag, ",")
						if len(bsonTagList) == 0 {
							continue
						}

						if bsonTagList[0] == tagList[0] {
							localField = bsonTag
						}
					}
					// walk through all fields in struct to find the bson tag
					for j := 0; j < fieldType.NumField(); j++ {
						field := fieldType.Field(j)
						if field.Type.Kind() == reflect.Struct {
							continue
						}
						bsonTag := field.Tag.Get("bson")
						// split bson tag to get the real tag
						bsonTagList := gostrings.Split(bsonTag, ",")
						if len(bsonTagList) == 0 {
							continue
						}
						if bsonTagList[0] == tagList[1] {
							foreignField = bsonTag
						}
					}

					// check if we get the localField and foreignField
					if localField == "" || foreignField == "" {
						return nil, errors.New("join tag error")
					}

					// add to lookup
					lookup = append(lookup, Lookup{
						From:         foreignCollection,
						LocalField:   localField,
						ForeignField: foreignField,
						As:           field.Name,
					})
				}
			}
		}

		// set lookup cache
		setLookupCache(collection, lookup)
	}

	return lookup, nil
}

func ModelGetJoin[T any, R any](filter MongoFilter) (*R, error) {
	var origin T
	collection := meta.GetTypeName(origin)
	if !checkAlive() {
		return nil, errors.New("mongodb connection lost")
	}
	coll := mongoConn.Database(conf.MongoDBName()).Collection(collection)

	// get lookup info
	lookup, err := getLookup[T, R]()
	if err != nil {
		return nil, err
	}

	// create lookup
	lookupList := make([]bson.D, 0)
	for _, item := range lookup {
		lookupList = append(lookupList, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: item.From},
				{Key: "localField", Value: item.LocalField},
				{Key: "foreignField", Value: item.ForeignField},
				{Key: "as", Value: item.As},
			}},
		})

		// add unwind
		lookupList = append(lookupList, bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$" + item.As},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}},
		})
	}

	// create pipeline
	pipeline := make([]bson.D, 0)
	pipeline = append(pipeline, bson.D{{Key: "$match", Value: filter}})
	pipeline = append(pipeline, lookupList...)
	pipeline = append(pipeline, bson.D{{Key: "$limit", Value: 1}})

	// exec pipeline
	cursor, err := coll.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	var result []*R
	err = cursor.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}

	// join data
	if len(result) > 0 {
		return result[0], nil
	}

	return nil, errors.New("not found")
}

// ModelGetAllJoin return all the model join with other model
// R is the result model with join data, T is the model we want
// foreign Model info must be set in lookup
func ModelGetAllJoin[T any, R any](filter MongoFilter) ([]*R, error) {
	var origin T
	collection := meta.GetTypeName(origin)
	if !checkAlive() {
		return nil, errors.New("mongodb connection lost")
	}
	coll := mongoConn.Database(conf.MongoDBName()).Collection(collection)

	// get lookup info
	lookup, err := getLookup[T, R]()
	if err != nil {
		return nil, err
	}

	// create lookup
	lookupList := make([]bson.D, 0)
	for _, item := range lookup {
		lookupList = append(lookupList, bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: item.From},
				{Key: "localField", Value: item.LocalField},
				{Key: "foreignField", Value: item.ForeignField},
				{Key: "as", Value: item.As},
			}},
		})

		// add unwind
		lookupList = append(lookupList, bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$" + item.As},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			}},
		})
	}

	// create pipeline
	pipeline := make([]bson.D, 0)
	pipeline = append(pipeline, bson.D{{Key: "$match", Value: filter}})
	pipeline = append(pipeline, lookupList...)

	// exec pipeline
	cursor, err := coll.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	var result []*R
	err = cursor.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
