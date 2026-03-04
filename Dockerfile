FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache git

# install air
RUN go install github.com/air-verse/air@latest

ENV PATH="/go/bin:${PATH}"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 9000

CMD ["air"]