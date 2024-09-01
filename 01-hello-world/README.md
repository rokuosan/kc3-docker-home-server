# Step:01 Hello World をしてみよう

Docker がインストールされているか確認するためにも、一度 Docker のコマンドを叩いてみましょう。

以下のコマンドを実行することで、`hello-world` というイメージのコンテナが作成されて標準出力に`Hello from Docker!` という文字列が返ってくるでしょう。

```bash
docker run --rm hello-world:latest
```

実行例を以下に示します。1 行目はターミナルのプロンプトです。

```shell
node@b0bb7f13ac51:/workspaces/kc3-docker-home-server$ docker run --rm  hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (arm64v8)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```

## コマンド解説

今回は Docker でコンテナを動作させるために`docker run` コマンドを利用しました。（このコマンドは実はエイリアスで、本当のコマンドは`docker container run` です。）

このコマンドは以下のような文法になっています。括弧で囲われた部分については省略できます。

```bash
docker container run [オプション] イメージ名 [コマンド] [引数]
```

今回の例を当てはめると、`--rm` がオプションとなり、`hello-world:latest` がイメージ名になります。

この`--rm` オプションは、実行終了後にコンテナを削除するコマンドです。この`hello-world` イメージのコンテナは使用後に再度利用することがないので実行後に削除しています。
