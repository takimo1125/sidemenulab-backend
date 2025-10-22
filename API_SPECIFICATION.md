# 📋 サイドメニュー研究所 API 仕様書

## 基本情報

- **ベース URL**: `http://localhost:8080/api/v1`
- **認証方式**: JWT Bearer Token
- **データ形式**: JSON
- **文字エンコーディング**: UTF-8

---

## 🔐 認証 API

### ユーザー登録

```http
POST /api/v1/auth/signup
```

**リクエストボディ:**

```json
{
  "email": "user@example.com",
  "password": "password123",
  "name": "ユーザー名"
}
```

**レスポンス:**

```json
{
  "message": "ユーザー登録が完了しました",
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "name": "ユーザー名",
      "created_at": "2025-10-22T14:21:36.795536007Z",
      "updated_at": "2025-10-22T14:21:36.795536007Z"
    },
    "token": {
      "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "expires_at": "2025-10-23T14:21:36.806781783Z",
      "token_type": "Bearer"
    }
  }
}
```

### ユーザーログイン

```http
POST /api/v1/auth/signin
```

**リクエストボディ:**

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**レスポンス:**

```json
{
  "message": "サインインに成功しました",
  "data": {
    "user": {
      "id": 1,
      "email": "user@example.com",
      "name": "ユーザー名",
      "created_at": "2025-10-22T14:21:36.795536007Z",
      "updated_at": "2025-10-22T14:21:36.795536007Z"
    },
    "token": {
      "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "expires_at": "2025-10-23T14:21:36.806781783Z",
      "token_type": "Bearer"
    }
  }
}
```

---

## 🏪 店舗管理 API

### 店舗一覧取得

```http
GET /api/v1/stores
```

**レスポンス:**

```json
{
  "data": [
    {
      "id": 1,
      "name": "サイドメニュー研究所 本店",
      "address": "東京都渋谷区",
      "phone": "03-1234-5678",
      "created_at": "2025-10-22T14:22:51.915685351Z",
      "updated_at": "2025-10-22T14:22:51.915685351Z",
      "deleted_at": null
    }
  ]
}
```

### 店舗詳細取得

```http
GET /api/v1/stores/:id
```

**パラメータ:**

- `id` (number): 店舗 ID

**レスポンス:**

```json
{
  "data": {
    "id": 1,
    "name": "サイドメニュー研究所 本店",
    "address": "東京都渋谷区",
    "phone": "03-1234-5678",
    "created_at": "2025-10-22T14:22:51.915685351Z",
    "updated_at": "2025-10-22T14:22:51.915685351Z",
    "deleted_at": null
  }
}
```

### 店舗作成

```http
POST /api/v1/stores
```

**リクエストボディ:**

```json
{
  "name": "サイドメニュー研究所 本店",
  "address": "東京都渋谷区",
  "phone": "03-1234-5678"
}
```

**バリデーション:**

- `name`: 必須、文字列
- `address`: 任意、文字列
- `phone`: 任意、文字列

**レスポンス:**

```json
{
  "message": "店舗が作成されました",
  "data": {
    "id": 1,
    "name": "サイドメニュー研究所 本店",
    "address": "東京都渋谷区",
    "phone": "03-1234-5678",
    "created_at": "2025-10-22T14:22:51.915685351Z",
    "updated_at": "2025-10-22T14:22:51.915685351Z",
    "deleted_at": null
  }
}
```

---

## 🍽️ サイドメニュー管理 API

### サイドメニュー一覧取得

```http
GET /api/v1/side-menus
```

**レスポンス:**

```json
{
  "data": [
    {
      "id": 1,
      "store_id": 1,
      "store": {
        "id": 1,
        "name": "サイドメニュー研究所 本店",
        "address": "東京都渋谷区",
        "phone": "03-1234-5678",
        "created_at": "2025-10-22T14:22:51.915685Z",
        "updated_at": "2025-10-22T14:22:51.915685Z",
        "deleted_at": null
      },
      "name": "特製サラダ",
      "description": "新鮮な野菜とドレッシングのサラダ",
      "price": 580,
      "created_at": "2025-10-22T14:23:04.706037Z",
      "updated_at": "2025-10-22T14:23:04.706037Z",
      "deleted_at": null
    }
  ]
}
```

