
# belajar-authentication

use existing libs :

- Chi Router

- Ozzo Validation, for input request validation

- Godotenv, for env loader

- Gorm, for ORM


## Setups
-  Clone Repo
```
$ 	git clone https://github.com/riyanda432/belajar-authentication.git
$ 	cd belajar-authentication
$ 	go mod tidy
   ```
- Install [Air](https://github.com/cosmtrek/air/blob/master/README.md) for Hot Reloading
- Install [soda-cli](https://gobuffalo.io/en/docs/db/toolbox) for migration tools
- Setup .env
- Setup database.yml
- [optional] Run sample migrations
```
$ 	soda migrate up
```
 - Running Application
 ```
$ air
 ```
 or
 ```
$ go run main.go
 ```
## Pre-Usage

Please install to verify unit test before commit

```
go install github.com/go-courier/husky/cmd/husky@latest
```

Then install git hooks

```
husky init
```