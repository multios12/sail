
## プログラムの構成
* convert
  給与支給明細書をデータファイルに変換する。popplerを使用するため、現在のところubuntsu上で利用する必要がある  
* src
* srv

## 必要なライブラリ
* poppler-utils
  https://texwiki.texjp.org/?Poppler#k603e696

## 各種インストール
### poppler
> sudo apt install poppler-utils poppler-data

### create new front react project
> create-react-app

### create new server golang project
> mkdir srv
> cd srv
> go mod init github.com/multios12/sail/srv
> go get modernc.org/sqlite

## popplerのPDF 変換コマンド
> pdftotext [pdfファイル] -opw [パスワード]
