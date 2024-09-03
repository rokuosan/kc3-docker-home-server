# Step:05 自作アプリケーションのデプロイをしてみよう

これまではすでにあるソフトウェアと配信されているイメージを使用してサーバーを立ててきました。
今回のステップでは自作アプリケーションのデプロイのために、独自イメージを使ってサーバーを立ててみましょう。

独自イメージを作成して効率的にサーバーを立てるには Dockerfile を作成して Compose に組み込みます。

## Dockerfileの作成

Dockerfile はイメージを作成するためのファイルです。Dockerfile はこのままの名前でファイルを作成します。拡張子はありません。

Dockerfile の詳細な書き方については触れませんが、以下のような内容のものを記述します。

```dockerfile
FROM python:3.11.2

ARG UID=1000
ARG USERNAME=admin
RUN useradd -m -u ${UID} admin && gpasswd -a ${USERNAME} staff

WORKDIR /usr/src/app

RUN pip install --no-cache-dir -r requirements.txt

RUN chown -R ${USERNAME}:${USERNAME} .

USER ${UID}

CMD ["/usr/bin/python3" "main.py"]
```

今回は`05-dockerfile/backend/Dockerfile` にすでに作成済みの Dockerfile があるのでそれを利用します。

## サービスを起動する

今回は`compose.yaml` が`05-dockerfile` ディレクトリにあります。この Compose ファイルを使ってサービスを起動してみましょう。

サーバーしたサーバーは`localhost:8080` で接続を待ち受けているので、そのアドレスへリクエストを飛ばしてみてください。
