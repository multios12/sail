# sail - 給与明細集計  

## 必要なライブラリ
### poppler-utils
  https://texwiki.texjp.org/?Poppler#k603e696

### popplerのインストール方法(Ubuntu)
下記のコマンドを実行
> sudo apt install poppler-utils poppler-data

### popplerのインストール方法(Windows)
1. 下記から、Windows用バイナリをダウンロード、任意のフォルダで回答
https://blog.alivate.com.au/poppler-windows/
2. 環境変数pathに解凍したバイナリを置いたフォルダを指定

### popplerのPDF→テキスト 変換コマンド
> pdftotext [pdfファイル] -opw [パスワード]

## プロジェクトの初期構築コマンド

### create new front react project
> create-react-app sail --template typescript\
> yarn add bulma
> yarn add react-router-dom 

### create new server golang project
> mkdir srv
> cd srv
> go mod init github.com/multios12/sail
> go get "github.com/glebarez/sqlite"
> go get "gorm.io/gorm"
