package instance

type Instance struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Size   string `json:"size"`
	Region string `json:"region"`
	Ram    string `json:"ram"`
	OS     string `json:"os"`
}

type InstanceCreate struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Size   string `json:"size"`
	Region string `json:"region"`
	Ram    string `json:"ram"`
	OS     string `json:"os"`
}

type InstanceUpdate struct {
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
