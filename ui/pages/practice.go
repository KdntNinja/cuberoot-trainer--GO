package pages

import (
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
)

// The Problem represents a cube root problem
type Problem struct {
	Cube       int    // The cube number to display
	CubeRoot   int    // The answer (cube root)
	UserAnswer int    // The user's answer
	Message    string // Feedback message
	IsCorrect  bool   // Whether the answer is correct
	Attempted  bool   // Whether an answer has been attempted
}

// GenerateProblem creates a new cube root problem
func GenerateProblem() Problem {
	// Generate a random number between 1 and 100
	min := 1
	max := 100
	cubeRoot := rand.Intn(max-min+1) + min

	// Cube the number
	cube := int(math.Pow(float64(cubeRoot), 3))

	return Problem{
		Cube:     cube,
		CubeRoot: cubeRoot,
	}
}

// CheckAnswer validates the user's answer against the correct answer
func (p *Problem) CheckAnswer(userAnswer int) {
	p.UserAnswer = userAnswer
	p.Attempted = true

	if userAnswer == p.CubeRoot {
		p.IsCorrect = true
		p.Message = "Correct! 🎉"
	} else {
		p.IsCorrect = false
		p.Message = "Incorrect. The cube root of " + strconv.Itoa(p.Cube) + " is " + strconv.Itoa(p.CubeRoot) + "."
	}
}

// HandlePractice processes form submissions for the practice page
func HandlePractice(w http.ResponseWriter, r *http.Request) {
	var problem Problem

	// Check if this is a form submission
	if r.Method == http.MethodPost {
		// Parse the form
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Get the cube and user answer from the form
		cubeStr := r.FormValue("cube")
		answerStr := r.FormValue("answer")

		cube, err := strconv.Atoi(cubeStr)
		if err != nil {
			http.Error(w, "Invalid cube value", http.StatusBadRequest)
			return
		}

		// Calculate the correct answer (cube root)
		cubeRoot := int(math.Round(math.Cbrt(float64(cube))))

		problem = Problem{
			Cube:     cube,
			CubeRoot: cubeRoot,
		}

		// If the answer field is not empty, check it
		if answerStr != "" {
			userAnswer, err := strconv.Atoi(answerStr)
			if err != nil {
				http.Error(w, "Invalid answer", http.StatusBadRequest)
				return
			}

			problem.CheckAnswer(userAnswer)

			// If the answer is correct, generate a new problem
			if problem.IsCorrect {
				problem = GenerateProblem()
				problem.Message = "Correct! 🎉" // Preserve the message
			}
		}
	} else {
		// For GET requests, just generate a new problem
		problem = GenerateProblem()
	}

	// Render the practice page with the problem
	templ.Handler(PracticeWithProblem(problem)).ServeHTTP(w, r)
}
