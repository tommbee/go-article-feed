FROM golang:latest AS build-env
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o main .

# final stage
FROM alpine:3.7
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
WORKDIR /app
COPY --from=build-env /app /app/
EXPOSE 8080
ENTRYPOINT ["/app/main"]
