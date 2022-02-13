FROM golang:1.17.7-alpine3.15 AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/pensionera-app

# populate the module cache
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build 
RUN go build -o ./out/pensionera-app .

# We need smaller image
FROM alpine:3.15 
RUN apk add ca-certificates

COPY --from=build_base /tmp/pensionera-app/out/pensionera-app /app/pensionera-app

# Expose the app
EXPOSE 8080
# Run 
CMD ["/app/pensionera-app"]