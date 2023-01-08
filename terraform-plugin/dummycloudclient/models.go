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
