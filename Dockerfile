FROM golang
WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

COPY . .
# RUN go get
RUN go mod download
RUN go build -o /golang-apis


EXPOSE 8080