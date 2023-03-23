package main

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"

	"github.com/go-xiaohei/pugo/app/asset"

	"net/http"
)



func main() {

	fs := assetfs.AssetFS{

		Asset: asset.Asset,

		AssetDir: asset.AssetDir,

		AssetInfo: asset.AssetInfo,

	}

	http.Handle("/", http.FileServer( fs))

	http.ListenAndServe(":12345", nil)

}