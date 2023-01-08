package network

import (
	"dummy-cloud/dummycloudclient"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func getIpList(network dummycloudclient.Network) (*basetypes.ListValue, error) {
	iplistValues := []attr.Value{}

	for _, v := range network.IPList {
		iplistValues = append(iplistValues, types.StringValue(v))
	}

	newiplist, err := types.ListValue(types.StringType, iplistValues)
	if err != nil {
		return nil, fmt.Errorf("Error in conversion getIpList.")
	}
	return &newiplist, nil
}

func getInstanceList(network dummycloudclient.Network) (*basetypes.ListValue, error) {
	iplistValues := []attr.Value{}
	objType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":   types.StringType,
			"region": types.StringType,
		},
	}

	for _, v := range network.InstanceList {
		val := types.ObjectValueMust(
			map[string]attr.Type{
				"name":   types.StringType,
				"region": types.StringType,
			},
			map[string]attr.Value{
				"name":   types.StringValue(v.Name),
				"region": types.StringValue(v.Region),
			},
		)
		iplistValues = append(iplistValues, val)
	}

	newiplist, err := types.ListValue(objType, iplistValues)
	if err != nil {
		return nil, fmt.Errorf("Error in conversion getInstanceList.")
	}
	return &newiplist, nil
}
