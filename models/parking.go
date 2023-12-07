package models

import "time"

func MoveCarsLogic() {
	for i := len(Cars) - 1; i >= 0; i-- {
		moveCar(&Cars[i])
	}
	ExitCarLogic()
}

func calculateParkingPosition(lane int, laneWidth float64) (float64, float64) {
	var targetX, targetY float64
	if lane < 10 {
		targetX = 100.0 + float64(lane)*laneWidth + laneWidth/2
		targetY = 400 + (500-350)/2
	} else {
		targetX = 100.0 + float64(lane-10)*laneWidth + laneWidth/2
		targetY = 100 + (250-100)/2
	}
	return targetX, targetY
}

func ExitCarLogic() {
	for i := len(Cars) - 1; i >= 0; i-- {
		exitParkedCar(&Cars[i])
	}
}

func exitParkedCar(car *Car) {
	if car.isParked && time.Now().After(car.ExitTime) && !car.IsEntering {
		handleExitingCar(car)
	}
}

func handleExitingCar(car *Car) {
	if !car.isTeleport {
		initiateTeleport(car)
	} else if time.Since(car.TeleportStartTime) >= time.Millisecond*500 {
		finalizeExit(car)
	}
}

func initiateTeleport(car *Car) {
	car.isTeleport = true
	car.TeleportStartTime = time.Now()
	car.Position.X = 50
	car.Position.Y = 400
}

func finalizeExit(car *Car) {
	updateLaneStatus(car.Lane, false)
	removeCarFromSlice(car)
}

func removeCarFromSlice(car *Car) {
	for i, c := range Cars {
		if c.ID == car.ID {
			Cars = append(Cars[:i], Cars[i+1:]...)
			break
		}
	}
}

func updateLaneStatus(lane int, status bool) {
	LaneMutex.Lock()
	defer LaneMutex.Unlock()
	LaneStatus[lane] = status
}
