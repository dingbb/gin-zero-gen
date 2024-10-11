/*
 * @Author: aladdin
 * @Date: 2024-10-11 15:13:45
 * @LastEditTime: 2024-10-11 15:16:34
 * @LastEditors: aladdin
 * @FilePath: /gin-zero-gen/generator/gen_logic.go
 */
package generator

import (
	"path"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/dingbb/gin-zero-gen/prepare"
	"github.com/dingbb/gin-zero-gen/tpl"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenLogic() error {
	for _, g := range prepare.ApiSpec.Service.Groups {
		for _, r := range g.Routes {
			err := genLogicByRoute(g, r)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func genLogicByRoute(group spec.Group, route spec.Route) error {
	logicName, err := format.FileNamingFormat(fileNameStyle, route.Handler)
	if err != nil {
		return err
	}

	logicFileName := strings.TrimSuffix(strings.TrimSuffix(logicName, "logic"), "_") + "_logic.go"

	subDir := group.GetAnnotation(groupProperty)
	subDir, err = format.FileNamingFormat(dirStyle, subDir)
	if err != nil {
		return err
	}

	logicPkg := path.Join("logic", subDir)
	logicBase := path.Base(logicPkg)

	respIsPrimitiveType, respTypeName := parseResponseType(route.ResponseType)

	return GenFile(
		logicFileName,
		tpl.LogicTemplate,
		WithSubDir(logicPkg),
		WithData(map[string]any{
			"rootPkg":           prepare.RootPkg,
			"pkgName":           logicBase,
			"comment":           parseComment(route),
			"logicName":         cases.Title(language.English, cases.NoLower).String(route.Handler),
			"requestType":       cases.Title(language.English, cases.NoLower).String(route.RequestTypeName()),
			"responseType":      respTypeName,
			"needImportTypePkg": len(route.RequestTypeName()) > 0 || (!respIsPrimitiveType && len(route.ResponseTypeName()) > 0),
			"hasReq":            len(route.RequestTypeName()) > 0,
			"hasResp":           len(route.ResponseTypeName()) > 0,
		}),
	)
}
