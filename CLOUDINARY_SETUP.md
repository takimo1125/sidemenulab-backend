# Cloudinary 設定ガイド

このプロジェクトでは、画像アップロードに Cloudinary を使用しています。

## Cloudinary アカウントの作成

1. [Cloudinary](https://cloudinary.com/)にアクセスしてアカウントを作成
2. ダッシュボードにログイン
3. 以下の情報を取得：
   - Cloud Name
   - API Key
   - API Secret

## 環境変数の設定

`env.example`ファイルをコピーして`.env`ファイルを作成し、以下の値を設定してください：

```bash
# Cloudinary設定
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret
```

## 機能

- **自動画像最適化**: アップロード時に自動でフォーマット変換と品質最適化が行われます
- **セキュア URL**: HTTPS で配信されるセキュアな URL が使用されます
- **フォルダ管理**: 年別にフォルダが作成され、画像が整理されます
- **エラーハンドリング**: アップロード失敗時の自動クリーンアップ機能

## フォルダ構造

Cloudinary では以下のフォルダ構造で画像が保存されます：

```
sidemenulab/reviews/2024/review_1_1234567890_0.jpg
```

- `sidemenulab`: プロジェクト名
- `reviews`: レビュー画像用フォルダ
- `2024`: 年別フォルダ
- `review_1_1234567890_0`: レビュー ID*タイムスタンプ*インデックス

## トラブルシューティング

### Cloudinary が利用できない場合

環境変数が設定されていない場合、アプリケーションはローカルファイルアップロードにフォールバックします。ログで確認できます：

```
Cloudinaryの環境変数が設定されていないため、ローカルファイルアップロードを使用します
```

### アップロードエラー

- API キーとシークレットが正しく設定されているか確認
- Cloudinary のアカウント制限に達していないか確認
- ネットワーク接続を確認
