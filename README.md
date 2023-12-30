# GO Api Boilerplate
![GitHub tag (with filter)](https://img.shields.io/github/v/tag/luigiescalante/go-api-template)
![GitHub last commit (branch)](https://img.shields.io/github/last-commit/luigiescalante/go-api-template/main)
![Static Badge](https://img.shields.io/badge/email-luigi.escalante%5Bat%5Dgmail.com-blue)
![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/luigi_escalante)
<p align="center">
<img src="github-logo.png" alt="logo" width="200" height="292">
</p>

A basic template schema to create an API with access to database with the basic tools preloaded for migrate,test and
document the api.

## Requirements
* GO V1.19
* Swagger V2.0
* Postgres V15.3
* Flyway
* Bearer token authorization
* SqlMock
* Testify
* PsqlX

## Environment vars
~~~~
DB_HOST=<database ip address>
DB_USER=<database user name>
DB_PASSWORD=<database user password>
DB_NAME=<database name>
APP_NAME=<app name>
APP_VERSION=<app version number>
PORT=<api port connection>
~~~~