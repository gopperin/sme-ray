package handler

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/context"

	mystrings "github.com/Eric-GreenComb/contrib/strings"
	"github.com/coocood/freecache"

	"github.com/gopperin/sme-ray/srv/version/bean"
	mylogger "github.com/gopperin/sme-ray/srv/version/logger"
	"github.com/gopperin/sme-ray/srv/version/persist"
	myproto "github.com/gopperin/sme-ray/srv/version/proto"
)

// Version Version
type Version struct{}

var cache *freecache.Cache

func init() {
	cacheSize := 1024
	cache = freecache.NewCache(cacheSize)
}

// GetVersion GetVersion
func (v *Version) GetVersion(ctx context.Context, req *myproto.VersionRequest, rsp *myproto.VersionResponse) error {

	rsp.Version = v.GetCacheVersion().Version

	return nil
}

// CheckVersion CheckVersion
func (v *Version) CheckVersion(ctx context.Context, req *myproto.CheckVersionRequest, rsp *myproto.CheckVersionResponse) error {

	rsp.Update = false
	rsp.Msg = "当前版本为最新版本"

	_version := v.GetCacheVersion()

	// big:1 smaill:2 equ:0 error:-1
	_bigger, err := mystrings.CompareVersion(req.Version, _version.Version)
	if err != nil {
		return nil
	}

	switch _bigger {
	case 2:
		rsp.Update = true
		rsp.Msg = _version.UpdateLog
	}

	fmt.Println("test debug")
	mylogger.ZapLogger.Debug("test debug")

	return nil
}

// GetCacheVersion GetCacheVersion
func (v *Version) GetCacheVersion() bean.VersionBase {

	var _ver bean.VersionBase

	_key := []byte("VER_" + "1")
	_got, err := cache.Get(_key)
	if err == nil {

		fmt.Println("cache : ", string(_got))

		json.Unmarshal(_got, &_ver)

		return _ver
	}

	_version, err := persist.GMariadb.GetVersion("1")
	if err != nil {
		_ver.Version = "0.0.0.0"
		return _ver
	}

	fmt.Println("persist : ", _version.Version)

	expire := 3600 // expire in 60*60 seconds
	_val, _ := json.Marshal(_version)
	cache.Set(_key, _val, expire)

	return _version
}
