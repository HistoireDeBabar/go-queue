##### Queue

A queue library to store data, and interact with it in an
first in, first out scenario.



####### Using a Queue

You can create and use a queue like so:

    import (
      "github.com/HistoireDeBabar/go-queue"
    )

    func main() {
      q := queue.Queue{}
      item, err := q.Pop() //item == nil, err == 'No Items in Queue'
      q.Push("hello")
      q.Push("world")
      length = q.Length() //length == 2
      item, _ = q.Pop() //item == "hello"
      item, _ = q.Pop() //item == "world"
      length = q.Length() //length == 0
    }



###### godoc

`type Queue struct {}

Queue Type for storing items in first in, first out order.

func (q *Queue) Length() int
func (q *Queue) Pop() (interface{}, error)
func (q *Queue) Push(item interface{})
`
