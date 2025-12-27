package handlers

import (
	"gRPC_school_api/pkg/utils"
	"reflect"
	"strings"

	pb "gRPC_school_api/proto/gen"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuildFilter(object interface{}, model interface{}) (bson.M, error) {
	filter := bson.M{}

	if object == nil || reflect.ValueOf(object).IsNil() {
		return filter, nil
	}

	// var ModelTeacher models.Teacher
	modelVal := reflect.ValueOf(model).Elem()
	modelType := modelVal.Type()

	reqVal := reflect.ValueOf(object).Elem()
	reqType := reqVal.Type()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		fieldName := reqType.Field(i).Name

		if fieldVal.IsValid() && !fieldVal.IsZero() {
			modelfield := modelVal.FieldByName(fieldName)
			if modelfield.IsValid() && modelfield.CanSet() {
				modelfield.Set(fieldVal)
			}
		}
	}

	// Now we iterate over the modelTeacher to build filter using bson.M
	for i := 0; i < modelVal.NumField(); i++ {
		fieldVal := modelVal.Field(i)
		fieldName := modelType.Field(i).Name

		if fieldVal.IsValid() && !fieldVal.IsZero() {
			bsonTag := modelType.Field(i).Tag.Get("bson")
			bsonTag = strings.TrimSuffix(bsonTag, ",omitempty")
			if bsonTag == "_id" {
				objId, err := primitive.ObjectIDFromHex(reqVal.FieldByName(fieldName).Interface().(string))
				if err != nil {
					return nil, utils.ErrorHandler(err, "Internal error")
				}
				filter[bsonTag] = objId
			} else {
				filter[bsonTag] = fieldVal.Interface().(string)
			}

		}
	}
	return filter, nil
}

func BuildSortOptions(sortfields []*pb.SortField) bson.D {
	var sortOptions bson.D

	for _, sortfield := range sortfields {
		order := 1
		if sortfield.GetOrder() == pb.Order_DESC {
			order = -1
		}
		sortOptions = append(sortOptions, bson.E{Key: sortfield.Field, Value: order})
	}
	return sortOptions
}
