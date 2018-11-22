# gocarts(go-CERT-alerts-summarizer)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/tomoyamachi/gocarts/blob/master/LICENSE)

`gocarts` checks alerts of X-CERT (e.g. [JPCERT](http://www.jpcert.or.jp/), [US-CERT](https://www.us-cert.gov/ncas/alerts)).<br/>
This project refers to [knqyf263/gost](https://github.com/knqyf263/gost).

<img src="img/gocarts.gif" width="700">

# Abstract
`gocarts` is written in Go, and therefore you can just grab the binary releases and drop it in your `$PATH`.

gocarts summarizes alerts by CVE ID. You can search alert's detail by CVE ID.

# Main features
`gocarts` has the following features.
- Summarizing X-CERT alarts
- Searching alerts by CVE ID

# Usage

```
$ gocarts
X-CERT alerts summarizer

Usage:
  gocarts [command]

Available Commands:
  fetch       Fetch X-CERT alerts
  help        Help about any command
  search      Search X-CERT alerts

Flags:
      --dbpath string   /path/to/sqlite3 or SQL connection string
      --dbtype string   Database type to store data in (sqlite3, mysql or postgres supported)
      --debug           debug mode (default: false)
      --debug-sql       SQL debug mode
  -h, --help            help for gocarts

X-CERT alerts summarizer
```

# Fetch JPCERT

## Fetch alarts updated after 2016

```
$ gocarts fetch jpcert --after 2016

....
```

# Fetch USCERT

TODO

# Search mode

You need to install selector command (fzf or peco).

```
$ search alert --select-after 2018-01-01
> 2018-01-10 | https://www.jpcert.or.jp/at/2018/at180001.txt | ...                                                          | Adobe Flash Player の脆弱性 (APSB18-01) に関する注意喚起
  2018-01-10 | https://www.jpcert.or.jp/at/2018/at180002.txt | CVE-2018-0758, CVE-2018-0762, CVE-2018-0767, CVE-2018-0769,… | 2018年 1月マイクロソフトセキュリティ更新プログラムに関する注意喚起
  2018-01-17 | https://www.jpcert.or.jp/at/2018/at180003.txt | ...                                                          | 2018年 1月 Oracle Java SE のクリティカルパッチアップデートに関する注意喚起
  2018-01-17 | https://www.jpcert.or.jp/at/2018/at180004.txt | CVE-2017-10271                                               | Oracle WebLogic Server の脆弱性 (CVE-2017-10271) に関する注意喚起
  2018-01-17 | https://www.jpcert.or.jp/at/2018/at180005.txt | CVE-2017-3145                                                | ISC BIND 9 の脆弱性に関する注意喚起
  2018-02-02 | https://www.jpcert.or.jp/at/2018/at180006.txt | CVE-2018-4878                                                | Adobe Flash Player の未修正の脆弱性 (CVE-2018-4878) に関する注意喚起
  2018-02-14 | https://www.jpcert.or.jp/at/2018/at180007.txt | ...                                                          | Adobe Reader および Acrobat の脆弱性 (APSB18-02) に関する注意喚起
  2018-02-14 | https://www.jpcert.or.jp/at/2018/at180008.txt | CVE-2018-0763, CVE-2018-0825, CVE-2018-0834, CVE-2018-0835,… | 2018年 2月マイクロソフトセキュリティ更新プログラムに関する注意喚起
  2018-02-27 | https://www.jpcert.or.jp/at/2018/at180009.txt | ...                                                          | memcached のアクセス制御に関する注意喚起
  2018-03-14 | https://www.jpcert.or.jp/at/2018/at180010.txt | ...                                                          | Adobe Flash Player の脆弱性 (APSB18-05) に関する注意喚起
  2018-03-14 | https://www.jpcert.or.jp/at/2018/at180011.txt | CVE-2018-0872, CVE-2018-0874, CVE-2018-0876, CVE-2018-0889,… | 2018年 3月マイクロソフトセキュリティ更新プログラムに関する注意喚起
  2018-03-29 | https://www.jpcert.or.jp/at/2018/at180012.txt | CVE-2018-7600                                                | Drupal の脆弱性 (CVE-2018-7600) に関する注意喚起
  2018-04-06 | https://www.jpcert.or.jp/at/2018/at180013.txt | CVE-2018-0171                                                | Cisco Smart Install Client を悪用する攻撃に関する注意喚起
  2018-04-10 | https://www.jpcert.or.jp/at/2018/at180014.txt | CVE-2018-1270, CVE-2018-1271, CVE-2018-1272, CVE-2018-1275   | Spring Framework の脆弱性に関する注意喚起
  2018-04-11 | https://www.jpcert.or.jp/at/2018/at180015.txt | ...                                                          | Adobe Flash Player の脆弱性 (APSB18-08) に関する注意喚起
  2018-04-11 | https://www.jpcert.or.jp/at/2018/at180016.txt | CVE-2018-0870, CVE-2018-0979, CVE-2018-0980, CVE-2018-0981,… | 2018年 4月マイクロソフトセキュリティ更新プログラムに関する注意喚起
  2018-04-17 | https://www.jpcert.or.jp/at/2018/at180017.txt | CVE-2018-1273, CVE-2018-1274                                 | Spring Data Commons の脆弱性に関する注意喚起
  2018-04-18 | https://www.jpcert.or.jp/at/2018/at180018.txt | ...                                                          | 2018年 4月 Oracle 製品のクリティカルパッチアップデートに関する注意喚起
  2018-04-26 | https://www.jpcert.or.jp/at/2018/at180019.txt | CVE-2018-7602                                                | Drupal の脆弱性 (CVE-2018-7602) に関する注意喚起
  2018-05-09 | https://www.jpcert.or.jp/at/2018/at180020.txt | ...                                                          | Adobe Flash Player の脆弱性 (APSB18-16) に関する注意喚起
  2018-05-09 | https://www.jpcert.or.jp/at/2018/at180021.txt | CVE-2018-0943, CVE-2018-0945, CVE-2018-0946, CVE-2018-0951,… | 2018年 5月マイクロソフトセキュリティ更新プログラムに関する注意喚起
  2018-05-15 | https://www.jpcert.or.jp/at/2018/at180022.txt | ...                                                          | Adobe Reader および Acrobat の脆弱性 (APSB18-09) に関する注意喚起
  2018-05-15 | https://www.jpcert.or.jp/at/2018/at180023.txt | ...                                                          | メールクライアントにおける OpenPGP および S/MIME のメッセージの取り扱いに関する注意喚起
  2018-06-08 | https://www.jpcert.or.jp/at/2018/at180024.txt | CVE-2018-5002                                                | Adobe Flash Player の脆弱性 (APSB18-19) に関する注意喚起
  2018-06-13 | https://www.jpcert.or.jp/at/2018/at180025.txt | CVE-2018-8110, CVE-2018-8111, CVE-2018-8213, CVE-2018-8225,… | 2018年 6月マイクロソフトセキュリティ更新プログラムに関する注意喚起
  2018-07-11 | https://www.jpcert.or.jp/at/2018/at180026.txt | ...                                                          | Adobe Reader および Acrobat の脆弱性 (APSB18-21) に関する注意喚起
  2018-07-11 | https://www.jpcert.or.jp/at/2018/at180027.txt | ...                                                          | Adobe Flash Player の脆弱性 (APSB18-24) に関する注意喚起
  2018-07-11 | https://www.jpcert.or.jp/at/2018/at180028.txt | CVE-2018-8242, CVE-2018-8262, CVE-2018-8274, CVE-2018-8275,… | 2018年 7月マイクロソフトセキュリティ更新プログラムに関する注意喚起
  2018-07-18 | https://www.jpcert.or.jp/at/2018/at180029.txt | CVE-2018-2893, CVE-2018-2894, CVE-2018-2998, CVE-2018-2933,… | 2018年 7月 Oracle 製品のクリティカルパッチアップデートに関する注意喚起
  2018-07-23 | https://www.jpcert.or.jp/at/2018/at180030.txt | CVE-2018-1336, CVE-2018-8034, CVE-2018-8037                  | Apache Tomcat における複数の脆弱性に関する注意喚起
  2018-08-09 | https://www.jpcert.or.jp/at/2018/at180031.txt | CVE-2018-5740                                                | ISC BIND 9 サービス運用妨害の脆弱性 (CVE-2018-5740) に関する注意喚起
  2018-08-15 | https://www.jpcert.or.jp/at/2018/at180032.txt | ...                                                          | Adobe Reader および Acrobat の脆弱性 (APSB18-29) に関する注意喚起
  2018-08-15 | https://www.jpcert.or.jp/at/2018/at180033.txt | ...                                                          | Adobe Flash Player の脆弱性 (APSB18-25) に関する注意喚起
  2018-08-15 | https://www.jpcert.or.jp/at/2018/at180034.txt | CVE-2018-8266, CVE-2018-8273, CVE-2018-8302, CVE-2018-8344,… | 2018年 8月マイクロソフトセキュリティ更新プログラムに関する注意喚起
 

```

# License
MIT

# Author
TOMOYA Amachi
