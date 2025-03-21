## Requirements/dependencies
- Docker
- Docker-compose

## Getting Started

- Environment variables

```sh
make init
```

- Starting API in development mode

```sh
make up
```

- Run tests in container

```sh
make test
```

- Run tests local (it is necessary to have golang installed)

```sh
make test-local
```

- Run coverage report

```sh
make test-report
```
```sh
make test-report-func
```

- View logs

```sh
make logs
```

## API Request

| Endpoint        | HTTP Method           | Description       |
| --------------- | :---------------------: | :-----------------: |
| `/v1/accounts` | `POST`                | `Create accounts` |
| `/v1/accounts` | `GET`                 | `List accounts`   |
| `/v1/accounts/{{account_id}}/balance`   | `GET`                |    `Find balance account` |
| `/v1/health`| `GET`                 | `Health check`  |

## Test endpoints API using curl

- #### Creating new account

`Request`
```bash
curl -i --request POST 'http://localhost:3001/v1/accounts' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Test",
    "cpf": "070.910.584-24",
    "balance": 100
}'
```

`Response`
```json
{
    "id":"5cf59c6c-0047-4b13-a118-65878313e329",
    "name":"Test",
    "cpf":"070.910.584-24",
    "balance":1,
    "created_at":"2020-11-02T14:50:46Z"
}
```
- #### Listing accounts

`Request`
```bash
curl -i --request GET 'http://localhost:3001/v1/accounts'
```

`Response`
```json
[
    {
        "id": "5cf59c6c-0047-4b13-a118-65878313e329",
        "name": "Test",
        "cpf": "070.910.584-24",
        "balance": 1,
        "created_at": "2020-11-02T14:50:46Z"
    }
]
```

- #### Fetching account balance

`Request`
```bash
curl -i --request GET 'http://localhost:3001/v1/accounts/{{account_id}}/balance'
```

`Response`
```json
{
    "balance": 1
}
```

## Git workflow
- Gitflow

## Code status
- Development

## localstck setting
```bash
gentaro@gentaro-SVT13119FJS:~/dev_somu_script$ aws configure --profile localstack
AWS Access Key ID [None]: dummy
AWS Secret Access Key [None]: dummy
Default region name [None]: ap-northeast-1
Default output format [None]: json
```
