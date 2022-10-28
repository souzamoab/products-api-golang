FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . ./
COPY main.go ./

RUN go build -o /server

EXPOSE 8080

CMD [ "/server" ]