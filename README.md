# 九州工業大学 オープンキャンパス 2023 年 C3 企画
## 環境構築
1. Dockerのインストール
- [https://www.docker.com/products/docker-desktop/](https://www.docker.com/products/docker-desktop/)ここからDockerDesktopをインストールしてください.
2. このリポジトリをクローン
- ```git clone git@github.com:Kyutech-C3/OC_2023_backend.git```
3. Dockerで起動
- `docker compose up --build`
- --buildは初回だけでOK
- -d オプションでバックグラウンド実行が可能

## 開発の手引き
-   ブランチ

    -   ブランチのルート

        ブランチは`develop`ブランチから生やす


    -   ブランチの命名規則

        ブランチは基本以下のように命名する
        `[feature | fix]/[issue番号]-[issue内容]`
        例えば issue 番号が`#14`で issue の内容が`headerの作成`だった場合以下のようになる

        `feature/#14-header`
        また不具合修正の場合`feature`の部分が`fix`となる

-   コミットメッセージ
    コミットメッセージは以下のように記述する

    `[add | change | fix | delete]:[メッセージ] [issue番号]`

    例えば issue 番号が`#14`で issue の内容が`headerの作成`だった場合以下のようになる

    `add:ヘッダーの追加 #14`

    機能追加の場合は`add`,機能変更の場合は`change`,不具合修正の場合は`fix`,機能削除の場合は`delete`となる。

    またメッセージの部分はわかりやすく簡潔にする。
