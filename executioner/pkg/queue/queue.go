/*Package queue provides queue functionality*/
package queue

import "gitlab.com/greenteam1/executioner/pkg/models"

//Queue interface for working with the queue data structure
type Queue interface {
	Push(solution *models.AggregatedSolution) error
	Pop() (models.AggregatedSolution, error)
}

//ChannelQueue provides an implementation of the queue on the channels
type ChannelQueue struct {
	ch chan *models.AggregatedSolution
}

//NewChannelQueue creates a new ChannelQueue object
func NewChannelQueue(maxTasksInChannel int) *ChannelQueue {
	return &ChannelQueue{
		ch: make(chan *models.AggregatedSolution, maxTasksInChannel),
	}
}

//Push pushes the value into the channel
func (q *ChannelQueue) Push(solution *models.AggregatedSolution) error {
	q.ch <- solution
	return nil
}

//Pop retrieves the value from the channel
func (q *ChannelQueue) Pop() (models.AggregatedSolution, error) {
	return *<-q.ch, nil
}
