package test

type SwapiTest struct {}

type Error struct {}

func (e *Error) Error() string {
	return "Error"
}

func (s *SwapiTest) NumOfAppearances(planet string) (int, error) {
	if planet == "Name" || planet =="Name2" {
		return 1, nil
	}

	a := Error{}
	return 0, &a

}

func (s *SwapiTest) ContainPlanet(planet string) bool {
	if planet == "Name" || planet =="Name2"{
		return true
	}
	return false
}