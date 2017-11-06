package commandLineArguments

import (
	"os"
	"fmt"
	"flag"
)

/**命令行参数*/
var p = fmt.Println

func main() {
	argsWithProgram := os.Args
	argsWithoutProgram := os.Args[1:]

	arg := os.Args[3]

	/**
	./demo a b c d
	*/
	p(argsWithProgram)
	p(argsWithoutProgram)
	p(arg)

	/**传参的值与用方法*/
	// flog.x(关键字,默认值,用法)
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")

	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	/**
	./demo -word=opt -numb=7 -fork -svar=flag
	./demo -word=opt
	./demo -word=opt a1 a2 a3
	./demo -word=opt a1 a2 a3 -numb=7
	./demo -h
	./demo -wat
	*/

	p("word:", *wordPtr)
	p("nomb:", *numbPtr)
	p("fork:", *boolPtr)
	p("svar:", svar)
	p("tail:", flag.Args())
}
