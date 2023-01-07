package bucket

type Bucket struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Size   string `json:"size"`
	Region string `json:"region"`
}

type BucketCreate struct {
	Name   string `json:"name" validate:"required"`
	Size   string `json:"size"`
	Region string `json:"region"`
}

type BucketUpdate struct {
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
