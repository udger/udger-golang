# Udger golang (format V3)

This package is a fork of github.com/udger/udger with the following changes:
- Using golang regexp instead of github.com/glenn-brown/golang-pkg-pcre
- Supporting partial UA parsing (device, browser, os)

# open issues

- Browser version is not supported since using golang regexp. it require a fic in findData func to support findDataWithVersion func
