FROM golang:1.24.4 AS builder
## コンテナでの作業を/appディレクトリに設定
WORKDIR /app
# Goのキャッシュを利用し、go.modとgo.sumを先にコピー
COPY go.mod go.sum ./
# Goの依存関係をダウンロード
RUN go mod download
# ソースコードをコンテナにコピー
COPY ../ .
# ビルドを実行
RUN go build -o main cmd/main.go
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["/app/main"]
