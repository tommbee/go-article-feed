FROM golang:latest AS build-env
ADD . /go/src/app/ 
WORKDIR /go/src/app
RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o main .

# final stage
FROM alpine:3.7
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
WORKDIR /app
COPY --from=build-env /app /app/
EXPOSE 8080
ENTRYPOINT ["/app/main"]
