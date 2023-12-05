/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var d1Cmd = &cobra.Command{
	Use:   "d1",
	Short: "day one",
	Long:  "day one, use [ -two | -t ] to run day two",
	Run: func(cmd *cobra.Command, args []string) {

		if !cmd.Flags().Changed(("two")) {
			calibrate()
		} else {
			betterCalibrate()
		}
	},
}

func init() {
	rootCmd.AddCommand(d1Cmd)

	d1Cmd.Flags().BoolP("two", "t", false, "Help message for toggle")
}

// part one
func firstNum(line string) (int, string) {
	for i, character := range line {
		if unicode.IsNumber(character) {
			return i, string(character)
		}
	}
	return -1, "!"
}

func lastNum(line string) string {
	_, num := firstNum(reverse(line))
	return num
}

func reverse(s string) string {
	var reversed string
	for _, character := range s {
		reversed = string(character) + reversed
	}
	return reversed
}

func calibrate() {
	sum := 0
	lines := Lines("1-1")
	for _, line := range lines {
		_, f := firstNum(line)
		num, err := strconv.Atoi(f + lastNum(line))
		if err != nil {
			fmt.Println("uh oh")
		}
		sum += num
	}
	fmt.Println(sum)
}

// part two
var numbers = [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var reverseNumbers = [...]string{"orez", "eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}

func firstWord(line string, numberWords [10]string) (int, int) {
	distances := make([]int, len(numbers))
	for i, numberWord := range numberWords {
		distances[i] = strings.Index(line, numberWord)
	}
	small := smallestNum(distances)
	return distances[small], small
}

// returns the position of the smallest number ("zero" is the 0th position)
func smallestNum(numbers []int) int {
	small := -1
	smallPos := -1
	for i, num := range numbers {
		if small == -1 || (num < small && num != -1) {
			small = num
			smallPos = i
		}
	}
	return smallPos
}

func firstNumOrWord(line string) string {
	pos, first := firstNum(line)
	wordPos, wordFirst := firstWord(line, numbers)

	if (pos < wordPos && pos != -1) || wordPos == -1 {
		return first
	} else {
		return fmt.Sprint(wordFirst)
	}
}

func lastNumOrWord(line string) string {
	reversed := reverse(line)

	pos, first := firstNum(reversed)
	wordPos, wordFirst := firstWord(reversed, reverseNumbers)

	if (pos < wordPos && pos != -1) || wordPos == -1 {
		return first
	} else {
		return fmt.Sprint(wordFirst)
	}
}

func betterCalibrate() {
	sum := 0
	lines := Lines("1-1")
	for _, line := range lines {
		first := firstNumOrWord(line)
		last := lastNumOrWord(line)
		num, err := strconv.Atoi(first + last)
		if err != nil {
			fmt.Println("error converting result to integer: ", first, ", ", last)
		}
		sum += num
	}
	fmt.Println(sum)
}
