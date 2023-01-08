package dummycloudclient

type Instance struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Size   string `json:"size"`
	Region string `json:"region"`
	Ram    string `json:"ram"`
	OS     string `json:"os"`
}

type InstanceRead struct {
	ID string `json:"id"`
}

type InstanceDelete struct {
	ID string `json:"id"`
}

type DeleteInstance struct {
	Deleted bool     `json:"deleted"`
	Data    Instance `json:"data"`
}

type Bucket struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Size   string `json:"size"`
	Region string `json:"region"`
}

type BucketRead struct {
	ID string `json:"id"`
}

type BucketDelete struct {
	ID string `json:"id"`
}

type DeleteBucket struct {
	Deleted bool   `json:"deleted"`
	Data    Bucket `json:"data"`
}

//----------------------------------------------------------------------------------------

type NetworkRegion struct {
	Name   string `json:"name"`
	Region string `json:"region"`
}

type Network struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	IPList       []string        `json:"iplist"`
	InstanceList []NetworkRegion `json:"instancelist"`
	IsActive     bool            `json:"isactive"`
}

type NetworkRead struct {
	ID string `json:"id"`
}

type NetworkDelete struct {
	ID string `json:"id"`
}

type DeleteNetwork struct {
	Deleted bool    `json:"deleted"`
	Data    Network `json:"data"`
}
