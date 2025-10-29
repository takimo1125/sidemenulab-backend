FROM golang:1.23-alpine
WORKDIR /app
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/air-verse/air@latest
CMD ["air"]