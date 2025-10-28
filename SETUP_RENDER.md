# Render デプロイガイド

## 📋 概要

このガイドでは、SideMenuLab Backend を Render にデプロイする手順を説明します。

## 🚀 デプロイ手順

### 1. Render アカウントの作成

1. [Render](https://render.com)にアクセス
2. GitHub アカウントでサインアップ

### 2. データベースの作成

1. Render ダッシュボードで「New」→「PostgreSQL」を選択
2. 以下の設定を行う：
   - **Name**: `sidemenulab-db`
   - **Database**: `sidemenulab`
   - **User**: `sidemenulab_user`
   - **Plan**: Free
   - **Region**: Ohio (または任意のリージョン)
3. 「Create Database」をクリック

### 3. Web サービスの作成

1. Render ダッシュボードで「New」→「Web Service」を選択
2. GitHub リポジトリを選択
3. 以下の設定を行う：
   - **Name**: `sidemenulab-backend`
   - **Runtime**: `Go`
   - **Region**: Ohio (データベースと同じリージョン)
   - **Branch**: `main`
   - **Root Directory**: (空白のまま)
   - **Build Command**: `go build -o main .`
   - **Start Command**: `./main`
   - **Plan**: Free

### 4. 環境変数の設定

Web サービスの「Environment」タブで以下の環境変数を設定：

#### 必須の環境変数

```
GIN_MODE=release
```

```
DATABASE_URL=<PostgreSQLデータベースのInternal Database URL>
```

データベースの「Connections」タブから「Internal Database URL」をコピーします。

```
JWT_SECRET=<ランダムな文字列>
```

#### Cloudinary 設定（任意）

```
CLOUDINARY_CLOUD_NAME=ducbz8w43
CLOUDINARY_API_KEY=<あなたのAPIキー>
CLOUDINARY_API_SECRETUND=<あなたのシークレット>
```

#### 自動設定される環境変数

- `PORT`: Render が自動的に設定（デフォルト: 10000）

### 5. デプロイ

1. 「Save Changes」をクリック
2. Render が自動的にデプロイを開始します
3. ビルドログを確認して、エラーがないか確認します

## 🔧 設定ファイルの説明

### render.yaml

Render Blueprint を使用する場合、`render.yaml`を使用して自動セットアップできます。

使用する場合：

1. Render ダッシュボードで「New」→「Blueprint」を選択
2. GitHub リポジトリを選択
3. 自動的に設定が適用されます

### Dockerfile.production

本番環境用の Dockerfile です。マルチステージビルドを使用して軽量なイメージを作成します。

Render では通常、Dockerfile は不要ですが、カスタムビルドが必要な場合に使用できます。

## 📝 環境変数の管理

### ローカル開発環境

`.env.local`ファイルを使用（gitignore されています）：

```bash
cp env.example .env.local
# .env.localを編集
```

### 本番環境（Render）

Render のダッシュボードまたは`render.yaml`で管理します。

## 🔍 デプロイ後の確認

### ヘルスチェック

```bash
curl https://your-app.onrender.com/health
```

期待されるレスポンス：

```json
{
  "status": "healthy",
  "database": "connected"
}
```

### データベース接続確認

Render の PostgreSQL ダッシュボードの「Connect」タブで接続を確認できます。

## 🐛 トラブルシューティング

### ビルドエラー

- Go モジュールの依存関係が正しくインストールされているか確認
- `go.mod`と`go.sum`がリポジトリに含まれているか確認

### データベース接続エラー

- `DATABASE_URL`が正しく設定されているか確認
- データベースが同じリージョンにあるか確認
- SSL モードの設定を確認（Render の PostgreSQL は SSL 必須）

### タイムアウトエラー

- Free プランでは、アプリケーションが一定時間アクセスされないとスリープします
- 最初のリクエスト後に起動します（約 30 秒）

## 📚 参考資料

- [Render Documentation](https://render.com/docs)
- [Go on Render](https://render.com/docs/deploy-go)
- [PostgreSQL on Render](https://render.com/docs/databases)
