package subparser

import "github.com/allanpk716/ChineseSubFinder/internal/types"

type FileInfo struct {
	FromWhereSite string          // 从那个网站下载的
	Name          string          // 字幕的名称，注意，这里需要额外的赋值，不会自动检测
	Ext           string          // 字幕的后缀名
	Lang          types.Language  // 识别出来的语言
	FileFullPath  string          // 字幕文件的全路径
	Data          []byte          // 字幕的二进制文件内容
	Dialogues     []OneDialogue   // 整个字幕文件的所有对话
	DialoguesEx   []OneDialogueEx // 整个字幕文件的所有对话，这里会把一句话中支持的 中、英、韩、日 四国语言给分离出来
	CHLines       []string        // 抽取出所有的中文对话
	OtherLines    []string        // 抽取出所有的第二语言对话，可能是英文、韩文、日文
}

// OneDialogue 一句对话
type OneDialogue struct {
	StartTime string   // 开始时间
	EndTime   string   // 结束时间
	StyleName string   // StyleName
	Lines     []string // 台词
}

// OneDialogueEx 一句对话，这里会把一句话中支持的 中、英、韩、日 四国语言给分离出来
type OneDialogueEx struct {
	StartTime string // 开始时间
	EndTime   string // 结束时间
	ChLine    string
	EnLine    string
	KrLine    string
	JpLine    string
}
