package network

type Network struct {
	ID           string                  `json:"id"`
	Name         string                  `json:"name"`
	IPList       []string                `json:"iplist"`
	InstanceList []NetworkInstanceMember `json:"instancelist"`
	IsActive     bool                    `json:"isactive"`
}

type NetworkInstanceMember struct {
	Name   string `json:"name"`
	Region string `json:"region"`
}

type NetworkCreate struct {
	Name         string                  `json:"name"`
	IPList       []string                `json:"iplist"`
	InstanceList []NetworkInstanceMember `json:"instancelist"`
	IsActive     bool                    `json:"isactive"`
}

type NetworkUpdate struct {
	Name         string                  `json:"name"`
	IPList       []string                `json:"iplist"`
	InstanceList []NetworkInstanceMember `json:"instancelist"`
	IsActive     bool                    `json:"isactive"`
}

type NetworkRead struct {
	ID string `json:"id"`
}

type NetworkDelete struct {
	ID string `json:"id"`
}
