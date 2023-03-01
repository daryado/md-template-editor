FROM golang:latest AS build-stage
WORKDIR /app
ENV GOPATH=/
COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client


RUN go mod download
RUN go build -o auth-server ./cmd/main.go


RUN chmod +x wait-for-postgres.sh
CMD [ "./auth-server" ]



