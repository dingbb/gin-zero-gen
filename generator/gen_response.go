/*
 * @Author: aladdin
 * @Date: 2024-10-11 11:22:52
 * @LastEditTime: 2024-10-11 11:42:02
 * @LastEditors: aladdin
 * @FilePath: /gin-zero-gen/generator/gen_response.go
 */
package generator

import (
	"github.com/dingbb/gin-zero-gen/prepare"
	"github.com/dingbb/gin-zero-gen/tpl"
)

func GenResponse() error {
	return GenFile(
		"response.go",
		tpl.ResponseTemplate,
		WithSubDir("internal/response"),
		WithData(map[string]string{
			"rootPkg": prepare.RootPkg,
		}),
	)
}
