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
> create-react-app sail --template typescript\
> yarn add bulma react-router-dom 

### create new server golang project
> mkdir srv
> cd srv
> go mod init github.com/multios12/sail
> go get "github.com/glebarez/sqlite" "gorm.io/gorm"
