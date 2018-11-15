# gocarts(go-CERT-alerts-summarizer)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/tomoyamachi/gocarts/blob/master/LICENSE)

`gocarts` checks alerts of X-CERT (e.g. [JPCERT](http://www.jpcert.or.jp/), [US-CERT](https://www.us-cert.gov/ncas/alerts).

After you register CVEs to watch list

# Abstract
`gocarts` is written in Go, and therefore you can just grab the binary releases and drop it in your $PATH.

go-cas summarizes alerts by CVE ID. You can search alert's detail by CVE ID.

# Main features
`go-cas` has the following features.
- Summarizing X-CERT alarts
- Searching alerts by CVE ID

# Usage

```
$ gocarts
X-CERT alerts summarizer

Usage:
  gost [command]

Available Commands:
  fetch     Fetch the data of the X-CERT alerts
  tui       Search alerts by CVE ID in TUI

Flags:
      --dbpath string       /path/to/sqlite3 or SQL connection string
      --dbtype string       Database type to store data in (sqlite3, mysql or postgres supported)
      --debug               debug mode (default: false)
      --debug-sql           SQL debug mode
  -h, --help                help for go-cas

Use "gocarts [command] --help" for more information about a command.
```

# Fetch JPCERT

## Fetch alarts updated after 2016

```
$ gocarts fetch jpcert --after 2016

....
```

# Fetch USCERT

TODO

# TUI mode

```
$ gocarts tui
...
```

# License
MIT

# Author
TOMOYA Amachi