### サイドメニュー詳細取得

```http
GET /api/v1/side-menus/:id
```

**パラメータ:**

- `id` (number): サイドメニュー ID

**レスポンス:**

```json
{
  "data": {
    "id": 1,
    "store_id": 1,
    "store": {
      "id": 1,
      "name": "サイドメニュー研究所 本店",
      "address": "東京都渋谷区",
      "phone": "03-1234-5678",
      "created_at": "2025-10-22T14:22:51.915685Z",
      "updated_at": "2025-10-22T14:22:51.915685Z",
      "deleted_at": null
    },
    "name": "特製サラダ",
    "description": "新鮮な野菜とドレッシングのサラダ",
    "price": 580,
    "created_at": "2025-10-22T14:23:04.706037Z",
    "updated_at": "2025-10-22T14:23:04.706037Z",
    "deleted_at": null
  }
}
```

### 店舗別サイドメニュー一覧取得

```http
GET /api/v1/side-menus/store/:storeId
```

**パラメータ:**

- `storeId` (number): 店舗 ID

**レスポンス:**

```json
{
  "data": [
    {
      "id": 1,
      "store_id": 1,
      "store": {
        "id": 1,
        "name": "サイドメニュー研究所 本店",
        "address": "東京都渋谷区",
        "phone": "03-1234-5678",
        "created_at": "2025-10-22T14:22:51.915685Z",
        "updated_at": "2025-10-22T14:22:51.915685Z",
        "deleted_at": null
      },
      "name": "特製サラダ",
      "description": "新鮮な野菜とドレッシングのサラダ",
      "price": 580,
      "created_at": "2025-10-22T14:23:04.706037Z",
      "updated_at": "2025-10-22T14:23:04.706037Z",
      "deleted_at": null
    }
  ]
}
```

### サイドメニュー作成

```http
POST /api/v1/side-menus
```

**リクエストボディ:**

```json
{
  "store_id": 1,
  "name": "特製サラダ",
  "description": "新鮮な野菜とドレッシングのサラダ",
  "price": 580
}
```

**バリデーション:**

- `store_id`: 必須、数値（存在する店舗 ID）
- `name`: 必須、文字列
- `description`: 任意、文字列
- `price`: 任意、数値

**レスポンス:**

```json
{
  "message": "サイドメニューが作成されました",
  "data": {
    "id": 1,
    "store_id": 1,
    "store": {
      "id": 1,
      "name": "サイドメニュー研究所 本店",
      "address": "東京都渋谷区",
      "phone": "03-1234-5678",
      "created_at": "2025-10-22T14:22:51.915685Z",
      "updated_at": "2025-10-22T14:22:51.915685Z",
      "deleted_at": null
    },
    "name": "特製サラダ",
    "description": "新鮮な野菜とドレッシングのサラダ",
    "price": 580,
    "created_at": "2025-10-22T14:23:04.706037Z",
    "updated_at": "2025-10-22T14:23:04.706037Z",
    "deleted_at": null
  }
}
```

---

## 📝 レビュー管理 API

### レビュー一覧取得

```http
GET /api/v1/reviews
```

**クエリパラメータ:**

- `side_menu_id` (number): サイドメニュー ID でフィルタリング
- `user_id` (number): ユーザー ID でフィルタリング

**レスポンス:**

