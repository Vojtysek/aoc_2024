package parts

import (
	"os"
	"project/utils"
	"regexp"
	"strconv"
	"strings"
)

var sum int

func calcMul(muls []string) {

	reg := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	muls = reg.FindAllString(strings.Join(muls, ""), -1)

	for _, mul := range muls {
		first := strings.Split(mul, ",")[0]
		sec := strings.Split(mul, ",")[1]

		first = first[4:]
		sec = sec[:len(sec)-1]

		f, err := strconv.Atoi(first)
		if err != nil {
			panic(err)
		}

		s, err := strconv.Atoi(sec)
		if err != nil {
			panic(err)
		}

		sum += f * s
	}
}

func calcUntilDont(muls string) {
	reg := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	var mulsArr []string = reg.FindAllString(muls, -1)

	for _, mul := range mulsArr {
		first := strings.Split(mul, ",")[0]
		sec := strings.Split(mul, ",")[1]

		first = first[4:]
		sec = sec[:len(sec)-1]

		f, err := strconv.Atoi(first)
		if err != nil {
			panic(err)
		}

		s, err := strconv.Atoi(sec)
		if err != nil {
			panic(err)
		}

		sum += f * s
	}

}

func Second() {

	dat, err := os.ReadFile("data/input.txt")
	utils.Check(err)

	untilDont := strings.Split(string(dat), "don't()")[0]
	rest := strings.Split(string(dat), "don't()")[1:]

	restString := strings.Join(rest, "don't()")
	doTillDont := regexp.MustCompile(`(?:do\((\)))[\W\w]*?(don't\(\))`)
	var muls []string = doTillDont.FindAllString(restString, -1)

	calcUntilDont(untilDont)
	calcMul(muls)

	println(sum)

}
