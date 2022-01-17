package main

import (
	"testing"
)

/*
Packets with type ID 4 represent a literal value. Literal value packets encode a single binary number. To do this,
the binary number is padded with leading zeroes until its length is a multiple of four bits, and then it is broken
into groups of four bits. Each group is prefixed by a 1 bit except the last group, which is prefixed by a 0 bit.
These groups of five bits immediately follow the packet header. For example, the hexadecimal string D2FE28 becomes:

110100101111111000101000
VVVTTTAAAAABBBBBCCCCC
Below each bit is a label indicating its purpose:

The three bits labeled V (110) are the packet version, 6.
The three bits labeled T (100) are the packet type ID, 4, which means the packet is a literal value.
The five bits labeled A (10111) start with a 1 (not the last group, keep reading) and contain the first four bits
of the number, 0111.
The five bits labeled B (11110) start with a 1 (not the last group, keep reading) and contain four more bits of the
number, 1110.
The five bits labeled C (00101) start with a 0 (last group, end of packet) and contain the last four bits of the
number, 0101.
The three unlabeled 0 bits at the end are extra due to the hexadecimal representation and should be ignored.
So, this packet represents a literal value with binary representation 011111100101, which is 2021 in decimal.
*/
func TestExample(t *testing.T) {
	b := ParseString("D2FE28")
	start := 0
	a := Parse(b, &start)
	if a.Version != 6 {
		t.Error("Version != 6")
	}
	if a.TypeID != 4 {
		t.Error("TypeID != 4")
	}
	if a.Value != 2021 {
		t.Error("Value != 2021")
	}
	if len(a.SubPackets) != 0 {
		t.Error("Expecting no children")
	}
}

/*
For example, here is an operator packet (hexadecimal string 38006F45291200) with length type ID 0 that contains two sub-packets:

00111000000000000110111101000101001010010001001000000000
VVVTTTILLLLLLLLLLLLLLLAAAAAAAAAAABBBBBBBBBBBBBBBB
The three bits labeled V (001) are the packet version, 1.
The three bits labeled T (110) are the packet type ID, 6, which means the packet is an operator.
The bit labeled I (0) is the length type ID, which indicates that the length is a 15-bit number representing the number of bits in the sub-packets.
The 15 bits labeled L (000000000011011) contain the length of the sub-packets in bits, 27.
The 11 bits labeled A contain the first sub-packet, a literal value representing the number 10.
The 16 bits labeled B contain the second sub-packet, a literal value representing the number 20.
*/
func TestExample1(t *testing.T) {
	b := ParseString("38006F45291200")
	start := 0
	a := Parse(b, &start)
	if a.Version != 1 {
		t.Error("Version != 1")
	}
	if a.TypeID != 6 {
		t.Error("TypeID != 6")
	}
	if len(a.SubPackets) != 2 {
		t.Error("Expecting 2 children")
	}
	if a.SubPackets[0].TypeID != 4 && a.SubPackets[0].Value != 10 {
		t.Error("First child value should be 10")
	}
	if a.SubPackets[1].TypeID != 4 && a.SubPackets[1].Value != 20 {
		t.Error("First child value should be 20")
	}
}

/*
As another example, here is an operator packet (hexadecimal string EE00D40C823060) with length type ID 1 that contains three sub-packets:

11101110000000001101010000001100100000100011000001100000
VVVTTTILLLLLLLLLLLAAAAAAAAAAABBBBBBBBBBBCCCCCCCCCCC
The three bits labeled V (111) are the packet version, 7.
The three bits labeled T (011) are the packet type ID, 3, which means the packet is an operator.
The bit labeled I (1) is the length type ID, which indicates that the length is a 11-bit number representing the number of sub-packets.
The 11 bits labeled L (00000000011) contain the number of sub-packets, 3.
The 11 bits labeled A contain the first sub-packet, a literal value representing the number 1.
The 11 bits labeled B contain the second sub-packet, a literal value representing the number 2.
The 11 bits labeled C contain the third sub-packet, a literal value representing the number 3.
After reading 3 complete sub-packets, the number of sub-packets indicated in L (3) is reached, and so parsing of this packet stops.
*/
func TestExample2(t *testing.T) {
	b := ParseString("EE00D40C823060")
	start := 0
	a := Parse(b, &start)
	if a.Version != 7 {
		t.Error("Version != 7")
	}
	if a.TypeID != 3 {
		t.Error("TypeID != 3")
	}
	if len(a.SubPackets) != 3 {
		t.Error("Expecting 3 children")
	}
	if a.SubPackets[0].TypeID != 4 && a.SubPackets[0].Value != 1 {
		t.Error("First child value should be 1")
	}
	if a.SubPackets[1].TypeID != 4 && a.SubPackets[1].Value != 2 {
		t.Error("First child value should be 2")
	}
	if a.SubPackets[2].TypeID != 4 && a.SubPackets[2].Value != 3 {
		t.Error("First child value should be 2")
	}
}

