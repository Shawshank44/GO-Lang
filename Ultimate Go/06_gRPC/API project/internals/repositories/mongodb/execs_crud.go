package mongodb

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"gRPC_school_api/internals/models"
	"gRPC_school_api/pkg/utils"
	pb "gRPC_school_api/proto/gen"
	"os"
	"strconv"
	"time"

	"github.com/go-mail/mail/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddExecsToDB(ctx context.Context, execsFromReq []*pb.Exec) ([]*pb.Exec, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	defer client.Disconnect(ctx)

	newExecs := make([]*models.Exec, len(execsFromReq))
	for i, pbExec := range execsFromReq {
		newExecs[i] = MapPbExecToModelExec(pbExec)
		hashedPassword, err := utils.HashPassword(newExecs[i].Password)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Internal error")
		}
		newExecs[i].Password = hashedPassword
		currentTime := time.Now().Format(time.RFC3339)
		newExecs[i].UserCreatedAt = currentTime
		newExecs[i].InactiveStatus = false
	}

	var addedExecs []*pb.Exec

	for _, exec := range newExecs {
		result, err := client.Database("School").Collection("execs").InsertOne(ctx, exec)
		if err != nil {
			return nil, utils.ErrorHandler(err, "error adding value to database")
		}

		objectid, ok := result.InsertedID.(primitive.ObjectID)
		if ok {
			exec.Id = objectid.Hex()
		}

		pbExec := MapModelExecToPb(*exec)
		addedExecs = append(addedExecs, pbExec)
	}
	return addedExecs, nil
}

func GetExecsfromDB(ctx context.Context, sortOptions primitive.D, filter primitive.M, pageNumber, pageSize uint32) ([]*pb.Exec, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal Error")
	}
	defer client.Disconnect(ctx)
	collection := client.Database("School").Collection("execs")

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

	execs, err := DecodeEntities(ctx, cursor, pbModelExec, newModelExec)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	return execs, nil
}

func pbModelExec() *pb.Exec { return &pb.Exec{} }

func newModelExec() *models.Exec { return &models.Exec{} }

func UpdateExecsinDB(ctx context.Context, PbExecs []*pb.Exec) ([]*pb.Exec, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	defer client.Disconnect(ctx)
	var updatedExecs []*pb.Exec
	for _, exec := range PbExecs {
		if exec.Id == "" {
			return nil, utils.ErrorHandler(errors.New("id Cannot be blank"), "id cannot be blank")
		}
		modelExec := MapPbExecToModelExec(exec)
		objId, err := primitive.ObjectIDFromHex(exec.Id)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Internal error")
		}
		// Convert ModelExec to BSON document:
		modelDoc, err := bson.Marshal(modelExec)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Internal error")
		}
		var updateDoc bson.M
		err = bson.Unmarshal(modelDoc, &updateDoc)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Internal error")
		}
		// Removing the Id field from the update document
		delete(updateDoc, "_id")

		_, err = client.Database("School").Collection("execs").UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updateDoc})
		if err != nil {
			return nil, utils.ErrorHandler(err, fmt.Sprintf("Error updating exec id : %v", exec.Id))
		}
		updatedExec := MapModelExecToPb(*modelExec)
		updatedExecs = append(updatedExecs, updatedExec)
	}
	return updatedExecs, nil
}

func DeleteExecsFromDB(ctx context.Context, execIdsToDelete []string) ([]string, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	defer client.Disconnect(ctx)

	objectIds := make([]primitive.ObjectID, len(execIdsToDelete))
	for i, id := range execIdsToDelete {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Invalid ID")
		}
		objectIds[i] = objectId
	}

	filter := bson.M{"_id": bson.M{"$in": objectIds}}

	res, err := client.Database("School").Collection("execs").DeleteMany(ctx, filter)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error in deleting the _Id")
	}

	if res.DeletedCount == 0 {
		return nil, utils.ErrorHandler(err, "No execs were deleted")
	}

	deletedIds := make([]string, res.DeletedCount)
	for i, id := range objectIds {
		deletedIds[i] = id.Hex()
	}
	return deletedIds, nil
}

func GetUserByUserName(ctx context.Context, username string) (*models.Exec, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}

	defer client.Disconnect(ctx)

	filter := bson.M{"username": username}
	var exec models.Exec
	err = client.Database("School").Collection("execs").FindOne(ctx, filter).Decode(&exec)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, utils.ErrorHandler(err, "Internal error")
		}
	}
	return &exec, nil
}

