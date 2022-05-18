# Build the manager binary
FROM golang:1.18.1 as builder

# Copy in the go src
WORKDIR /go/src/db_agent
COPY . .

#
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GO111MODULE=on
RUN go mod tidy

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o db_agent cmd/main.go

# Copy the controller-manager into a thin image
FROM alpine:3.12.0
RUN apk add --no-cache tzdata
WORKDIR /root/
COPY --from=builder /go/src/db_agent/db_agent .
COPY docker-entrypoint.sh .

WORKDIR /root/conf
# COPY --from=builder /go/src/db_agent/conf .
RUN chmod +x /root/docker-entrypoint.sh

ENTRYPOINT  ["/root/docker-entrypoint.sh"]

# docker run -itd -v ./data/data.db:/data/data.db db_agent --restart=always
CMD ["/root/db_agent"]