```json
{
  "data": [
    {
      "id": 1,
      "side_menu_id": 1,
      "side_menu": {
        "id": 1,
        "store_id": 1,
        "store": {
          "id": 1,
          "name": "サイドメニュー研究所 本店",
          "address": "東京都渋谷区",
          "phone": "03-1234-5678",
          "created_at": "2025-10-22T14:22:51.915685Z",
          "updated_at": "2025-10-22T14:22:51.915685Z",
          "deleted_at": null
        },
        "name": "特製サラダ",
        "description": "新鮮な野菜とドレッシングのサラダ",
        "price": 580,
        "created_at": "2025-10-22T14:23:04.706037Z",
        "updated_at": "2025-10-22T14:23:04.706037Z",
        "deleted_at": null
      },
      "user_id": 1,
      "user": {
        "id": 1,
        "email": "user@example.com",
        "name": "ユーザー名",
        "created_at": "2025-10-22T14:21:36.795536007Z",
        "updated_at": "2025-10-22T14:21:36.795536007Z"
      },
      "rating": 5,
      "title": "とても美味しかった！",
      "comment": "新鮮な野菜で、ドレッシングも絶品でした。また食べたいです。",
      "is_verified": true,
      "created_at": "2025-10-22T15:00:00.000000Z",
      "updated_at": "2025-10-22T15:00:00.000000Z",
      "deleted_at": null
    }
  ]
}
```

### レビュー詳細取得

```http
GET /api/v1/reviews/:id
```

**パラメータ:**

- `id` (number): レビュー ID

**レスポンス:**

```json
{
  "data": {
    "id": 1,
    "side_menu_id": 1,
    "side_menu": {
      "id": 1,
      "store_id": 1,
      "store": {
        "id": 1,
        "name": "サイドメニュー研究所 本店",
        "address": "東京都渋谷区",
        "phone": "03-1234-5678",
        "created_at": "2025-10-22T14:22:51.915685Z",
        "updated_at": "2025-10-22T14:22:51.915685Z",
        "deleted_at": null
      },
      "name": "特製サラダ",
      "description": "新鮮な野菜とドレッシングのサラダ",
      "price": 580,
      "created_at": "2025-10-22T14:23:04.706037Z",
      "updated_at": "2025-10-22T14:23:04.706037Z",
      "deleted_at": null
    },
    "user_id": 1,
    "user": {
      "id": 1,
      "email": "user@example.com",
      "name": "ユーザー名",
      "created_at": "2025-10-22T14:21:36.795536007Z",
      "updated_at": "2025-10-22T14:21:36.795536007Z"
    },
    "rating": 5,
    "title": "とても美味しかった！",
    "comment": "新鮮な野菜で、ドレッシングも絶品でした。また食べたいです。",
    "is_verified": true,
    "created_at": "2025-10-22T15:00:00.000000Z",
    "updated_at": "2025-10-22T15:00:00.000000Z",
    "deleted_at": null
  }
}
```

### サイドメニュー別レビュー一覧取得

```http
GET /api/v1/reviews/side-menu/:sideMenuId
```

**パラメータ:**

- `sideMenuId` (number): サイドメニュー ID

**レスポンス:**

```json
{
  "data": [
    {
      "id": 1,
      "side_menu_id": 1,
      "side_menu": {
        "id": 1,
        "store_id": 1,
        "store": {
          "id": 1,
          "name": "サイドメニュー研究所 本店",
          "address": "東京都渋谷区",
          "phone": "03-1234-5678",
          "created_at": "2025-10-22T14:22:51.915685Z",
          "updated_at": "2025-10-22T14:22:51.915685Z",
          "deleted_at": null
        },
        "name": "特製サラダ",
        "description": "新鮮な野菜とドレッシングのサラダ",
        "price": 580,
        "created_at": "2025-10-22T14:23:04.706037Z",
        "updated_at": "2025-10-22T14:23:04.706037Z",
        "deleted_at": null
      },
      "user_id": 1,
      "user": {
        "id": 1,
        "email": "user@example.com",
        "name": "ユーザー名",
        "created_at": "2025-10-22T14:21:36.795536007Z",
        "updated_at": "2025-10-22T14:21:36.795536007Z"
      },
      "rating": 5,
      "title": "とても美味しかった！",
      "comment": "新鮮な野菜で、ドレッシングも絶品でした。また食べたいです。",
      "is_verified": true,
      "created_at": "2025-10-22T15:00:00.000000Z",
      "updated_at": "2025-10-22T15:00:00.000000Z",
      "deleted_at": null
    }
  ]
}
```

