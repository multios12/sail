# sail - 給与明細集計  


## 開発環境の立ち上げ
必須ソフトウェア：VS Code, Docker Desktop
VS Codeに[Remote - Containers]Extensionのインストールが必要
Vscode上でCTRL+SHIFT+P押下、[Reopen in Container]選択で開発環境の立ち上げが可能

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
> mkdir srv
> cd srv
> go mod init github.com/multios12/sail
> go get "github.com/glebarez/sqlite" "gorm.io/gorm"

### goreleaserのインストール
>  wget https://github.com/goreleaser/goreleaser/releases/download/v1.22.1/goreleaser_1.22.1_x86_64.apk
>  apk add --allow-untrusted --no-network --repositories-file=repo.list goreleaser_1.22.1_x86_64.apk

git tag -a v0.8.10 -m ''; git push origin --tags