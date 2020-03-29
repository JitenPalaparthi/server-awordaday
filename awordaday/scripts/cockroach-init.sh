#!/bin/bash
/cockroach/cockroach.sh start --join=node_1 --listen-addr=localhost --insecure
/cockroach/cockroach.sh user set $COCKROACH_USER --insecure -u root
/cockroach/cockroach.sh sql -e "CREATE DATABASE $COCKROACH_DB;" --insecure -u root
/cockroach/cockroach.sh sql -e "GRANT ALL ON DATABASE recipes TO $COCKROACH_USER;" --insecure -u root
/cockroach/cockroach.sh sql -e "CREATE TABLE IF NOT EXISTS  words
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    word STRING unique NOT NULL,
    meaning STRING NOT NULL,
    type STRING NOT NULL,
    status STRING NOT NULL DEFAULT 'NOT-ACTIVE',
    last_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by STRING NOT NULL
);" --insecure -u root
/cockroach/cockroach.sh sql -e "
CREATE TABLE IF NOT EXISTS  sentences
(
     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
     word_id UUID NOT NULL REFERENCES words (id),
     sentence STRING NOT NULL,
     status STRING NOT NULL DEFAULT 'NOT-ACTIVE',
    last_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by STRING NOT NULL
);" --insecure -u root
/cockroach/cockroach.sh sql -e "Create Table IF NOT EXISTS request_words
(   
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    word STRING NOT NULL,
    status STRING NOT NULL DEFAULT 'Created',
    last_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    requested_by STRING NOT NULL
);" --insecure -u root
/cockroach/cockroach.sh sql -e "CREATE TABLE IF NOT EXISTS  audits
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    data STRING NOT NULL,
    ip STRING NOT NULL,
    device STRING NOT NULL,
    url_path STRING NOT NULL ,
    headers STRING NOT NULL ,
    date_time TIMESTAMP NOT NULL
);" --insecure -u root
