package cmd

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/greenteam1/executioner/pkg/models"
)

//supported languages
const (
	cpp    = "c++"
	golang = "golang"
)

func init() {
	os.Setenv("SOLUTION_FOLDER_PATH", "../../../solutions")
}

func TestExecutorOK(t *testing.T) {
	e, err := NewExecutor(
		Config("../../../configs/build_configs.json"),
	)
	if err != nil {
		log.Println(err)
	}
	var memoryLimit, timeLimit int64
	var code, language string
	timeLimit = 1000
	memoryLimit = 64 * 1024
	var testCase1 = &models.TestCase{
		TestData: []byte("3 3"),
		Answer:   []byte("6"),
	}
	var testCase2 = &models.TestCase{
		TestData: []byte("4 3"),
		Answer:   []byte("7"),
	}
	var testCase3 = &models.TestCase{
		TestData: []byte("4 4"),
		Answer:   []byte("8"),
	}
	code = `#include <iostream>
#include <vector>
int main()
{
	int a, b;
	std::cin >> a >> b;
	std::cout << a + b;
	return 0;
}`
	language = cpp
	var req = &models.Solution{
		Solution:    []byte(code),
		MemoryLimit: memoryLimit,
		TimeLimit:   timeLimit,
		Language:    language,
		TestCases:   []*models.TestCase{testCase1, testCase2, testCase3},
	}

	var ans *models.TestResults
	ans, _ = e.TestUserSolution(&models.AggregatedSolution{
		ID:       "1",
		Solution: req,
	})
	var correctAnswers = []string{testPassed, testPassed, testPassed}
	assert.Equal(t, int64(3), ans.PassedTestsCount)
	for i, item := range ans.TestResults {
		if item != nil {
			assert.Equal(t, correctAnswers[i], item.Result)
		}
	}

	language = golang
	code = `package main
	import (
		"fmt"
		"os"
	)
	func main() {
		var a, b int
		fmt.Fscan(os.Stdin, &a)
		fmt.Fscan(os.Stdin, &b)

		fmt.Print(a + b)
	}`
	var req2 = &models.Solution{
		Solution:    []byte(code),
		MemoryLimit: memoryLimit,
		TimeLimit:   timeLimit,
		Language:    language,
		TestCases:   []*models.TestCase{testCase1, testCase2, testCase3},
	}
	ans, _ = e.TestUserSolution(&models.AggregatedSolution{
		ID:       "2",
		Solution: req2,
	})
	assert.Equal(t, int64(3), ans.PassedTestsCount)
	for i, item := range ans.TestResults {
		if item != nil {
			assert.Equal(t, correctAnswers[i], item.Result)
		}
	}
}

func TestExecutorWA(t *testing.T) {
	e, err := NewExecutor(
		Config("../../../configs/build_configs.json"),
	)
	if err != nil {
		log.Println(err)
	}
	var memoryLimit, timeLimit int64
	var code, language string
	timeLimit = 1000
	memoryLimit = 64 * 1024
	var testCase1 = &models.TestCase{
		TestData: []byte("3 3"),
		Answer:   []byte("6"),
	}
	var testCase2 = &models.TestCase{
		TestData: []byte("4 3"),
		Answer:   []byte("7"),
	}
	var testCase3 = &models.TestCase{
		TestData: []byte("4 4"),
		Answer:   []byte("9"),
	}
	code = `#include <iostream>
#include <vector>
int main()
{
	int aa, bb;
	std::cin >> aa >> bb;
	std::cout << aa + bb;
	return 0;
}`
	language = cpp
	var req = &models.Solution{
		Solution:    []byte(code),
		MemoryLimit: memoryLimit,
		TimeLimit:   timeLimit,
		Language:    language,
		TestCases:   []*models.TestCase{testCase1, testCase2, testCase3},
	}
	var ans *models.TestResults
	ans, _ = e.TestUserSolution(&models.AggregatedSolution{
		ID:       "3",
		Solution: req,
	})
	var correctAnswers = []string{testPassed, testPassed, wrongAnswer}
	assert.Equal(t, int64(2), ans.PassedTestsCount)
	for i, item := range ans.TestResults {
		if item != nil {
			assert.Equal(t, correctAnswers[i], item.Result)
		}
	}
}

