package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/srowles/adventofcode2021"
)

func main() {
	do1()
	do2()
}

func do1() {
	input := adventofcode2021.MustInputFromWebsite("16")
	start := time.Now().UTC()
	versionSum := versionSum(strings.TrimSpace(input))
	fmt.Println(versionSum)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func do2() {
	input := adventofcode2021.MustInputFromWebsite("16")
	start := time.Now().UTC()
	value := value(strings.TrimSpace(input))
	fmt.Println(value)
	fmt.Println("took:", time.Now().UTC().Sub(start))
}

func versionSum(input string) int {
	decoded := binaryDecode(input)
	versionSum, _ := decode(decoded, 0)
	return int(versionSum)
}

func value(input string) int {
	decoded := binaryDecode(input)
	versionSum, _ := decodeValue(decoded, 0)
	return int(versionSum)
}

func decode(decoded []uint64, versionSum uint64) (uint64, int) {
	version := decoded[0:3]
	versionSum += sliceToInt(version)
	typeID := decoded[3:6]
	idx := 6

	switch sliceToInt(typeID) {
	case 4:
		for {
			idx += 5
			if decoded[idx-5] == 0 {
				return versionSum, idx
			}
		}
	default:
		if decoded[idx] == 0 {
			subPacketsLen := sliceToInt(decoded[idx+1 : idx+16])
			start := idx + 16
			packetCount := 0
			read := 0
			for {
				subSum, upTo := decode(decoded[start:start+int(subPacketsLen)-read], 0)
				versionSum += subSum
				start += upTo
				read += upTo
				packetCount += upTo
				if packetCount >= int(subPacketsLen) {
					return versionSum, start
				}
			}
		} else if decoded[idx] == 1 {
			numberOfPackets := sliceToInt(decoded[idx+1 : idx+12])
			start := idx + 12
			for i := 0; i < int(numberOfPackets); i++ {
				subSum, upTo := decode(decoded[start:], 0)
				versionSum += subSum
				start += upTo
			}
			return versionSum, start
		}
	}

	return versionSum, 0
}

func decodeValue(decoded []uint64, value uint64) (uint64, int) {
	typeID := decoded[3:6]
	idx := 6

	switch sliceToInt(typeID) {
	case 4:
		var literals []uint64
		for {
			literals = append(literals, decoded[idx+1:idx+5]...)
			idx += 5
			if decoded[idx-5] == 0 {
				return sliceToInt(literals), idx
			}
		}
	default:
		if decoded[idx] == 0 {
			subPacketsLen := sliceToInt(decoded[idx+1 : idx+16])
			start := idx + 16
			packetCount := 0
			read := 0
			var subValues []uint64
			for {
				subSum, upTo := decodeValue(decoded[start:start+int(subPacketsLen)-read], 0)
				subValues = append(subValues, subSum)
				start += upTo
				read += upTo
				packetCount += upTo
				if packetCount >= int(subPacketsLen) {
					return doMaths(sliceToInt(typeID), subValues...), start
				}
			}
		} else if decoded[idx] == 1 {
			numberOfPackets := sliceToInt(decoded[idx+1 : idx+12])
			start := idx + 12
			var subValues []uint64
			for i := 0; i < int(numberOfPackets); i++ {
				subSum, upTo := decodeValue(decoded[start:], 0)
				subValues = append(subValues, subSum)
				start += upTo
			}
			return doMaths(sliceToInt(typeID), subValues...), start
		}
	}

	return value, 0
}

func doMaths(typeid uint64, values ...uint64) uint64 {
	var result uint64
	switch typeid {
	case 0:
		for _, v := range values {
			result += v
		}
	case 1:
		result = 1
		for _, v := range values {
			result *= v
		}
	case 2:
		result = math.MaxUint64
		for _, v := range values {
			if v < result {
				result = v
			}
		}
	case 3:
		result = 0
		for _, v := range values {
			if v > result {
				result = v
			}
		}
	case 5:
		if values[0] > values[1] {
			return 1
		} else {
			return 0
		}
	case 6:
		if values[0] < values[1] {
			return 1
		} else {
			return 0
		}
	case 7:
		if values[0] == values[1] {
			return 1
		} else {
			return 0
		}
	}
	return result
}

func sliceToInt(slice []uint64) uint64 {
	var val uint64 = 0
	for _, v := range slice {
		val += v
		val = val << 1
	}

	val = val >> 1

	return val
}

func binaryDecode(input string) []uint64 {
	var bits []uint64
	for _, c := range input {
		i, err := strconv.ParseInt(string(c), 16, 32)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse %s as hex: %v", string(c), err))
		}
		bits = append(bits, (uint64(i)&0b1000)>>3)
		bits = append(bits, (uint64(i)&0b0100)>>2)
		bits = append(bits, (uint64(i)&0b0010)>>1)
		bits = append(bits, (uint64(i) & 0b0001))
	}
	return bits
}
