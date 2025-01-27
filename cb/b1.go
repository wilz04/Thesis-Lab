package cb

import (
	"lab3/ca"
	"os"
)

const OBS = 16

var serie = "1"

func Main() {
	serie = os.Args[1]
	ca.REWARDENABLED = serie == "1"

	var treats = getTreatments()
	var i, j int
	for i = 0; i < len(treats); i++ {
		for j = 0; j < OBS; j++ {
			run(treats[i], i, j)
		}
	}
}