### レビュー作成

```http
POST /api/v1/reviews
```

**リクエストボディ:**

```json
{
  "side_menu_id": 1,
  "rating": 5,
  "title": "とても美味しかった！",
  "comment": "新鮮な野菜で、ドレッシングも絶品でした。また食べたいです。"
}
```

**バリデーション:**

- `side_menu_id`: 必須、数値（存在するサイドメニュー ID）
- `rating`: 必須、数値（1-5 の範囲）
- `title`: 任意、文字列
- `comment`: 任意、文字列

**レスポンス:**

```json
{
  "message": "レビューが作成されました",
  "data": {
    "id": 1,
    "side_menu_id": 1,
    "side_menu": {
      "id": 1,
      "store_id": 1,
      "store": {
        "id": 1,
        "name": "サイドメニュー研究所 本店",
        "address": "東京都渋谷区",
        "phone": "03-1234-5678",
        "created_at": "2025-10-22T14:22:51.915685Z",
        "updated_at": "2025-10-22T14:22:51.915685Z",
        "deleted_at": null
      },
      "name": "特製サラダ",
      "description": "新鮮な野菜とドレッシングのサラダ",
      "price": 580,
      "created_at": "2025-10-22T14:23:04.706037Z",
      "updated_at": "2025-10-22T14:23:04.706037Z",
      "deleted_at": null
    },
    "user_id": 1,
    "user": {
      "id": 1,
      "email": "user@example.com",
      "name": "ユーザー名",
      "created_at": "2025-10-22T14:21:36.795536007Z",
      "updated_at": "2025-10-22T14:21:36.795536007Z"
    },
    "rating": 5,
    "title": "とても美味しかった！",
    "comment": "新鮮な野菜で、ドレッシングも絶品でした。また食べたいです。",
    "is_verified": false,
    "created_at": "2025-10-22T15:00:00.000000Z",
    "updated_at": "2025-10-22T15:00:00.000000Z",
    "deleted_at": null
  }
}
```

### レビュー画像アップロード

```http
POST /api/v1/reviews/:id/images
```

**パラメータ:**

- `id` (number): レビュー ID

**リクエストボディ:**

```json
{
  "image_url": "https://example.com/image.jpg",
  "image_order": 1
}
```

**バリデーション:**

- `image_url`: 必須、文字列（有効な URL）
- `image_order`: 任意、数値（表示順序）

**レスポンス:**

```json
{
  "message": "レビュー画像がアップロードされました",
  "data": {
    "id": 1,
    "review_id": 1,
    "image_url": "https://example.com/image.jpg",
    "image_order": 1,
    "created_at": "2025-10-22T15:05:00.000000Z"
  }
}
```

### レビュー画像一覧取得

```http
GET /api/v1/reviews/:id/images
```

**パラメータ:**

- `id` (number): レビュー ID

**レスポンス:**

```json
{
  "data": [
    {
      "id": 1,
      "review_id": 1,
      "image_url": "https://example.com/image1.jpg",
      "image_order": 1,
      "created_at": "2025-10-22T15:05:00.000000Z"
    },
    {
      "id": 2,
      "review_id": 1,
      "image_url": "https://example.com/image2.jpg",
      "image_order": 2,
      "created_at": "2025-10-22T15:06:00.000000Z"
    }
  ]
}
```

### レビューにイイネ

```http
POST /api/v1/reviews/:id/like
```

**パラメータ:**

- `id` (number): レビュー ID

**レスポンス:**

```json
{
  "message": "レビューにイイネしました",
  "data": {
    "id": 1,
    "review_id": 1,
    "user_id": 2,
    "created_at": "2025-10-22T15:10:00.000000Z"
  }
}
```

### レビューのイイネ取り消し

```http
DELETE /api/v1/reviews/:id/like
```

**パラメータ:**

- `id` (number): レビュー ID

**レスポンス:**

```json
{
  "message": "レビューのイイネを取り消しました"
}
```

### レビューのイイネ一覧取得

```http
GET /api/v1/reviews/:id/likes
```

