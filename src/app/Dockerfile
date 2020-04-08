# ベースとなるDockerイメージ指定
FROM golang:1.13.9-alpine3.10

# コンテナ内に作業ディレクトリを作成
RUN mkdir /app

# コンテナログイン時のディレクトリ指定
WORKDIR /app

# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /app
