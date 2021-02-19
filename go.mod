module makesite

go 1.15

replace gopkg.in/russross/blackfriday.v2 v2.0.1 => github.com/russross/blackfriday/v2 v2.0.1

require (
	github.com/bregydoc/gtranslate v0.0.0-20200913051839-1bd07f6c1fc5
	github.com/robertkrimen/otto v0.0.0-20200922221731-ef014fd054ac // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/text v0.3.5
	gopkg.in/russross/blackfriday.v2 v2.0.1
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
)
