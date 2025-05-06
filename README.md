# gin_todo

## 使用している主な技術

- **バックエンド**: Go, Gin, gqlgen, Docker
- **フロントエンド**: Next.js, TypeScript, Tailwind CSS
- **データベース**: PostgreSQL
- **その他**: GraphQL, Docker Compose, Nginx

## プロジェクトの概要

このプロジェクトは、タスク管理アプリケーションのサンプル実装です。バックエンドは Go を使用して GraphQL API を提供し、フロントエンドは Next.js を使用してユーザーインターフェースを構築しています。Docker を利用して、バックエンド、フロントエンド、データベースをコンテナ化しています。

## ディレクトリ構成

```
gin_todo/
├── back/                     # バックエンド関連ファイル
│   ├── server.go             # サーバーのエントリーポイント
│   ├── gqlgen.yml            # gqlgenの設定ファイル
│   ├── graph/                # GraphQL関連のコード
│   │   ├── schema.graphqls   # GraphQLスキーマ
│   │   ├── resolver.go       # リゾルバ
│   │   └── middlewares/      # ミドルウェア
│   ├── internal/             # 内部ロジック
│   │   ├── entity/           # エンティティ定義
│   │   ├── infra/            # インフラ層（DB接続など）
│   │   └── usecase/          # ユースケース層
│   └── containers/           # Docker関連
│       └── nginx/            # Nginx設定
├── front/                    # フロントエンド関連ファイル
│   ├── src/
│   │   ├── app/              # app router
│   │   ├── components/       # shadcn/uiのコンポーネント
│   │   └── features/         # 特定の機能に関連するコード
│   ├── package.json          # npmパッケージ情報
│   └── next.config.ts        # Next.jsの設定
└── README.md                 # プロジェクトの説明
```
