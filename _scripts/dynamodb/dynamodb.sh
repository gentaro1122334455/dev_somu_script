#!/bin/bash
region=ap-northeast-1
endpoint=http://localstack:4566

aws dynamodb create-table --endpoint-url $endpoint --region $region --table-name user_data --billing-mode PAY_PER_REQUEST --cli-input-json file://create_user_data.json