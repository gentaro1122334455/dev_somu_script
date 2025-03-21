#!/bin/bash
region=ap-northeast-1
endpoint=http://0.0.0.0:4566

aws dynamodb create-table --endpoint-url $endpoint --region $region --table-name user_data --cli-input-json file:///json/userData.json