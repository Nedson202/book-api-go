FROM golang:alpine AS build-env
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR $GOPATH/src/github.com/user-service

COPY . .
RUN apk update -qq && apk add git

RUN go get -d -v

RUN go build -o main .

# RUN go get github.com/githubnemo/CompileDaemon

FROM alpine
WORKDIR /app

COPY --from=build-env /go/src/github.com/user-service /app

# ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o user-service" -command="./user-service"

EXPOSE 7000

CMD [ "main" ]
