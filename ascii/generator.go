package ascii

import (
	"bufio"
	"os"
	"io"
	"strings"
)

func GenerateAscii(text, banner string) string{
	file, err := os.Open("ascii/banners/" + banner + ".txt")
	if err != nil {
		return "Error opening banner"
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var asciiRead []string
	for{
		line, err := reader.ReadString('\n')
		line1 := strings.TrimRight(line, "\r\n")
		asciiRead = append(asciiRead, line1)
		if err != nil{
			if err == io.EOF{
				break
			}
			return "Error loading banner"
		}
	}

	var result []string
		for j := 0; j < 8; j++{
			var oneLineString strings.Builder
			
			for _, char := range text{
				if char < 32 || char > 126{
					continue
				}
				start := int(char-32) * 9 + 1
				oneLineString.WriteString((asciiRead[start + j]))
			}
			result = append(result, oneLineString.String())
		}
	return strings.Join(result, "\n") + "\n"
}