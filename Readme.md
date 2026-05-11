# sail - 給与明細集計  


## 開発環境の立ち上げ
必須ソフトウェア: VS Code, Docker Desktop
VS Code に Remote - Containers 拡張のインストールが必要です。
`mu.code-workspace` を開き、VS Code 上で `Ctrl+Shift+P` を押して `Reopen in Container` を選ぶと開発環境を起動できます。

## 使用ライブラリ
### poppler-utils
  https://texwiki.texjp.org/?Poppler#k603e696

### popplerのPDF→テキスト 変換コマンド
> pdftotext [pdfファイル] -opw [パスワード]

## プロジェクトの初期構築コマンド

### create new front react project
> npm create vite@latest 
> yarn add bulma

### create new server golang project
> mkdir api
> cd api
> go mod init github.com/multios12/sail
> go get "github.com/glebarez/sqlite" "gorm.io/gorm"

### goreleaserのインストール
>  wget https://github.com/goreleaser/goreleaser/releases/download/v1.22.1/goreleaser_1.22.1_x86_64.apk
>  apk add --allow-untrusted --no-network --repositories-file=repo.list goreleaser_1.22.1_x86_64.apk

git tag -a v0.8.10 -m ''; git push origin --tags
