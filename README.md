# Go API Example
[![Build Status](https://drone.nathanmalishev.com/api/badges/nathanmalishev/go_api_example/status.svg)](https://drone.nathanmalishev.com/nathanmalishev/go_api_example)
## Overview
Being new to golang there wasn't a whole lot of complete examples out there. This aims to be a nice restfull golang api, that can get you started with.  

## Running
Make sure the environment variables `JWT_SECRET`, `MONGO_USERNAME`, `MONGO_PASSWORD` are set.  
The other configs are configured in `common/config.json` at the moment.  
To run the API use the following command  
`MONGO_USERNAME= MONGO_PASSWORD= JWT_SECRET="super_secret_pls_change_me" go run router.go main.go`

## Testing
To run all of the tests run `JWT_SECRET="test" go test ./...`


## Design choices
Throughout this API i tried to avoid global objects & make each function as functional as possible.  
Trying to have every function, explicitly have everything it needs as arguments.  

### Router.go
I like having one router file were i can easily see everysingle route.

### Main.go
I like seeing everything that my app depends on, only using init in the common package to start running configs.


## Contribute
Please make PR's contribute and leave feedback
