package bucket

import (
	bucketModel "api/models/bucket"
	"fmt"

	"github.com/google/uuid"
)

type bucketTable struct {
	Buckets []bucketModel.Bucket
}

var table *bucketTable

func Create(i bucketModel.BucketCreate) bucketModel.Bucket {
	tbl := New()
	empty := bucketModel.Bucket{}
	empty.ID = uuid.NewString()

	empty.Name = i.Name
	empty.Region = i.Region
	empty.Size = i.Size

	tbl.Buckets = append(tbl.Buckets, empty)
	return empty
}

func Delete(id string) bool {
	tbl := New()
	response := false
	index := -1
	for i, v := range tbl.Buckets {
		if v.ID == id {
			index = i
		}
	}
	if index >= 0 {
		tbl.Buckets = append(tbl.Buckets[:index], tbl.Buckets[index+1:]...)
		response = true
	}
	return response
}

func Update(replace bucketModel.Bucket) bucketModel.Bucket {
	tbl := New()
	empty := bucketModel.Bucket{}
	empty.ID = replace.ID
	index := -1
	for i, v := range tbl.Buckets {
		if v.ID == replace.ID {
			index = i
			empty.Name = replace.Name
			empty.Region = replace.Region
			empty.Size = replace.Size
		}
	}

	if index >= 0 {
		tbl.Buckets = append(tbl.Buckets[:index], tbl.Buckets[index+1:]...)
		tbl.Buckets = append(tbl.Buckets, empty)
	}

	return empty
}

func Read(id string) []bucketModel.Bucket {

	tbl := New()
	if id == "" {
		return tbl.Buckets
	}

	fmt.Println("len(id) id ", len(id), id)

	empty := []bucketModel.Bucket{}
	for _, v := range tbl.Buckets {
		if v.ID == id {
			empty = append(empty, v)
		}
	}
	return empty
}

func New() *bucketTable {
	if table == nil {
		table = &bucketTable{
			Buckets: []bucketModel.Bucket{},
		}
	}

	return table
}
