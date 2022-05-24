# Build the manager binary
FROM golang:1.18.1-alpine as builder

# Copy in the go src
WORKDIR /go/src/db_agent
COPY . .

RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.15/main/ > /etc/apk/repositories
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update
RUN apk add build-base

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GO111MODULE=on
RUN go mod tidy

# Build
RUN CGO_ENABLED=1 GOOS=linux go build --ldflags "-extldflags -static" -a -o db_agent cmd/main.go

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

# docker build --network=host -t db_agent . --no-cache
# docker run -itd -v /home/pi/data/data.db:/data/data.db --name=DB_AGENT --restart=always db_agent
# docker run -it -v /home/pi/data/data.db:/data/data.db  db_agent
CMD ["/root/db_agent"]