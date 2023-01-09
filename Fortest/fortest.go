// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func routine() {
// 	for {
// 		select {
// 		case <-pause:
// 			fmt.Println("pause")
// 			select {
// 			case <-play:
// 				fmt.Println("play")
// 				work()
// 			case <-quit:
// 				wg.Done()
// 				return
// 			}
// 		case <-quit:
// 			wg.Done()
// 			return
// 		default:
// 			work()
// 		}
// 	}
// }

// func main() {
// 	wg.Add(1)
// 	go routine()

// 	time.Sleep(1 * time.Second)
// 	pause <- struct{}{}

// 	time.Sleep(1 * time.Second)
// 	play <- struct{}{}

// 	// time.Sleep(1 * time.Second)
// 	// pause <- struct{}{}

// 	// time.Sleep(1 * time.Second)
// 	// play <- struct{}{}

// 	time.Sleep(1 * time.Second)
// 	close(quit)

// 	wg.Wait()
// 	fmt.Println("done")
// }

// func work() {
// 	time.Sleep(250 * time.Millisecond)
// 	i++
// 	fmt.Println(i)
// }

// var play = make(chan struct{})
// var pause = make(chan struct{})
// var quit = make(chan struct{})
// var wg sync.WaitGroup
// var i = 0

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Stop and remove a container
func stopAndRemoveContainer(client *client.Client, containername string) error {
	ctx := context.Background()

	if err := client.ContainerStop(ctx, containername, nil); err != nil {
		log.Printf("Unable to stop container %s: %s", containername, err)
	}

	removeOptions := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}

	if err := client.ContainerRemove(ctx, containername, removeOptions); err != nil {
		log.Printf("Unable to remove container: %s", err)
		return err
	}

	return nil
}

func main() {
	client, err := client.NewEnvClient()
	if err != nil {
		fmt.Printf("Unable to create docker client: %s", err)
	}

	// Stops and removes a container
	stopAndRemoveContainer(client, "containername")
}
