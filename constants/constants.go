package constants

type Type int

const (
	MSGTEXT      Type = 1   // 文字消息
	MSGIMAGE     Type = 2   // 图片消息
	MSGVIDEO     Type = 3   // 视频消息
	MSGFILE      Type = 4   // 文件消息
	MSGVOICE     Type = 8   // 音频消息
	MSGKMARKDOWN Type = 9   // KMarkdown
	MSGCARD      Type = 10  // card 消息
	MSGSYSTEM    Type = 255 // 系统消息
)

type ChannelType string

const (
	GROUP     ChannelType = "GROUP"     // 组播消息
	PERSON    ChannelType = "PERSON"    // 单播消息
	BROADCAST ChannelType = "BROADCAST" // 广播消息
)

// 卡片消息主题常量
type Theme string

const (
	ThemePrimary   Theme = "primary"
	ThemeSuccess   Theme = "success"
	ThemeDanger    Theme = "danger"
	ThemeWarning   Theme = "warning"
	ThemeInfo      Theme = "info"
	ThemeSecondary Theme = "secondary"
	ThemeNone      Theme = "none"
)

// 卡片消息大小
type Size string

const (
	SizeXS Size = "xs"
	SizeSM Size = "sm"
	SizeMD Size = "md"
	SizeLG Size = "lg"
)
