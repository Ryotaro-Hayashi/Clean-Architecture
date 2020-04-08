# ベースとなるDockerイメージ指定
FROM golang:1.14-alpine

# コンテナ内に作業ディレクトリを作成
RUN mkdir /api

# コンテナログイン時のディレクトリ指定
WORKDIR /api

# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /api
