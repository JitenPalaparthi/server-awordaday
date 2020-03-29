#!/bin/bash
echo Wait for servers to be up
sleep 30

HOSTPARAMS="--host db-1 --insecure"
SQL="/cockroach/cockroach.sh sql $HOSTPARAMS"

$SQL -e "CREATE DATABASE a_word_a_day;"
$SQL -d a_word_a_day -e "CREATE TABLE IF NOT EXISTS  audits
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    data STRING NOT NULL,
    ip STRING NOT NULL,
    device STRING NOT NULL,
    url_path STRING NOT NULL ,
    headers STRING NOT NULL ,
    date_time STRING NOT NULL
);"