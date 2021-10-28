# echo-trial
Personal trial of echo - light weight web framework

## install
```
git clone github.com/tenkoh/echo-trial
cd echo-trial
go mod tidy
go run main.go
```

If you want to shutdown the app, press `ctrl[^] + C`.

## methods

api root: `localhost:1323`

| URI | Method | description |
| --- | --- | --- |
| /user/:id | GET | return `id` |
| /users | POST | save user |
| /users | GET | show user list |

post example:
```
curl -F "name=John" http://localhost:1323/users
```