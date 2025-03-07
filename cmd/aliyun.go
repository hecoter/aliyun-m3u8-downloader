package cmd

import (
	"fmt"
	"log"

	"github.com/lbbniu/aliyun-m3u8-downloader/pkg/download"

	"github.com/lbbniu/aliyun-m3u8-downloader/pkg/tool"
	"github.com/spf13/cobra"
)

// aliyunCmd represents the aliyun command
var aliyunCmd = &cobra.Command{
	Use:   "aliyun",
	Short: "阿里云私有m3u8加密下载工具",
	Long: `阿里云私有m3u8加密下载工具. 使用示例:
aliyun-m3u8-downloader aliyun -p "WebPlayAuth" -v 视频id -o=/data/example --chanSize 1 -f 文件名`,
	Run: func(cmd *cobra.Command, args []string) {
		playAuth, _ := cmd.Flags().GetString("playAuth")
		videoId, _ := cmd.Flags().GetString("videoId")
		filename, _ := cmd.Flags().GetString("filename")
		output, _ := cmd.Flags().GetString("output")
		chanSize, _ := cmd.Flags().GetInt("chanSize")
		if playAuth == "" {
			tool.PanicParameter("playAuth")
		}
		if videoId == "" {
			tool.PanicParameter("videoId")
		}
		if output == "" {
			tool.PanicParameter("output")
		}
		if chanSize <= 0 {
			panic("parameter 'chanSize' must be greater than 0")
		}
		if err := download.AliyunDownload(output, filename, chanSize, videoId, playAuth); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(aliyunCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aliyunCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aliyunCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	aliyunCmd.Flags().StringP("playAuth", "p", "", "web播放认证信息")
	aliyunCmd.Flags().StringP("videoId", "v", "", "视频id")
	aliyunCmd.Flags().StringP("output", "o", "", "下载保存位置")
	aliyunCmd.Flags().IntP("chanSize", "c", 1, "下载并发数")
	aliyunCmd.Flags().StringP("filename", "f", "", "保存文件名")
	_ = aliyunCmd.MarkFlagRequired("playAuth")
	_ = aliyunCmd.MarkFlagRequired("videoId")
	_ = aliyunCmd.MarkFlagRequired("output")
	_ = aliyunCmd.MarkFlagRequired("chanSize")
}
