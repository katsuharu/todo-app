FROM golang:1.20

ENV TZ="Asia/Tokyo"
WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@v1.42.0

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -o ./bin/app ./cmd/todo-app

CMD [ "./bin/app" ]
