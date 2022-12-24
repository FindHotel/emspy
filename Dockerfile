FROM golang:1.19-alpine as builder

ARG gh_token

RUN echo "https://dl-cdn.alpinelinux.org/alpine/v3.11/community" >> /etc/apk/repositories

RUN apk --no-cache add ca-certificates git gcc libc-dev upx

WORKDIR /emspy

COPY go.mod .
COPY go.sum .

RUN go mod download


COPY cmd/ ./cmd
COPY internal/ ./internal
# COPY pkg/ ./pkg

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=readonly -ldflags='-s -w' -o /app/emspy ./cmd/emspy

RUN upx -q /app/emspy

# final stage
FROM alpine:3.17

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/emspy /usr/local/bin

WORKDIR /emspy
CMD emspy
