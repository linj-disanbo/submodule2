package main

import "github.com/linj-disanbo/submodule1/tx/sdk"

func main() {
	tx := sdk.CreateTx("coins")
	sdk.SignTx(tx)
}
