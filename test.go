package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// func main() {
// 	randomTakeBall(10, 10000*time.Millisecond)
// }

// func randomTakeBall(numGoroutines int, duration time.Duration) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	ballChan := make(chan struct{}, 1)
// 	var wg sync.WaitGroup
// 	wg.Add(numGoroutines)

// 	for idx := range numGoroutines {
// 		go func(goroutineID int) {
// 			defer wg.Done()
// 			isHaveBall := false
// 			for {
// 				select {
// 				case ballChan <- struct{}{}:
// 					fmt.Printf("Goroutine %d took the ball\n", goroutineID+1)
// 					isHaveBall = timeTakeBall(ctx, ballChan)
// 				case <-ctx.Done():
// 					if isHaveBall {
// 						fmt.Printf("Goroutine %d finished with ball\n", goroutineID+1)
// 					} else {
// 						// fmt.Printf("Goroutine %d has no ball\n", goroutineID+1)
// 					}
// 					return
// 				}
// 			}
// 		}(idx)
// 	}

// 	time.Sleep(duration)
// 	cancel()
// 	wg.Wait()
// }

// func timeTakeBall(ctx context.Context, ball chan struct{}) bool {
// 	select {
// 	case <-time.After(time.Second):
// 		<-ball
// 		return false
// 	case <-ctx.Done():
// 		return true
// 	}
// }

func main() {
	rand.Seed(time.Now().UnixNano())
	playGame(10, 10*time.Second) // Chỉ truyền vào một timeout duy nhất
}

func playGame(numPlayers int, timeout time.Duration) {
	var mu sync.Mutex
	ballOwner := randomNewOwner(0, numPlayers)
	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(numPlayers)

	// Tạo goroutines cho từng người chơi
	for i := range numPlayers {
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-stop:
					fmt.Printf("\nPlayer %d %s ball\n", id, map[bool]string{true: "has", false: "has no"}[ballOwner == id])
					return
				default:
					if ballOwner == id {
						fmt.Printf("Player %d took the ball\n", id)
						time.Sleep(1 * time.Second)
						mu.Lock()
						ballOwner = randomNewOwner(id, numPlayers)
						mu.Unlock()
					}
				}
			}
		}(i + 1)
	}

	time.Sleep(timeout)
	close(stop)
	wg.Wait()
}

func randomNewOwner(current, numPlayers int) int {
	for {
		newOwner := rand.Intn(numPlayers) + 1
		if newOwner != current {
			return newOwner
		}
	}
}
