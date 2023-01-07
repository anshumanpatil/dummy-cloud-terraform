package instance

import (
	instanceModel "api/models/instance"
	"fmt"

	"github.com/google/uuid"
)

type instanceTable struct {
	Instances []instanceModel.Instance
}

var table *instanceTable

func Create(i instanceModel.InstanceCreate) instanceModel.InstanceCreate {
	tbl := New()
	empty := instanceModel.Instance{}
	empty.ID = uuid.NewString()

	empty.Name = i.Name
	empty.OS = i.OS
	empty.Ram = i.Ram
	empty.Region = i.Region
	empty.Size = i.Size

	tbl.Instances = append(tbl.Instances, empty)
	return i
}

func Delete(id string) bool {
	tbl := New()
	response := false
	index := -1
	for i, v := range tbl.Instances {
		if v.ID == id {
			index = i
		}
	}
	if index >= 0 {
		tbl.Instances = append(tbl.Instances[:index], tbl.Instances[index+1:]...)
		response = true
	}
	return response
}

func Update(replace instanceModel.Instance) instanceModel.Instance {
	tbl := New()
	empty := instanceModel.Instance{}
	empty.ID = replace.ID
	index := -1
	for i, v := range tbl.Instances {
		if v.ID == replace.ID {
			index = i
			empty.Name = replace.Name
			empty.OS = replace.OS
			empty.Ram = replace.Ram
			empty.Region = replace.Region
			empty.Size = replace.Size
		}
	}

	if index >= 0 {
		tbl.Instances = append(tbl.Instances[:index], tbl.Instances[index+1:]...)
		tbl.Instances = append(tbl.Instances, empty)
	}

	return empty
}

func Read(id string) []instanceModel.Instance {

	tbl := New()
	if id == "" {
		return tbl.Instances
	}

	fmt.Println("len(id) id ", len(id), id)

	empty := []instanceModel.Instance{}
	for _, v := range tbl.Instances {
		if v.ID == id {
			empty = append(empty, v)
		}
	}
	return empty
}

func New() *instanceTable {
	if table == nil {
		table = &instanceTable{
			Instances: []instanceModel.Instance{},
		}
	}

	return table
}
