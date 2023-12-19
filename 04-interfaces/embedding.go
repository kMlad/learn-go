package interfaces

import "fmt"

type Bytes int
type Celsius float32

type CpuTemp struct {
	temp []Celsius
}

func (cpu *CpuTemp) CalcAverage() float32 {
	temps := cpu.temp
	var sum float32 = 0.0
	for _, temp := range temps {
		sum += float32(temp)
	}
	return sum / float32(len(temps))
}

type MemoryUsage struct {
	amount []Bytes
}

func (mem *MemoryUsage) CalcAverage() int {
	bytes := mem.amount
	sum := 0
	for _, singleAmount := range bytes {
		sum += int(singleAmount)
	}
	return sum / len(bytes)
}

type BandwidthUsage struct {
	amount []Bytes
}

func (bandwidth *BandwidthUsage) CalcAverage() int {
	bytes := bandwidth.amount
	sum := 0
	for _, singleAmount := range bytes {
		sum += int(singleAmount)
	}
	return sum / len(bytes)
}

type Dashboard struct {
	CpuTemp
	MemoryUsage
	BandwidthUsage
}

func Embedding() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celsius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{
		CpuTemp:        temp,
		MemoryUsage:    memory,
		BandwidthUsage: bandwidth,
	}

	fmt.Printf("Average CPU Temperature is %v\n", dashboard.CpuTemp.CalcAverage())
	fmt.Printf("Average Memory Usage is %v\n", dashboard.MemoryUsage.CalcAverage())
	fmt.Printf("Average Bandwidth Usage is %v\n", dashboard.BandwidthUsage.CalcAverage())
}
