package mongoApi

import (
	"mime/multipart"
	"os"

	"github.com/big-larry/mgo"
	"github.com/big-larry/mgo/bson"
)

func UpdateById(bucket *mgo.GridFS, id string, file *os.File, path []string) (ok bool, err error) {
	if !bson.IsObjectIdHex(id) {
		return false, nil
	}

	objId := bson.ObjectIdHex(id)

	err = bucket.Files.RemoveId(objId)
	if err != nil {
		return false, err
	}

	_, err = bucket.Chunks.RemoveAll(bson.M{"files_id": objId})
	if err != nil {
		return false, err
	}

	_, err = InsertWithId(bucket, file, path, objId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func UpdateByIdFromMultipart(bucket *mgo.GridFS, id string, fileHeader *multipart.FileHeader, path []string) (ok bool, err error) {
	if !bson.IsObjectIdHex(id) {
		return false, nil
	}

	objId := bson.ObjectIdHex(id)

	err = bucket.Files.RemoveId(objId)
	if err != nil {
		return false, err
	}

	_, err = bucket.Chunks.RemoveAll(bson.M{"files_id": objId})
	if err != nil {
		return false, err
	}

	_, err = InsertWithIdFromMultipart(bucket, fileHeader, path, objId)
	if err != nil {
		return false, err
	}

	return true, nil
}
