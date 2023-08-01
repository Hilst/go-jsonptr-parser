module github.com/Hilst/app-ui

go 1.20

replace github.com/Hilst/app-ui/lexer => ./lexer

replace github.com/Hilst/app-ui/model => ./model

replace github.com/Hilst/app-ui/syntax => ./syntax

require (
	github.com/Hilst/app-ui/lexer v0.0.0-00010101000000-000000000000
	github.com/Hilst/app-ui/syntax v0.0.0-00010101000000-000000000000
)
