FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
# get gin lib
RUN go get github.com/gin-gonic/gin
#copy src and db.json file
COPY src/ ./src/
COPY src/db.json ./db.json

# **compiler index.go**
RUN go build -o index ./src/index.go
# define port 3001
EXPOSE 3001

# **run index.go**
CMD ["./index"]
