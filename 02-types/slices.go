package types

import "fmt"

type Part string

func printAssemblyLine(line []Part) {
	text := "The assembly line consist of:"
	for _, part := range line {
		text += " " + string(part) + ","
	}
	fmt.Println(text)
}

func Slices() {
	assemblyLine := make([]Part, 0, 5)
	assemblyLine = append(assemblyLine, "disks", "callipers", "springs")
	printAssemblyLine(assemblyLine)
	assemblyLine = append(assemblyLine, "shock absorbers", "linkages")
	printAssemblyLine(assemblyLine)
	assemblyLine = assemblyLine[3:]
	printAssemblyLine(assemblyLine)
}
