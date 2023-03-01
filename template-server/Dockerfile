FROM golang:latest AS build

RUN go version
ENV GOPATH=/
COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client


RUN go mod download
RUN go build -o template-server ./cmd/main.go

#FROM alpine:3.16 AS product

#COPY --from=build template_server .
#COPY --from=build env .
#COPY --from=build migrations ./migrations
#COPY --from=build migrations ./migrations
#COPY --from=build wait-for-postgres.sh .
#RUN apk add --no-cache bash

RUN chmod +x wait-for-postgres.sh


CMD [ "./template-server" ]



