package mongodb

import (
	"context"
	"gRPC_school_api/internals/models"
	"gRPC_school_api/pkg/utils"
	pb "gRPC_school_api/proto/gen"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
)

func DecodeEntities[T any, M any](ctx context.Context, cursor *mongo.Cursor, newEntity func() *T, newModel func() *M) ([]*T, error) {
	var Entities []*T
	for cursor.Next(ctx) {
		model := newModel()
		err := cursor.Decode(&model)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Internal error")
		}
		entity := newEntity()
		modelVal := reflect.ValueOf(model).Elem()
		pbVal := reflect.ValueOf(entity).Elem()

		for i := 0; i < modelVal.NumField(); i++ {
			modelField := modelVal.Field(i)
			modelFieldName := modelVal.Type().Field(i).Name

			pbField := pbVal.FieldByName(modelFieldName)
			if pbField.IsValid() && pbField.CanSet() {
				pbField.Set(modelField)
			}
		}
		Entities = append(Entities, entity)
	}

	err := cursor.Err()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	return Entities, nil
}

func MapModelToPb[M, P any](model M, newPb func() *P) *P {
	pbStruct := newPb()
	modelVal := reflect.ValueOf(model)
	pbVal := reflect.ValueOf(pbStruct).Elem()
	for i := 0; i < modelVal.NumField(); i++ {
		modelField := modelVal.Field(i)
		modelFieldtype := modelVal.Type().Field(i)
		pbField := pbVal.FieldByName(modelFieldtype.Name)
		if pbField.IsValid() && pbField.CanSet() {
			pbField.Set(modelField)
		}
	}
	return pbStruct
}

func MapModelTeacherToPb(teacherModel models.Teacher) *pb.Teacher {
	return MapModelToPb(teacherModel, func() *pb.Teacher {
		return &pb.Teacher{}
	})
}

// TODO:
func MapModelStudentToPb(studentModel models.Student) *pb.Student {
	return MapModelToPb(studentModel, func() *pb.Student {
		return &pb.Student{}
	})
}

// TODO:
func MapModelExecToPb(execModel models.Exec) *pb.Exec {
	return MapModelToPb(execModel, func() *pb.Exec {
		return &pb.Exec{}
	})
}

func MapPbToModel[M, P any](pbStruct P, newModel func() *M) *M {
	modelStruct := newModel()
	pbval := reflect.ValueOf(pbStruct).Elem()
	modelVal := reflect.ValueOf(modelStruct).Elem()
	for i := 0; i < pbval.NumField(); i++ {
		pbField := pbval.Field(i)
		fieldName := pbval.Type().Field(i).Name
		modelfield := modelVal.FieldByName(fieldName)
		if modelfield.IsValid() && modelfield.CanSet() {
			modelfield.Set(pbField)
		}
	}
	return modelStruct
}

func MapPbTeacherToModelTeacher(pbTeacher *pb.Teacher) *models.Teacher {
	return MapPbToModel(pbTeacher, func() *models.Teacher {
		return &models.Teacher{}
	})
}

// TODO:
func MapPbStudentToModelStudent(pbStudent *pb.Student) *models.Student {
	return MapPbToModel(pbStudent, func() *models.Student {
		return &models.Student{}
	})
}

// TODO:
func MapPbExecToModelExec(pbExec *pb.Exec) *models.Exec {
	return MapPbToModel(pbExec, func() *models.Exec {
		return &models.Exec{}
	})
}
