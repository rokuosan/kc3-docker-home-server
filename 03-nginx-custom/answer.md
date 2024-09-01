# 解答例

解答例ですので、オプションの位置などは違っても構いません。

## 1. バックグラウンドで動作させる

ホストマシンの`localhost:8080` から接続できるようなコンテナを実行するには以下のようにします。

```shell
docker run --rm -d -p 8080:80 nginx:latest
```


## 2. nginx のコンテナにマウントする

nginx のコンテナで任意のファイルを配信するには以下のようにします。リポジトリのルートから以下のコマンドを実行してください。

```shell
docker run --rm -v ./03-nginx-custom/data:/usr/share/nginx/html -d -p 8080:80 nginx:latest
```

実行したのちに、`localhost:8080` にアクセスすると前回とは違ったものが表示されているはずです。

## 3. nginx のコンフィグを書いてマウントしてみよう

前のセクションと同じような形になっていますが、徐々にコマンドが長くなってきます。

```shell
docker run --rm -d -p 8080:80 -v ./03-nginx-custom/data:/usr/share/nginx/html -v ./03-nginx-custom/config:/etc/nginx/conf.d -v ./03-nginx-custom/secret/.htpasswd:/etc/nginx/.htpasswd nginx:latest
```
