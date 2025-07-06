
FROM golang:1.24.4 AS builder
## コンテナでの作業を/appディレクトリに設定
WORKDIR /app
# Goのキャッシュを利用し、go.modとgo.sumを先にコピー
COPY go.mod go.sum ../
# Goの依存関係をダウンロード
RUN go mod download

# --- ステージ 2: dev ---
# 開発用ステージ air環境を構築
FROM builder AS dev
# airをインストール
RUN go install github.com/air-verse/air@latest
# ソースコードをコピー
COPY . .
# airを実行
CMD ["air", "-c", ".air.toml"]
# ソースコードをコンテナにコピー
COPY ../ .
# ビルドを実行
RUN go build -o main cmd/main.go

# --- ステージ 3: build ---
FROM alpine:latest AS production

WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["/app/main"]
