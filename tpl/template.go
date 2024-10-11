/*
 * @Author: aladdin
 * @Date: 2024-10-10 18:22:09
 * @LastEditTime: 2024-10-11 15:13:07
 * @LastEditors: aladdin
 * @FilePath: /gin-zero-gen/tpl/template.go
 */
package tpl

import _ "embed"

var (
	//go:embed types.tpl
	TypesTemplate string

	//go:embed response.tpl
	ResponseTemplate string

	//go:embed handler.tpl
	HandlerTemplate string

	//go:embed logic.tpl
	LogicTemplate string

	//go:embed routes.tpl
	RoutesTemplate string

	//go:embed routes_setup.tpl
	RoutesSetupTemplate string
)
