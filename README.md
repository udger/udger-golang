# Udger golang (format V3):

This package is a fork of github.com/udger/udger with the following changes:
- Using golang regexp instead of github.com/glenn-brown/golang-pkg-pcre
- Supporting partial UA parsing by constructor flags (device, browser, os)

# Usage:

```
package main

import (
  "github.com/yoavfeld/udger"
)

func main() {
  client, err := udger.New("udgerDBv3FilePath", &udger.Flags{Device: true})
  if err != nil {
     log.Fatal(err)
  }
  ua := "Mozilla/5.0 (Linux; Android 4.4.4; MI PAD Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/33.0.0.0 Safari/537.36"
  res,err := client.Lookup(ua)
  if err != nil {
    log.Fatal(err)
  }
}


```

# Open issues:

- Browser version is not supported since using golang regexp. it require a fic in findData func to support findDataWithVersion func
