package mongodb

import (
	"context"
	"gRPC_school_api/internals/models"
	"gRPC_school_api/pkg/utils"
	pb "gRPC_school_api/proto/gen"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddStudentsToDB(ctx context.Context, studentsFromReq []*pb.Student) ([]*pb.Student, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	defer client.Disconnect(ctx)

	newStudents := make([]*models.Student, len(studentsFromReq))
	for i, pbStudent := range studentsFromReq {
		newStudents[i] = MapPbStudentToModelStudent(pbStudent)
	}

	var addedStudents []*pb.Student

	for _, student := range newStudents {
		result, err := client.Database("School").Collection("students").InsertOne(ctx, student)
		if err != nil {
			return nil, utils.ErrorHandler(err, "error adding value to database")
		}

		objectid, ok := result.InsertedID.(primitive.ObjectID)
		if ok {
			student.Id = objectid.Hex()
		}

		pbStudent := MapModelStudentToPb(*student)
		addedStudents = append(addedStudents, pbStudent)
	}
	return addedStudents, nil
}

func GetStudentsfromDB(ctx context.Context, sortOptions primitive.D, filter primitive.M, pageNumber, pageSize uint32) ([]*pb.Student, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal Error")
	}
	defer client.Disconnect(ctx)
	collection := client.Database("School").Collection("students")

	findOptions := options.Find()
	findOptions.SetSkip(int64((pageNumber - 1) * pageSize))
	findOptions.SetLimit(int64(pageSize))

	if len(sortOptions) > 0 {
		findOptions.SetSort(sortOptions)
	}
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal Error")
	}
	defer cursor.Close(ctx)

	students, err := DecodeEntities(ctx, cursor, pbModelStudent, newModelStudent)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	return students, nil
}

func pbModelStudent() *pb.Student { return &pb.Student{} }

func newModelStudent() *models.Student { return &models.Student{} }
