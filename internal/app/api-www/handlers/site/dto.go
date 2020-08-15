package site

type ReqModel struct {
	ID          int64       ` json:"id"`           // id
	LanguageID  int         `json:"languageID"  `  // 语言ID:1中文,2英语
	Title       string      `json:"title"  `       //标题
	Email       string      `json:"email"  `       //邮箱
	Description string      `json:"description"  ` //描述
	Extended    interface{} `json:"extended" `     // 扩展的
}

type RespModel struct {
	ID          int64       ` json:"id"`           // id
	CreatedAt   int64       ` json:"createdAt"`    // 创建时间戳
	UpdatedAt   int64       ` json:"updatedAt"`    // 更新时间戳
	Status      int8        `json:"status"  `      // 44 删除, 1 启用, 4 禁用
	LanguageID  int         `json:"languageID"  `  // 语言ID:1中文,2英语
	Title       string      `json:"title"  `       //标题
	Email       string      `json:"email"  `       //邮箱
	Description string      `json:"description"  ` //描述
	Extended    interface{} `json:"extended" `     // 扩展的
}
