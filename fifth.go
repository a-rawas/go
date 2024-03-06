package main

import "fmt"

func main(){

	FavLocations := make(map[string] int)

	FavLocations["Home"] = 1
	FavLocations["Office"] = 2
	//possible to use key as int
	//FavLocations[1] = 1


	fmt.Println(FavLocations)
	fmt.Println("Favorite location for Home", FavLocations["Home"], "out of ", len(FavLocations))

	superhero := map[string]map[string]string{
		"Superman":map[string]string{
			"RealName":"Clark Kent",
			"City":"Metropolis",
		},
		"Batman":map[string]string{
			"RealName":"Bruce Wayne",
			"City":"Gotham",
		},
	}

	if temp, hero:=superhero["Superman"]; hero{
		fmt.Println(temp["RealName"], temp["City"])
	}

}