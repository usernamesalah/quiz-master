# lightweight container for go
FROM golang:alpine

# update container's packages and install git
RUN apk update && apk add --no-cache git

# set /app to be the active directory
WORKDIR /github.com/usernamesalah/quiz-master

# copy all files from outside container, into the container
COPY . .

# download dependencies
RUN go mod tidy -v
RUN go get -v ./...
RUN go get github.com/swaggo/swag/cmd/swag
# RUN go get -u github.com/swaggo/swag/cmd/swag

# generate swagger docs
RUN rm -rf api/v1/docs
RUN swag init -o api/v1/docs

# build binary
RUN go build -o github.com/usernamesalah/quiz-master

# set the entry point of the binary
ENTRYPOINT ["/github.com/usernamesalah/quiz-master"]