package mongoApi

import (
	"fmt"

	"github.com/big-larry/mgo"
	"github.com/big-larry/mgo/bson"
)

func GetAllFiles(bucket *mgo.GridFS) (gfsFiles []GridFSFile, err error) { // check with mgo.GridFile
	query := bucket.Files.Find(bson.M{})
	err = query.All(&gfsFiles)
	if err != nil {
		return nil, err
	}

	return gfsFiles, nil
}

func GetFileById(bucket *mgo.GridFS, idString string) (gfsFile *mgo.GridFile, err error) {
	if bson.IsObjectIdHex(idString) {
		idHex := bson.ObjectIdHex(idString)
		gfsFile, err = bucket.OpenId(idHex)
	} else {
		gfsFile, err = bucket.OpenId(idString)
	}

	if err != nil && err != mgo.ErrNotFound {
		return nil, err
	}

	return gfsFile, nil
}

func GetAllChunks(bucket *mgo.GridFS) {
	var fileMeta struct {
		Id       interface{} `bson:"_id"`
		Filename string      `bson:"filename"`
	}

	bucket.Files.Find(bson.M{"filename": "./testFile.txt"}).One(&fileMeta)
	// bucket.Files.Find(bson.M{"filename": "./test.msi"}).One(&fileMeta)
	fmt.Printf("Id for %v is: %v\n", fileMeta.Filename, fileMeta.Id)

	query := bucket.Chunks.Find(bson.M{"files_id": fileMeta.Id})

	var results any
	query.One(&results)

	fmt.Println(results)
}
