package counter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadAndIncrementCounter(filename string) (int, error) {
	currentNum := 0
	content, err := os.ReadFile(filename)
	if err == nil {
		numStr := strings.TrimSpace(string(content))
		currentNum, err = strconv.Atoi(numStr)
		if err != nil {
			return 0, fmt.Errorf("error parsing number from file: %v", err)
		}
	} else if !os.IsNotExist(err) {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	currentNum++
	newContent := fmt.Sprintf("%d", currentNum)

	err = os.WriteFile(filename, []byte(newContent), 0644)
	if err != nil {
		return 0, fmt.Errorf("error writing to file: %v", err)
	}

	return currentNum, nil
}
