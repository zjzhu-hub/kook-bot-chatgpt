package queues

type RequestWithCommand struct {
    Command string
    Body    map[string]interface{}
}

type ConcurrentQueue struct {
    queue chan RequestWithCommand
}

func NewConcurrentQueue(capacity int) *ConcurrentQueue {
    return &ConcurrentQueue{
        queue: make(chan RequestWithCommand, capacity),
    }
}

func (q *ConcurrentQueue) Push(c RequestWithCommand) {
    q.queue <- c
}

func (q *ConcurrentQueue) Pop() RequestWithCommand {
    return <-q.queue
}

func (q *ConcurrentQueue) Len() int {
    return len(q.queue)
}
