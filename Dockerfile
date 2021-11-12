FROM golang:latest
WORKDIR /app
COPY . .
EXPOSE 4000
RUN go mod init example.com/m/v2
RUN go get -u github.com/gin-gonic/gin
CMD go run main.go