**パラメータ:**

- `id` (number): レビュー ID

**レスポンス:**

```json
{
  "data": [
    {
      "id": 1,
      "review_id": 1,
      "user_id": 2,
      "user": {
        "id": 2,
        "email": "user2@example.com",
        "name": "ユーザー2",
        "created_at": "2025-10-22T14:30:00.000000Z",
        "updated_at": "2025-10-22T14:30:00.000000Z"
      },
      "created_at": "2025-10-22T15:10:00.000000Z"
    }
  ]
}
```

---

## 🏥 ヘルスチェック API

### ヘルスチェック

```http
GET /health
```

**レスポンス:**

```json
{
  "status": "healthy",
  "database": "connected"
}
```

### API 情報取得

```http
GET /
```

**レスポンス:**

```json
{
  "message": "Sidemenulab Backend API",
  "status": "running"
}
```

---

## 📊 エラーレスポンス

### バリデーションエラー (400)

```json
{
  "error": "Key: 'CreateStoreRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
}
```

### 認証エラー (401)

```json
{
  "error": "メールアドレスまたはパスワードが正しくありません"
}
```

### リソース未発見 (404)

```json
{
  "error": "指定された店舗が見つかりません: record not found"
}
```

### サーバーエラー (500)

```json
{
  "error": "店舗の作成に失敗しました: duplicate key value violates unique constraint"
}
```

---

## 🔧 使用例

### cURL での使用例

**店舗作成:**

```bash
curl -X POST http://localhost:8080/api/v1/stores \
  -H "Content-Type: application/json" \
  -d '{"name":"サイドメニュー研究所 本店","address":"東京都渋谷区","phone":"03-1234-5678"}'
```

**サイドメニュー作成:**

```bash
curl -X POST http://localhost:8080/api/v1/side-menus \
  -H "Content-Type: application/json" \
  -d '{"store_id":1,"name":"特製サラダ","description":"新鮮な野菜とドレッシングのサラダ","price":580}'
```

**サイドメニュー一覧取得:**

```bash
curl -X GET http://localhost:8080/api/v1/side-menus
```

**レビュー作成:**

```bash
curl -X POST http://localhost:8080/api/v1/reviews \
  -H "Content-Type: application/json" \
  -d '{"side_menu_id":1,"rating":5,"title":"とても美味しかった！","comment":"新鮮な野菜で、ドレッシングも絶品でした。"}'
```

**レビュー一覧取得:**

```bash
curl -X GET http://localhost:8080/api/v1/reviews
```

**レビューにイイネ:**

```bash
curl -X POST http://localhost:8080/api/v1/reviews/1/like
```

---

## 📝 注意事項

1. **認証**: 現在の実装では認証は必須ではありませんが、本番環境では認証を必須にすることを推奨します
2. **CORS**: フロントエンドからのアクセスを許可するため、CORS 設定が有効です
3. **データベース**: PostgreSQL を使用し、GORM でマイグレーションが自動実行されます
4. **ソフトデリート**: 削除されたデータは`deleted_at`フィールドで管理されます
5. **日時形式**: すべての日時は ISO 8601 形式（UTC）で返却されます

---

## 📋 データベーススキーマ

### stores テーブル

| カラム名   | データ型     | 制約                        | 説明                       |
| ---------- | ------------ | --------------------------- | -------------------------- |
| id         | uint         | PRIMARY KEY, AUTO_INCREMENT | 店舗 ID                    |
| name       | varchar(255) | NOT NULL, UNIQUE            | 店舗名                     |
| address    | varchar(500) | NULL                        | 住所                       |
| phone      | varchar(20)  | NULL                        | 電話番号                   |
| created_at | timestamp    | NOT NULL                    | 作成日時                   |
| updated_at | timestamp    | NOT NULL                    | 更新日時                   |
| deleted_at | timestamp    | NULL                        | 削除日時（ソフトデリート） |

### side_menus テーブル

