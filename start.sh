#!/bin/sh

# Author : Anshuman Patil
# rm -rfv cloud-console/node_modules
rm -rfv cloud-console/build
cd cloud-console
# npm install --verbose
npm run build
docker-compose -f single-kafka-single.yml up -d
docker-compose up -d
