# Step:04 Docker Compose を使ってみよう

前回のステップで静的サイトなら立てられそうな雰囲気を感じてきました。
しかしながら、何度もあの長ったらしいコマンドを打つのは面倒です。そこで登場するのが Docker Compose です。

Docker Compose では前回までに実行していたコマンドを YAML 形式のファイルで保存し、1つのコマンドでコンテナを操作できます。

## Docker Compose のファイルを書いてみる

Docker Compose のファイルは`compose.yaml` という名前で作成します。
カレントディレクトリにそのファイル名を定義しておくと、`docker compose` コマンドを実行した際に`-f` オプションをつける必要がなくなるというメリットが存在します。
その他の名前でも構いませんが、以下のリストに示す名前以外を利用した場合カレントディレクトリにあったとしても`-f` オプションを使用する必要があります。

- `-f` オプションが不要なファイル名
  - `compose.yaml` (推奨)
  - `compose.yml`
  - `docker-compose.yaml`
  - `docker-compose.yml`

今回は以下の内容を定義した`compose.yaml` がこの README と同じディレクトリ(`./04-compose`)にあります。

```yaml
services:
  web:
    image: nginx:latest
    ports:
      - "8080:80"
    volumes:
      - ../03-nginx-custom/data:/usr/share/nginx/html
      - ../03-nginx-custom/config:/etc/nginx/conf.d
      - ../03-nginx-custom/secret/.htpasswd:/etc/nginx/.htpasswd
```

## サービスの起動

Docker Compose では定義した各コンテナのことをサービスと呼びます。
上記の例の場合、`services` に`web` という名前で登録しているため、この nginx のコンテナは`web` というサービスとして扱われます。

上記のファイルを使ってサービスを起動するために、一度`04-compose` ディレクトリに移動してから以下のコマンドを実行してください。
このとき、前回作成したコンテナが削除されていることを確認してから実行してください。（「ポートが使用中です」といったエラーが発生するため）

``` shell
cd 04-compose
docker compose up -d
```

実行すると以下のような表示が得られるでしょう。

```
$ docker compose up -d
[+] Running 2/2
 ✔ Network 04-compose_default    Created
 ✔ Container 04-compose-web-1  Started
```

この状態で`localhost:8080` に接続すると、前回に作成した Web ページがそのまま表示されているはずです。

起動の際に指定した`-d` オプションは`docker run` と同様に、バックグラウンドでの動作を意味します。

## ログを見てみる

実行されているコンテナのログを見る場合は以下のコマンドを実行します。

```bash
docker compose logs
```

`-f` オプションを利用することで、継続的にログを見ることができます。見終わった時は`CTRL+C` で終了できます。


## サービスの停止

サービスを停止する場合は以下のコマンドを実行します。

```shell
docker compose down
```

## サービスを増やしてみる

nginx のサービスをもう1つ起動してみましょう。

`compose.yaml` の内容を一度全て消去し、以下の内容に上書きしてください。

```yaml
services:
  web:
    image: nginx:latest
    ports:
      - "8080:80"
    volumes:
      - ../03-nginx-custom/data:/usr/share/nginx/html
      - ../03-nginx-custom/config:/etc/nginx/conf.d
      - ../03-nginx-custom/secret/.htpasswd:/etc/nginx/.htpasswd

  web2:
    image: nginx:latest
    ports:
      - "8081:80"
    volumes:
      - ../03-nginx-custom/data:/usr/share/nginx/html
      - ../03-nginx-custom/config:/etc/nginx/conf.d
      - ../03-nginx-custom/secret/.htpasswd:/etc/nginx/.htpasswd
```

対応するポート番号が異なるだけの nginx のコンテナを`web2` というサービスとして定義しました。

この状態で`docker compose up -d` を実行します。

実行後に`localhost:8080` と`localhost:8081` の両方で同じコンテンツが帰ってきていることを確認しましょう。

このように定義を書くだけでお手軽にサービスを増やすことができます。
