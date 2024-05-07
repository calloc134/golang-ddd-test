# golang-ddd-test

## 概要

DDD の勉強を元に、Go 言語でバックエンドの API を作成しています。

## ディレクトリ構成

```
.
├── cmd
│   ├── api: API サーバー
│   │   ├── graph: 具体的なgraphqlサーバ実装
│   │   │   ├── schema.graphqls: GraphQLのスキーマ
│   │   │   ├── resolver.go: リゾルバのDIなど
│   │   │   ├── schema.resolvers.go: スキーマに対応するリゾルバ
│   │   │   └── model: gqlgenにより生成されたモデル
│   │   └── main.go: APIのエントリーポイント
│   └── migrate: マイグレーションツール
│       ├── migrations: マイグレーションファイル
│       ├── schemas: DBスキーマ
│       └── main.go: マイグレーションツールのエントリーポイント
├── src
│   ├── mutation: ミューテーションで参照される実装
│   │   ├── domain: ドメインモデル
│   │   ├── application: ユースケース
│   │   └── repository: リポジトリ
│   └── query: クエリで参照される実装
│       ├── view: ビューモデル
│       └── gateway: ゲートウェイ(dataloaderなど)
├── gqlgen.yml: gqlgen の設定ファイル
└── test.s3db: 生成されるSQLite3のデータベース(当リポジトリには含まれない)

```

## 採用技術

- Go
  - gqlgen
  - uptrace/bun

## 改善点

issue 参照
