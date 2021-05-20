package golibs

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gitee.com/ha666/golibs"
	"net/url"
	"sort"
	"strings"
)

func SignMd5(validateParams url.Values, secret string) string {
	keys := make([]string, 0, len(validateParams))
	for key := range validateParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	nsb := golibs.NewStringBuilder()
	nsb.Append(secret)
	for _, k := range keys {
		nsb.Append(k).Append(validateParams[k][0])
	}
	nsb.Append(secret)

	h := md5.New()
	h.Write([]byte(nsb.ToString()))
	by := h.Sum(nil)

	sb := golibs.NewStringBuilder()
	for i := 0; i < len(by); i++ {
		sb.Append(fmt.Sprintf("%02X", by[i]))
	}
	return sb.ToString()
}

func HexSignMd5(validateParams url.Values, secret string) string {
	keys := make([]string, 0, len(validateParams))
	for key := range validateParams {
		if strings.EqualFold(key, "sign") {
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)
	nsb := golibs.NewStringBuilder()
	for _, k := range keys {
		nsb.Append(k).Append(validateParams[k][0])
	}
	nsb.Append(secret)
	h := md5.New()
	h.Write([]byte(nsb.ToString()))
	by := h.Sum(nil)
	return hex.EncodeToString(by)
}


