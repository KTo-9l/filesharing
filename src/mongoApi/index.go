package mongoApi

import (
	"github.com/big-larry/mgo"
)

func EnsurePathIndex(bucket *mgo.GridFS) error {
	return bucket.Files.EnsureIndexKey("path")
}
