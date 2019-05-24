package gamelog

import "log"

// ToConsole func
func ToConsole(gameInfo informer, request string, err error) {
	log.Printf("New Action by %s\n", gameInfo.LastPlayer().Name())
	log.Printf("Action is %s\n", request)
	log.Printf("Any error raised: %+v\n", err)
	log.Printf("Player info after action: %+v\n", gameInfo.LastPlayer())
	log.Printf("Game info after action: %+v\n", gameInfo)
}
