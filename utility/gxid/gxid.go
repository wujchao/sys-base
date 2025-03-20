package gxid

import "github.com/rs/xid"

func Gen() string {
	guid := xid.New()
	return guid.String()
}
