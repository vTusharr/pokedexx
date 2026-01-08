package main

import "fmt"


func  commandExplore(cfg *config,s []string) error{

	var curr_name *string
	curr_name=&s[0]


	ExloreResp,err:=cfg.pokeapiClient.GetLocation(curr_name)


	if err!=nil{
		return err
	}
   fmt.Println("Exploring",ExloreResp.Name)

   fmt.Println("Found Pokemon:")
   for _, j := range ExloreResp.PokemonEncounters {
   	fmt.Println(j.Pokemon.Name)
   }




	return nil
}
