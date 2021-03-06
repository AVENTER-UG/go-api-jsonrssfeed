FROM golang:alpine as builder

WORKDIR /build

COPY . /build/

RUN apk add git && \
    go get -d

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.MinVersion=`date -u +%Y%m%d%.H%M%S` -extldflags \"-static\"" -o main app.go init.go


FROM alpine
LABEL maintainer="Andreas Peters <support@aventer.biz>"

ENV API_PORT=10777 
ENV API_BIND=0.0.0.0
ENV LOGLEVEL=info

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/main /app/

EXPOSE 10777

WORKDIR "/app"

CMD ["./main"]