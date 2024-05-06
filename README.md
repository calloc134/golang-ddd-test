# golang-ddd-test

## 概要

DDD の勉強を元に、Go 言語でバックエンドの API を作成しています。

## ディレクトリ構成

```

root
├── cmd
│   ├── api
│   │   ├── graph: graphqlのスキーマとリゾルバ定義
│   │   └── main.go: APIのエントリーポイント
│   └── migrate
│       ├── migrations: マイグレーションファイル
│       ├── schemas: DBのスキーマ定義
│       └── main.go: マイグレーションのエントリーポイント
└── src
     ├── domain: ドメイン層
     ├── application: アプリケーション層
     └── repositories: リポジトリ層

```

ディレクトリ構成変えたほうがいいか？わからない

## 採用技術

- Go
  - gqlgen
  - uptrace/bun

## 改善点

issue 参照
