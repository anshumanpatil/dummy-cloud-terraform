package network

import (
	networkModel "api/models/network"
	"fmt"

	"github.com/google/uuid"
)

type networkTable struct {
	Networks []networkModel.Network
}

var table *networkTable

func Create(i networkModel.NetworkCreate) networkModel.NetworkCreate {
	tbl := New()
	empty := networkModel.Network{}
	empty.ID = uuid.NewString()

	empty.Name = i.Name

	empty.IPList = i.IPList
	empty.InstanceList = i.InstanceList
	empty.IsActive = i.IsActive

	tbl.Networks = append(tbl.Networks, empty)
	i.ID = empty.ID
	return i
}

func Delete(id string) bool {
	tbl := New()
	response := false
	index := -1
	for i, v := range tbl.Networks {
		if v.ID == id {
			index = i
		}
	}
	if index >= 0 {
		tbl.Networks = append(tbl.Networks[:index], tbl.Networks[index+1:]...)
		response = true
	}
	return response
}

func Update(replace networkModel.Network) networkModel.Network {
	tbl := New()
	empty := networkModel.Network{}
	empty.ID = replace.ID
	index := -1
	for i, v := range tbl.Networks {
		if v.ID == replace.ID {
			index = i
			empty.Name = replace.Name
			empty.IPList = replace.IPList
			empty.InstanceList = replace.InstanceList
			empty.IsActive = replace.IsActive
		}
	}

	if index >= 0 {
		tbl.Networks = append(tbl.Networks[:index], tbl.Networks[index+1:]...)
		tbl.Networks = append(tbl.Networks, empty)
	}

	return empty
}

func Read(id string) []networkModel.Network {

	tbl := New()
	if id == "" {
		return tbl.Networks
	}

	fmt.Println("len(id) id ", len(id), id)

	empty := []networkModel.Network{}
	for _, v := range tbl.Networks {
		if v.ID == id {
			empty = append(empty, v)
		}
	}
	return empty
}

func New() *networkTable {
	if table == nil {
		table = &networkTable{
			Networks: []networkModel.Network{},
		}
	}

	return table
}
