# Development Docker file 
# we will have two stage docker file for production 
FROM golang:1.23-alpine

# Set the Current Working Directory inside the container
WORKDIR /app 

RUN apk add --no-cache make

# air 
RUN go install github.com/air-verse/air@latest

# in production we will only copy go binary files but for now this is fine
COPY go.* ./

# in prod we build
RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# this is for migrations but will prb rm soon bc i want to use go-migrate in code
# RUN apk add curl

# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz


EXPOSE 8080

# now run "make air"
CMD ["make", "watch"]