func UpdatePasswordInDB(ctx context.Context, req *pb.UpdatePasswordRequest) (string, string, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return "", "", utils.ErrorHandler(err, "Internal error")
	}
	defer client.Disconnect(ctx)

	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return "", "", utils.ErrorHandler(err, "Internal error")
	}
	var user models.Exec
	err = client.Database("School").Collection("execs").FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return "", "", utils.ErrorHandler(err, "User not found")
	}
	err = utils.VerifyPassword(req.GetCurrentPassword(), user.Password)
	if err != nil {
		return "", "", utils.ErrorHandler(err, "Internal error")
	}
	HashedPassword, err := utils.HashPassword(req.GetNewPassword())
	if err != nil {
		return "", "", utils.ErrorHandler(err, "Internal error")
	}
	update := bson.M{
		"$set": bson.M{
			"password":            HashedPassword,
			"password_changed_at": time.Now().Format(time.RFC3339),
		},
	}

	_, err = client.Database("School").Collection("execs").UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return "", "", utils.ErrorHandler(err, "Internal error")
	}
	return user.Username, user.Role, nil
}

func DeactivateUserInDB(ctx context.Context, ids []string) (*mongo.UpdateResult, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	defer client.Disconnect(ctx)

	var objectIds []primitive.ObjectID
	for _, id := range ids {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Internal error")
		}
		objectIds = append(objectIds, objId)
	}

	filter := bson.M{
		"_id": bson.M{
			"$in": objectIds,
		},
	}

	update := bson.M{
		"$set": bson.M{
			"inactive_status": true,
		},
	}

	res, err := client.Database("School").Collection("execs").UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Failed to decativate users")
	}
	return res, nil
}

func ForgotPasswordDB(ctx context.Context, email string) (string, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return "", utils.ErrorHandler(err, "Internal server error")
	}
	defer client.Disconnect(ctx)

	var exec models.Exec
	err = client.Database("School").Collection("execs").FindOne(ctx, bson.M{"email": email}).Decode(&exec)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", utils.ErrorHandler(err, "Internal server error")
		}
		return "", utils.ErrorHandler(err, "Internal server error")
	}

	tokenbyte := make([]byte, 32)
	_, err = rand.Read(tokenbyte)
	if err != nil {
		return "", utils.ErrorHandler(err, "Internal server error")
	}

	token := hex.EncodeToString(tokenbyte)
	hashedToken := sha256.Sum256(tokenbyte)
	hashedTokenString := hex.EncodeToString(hashedToken[:])
	duration, err := strconv.Atoi(os.Getenv("RESET_TOKEN_EXP_DURATION"))
	if err != nil {
		return "", utils.ErrorHandler(err, "Internal server error")
	}
	mins := time.Duration(duration)
	expiry := time.Now().Add(mins * time.Minute).Format(time.RFC3339)

	update := bson.M{
		"$set": bson.M{
			"password_reset_token":   hashedTokenString,
			"password_token_expires": expiry,
		},
	}
	_, err = client.Database("School").Collection("exec").UpdateOne(ctx, bson.M{"email": email}, update)
	if err != nil {
		return "", utils.ErrorHandler(err, "Internal server error")
	}

	resetURL := fmt.Sprintf("https://localhost:50051/execs/resetpassword/reset/%s", token)
	message := fmt.Sprintf("forget your password? Reset your password using the following link \n %s\n Please use the reset code :: %s along with your request to change password.\n if you didn't request a password reset, Please ignore this mail,\n This link is only valid for %v minutes.", resetURL, token, mins)

	subject := "**ACTION NEEDED ON PASSWORD RESET Edia.Edu.Admins**"

	m := mail.NewMessage()
	m.SetHeader("From", "noreply@admins.edia.edu")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", message)

	d := mail.NewDialer("localhost", 1025, "", "")
	err = d.DialAndSend(m)
	if err != nil {
		clean := bson.M{
			"$set": bson.M{
				"password_reset_token":   nil,
				"password_token_expires": nil,
			},
		}
		_, err = client.Database("School").Collection("exec").UpdateOne(ctx, bson.M{"email": email}, clean)
		return "", utils.ErrorHandler(err, "Internal server error")
	}
	return message, nil
}
