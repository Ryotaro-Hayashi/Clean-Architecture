# ベースとなるDockerイメージ指定
FROM golang:1.14-alpine

# コンテナログイン時のディレクトリ指定
WORKDIR /CleanArchitecture/api

# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /CleanArchitecture/api

# GOMUDULE を有効化
ENV GO111MODULE=on

# モジュールインストール
RUN go mod download

EXPOSE 8080
