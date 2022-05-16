package annalyn

//import "fmt"

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
    return !knightIsAwake
    //fmt.Println(CanFastAttack(knightIsAwake)) // Output: false
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
    return knightIsAwake || archerIsAwake || prisonerIsAwake
    //fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake)) // Output: true
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return !archerIsAwake && prisonerIsAwake
    //fmt.Println(CanSignalPrisoner(archerIsAwake, prisonerIsAwake)) // Output: true
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool { 
    case1 := !knightIsAwake && !archerIsAwake && prisonerIsAwake
    case2 := !archerIsAwake && petDogIsPresent
    return case1 || case2
    //fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent)) // Output: false
}

