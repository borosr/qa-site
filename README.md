# QA site

A test project to try out CockroachDB

## Requirements

- Go 1.17 or higher
- Docker or pre-installed cockroachdb

## Start project (locally)

```shell
cd infrastructure/local
./start.sh
cd ../..
go run main.go
```

## Settings

For custom configuration change settings.json, following options are available:

| Name               | Type   | Description                                                                                                                                                                                                             |
|--------------------|--------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| DBConnectionString | string | Database connection string to connect to CockroachDB                                                                                                                                                                    |
| Port               | string | Port to listen for the HTTP server                                                                                                                                                                                      |
| JwtHMAC            | string | JWT token's HMAC key                                                                                                                                                                                                    |
| JwtSecret          | string | JWT token secret, can be any random string                                                                                                                                                                              |
| RevokeHMAC         | string | Revoke token's HMAC key                                                                                                                                                                                                 |
| RevokeSecret       | string | Revoke token's secret, can be any random string                                                                                                                                                                         |
| Hostname           | string | Hostname of the instance                                                                                                                                                                                                |
| BadgerPath         | string | Path for the BadgerDB data files                                                                                                                                                                                        |
| Visibility         | string | Visibility defines who can read the questions and answers, can be `public` or `private`. Public means, no need for authentication to read Q&A. Private requires authentication to read Q&A. Default value is `private`. |

## Api docs

|Method| Path                                                      |Request| Response                                                                                                                                                                                             |
|---|-----------------------------------------------------------|---|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|GET| /api/status                                               |none| Status 200, `{"status":"ok"}`                                                                                                                                                                        |
|GET| /api/info                                                 |none| Status 200, `{"visibility":"private","oauth_providers":{"github": false}}`                                                                                                                                          |
|POST| /api/users                                                |`{"username":"user","password":"secret","full_name":"Sir John"}`| Status 200,`{"Msg":"user successfully created"}`                                                                                                                                                     |
|POST| /api/login                                                |`{"username":"user","password":"secret"}`| Status 200, `{"token":"jwtToken","revoke_token":"revoke token","auth_kind":"DefaultLogin"}`                                                                                                          |
|GET| /api/login/github                                         |none| Status 301                                                                                                                                                                                           |
|GET| /api/login/github/callback                                |?code=the_code_from_github| Status 200, `{"token":"jwtToken","revoke_token":"revoke token","auth_kind":"Github"}`                                                                                                                |
|GET| /api/users                                                |none| Status 200, `[{"id":"user_id","username":"user","full_name":"Sir John"}]`                                                                                                                            |
|GET| /api/users/{user_id}                                      |none| Status 200, `{"id":"user_id","username":"user","full_name":"Sir John"}`                                                                                                                              |
|PUT| /api/users/{user_id}                                      |`{"username":"user","password":"secret2","full_name":"Sir John"}`| Status 200, `{"username":"user","password":"secret2","full_name":"Sir John"}`                                                                                                                        |
|DELETE| /api/users/{user_id}                                      |none| Status 200, `{"Msg":"success"}`                                                                                                                                                                      |
|GET| /api/questions                                            |?limit=10&offset=0&sort=created_at| Status 200, `{"count":1,data":[{"id":"question_id","title":"short_teext","description":"long_text","created_by":"user_id","created_at":"2021-01-01T00:00:00.0Z","status":"published","rating":10}]}` |
|GET| /api/questions/{question_id}                              |none| Status 200, `{"id":"question_id","title":"short_teext","description":"long_text","created_by":"user_id","created_at":"2021-01-01T00:00:00.0Z","status":"published","rating":10}`                     |
|DELETE| /api/questions/{question_id}                              |none| Status 200, `{"Msg":"OK"}`                                                                                                                                                                           |
|POST| /api/questions                                            |`{"title":"short_text","description":"long_text"}`| Status 200, `{"id":"question_id","title":"short_teext","description":"long_text","created_by":"user_id","created_at":"2021-01-01T00:00:00.0Z","status":"published","rating":0}`                      |
|PUT| /api/questions/{question_id}                              |`{"title":"new_short_text","description":"new_long_text"}`| Status 200, `{"id":"question_id","title":"new_short_teext","description":"new_long_text","created_by":"user_id","created_at":"2021-01-01T00:00:00.0Z","status":"published","rating":0}`              |
|GET| /api/questions/{question_id}/answers                      |none| Status 200, `[{"id":"answer_id","question_id":"question_id","created_by":"user_id","answer":"text","created_at":"2021-01-01T00:00:00.0Z","answered":false,"rating":0}]`                              |
|PUT| /api/questions/{question_id}/answers/{answer_id}/answered |none| Status 200, `{"id":"answer_id","question_id":"question_id","created_by":"user_id","answer":"text","created_at":"2021-01-01T00:00:00.0Z","answered":true,"rating":0}`                                 |
|GET| /api/answers                                              |none| Status 200, `[{"id":"answer_id","question_id":"question_id","created_by":"user_id","answer":"text","created_at":"2021-01-01T00:00:00.0Z","answered":true,"rating":0}]`                               |
|POST| /api/answers                                              |`{"question_id":"question_id",answer":"text"}`| Status 200, `{"id":"answer_id","question_id":"question_id","created_by":"user_id","answer":"text","created_at":"2021-01-01T00:00:00.0Z","answered":true,"rating":0}`                                 |
|PUT| /api/answers/{answer_id}                                  |`{"answer":"text"}`| Status 200, `{"id":"answer_id","question_id":"question_id","created_by":"user_id","answer":"text","created_at":"2021-01-01T00:00:00.0Z","answered":true,"rating":0}`                                 |
|PUT| /api/answers/{answer_id}/rate                             |none| Status 200, `{"value":1}`                                                                                                                                                                            |
|PUT| /api/answers/{answer_id}/unrate                           |none| Status 200, `{"value":-1}`                                                                                                                                                                           |
|PUT| /api/answers/{answer_id}/rate/dismiss                     |none| Status 200, `{"value":0}`                                                                                                                                                                            |
|PUT| /api/questions/{question_id}/rate                         |none| Status 200, `{"value":1}`                                                                                                                                                                            |
|PUT| /api/questions/{question_id}/unrate                       |none| Status 200, `{"value":-1}`                                                                                                                                                                           |
|PUT| /api/questions/{question_id}/rate/dismiss                 |none| Status 200, `{"value":0}`                                                                                                                                                                            |
|DELETE| /api/logout                                               |none| Status 200, `{"Msg":"Logout success"}`                                                                                                                                                               |
|POST| /api/revoke                                               |`{"revoke_token":"revoke token"}`| Status 200, `{"token":"jwtToken","revoke_token":"revoke token","auth_kind":"DefaultLogin"}`                                                                                                          |
