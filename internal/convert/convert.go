package convert

func MarkdownImgUrl(u string) string {
	return "![img](" + u + ")"
}

func HtmlImgUrl(u string) string {
	return "<img src=\"" + u + "\" />"
}