/*
8A004A801A8002F4788A004A801A8002F478 represents an operator packet (version 4) which contains an operator packet (version 1)
which contains an operator packet (version 5) which contains a literal value (version 6); this packet has a version sum of 16.
*/
func TestExample3(t *testing.T) {
	b := ParseString("8A004A801A8002F4788A004A801A8002F478")
	start := 0
	a := Parse(b, &start)
	if a.Version != 4 {
		t.Error("Version != 4")
	}
	if len(a.SubPackets) != 1 {
		t.Error("Expecting 1 children")
	}
	version1Packet := a.SubPackets[0]
	if version1Packet.Version != 1 && len(version1Packet.SubPackets) != 1 {
		t.Error("Top child value should be version 1 and have one child")
	}
	version5Packet := version1Packet.SubPackets[0]
	if version5Packet.Version != 5 && len(version5Packet.SubPackets) != 1 {
		t.Error("v1 child value should be version 5 and have one child")
	}
	version6Packet := version5Packet.SubPackets[0]
	if version6Packet.TypeID != 4 && version6Packet.Value != 16 {
		t.Error("v5 child value should be a literal with value 16")
	}

	versionSum := a.VersionSum()
	if versionSum != 16 {
		t.Error("expected version sum of 12, got ", versionSum)
	}
}

/*
620080001611562C8802118E34 represents an operator packet (version 3) which contains two sub-packets; each sub-packet is an
operator packet that contains two literal values. This packet has a version sum of 12.
*/
func TestExample4(t *testing.T) {
	b := ParseString("620080001611562C8802118E34")
	start := 0
	a := Parse(b, &start)
	if a.Version != 3 {
		t.Error("Version != 3")
	}
	if len(a.SubPackets) != 2 {
		t.Error("Expecting 2 children")
	}
	firstChild := a.SubPackets[0]
	secondChild := a.SubPackets[1]
	if len(firstChild.SubPackets) != 2 {
		t.Error("first child should have 2 children")
	}
	if len(secondChild.SubPackets) != 2 {
		t.Error("second child should have 2 children")
	}

	versionSum := a.VersionSum()
	if versionSum != 12 {
		t.Error("expected version sum of 12, got ", versionSum)
	}
}

/*
C0015000016115A2E0802F182340 has the same structure as the previous example, but the outermost packet uses a different length
type ID. This packet has a version sum of 23.
*/
func TestExample5(t *testing.T) {
	b := ParseString("C0015000016115A2E0802F182340")
	start := 0
	a := Parse(b, &start)
	if len(a.SubPackets) != 2 {
		t.Error("Expecting 2 children")
	}
	firstChild := a.SubPackets[0]
	secondChild := a.SubPackets[1]
	if len(firstChild.SubPackets) != 2 {
		t.Error("first child should have 2 children")
	}
	if len(secondChild.SubPackets) != 2 {
		t.Error("second child should have 2 children")
	}
	versionSum := a.VersionSum()
	if versionSum != 23 {
		t.Error("expected version sum of 23, got ", versionSum)
	}
}

/*
A0016C880162017C3686B18A3D4780 is an operator packet that contains an operator packet that contains an operator packet that
contains five literal values; it has a version sum of 31.
*/
func TestExample6(t *testing.T) {
	b := ParseString("A0016C880162017C3686B18A3D4780")
	start := 0
	a := Parse(b, &start)
	literalValues := a.SubPackets[0].SubPackets[0].SubPackets
	if len(literalValues) != 5 {
		t.Error("expecting 5 literal values")
	}
	versionSum := a.VersionSum()
	if versionSum != 31 {
		t.Error("expected version sum of 31, got ", versionSum)
	}
}

func TestEvalSum(t *testing.T) {
	b := ParseString("C200B40A82")
	start := 0
	a := Parse(b, &start)
	result := a.Evaluate()
	if result != 3 {
		t.Error("expected eval of 3, got ", result)
	}
}

func TestEvalProduct(t *testing.T) {
	b := ParseString("04005AC33890")
	start := 0
	a := Parse(b, &start)
	result := a.Evaluate()
	if result != 54 {
		t.Error("expected eval of 54, got ", result)
	}
}

func TestEvalMin(t *testing.T) {
	b := ParseString("880086C3E88112")
	start := 0
	a := Parse(b, &start)
	result := a.Evaluate()
	if result != 7 {
		t.Error("expected eval of 7, got ", result)
	}
}

func TestEvalMax(t *testing.T) {
	b := ParseString("CE00C43D881120")
	start := 0
	a := Parse(b, &start)
	result := a.Evaluate()
	if result != 9 {
		t.Error("expected eval of 9, got ", result)
	}
}

func TestEvalLessThan(t *testing.T) {
	b := ParseString("D8005AC2A8F0")
	start := 0
	a := Parse(b, &start)
	result := a.Evaluate()
	if result != 1 {
		t.Error("expected eval of 1, got ", result)
	}
}

func TestEvalGreaterThan(t *testing.T) {
	b := ParseString("F600BC2D8F")
	start := 0
	a := Parse(b, &start)
	result := a.Evaluate()
	if result != 0 {
		t.Error("expected eval of 0, got ", result)
	}
}

func TestEvalEqual(t *testing.T) {
	b := ParseString("9C005AC2F8F0")
	start := 0
	a := Parse(b, &start)
	result := a.Evaluate()
	if result != 0 {
		t.Error("expected eval of 0, got ", result)
	}
}

func TestEvalSumsEqual(t *testing.T) {
	b := ParseString("9C0141080250320F1802104A08")
	start := 0
	a := Parse(b, &start)
	result := a.Evaluate()
	if result != 1 {
		t.Error("expected eval of 1, got ", result)
	}
}
