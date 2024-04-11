package athletes

import (
	"fmt"
	"testing"
)

func TestAthleteService_GetAthletes(t *testing.T) {
	var service AthleteService
	res := service.GetAthletes()
	fmt.Println(res)
}