func TestExecutorTL(t *testing.T) {
	e, err := NewExecutor(
		Config("../../../configs/build_configs.json"),
	)
	if err != nil {
		log.Println(err)
	}
	var memoryLimit, timeLimit int64
	var code, language string
	timeLimit = 2
	memoryLimit = 64 * 1024
	var testCase1 = &models.TestCase{
		TestData: []byte("3 3"),
		Answer:   []byte("6"),
	}
	var testCase2 = &models.TestCase{
		TestData: []byte("4 3"),
		Answer:   []byte("7"),
	}
	var testCase3 = &models.TestCase{
		TestData: []byte("4 4"),
		Answer:   []byte("9"),
	}
	code = `package main
import (
	"fmt"
	"os"
)
func main() {
	var a, b int
	var tmp []int
	fmt.Fscan(os.Stdin, &a)
	fmt.Fscan(os.Stdin, &b)
	for i := 0; i < 10000000; i++{
		tmp = append(tmp, 1)
	}
	fmt.Print(a + b)
}`
	language = golang
	var req = &models.Solution{
		Solution:    []byte(code),
		MemoryLimit: memoryLimit,
		TimeLimit:   timeLimit,
		Language:    language,
		TestCases:   []*models.TestCase{testCase1, testCase2, testCase3},
	}
	var ans *models.TestResults
	ans, _ = e.TestUserSolution(&models.AggregatedSolution{
		ID:       "4",
		Solution: req,
	})
	var correctAnswers = []string{"TL"}
	assert.Equal(t, int64(0), ans.PassedTestsCount)
	for i, item := range ans.TestResults {
		if item != nil {
			assert.Equal(t, correctAnswers[i], item.Result)
		}
	}
}

func TestExecutorCE(t *testing.T) {
	e, err := NewExecutor(
		Config("../../../configs/build_configs.json"),
	)
	if err != nil {
		log.Println(err)
	}
	var memoryLimit, timeLimit int64
	var code, language string
	timeLimit = 2
	memoryLimit = 64 * 1024
	var testCase1 = &models.TestCase{
		TestData: []byte("3 3"),
		Answer:   []byte("6"),
	}
	var testCase2 = &models.TestCase{
		TestData: []byte("4 3"),
		Answer:   []byte("7"),
	}
	var testCase3 = &models.TestCase{
		TestData: []byte("4 4"),
		Answer:   []byte("9"),
	}
	code = `package main
import (
	"fmt"
	"os"
	"time"
)
func main() {
	var a, b int
	var tmp []int
	fmt.Fscan(os.Stdin, &a)
	fmt.Fscan(os.Stdin, &b)
	for i := 0; i < 100000000; i++{
		tmp = append(tmp, 1)
	}
	fmt.Print(a + b)
}`
	language = golang
	var req = &models.Solution{
		Solution:    []byte(code),
		MemoryLimit: memoryLimit,
		TimeLimit:   timeLimit,
		Language:    language,
		TestCases:   []*models.TestCase{testCase1, testCase2, testCase3},
	}
	var ans *models.TestResults
	ans, _ = e.TestUserSolution(&models.AggregatedSolution{
		ID:       "5",
		Solution: req,
	})
	var correctAnswers = []string{compilationError}
	assert.Equal(t, int64(0), ans.PassedTestsCount)
	for i, item := range ans.TestResults {
		if item != nil {
			assert.Equal(t, correctAnswers[i], item.Result)
		}
	}
}

func TestWrongLanguage(t *testing.T) {
	e, err := NewExecutor(
		Config("../../../configs/build_configs.json"),
	)
	if err != nil {
		log.Println(err)
	}
	var memoryLimit, timeLimit int64
	var code, language string
	timeLimit = 1000
	memoryLimit = 64 * 1024
	var testCase1 = &models.TestCase{
		TestData: []byte("3 3"),
		Answer:   []byte("6"),
	}
	var testCase2 = &models.TestCase{
		TestData: []byte("4 3"),
		Answer:   []byte("7"),
	}
	var testCase3 = &models.TestCase{
		TestData: []byte("4 4"),
		Answer:   []byte("8"),
	}
	code = `#include <iostream>
#include <vector>
int main()
{
	int aaa, bbb;
	std::cin >> aaa >> bbb;
	std::cout << aaa + bbb;
	return 0;
}`
	language = "asvsafsada"
	var req = &models.Solution{
		Solution:    []byte(code),
		MemoryLimit: memoryLimit,
		TimeLimit:   timeLimit,
		Language:    language,
		TestCases:   []*models.TestCase{testCase1, testCase2, testCase3},
	}
	_, err = e.TestUserSolution(&models.AggregatedSolution{
		ID:       "6",
		Solution: req,
	})
	assert.Equal(t, err, fmt.Errorf("ERROR: %s is not supported language by executor", language))
}
