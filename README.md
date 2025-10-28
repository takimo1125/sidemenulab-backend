# SideMenuLab Backend

SideMenuLab のバックエンド API（Go + Gin + GORM + PostgreSQL）

## 🚀 ローカル開発環境のセットアップ

### 必要要件

- Go 1.25 以上
- Docker & Docker Compose
- PostgreSQL 14 以上

### セットアップ手順

1. **リポジトリのクローン**

   ```bash
   git clone https://github.com/yourusername/sidemenulab-backend.git
   cd sidemenulab-backend
   ```

2. **環境変数の設定**

   ```bash
   cp env.example .env.local
   # .env.localを編集して必要な設定を行う
   ```

3. **Docker Compose でデータベースを起動**

   ```bash
   docker-compose up -d
   ```

4. **依存パッケージのインストール**

   ```bash
   go mod download
   ```

5. **アプリケーションの起動**

   ```bash
   go run main.go
   ```

   Air を使ったホットリロードの場合：

   ```bash
   air
   ```

6. **API の確認**
   ```bash
   curl http://localhost:8080/health
   ```

## 🌐 Render へのデプロイ

### 方法 1: Render Web UI を使用

1. Render のダッシュボードから「New Web Service」を選択
2. GitHub リポジトリを接続
3. 以下の設定を行う：

   - **Build Command**: `go build -o main .`
   - **Start Command**: `./main`
   - **Environment**: `go`
   - **Region**: `ohio` (または任意のリージョン)

4. 環境変数を設定：
   - `DATABASE_URL`: PostgreSQL データベースの接続文字列
   - `JWT_SECRET`: JWT 署名用の秘密鍵
   - `CLOUDINARY_CLOUD_NAME`: Cloudinary クラウド名
   - `CLOUDINARY_API_KEY`: Cloudinary API キー
   - `CLOUDINARY_API_SECRET`: Cloudinary API シークレット
   - `GIN_MODE`: `release`
   - `PORT`: `10000` (Render が自動設定)

### 方法 2: render.yaml を使用

1. リポジトリの`render.yaml`を確認
2. Render のダッシュボードで「New Blueprint」を選択
3. リポジトリを接続すると自動的に設定が適用されます

## 📁 ディレクトリ構造

```
.
├── internal/
│   ├── delivery/          # HTTPレイヤー
│   │   └── http/
│   ├── domain/            # ドメイン層
│   │   ├── entity/       # エンティティ
│   │   └── repository/   # リポジトリインターフェース
│   ├── infrastructure/    # インフラ層
│   │   ├── cloudinary/   # Cloudinaryサービス
│   │   └── database/     # データベース実装
│   └── usecase/          # ユースケース層
├── main.go                # エントリーポイント
├── render.yaml           # Renderデプロイ設定
└── Dockerfile.production # プロダクション用Dockerfile
```

## 🔧 主要な機能

- JWT 認証（サインアップ/サインイン）
- レビュー管理（CRUD）
- レビューコメント
- レビュー画像アップロード（Cloudinary）
- レビューいいね機能
- データベースマイグレーション
- 初期データのシーディング

## 📚 API 仕様

詳細な API 仕様は [API_SPECIFICATION.md](./API_SPECIFICATION.md) を参照してください。

## 🛠️ 開発ツール

### Air (Hot Reload)

```bash
# Airのインストール
go install github.com/air-verse/air@latest

# Airの起動
air
```

### データベース接続確認

```bash
docker-compose exec postgres psql -U postgres -d sidemenulab
```

## 📝 環境変数

| 変数名                  | 説明                                | デフォルト        |
| ----------------------- | ----------------------------------- | ----------------- |
| `DATABASE_URL`          | PostgreSQL 接続文字列               | -                 |
| `JWT_SECRET`            | JWT 署名用の秘密鍵                  | `your-secret-key` |
| `CLOUDINARY_CLOUD_NAME` | Cloudinary クラウド名               | -                 |
| `CLOUDINARY_API_KEY`    | Cloudinary API キー                 | -                 |
| `CLOUDINARY_API_SECRET` | Cloudinary API シークレット         | -                 |
| `PORT`                  | サーバーポート                      | `8080`            |
| `GIN_MODE`              | Gin のモード (`debug` or `release`) | `debug`           |

## 🐛 トラブルシューティング

### データベース接続エラー

```bash
# Dockerコンテナの状態を確認
docker-compose ps

# コンテナを再起動
docker-compose restart
```

### ポートが既に使用されている

```bash
# PORT環境変数を変更するか、他のプロセスを停止
PORT=8081 go run main.go
```

## 📄 ライセンス

MIT
