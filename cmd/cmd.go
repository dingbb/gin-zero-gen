/*
 * @Author: aladdin
 * @Date: 2024-10-10 17:20:37
 * @LastEditTime: 2024-10-10 18:21:00
 * @LastEditors: aladdin
 * @FilePath: /gin-zero-gen/cmd/cmd.go
 */

package cmd

import (
	"log"
	"os"

	"github.com/dingbb/gin-zero-gen/generator"
	"github.com/dingbb/gin-zero-gen/prepare"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "gengin",
		Short:   "生成基于 GIN 框架的 WEB 服务的相关文件",
		Example: "gin-zero-gen --dir=. user.api",
		Args:    cobra.ExactValidArgs(1),
		RunE:    GenGinCode,
	}
)

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&prepare.OutputDir, "dir", ".", "生成项目目录")
}

func GenGinCode(cmd *cobra.Command, args []string) error {
	prepare.ApiFile = args[0]
	prepare.Setup()
	Must(generator.GenTypes())
	return nil
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
