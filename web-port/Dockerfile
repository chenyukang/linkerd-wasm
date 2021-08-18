FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN mkdir -p ./static
COPY static ./static/
RUN go build -o ./web_port

EXPOSE 8085
CMD [ "./web_port" ]