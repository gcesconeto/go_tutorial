module github.com/gcesconeto/go_tutorial/hello

go 1.18

require (
	github.com/gcesconeto/go_tutorial/greetings v0.0.0-20220719194009-cadd7f95b752
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.3.7 // indirect
	rsc.io/sampler v1.99.99 // indirect
)

replace github.com/gcesconeto/go_tutorial/greetings => ../greetings
