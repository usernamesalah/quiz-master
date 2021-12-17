# Quiz Master API DOCUMENTATION 

This repository contains source code for Quiz Master API.

## Getting Started

To run the project localy, make sure minimum requirements are fulfilled.

- Go version 1.10 or higher
- MYSQL or MariaDB


## How To Run this Project

>Since the project already use Go Module, I recommend to put the source code in any folder but GOPATH.

### Running Without Docker

1. Make sure Go is installed as global command (first time only)

2. Clone this project and go to the root project to install all dependencies (first time only)
    ```bash
    // copy project into your GOPATH
    > cd $GOPATH/src

    // change directory to root project folder
    > cd user-profile
    
    // install all the dependencies
    > make init   
    ```
3. Running your MYSQL
4. While still in root project build and run the app
    ```bash
    // source environmant setup
    > source .env
    
    // build and run project
    > make run

    // now go to http://localhost:8080/ in your browser to check the app.
    ```

### Run the testing

```bash
    make test
```

## API Documentation

We use [swag](https://github.com/swaggo/swag) to generate necearry Swagger files for API documentation. Everytime we run `make build`, the Swagger documentation will be updated.

To configure your API documentation, please refer to [swag's declarative comments format](https://github.com/swaggo/swag#declarative-comments-format) and [examples](https://github.com/swaggo/swag#examples).

To access the documentation, please visit [API DOCUMENTATION](http://localhost:8080/v1/docs/index.html).


## Repository Content

- **/app/controllers** contains controllers and API route registry
- **/pkg**
  - **/pkg/models** contains table models
  - **/pkg/repository** contains database transactions
 

## Tools Used

>In this project, I use some tools listed below

- All libraries listed in [go.mod](https://github.com/usernamesalah/soccer-api/tree/master/go.mod)
- ["github.com/vektra/mockery".](https://github.com/vektra/mockery) To Generate Mocks for testing needs.