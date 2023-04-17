package basic

var InitInfo InitInfoS

func Init(input string, output string, core int64) {
	InitInfo.InputFile = input
	InitInfo.OutputFile = output
	InitInfo.CPU = core
}

func Run() {
	dataBlock := Read()
	quickSort(dataBlock, 0, len(dataBlock)-1)
	Write(dataBlock)
}
