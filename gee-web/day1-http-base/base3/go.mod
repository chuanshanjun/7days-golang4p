module github.com/chuanshan/example

go 1.16

require gee v0.0.0

//在 go.mod 中使用 replace 将 gee 指向 ./gee
replace (
	gee => ./gee
)
