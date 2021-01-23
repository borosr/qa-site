# QA site
A test project to try out CockroachDB

## Api docs
|Method|Path|Request|Response|
|---|---|---|---|
|GET|/api/status|none|Status 200, `{"Status":"ok"}`|
|POST|/api/users|`{"username":"user","password":"secret","full_name":"Sir John"}`|Status 200,`{"Msg":"user successfully created"}`|
|POST|/api/login|`{"username":"user","password":"secret"}`|Status 200, `{"token":"jwtToken","revoke_token":"revoke token","auth_kind":"DefaultLogin"}`|
|GET|/api/login/github|none|Status 301|
|GET|/api/login/github/callback|?code=the_code_from_github|Status 200, `{"token":"jwtToken","revoke_token":"revoke token","auth_kind":"Github"}`|
|GET|/api/users|none|Status 200, `[{"id":"user_id","username":"user","full_name":"Sir John"}]`|
|GET|/api/users/{user_id}|none|Status 200, `{"id":"user_id","username":"user","full_name":"Sir John"}`|
|PUT|/api/users/{user_id}|`{"username":"user","password":"secret2","full_name":"Sir John"}`|Status 200, `{"username":"user","password":"secret2","full_name":"Sir John"}`|
|DELETE|/api/users/{user_id}|none|Status 200, `{"Msg":"success"}`|
|GET|/api/questions|?limit=10&offset=0&sort=created_at|Status 200, `{"count":1,data":[{"id":"question_id","title":"short_teext","description":"long_text","created_by":"user_id","created_at":"2021-01-01T00:00:00.0Z","status":"published","rating":10}]}`|
|GET|/api/questions/{question_id}|none|Status 200, `{"id":"question_id","title":"short_teext","description":"long_text","created_by":"user_id","created_at":"2021-01-01T00:00:00.0Z","status":"published","rating":10}`|
|DELETE|/api/questions/{question_id}|none|Status 200, `{"Msg":"OK"}`|
|POST|/api/questions|`{"title":"short_text","description":"long_text"}`|Status 200, `{"id":"question_id","title":"short_teext","description":"long_text","created_by":"user_id","created_at":"2021-01-01T00:00:00.0Z","status":"published","rating":0}`|
|PUT|/api/questions/{question_id}|`{"title":"new_short_text","description":"new_long_text"}`|Status 200, `{"id":"question_id","title":"new_short_teext","description":"new_long_text","created_by":"user_id","created_at":"2021-01-01T00:00:00.0Z","status":"published","rating":0}`|
|GET|/api/questions/{question_id}/answers|none|Status 200, `[{"id":"answer_id","question_id":"question_id","created_by":"user_id","answer":"text","created_at":"2021-01-01T00:00:00.0Z","answered":false,"rating":0}]`|
|PUT|/api/questions/{question_id}/answers/{answer_id}/answered|none|Status 200, `{"id":"answer_id","question_id":"question_id","created_by":"user_id","answer":"text","created_at":"2021-01-01T00:00:00.0Z","answered":true,"rating":0}`|
|GET|/api/answers|none|Status 200, `[{"id":"answer_id","question_id":"question_id","created_by":"user_id","answer":"text","created_at":"2021-01-01T00:00:00.0Z","answered":true,"rating":0}]`|
|POST|/api/answers|`{"question_id":"question_id",answer":"text"}`|Status 200, `{"id":"answer_id","question_id":"question_id","created_by":"user_id","answer":"text","created_at":"2021-01-01T00:00:00.0Z","answered":true,"rating":0}`|
|PUT|/api/answers/{answer_id}|`{"answer":"text"}`|Status 200, `{"id":"answer_id","question_id":"question_id","created_by":"user_id","answer":"text","created_at":"2021-01-01T00:00:00.0Z","answered":true,"rating":0}`|
|PUT|/api/answers/{answer_id}/rate|none|Status 200, `{"value":1}`|
|PUT|/api/answers/{answer_id}/unrate|none|Status 200, `{"value":-1}`|
|PUT|/api/answers/{answer_id}/rate/dismiss|none|Status 200, `{"value":0}`|
|PUT|/api/questions/{question_id}/rate|none|Status 200, `{"value":1}`|
|PUT|/api/questions/{question_id}/unrate|none|Status 200, `{"value":-1}`|
|PUT|/api/questions/{question_id}/rate/dismiss|none|Status 200, `{"value":0}`|
|DELETE|/api/logout|none|Status 200, `{"Msg":"Logout success"}`|
|POST|/api/revoke|`{"revoke_token":"revoke token"}`|Status 200, `{"token":"jwtToken","revoke_token":"revoke token","auth_kind":"DefaultLogin"}`|
