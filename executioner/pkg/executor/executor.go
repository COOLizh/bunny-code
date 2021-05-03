//Package executor provides executor functionalities
package executor

import "gitlab.com/greenteam1/executioner/pkg/models"

//Executor provides executor functionality
type Executor interface {
	TestUserSolution(sol *models.AggregatedSolution) (*models.TestResults, error)
	GetLanguages() ([]string, error)
}
