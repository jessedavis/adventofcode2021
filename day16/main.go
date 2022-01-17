package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// I also really need to understand packages in Go.

type Packet struct {
	Version      int
	TypeID       int
	Value        int
	LengthTypeID int
	SubPackets   []*Packet
}

func (p Packet) VersionSum() int {
	sum := p.Version
	for _, sp := range p.SubPackets {
		sum += sp.VersionSum()
	}
	return sum
}

// this part copied from lizthegrey, I'm still learning Go
// and was having a VERY hard time implmenting bit
// manipulation in Go.  I definitely need to practice this.

// Also credit to lizthegrey for help all over this solution,
// I definitely need to practice programming in general :/

type BitArray []bool

func (b BitArray) Read(start *int, length int) uint64 {
	if length > 64 || *start+length > len(b) {
		// explode
		fmt.Printf("should never get here: called with start %d and length %d\n", *start, length)
		return 0
	}
	var ret uint64
	for i := 0; i < length; i++ {
		ret = ret << 1
		if b[i+*start] {
			ret += 1
		}
	}
	*start += length
	return ret
}

func (b BitArray) Print() {
	for _, v := range b {
		if v {
			fmt.Printf("1")
		} else {
			fmt.Printf("0")
		}
	}
	fmt.Println()
}

func ParseString(s string) BitArray {
	result := make(BitArray, 0)

	for i := 0; i < len(s); i += 2 {
		// TODO: i don't really understand what this is doing, need to print it out
		// chunk := (s[i] << 4) + s[i+1]
		chunk, _ := strconv.ParseInt(s[i:i+2], 16, 16)
		for c := 0; c < 8; c++ {
			bit := (chunk >> (7 - c)) & 1
			result = append(result, bit == 1)
		}
	}

	return result
}

func Parse(b BitArray, pos *int) *Packet {
	version := int(b.Read(pos, 3))
	typeID := int(b.Read(pos, 3))
	subPackets := make([]*Packet, 0)

	if typeID == 4 {
		// literal
		var literalValue uint64
		for b.Read(pos, 1) == 1 {
			literalValue = literalValue << 4
			next := b.Read(pos, 4)
			literalValue += next
		}
		// and 0 part
		literalValue = literalValue << 4
		literalValue += b.Read(pos, 4)
		return &Packet{Version: version, TypeID: typeID, Value: int(literalValue), SubPackets: subPackets}
	} else {
		// operator packet
		lengthTypeID := int(b.Read(pos, 1))
		switch lengthTypeID {
		case 0:
			subPacketBitLength := int(b.Read(pos, 15))
			subPacketsPosEnd := *pos + subPacketBitLength
			for *pos < subPacketsPosEnd {
				subPackets = append(subPackets, Parse(b, pos))
			}
		case 1:
			numSubPackets := int(b.Read(pos, 11))
			for len(subPackets) < numSubPackets {
				subPackets = append(subPackets, Parse(b, pos))
			}
		}
		// REALLY need to get better at specifying via key, don't assume default values will be correct
	}
	return &Packet{Version: version, TypeID: typeID, SubPackets: subPackets}
}

func (p Packet) Evaluate() int {
	var result int

	switch p.TypeID {
	case 0:
		sp := p.SubPackets
		sum := 0
		if len(sp) == 1 {
			sum = sp[0].Evaluate()
		} else {
			for _, x := range sp {
				sum += x.Evaluate()
			}
		}
		result = sum
	case 1:
		sp := p.SubPackets
		product := 1
		if len(sp) == 1 {
			product = sp[0].Evaluate()
		} else {
			for _, x := range sp {
				product *= x.Evaluate()
			}
		}
		result = product
	case 2:
		var min int
		for _, sp := range p.SubPackets {
			v := sp.Evaluate()
			if v < min || min == 0 {
				min = v
			}
		}
		result = min
	case 3:
		var max int
		for _, sp := range p.SubPackets {
			v := sp.Evaluate()
			if v > max || max == 0 {
				max = v
			}
		}
		result = max
	case 4:
		result = p.Value
	case 5:
		first := p.SubPackets[0]
		second := p.SubPackets[1]
		if first.Evaluate() > second.Evaluate() {
			result = 1
		} else {
			result = 0
		}
	case 6:
		first := p.SubPackets[0]
		second := p.SubPackets[1]
		if first.Evaluate() < second.Evaluate() {
			result = 1
		} else {
			result = 0
		}
	case 7:
		first := p.SubPackets[0]
		second := p.SubPackets[1]
		if first.Evaluate() == second.Evaluate() {
			result = 1
		} else {
			result = 0
		}
	}

	return result
}

func main() {
	f, err := os.Open("input.txt")
	//f, err := os.Open("part1exampledata.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	b := ParseString(line)
	start := 0
	p := Parse(b, &start)

	fmt.Printf("Version sum: %d\n", p.VersionSum())
	fmt.Printf("Evaulation expression result: %d\n", p.Evaluate())
}
