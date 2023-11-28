package models

import (
	"math/rand"
	"time"
)

var (
	RandomSource = rand.NewSource(time.Now().UnixNano())
)

func WaitForPosition(Id int, targetX float64) {
	for {
		carPos := FindCarPosition(Id)
		if carPos.X >= targetX {
			break
		}
		time.Sleep(16 * time.Millisecond)
	}
}

func FindAvailableLane() (int, bool) {
	LaneMutex.Lock()
	defer LaneMutex.Unlock()

	rand := rand.New(RandomSource)
	lanes := rand.Perm(numLanes)

	for _, lane := range lanes {
		if !LaneStatus[lane] {
			LaneStatus[lane] = true
			return lane, true
		}
	}

	return -1, false
}

func Lane(Id int) {
	CreateCar(Id)
	WaitForPosition(Id, 100)

	lane, foundLane := FindAvailableLane()

	if !foundLane {
		ResetCarPosition(Id)
		return
	}

	AssignLaneToCar(Id, lane)
}
