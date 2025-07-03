package mongoApi

import (
	"fmt"

	"github.com/big-larry/mgo"
	"github.com/big-larry/mgo/bson"
)

func ListFilesForPath(bucket *mgo.GridFS, path []string) (gfsFiles []GridFSFile, err error) {
	// query := bucket.Files.Find(bson.M{"path": bson.M{"$all": path}})
	query := bucket.Files.Find(bson.M{"path": path[0]})
	err = query.All(&gfsFiles)
	if err != nil {
		return nil, err
	}

	return gfsFiles, nil
}

func ListAllFiles(bucket *mgo.GridFS) (gfsFiles []GridFSFile, err error) {
	query := bucket.Files.Find(bson.M{})
	err = query.All(&gfsFiles)
	if err != nil {
		return nil, err
	}

	return gfsFiles, nil
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
