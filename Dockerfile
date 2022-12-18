FROM golang:1.19-buster as builder

COPY . /app

WORKDIR /app

RUN go build \
    -ldflags '-s -w -extldflags "-static"' \
    -tags osusergo,netgo,sqlite_omit_load_extension \
    -o /usr/local/bin/myapp ./cmd/app

FROM alpine

COPY --from=builder /usr/local/bin/myapp /usr/local/bin/myapp

RUN mkdir -p /database && touch /database/data.db

CMD ["myapp"]
