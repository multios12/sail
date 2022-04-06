
## 必要なライブラリ
* poppler-utils
  https://texwiki.texjp.org/?Poppler#k603e696


## 各種インストール
### poppler
> sudo apt install poppler-utils poppler-data

# PDF 変換コマンド
> pdftotext -layout 2021年11月給与_177_小澤　和浩.pdf -opw H0812109
> pdftotext 2021年11月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 85 -y 150 -W 200 -H 100

## 画像変換
pdftocairo 2021年11月給与_177_小澤　和浩.pdf -png -opw H0812109  

## 出勤
pdftotext 2021年11月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 70 -y 100 -W 30 -H 60
## 休出
pdftotext 2021年11月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 100 -y 100 -W 30 -H 60
## 特休
pdftotext 2021年11月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 130 -y 100 -W 30 -H 60
## 有休
pdftotext 2021年11月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 160 -y 100 -W 20 -H 60
## 欠勤
pdftotext 2021年11月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 190 -y 100 -W 30 -H 60

## 出勤時間/遅早時間/残業 1(時間外)時間
pdftotext 2021年11月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 220 -y 90 -W 190 -H 60
## 残業 2(深夜)時間/所定休日残業時間/法定休日残業時間/ｸﾞﾚｰﾄﾞﾗﾝｸ
pdftotext 2021年11月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 410 -y 90 -W 300 -H 60

## 支給１
pdftotext 2021年11月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 85 -y 150 -W 600 -H 60
pdftotext 2022年2月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 85 -y 150 -W 600 -H 60
pdftotext 2021年2月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 85 -y 150 -W 600 -H 60

## 支給2
pdftotext 2021年11月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 85 -y 200 -W 600 -H 60

## 控除
pdftotext 2021年11月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 75 -y 330 -W 600 -H 60

## 合計
pdftotext 2021年11月給与_177_小澤　和浩.pdf 1.txt -opw H0812109 -x 75 -y 420 -W 600 -H 60
