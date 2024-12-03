package parts

import (
	"os"
	"project/utils"
	"regexp"
	"strconv"
	"strings"
)

func First() {

	sum := 0

	dat, err := os.ReadFile("data/input.txt")
	utils.Check(err)
	reg := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	muls := reg.FindAllString(string(dat), -1)

	// [\W\w]*don't\(\)
	// (?:do\((\)))[\W\w]*don't\(\)
	// (?:do\((\)))[\W]*(mul\([0-9]+,[0-9]+\))

	for _, mul := range muls {
		/* cut at index 4 */
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

	println(sum)

}
