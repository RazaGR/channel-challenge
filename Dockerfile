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

# We need maller image
FROM alpine:3.15 
RUN apk add ca-certificates

COPY --from=build_base /tmp/pensionera-app/out/pensionera-app /app/pensionera-app

# Need mobgodb inside? let's not, comment it out

# RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/main' >> /etc/apk/repositories
# RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/community' >> /etc/apk/repositories
# RUN  apk update
# RUN  apk add openrc
# RUN  apk add mongodb
# RUN  apk add mongodb-tools
# RUN mkdir -p /data/db/
# RUN chown root /data/db
# RUN chmod 777 /data/db
# RUN rc-update add mongodb default
# RUN rc-service mongodb start


# Expose the app
EXPOSE 8080
# Run 
CMD ["/app/pensionera-app"]