package mongoApi

import (
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/big-larry/mgo"
)

func Insert(bucket *mgo.GridFS, file *os.File) (interface{}, error) {
	filename := file.Name()

	gridFSFile, err := bucket.Create(filename)
	if err != nil {
		log.Println("Error creating gfs file:", err)
		return nil, err
	}

	_, err = io.Copy(gridFSFile, file)
	if err != nil {
		log.Println("Error copying file to gridFS:", err)
		return nil, err
	}

	err = file.Close()
	if err != nil {
		log.Println("Error closing file:", err)
		return nil, err
	}

	err = gridFSFile.Close()
	if err != nil {
		log.Println("Error closing gridFS File:", err)
		return nil, err
	}

	return gridFSFile.Id(), nil
}

func InsertFromMultipart(bucket *mgo.GridFS, fileHeader *multipart.FileHeader) (interface{}, error) {
	gridFSFile, err := bucket.Create(fileHeader.Filename)
	if err != nil {
		log.Println("Error creating gfs file:", err)
		return nil, err
	}

	file, err := fileHeader.Open()
	if err != nil {
		log.Println("Error open file:", err)
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(gridFSFile, file)
	if err != nil {
		log.Println("Error copying file to gridFS:", err)
		return nil, err
	}

	err = file.Close()
	if err != nil {
		log.Println("Error closing file:", err)
		return nil, err
	}

	err = gridFSFile.Close()
	if err != nil {
		log.Println("Error closing gridFS File:", err)
		return nil, err
	}

	return gridFSFile.Id(), nil
}