| カラム名    | データ型      | 制約                        | 説明                       |
| ----------- | ------------- | --------------------------- | -------------------------- |
| id          | uint          | PRIMARY KEY, AUTO_INCREMENT | サイドメニュー ID          |
| store_id    | uint          | NOT NULL, FOREIGN KEY       | 店舗 ID                    |
| name        | varchar(255)  | NOT NULL                    | サイドメニュー名           |
| description | text          | NULL                        | 説明文                     |
| price       | decimal(10,2) | NULL                        | 価格（円）                 |
| created_at  | timestamp     | NOT NULL                    | 作成日時                   |
| updated_at  | timestamp     | NOT NULL                    | 更新日時                   |
| deleted_at  | timestamp     | NULL                        | 削除日時（ソフトデリート） |

### users テーブル

| カラム名   | データ型     | 制約                        | 説明                       |
| ---------- | ------------ | --------------------------- | -------------------------- |
| id         | uint         | PRIMARY KEY, AUTO_INCREMENT | ユーザー ID                |
| email      | varchar(255) | NOT NULL, UNIQUE            | メールアドレス             |
| password   | varchar(255) | NOT NULL                    | パスワード（ハッシュ化）   |
| name       | varchar(255) | NOT NULL                    | ユーザー名                 |
| created_at | timestamp    | NOT NULL                    | 作成日時                   |
| updated_at | timestamp    | NOT NULL                    | 更新日時                   |
| deleted_at | timestamp    | NULL                        | 削除日時（ソフトデリート） |

### side_menu_reviews テーブル

| カラム名     | データ型     | 制約                        | 説明                       |
| ------------ | ------------ | --------------------------- | -------------------------- |
| id           | uint         | PRIMARY KEY, AUTO_INCREMENT | レビュー ID                |
| side_menu_id | uint         | NOT NULL, FOREIGN KEY       | サイドメニュー ID          |
| user_id      | uint         | NOT NULL, FOREIGN KEY       | ユーザー ID                |
| rating       | int          | NOT NULL, CHECK (1-5)       | 評価（1-5 の星評価）       |
| title        | varchar(255) | NULL                        | レビュータイトル           |
| comment      | text         | NULL                        | レビューコメント           |
| is_verified  | boolean      | NOT NULL, DEFAULT false     | 購入確認済みフラグ         |
| created_at   | timestamp    | NOT NULL                    | 作成日時                   |
| updated_at   | timestamp    | NOT NULL                    | 更新日時                   |
| deleted_at   | timestamp    | NULL                        | 削除日時（ソフトデリート） |

### side_menu_review_images テーブル

| カラム名    | データ型     | 制約                        | 説明        |
| ----------- | ------------ | --------------------------- | ----------- |
| id          | uint         | PRIMARY KEY, AUTO_INCREMENT | 画像 ID     |
| review_id   | uint         | NOT NULL, FOREIGN KEY       | レビュー ID |
| image_url   | varchar(500) | NOT NULL                    | 画像 URL    |
| image_order | int          | NOT NULL, DEFAULT 0         | 表示順序    |
| created_at  | timestamp    | NOT NULL                    | 作成日時    |

### side_menu_review_likes テーブル

| カラム名   | データ型  | 制約                        | 説明        |
| ---------- | --------- | --------------------------- | ----------- |
| id         | uint      | PRIMARY KEY, AUTO_INCREMENT | イイネ ID   |
| review_id  | uint      | NOT NULL, FOREIGN KEY       | レビュー ID |
| user_id    | uint      | NOT NULL, FOREIGN KEY       | ユーザー ID |
| created_at | timestamp | NOT NULL                    | 作成日時    |

---

## 🚀 開発・デプロイ

### ローカル開発環境の起動

```bash
# Docker Composeで起動
docker-compose up --build

# ヘルスチェック
curl http://localhost:8080/health
```

### 環境変数

- `DATABASE_URL`: PostgreSQL 接続文字列
- `JWT_SECRET`: JWT 署名用の秘密鍵
- `PORT`: サーバーポート（デフォルト: 8080）

この API 仕様書を参考に、フロントエンドの実装を進めてください。
