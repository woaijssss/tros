package utils

import (
	"bytes"
	"github.com/yuin/goldmark"
	"regexp"
)

// IsMarkdown Check if the text contains Markdown features
func IsMarkdown(text string) bool {
	// Use raw string literals (`) to avoid escape character issues
	patterns := []string{
		`^#{1,6}\s`,        // 标题 title
		`^\*+\s`,           // 无序列表 unordered list
		`^\d+\.\s`,         // 有序列表 ordered list
		`\*\*.*\*\*`,       // 粗体 bold
		`\*.*\*`,           // 斜体 italic
		"```[\\s\\S]*?```", // 代码块（使用双引号以支持转义） Code block (use double quotation marks to support escaping)
		`!\[.*\]\(.*\)`,    // 图片 image
		`\[.*\]\(.*\)`,     // 链接 link
	}

	for _, pattern := range patterns {
		if matched, _ := regexp.MatchString(pattern, text); matched {
			return true
		}
	}
	return false
}

// MarkdownToHTML Convert Markdown to HTML
func MarkdownToHTML(markdown string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
