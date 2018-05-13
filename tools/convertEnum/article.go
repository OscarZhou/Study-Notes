package types

type ArticleType int

const (
	ArticlePage ArticleType = iota
	ArticleNews
	ArticleBlog
	ArticleUser
	ArticleAgent
)

//go:generate jsonenums -type=ArticleStatusType

type ArticleStatusType int

const (
	ArticleAuditing ArticleStatusType = iota
	ArticlePublished
)

//go:generate jsonenums -type=ArticleAccessStatus

type ArticleAccessStatus int

const (
	ArticleDraft ArticleAccessStatus = iota
	ArticlePublic
	ArticlePrivate
)
