package trail

type Trail struct {
	ID         string `json:"id"`
	Name       string `json:"AccessName"`
	BikeTrail  string `json:"bikeTrail"`
	Fishing    string `json:"Fishing"`
	Difficulty string `json:"ADAtrail"`
	Fee        string `json:"Fee"`
	Picnic     string `json:"Picnic"`
	Restrooms  string `json:"RestRooms"`
	Address    string `json:"Address"`
}

var Trails []Trail

func GetTrailByID(id string) *Trail {
	for _, trail := range Trails {
		if trail.ID == id {
			return &trail
		}
	}
	return nil
}

func GetFreeBikingTrails() []Trail {
	var freeBikeTrails []Trail
	for _, trail := range Trails {
		if trail.BikeTrail == "Yes" && trail.Fee == "No" {
			freeBikeTrails = append(freeBikeTrails, trail)
		}
	}
	return freeBikeTrails
}

func GetDifficultBikeTrails() []Trail {
	var diffBikeTrails []Trail
	for _, trail := range Trails {
		if trail.BikeTrail == "Yes" && trail.Difficulty == "Most Difficult" {
			diffBikeTrails = append(diffBikeTrails, trail)
		}
	}
	return diffBikeTrails
}

func GetFreeFishingTrails() []Trail {
	var freeFishingTrails []Trail
	for _, trail := range Trails {
		if trail.Fishing == "Yes" && trail.Fee == "No" {
			freeFishingTrails = append(freeFishingTrails, trail)
		}
	}
	return freeFishingTrails
}

func GetPicnicWithRestrooms() []Trail {
	var picnicSpots []Trail
	for _, trail := range Trails {
		if trail.Picnic == "Yes" && trail.Restrooms == "Yes" {
			picnicSpots = append(picnicSpots, trail)
		}
	}
	return picnicSpots
}
