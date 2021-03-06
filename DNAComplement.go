
package main

import (
	"fmt"
	"bytes"
)

/*'Reverse' function pulled from 'stringutil' package on github*/
func Reverse (s string) string {
	r := []rune(s)

	for i,j := 0, len(r) - 1; i < len(r)/2; i,j = i + 1, j - 1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main(){
	dna := []byte(`AGTCGCATCTTTCCATCCATGAACGGAATTAGATGTAAATGGCGTCCACTGTATGTCTGGCGTGCCTAATTCGCCGAACTGTACGCGAGGGCCATGCAGAATAAAGGGATATCGTGGTTACTATGGGAGAACACCATAGTCGGGTTGGCCGGACTTGCTATCGTATCACGTCTGATGGCTGTAAGCACTGAACTATCCCCTTCACCGCAAATTGAGCGCTATTTTTGAACACTTGAGTATGGTCGGAAAATGAACCTTAGATCTGTGTCTCCGATCCAGTAGTGCGACACGCACTATGAGGAGCAACCTACTAGGGTGACCCATTGCTGCGTGAAGCAGGTGCCCCGAGTAGCCTAGGAAGTACGTGGCCAATGGGTATGAAAATTGTGTACCGTCCTTTAAGGTCTTTCAGTCAGGCCTGAAGACTGTTAAGGAACGTGTACATGTCCAGCCTCAGCAGACTACGTTTCGTATCTTCGCGGAACAGGTTTCAGATACAAGGGCACCTAAATGATTAGTGAGAGCGAATTAGTTAATATTTAAGATTTTATGACGGCGCCTCAGTTCTTATCGTACGCAGACGCCGGCGATGCAGAGGGGGACTGAGGGCGAAGCGAGGTGCTTTTTGGTGAATCGAGCTGCGAGGTGGAAGACATCGTCGGCCTAGGAGAAAGTCAAGTTACCGCTCATTGACTCATGCGTCCCGACCCGTTTAATAATGTGCATCAATCCAGGCGCGGGCATAGAGTTGATTCCTCTAGTGACGTGTTTACCTTCGATGATCCAATCGTACCACTGGTCATTTAGGGGGTTCTTACCTAAGTCCCCGCCCAGTGGGTATCCCAGTAAAAATGAATCATCCCGATCTCCCCTCGGCGGTATCAACTGTGCACTAAGGGTTACCTTAT`)

	/*Probably should abstract this out into a function... and make it more flexible...*/
	dna = bytes.Replace(dna, []byte("T"), []byte("a"), len(dna))
	dna = bytes.Replace(dna, []byte("A"), []byte("t"), len(dna))
	dna = bytes.Replace(dna, []byte("C"), []byte("g"), len(dna))
	dna = bytes.Replace(dna, []byte("G"), []byte("c"), len(dna))

	dna = bytes.ToUpper(dna)

	compliment := Reverse(string(dna))
	fmt.Println(compliment)
}
