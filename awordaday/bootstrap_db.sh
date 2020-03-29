#!/bin/bash
#docker network create -d bridge awordadaynet

#docker run -d --name=awordaday1 --hostname=awordaday1 --net=awordadaynet -p 26257:26257 -p 8080:8080  -v "${PWD}/data/awordaday1:/cockroach/data"  cockroachdb/cockroach:v19.2.5 start --insecure --join=awordaday1,awordaday2,awordaday3

#docker run -d --name=awordaday2 --hostname=awordaday2 --net=awordadaynet -v "${PWD}/data/awordaday2:/cockroach/data" cockroachdb/cockroach:v19.2.5 start --insecure --join=awordaday1,awordaday2,awordaday3

#docker run -d --name=awordaday3 --hostname=awordaday3 --net=awordadaynet -v "${PWD}/data/awordaday3:/cockroach/data" cockroachdb/cockroach:v19.2.5 start --insecure --join=awordaday1,awordaday2,awordaday3

docker-compose -f ./docker-compose.yml up --build -d

docker exec -it awordaday1 ./cockroach init --insecure
docker exec -ti awordaday1 \
sh -c "/cockroach/cockroach sql --insecure < scripts/scripts.sql"

docker-compose -f ./app.yml up --build -d

