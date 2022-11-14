FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . ./
COPY main/server/main.go ./

RUN go get -u github.com/cosmtrek/air

ENTRYPOINT ["air"]

RUN go build -o /server

EXPOSE 8080

CMD [ "/server" ]