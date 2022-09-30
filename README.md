# Basic Backend ( Cloud Function ) for The Blog System
## Introduction
- Uses `GCP  Cloud Functions` platform and `Golang` as it's primary language.
- Deployed at [BlogsAPI](https://us-central1-leafy-ai.cloudfunctions.net/BlogsAPI/)
- Depends on Authenticator.

## Setup
- Clone Repository
- Install dependencies by running `go get`
- Create a `.env` file at project root
- After Following the steps to properly set up the `.env` file , please execute ./source to run the program locally.
- Build the application by running `go build ./cmd`
- Execute the binary Application `./cmd`
- The application will be running on `localhost:{PORT}` [Where PORT is the port number specified in the .env file]

## Environment Variables
- `port` : optional -> The port number the application will run on. Defaults to `:8080`
- `FUNCTION_TARGET` : required -> The name of the function to execute
- `GIN_MODE` : required -> The mode the application will run in. `debug` or `release`
- `SECRET_KEY` : required -> microservice wide secret key for jwt auth
- `TOKEN_HOUR_LIFESPAN` : required -> For generation of token , numeric value, set as 48 [hours]

## Dependencies
- GCP Cloud Function Platform
- Golang
- FireStore Database
- gin/gonic
- jwt-go

## Implemented EndPoints:
- `GET /` : Returns a welcome message
- `GET /blogs` : Returns the paginated Feed [AUTH OPTIONAL] [TODO]
- `GET /blogs/{id}` : Returns a blog with the specified id [TODO]
- `POST /blogs/create` : Creates a new blog [AUTH REQUIRED]
- `GET /blogs/all` : RETURNS ALL THE BLOGS IN THE DATABASE