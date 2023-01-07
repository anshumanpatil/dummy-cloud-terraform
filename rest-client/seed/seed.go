package seed

import (
	buckettable "api/database/bucket"
	instancetable "api/database/instance"
	bucketModel "api/models/bucket"
	instanceModel "api/models/instance"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Seed() {
	instance()
	bucket()
}

func instance() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	jsonFile, err := os.Open("./seed/instance.json")
	if err != nil {
		log.Err(err)
	}
	log.Print("Successfully Opened instance.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var instances []instanceModel.InstanceCreate
	if err := json.Unmarshal(byteValue, &instances); err == nil {
		for _, v := range instances {
			created := instancetable.Create(v)
			log.Print("instance created - ", created)
		}
	}
	defer jsonFile.Close()
}

func bucket() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	jsonFile, err := os.Open("./seed/bucket.json")
	if err != nil {
		log.Err(err)
	}
	log.Print("Successfully Opened bucket.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var buckets []bucketModel.BucketCreate
	if err := json.Unmarshal(byteValue, &buckets); err == nil {
		for _, v := range buckets {
			created := buckettable.Create(v)
			log.Print("bucket created - ", created)
		}
	}
	defer jsonFile.Close()
}
