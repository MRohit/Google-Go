package main

import "fmt"

func main() {
   num := []int{11,32,34,21,23,12,55,44} 
   
   for i:= range num {
      fmt.Println("Slice item",i,"is",num[i])
   }
   countryCapitalMap := map[string] string {"India":"Delhi","Italy":"Rome","England":"London"}
   
   fmt.Println("\nPrinting map using Keys")
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
   
   fmt.Println("\nPrinting map using key-value pairs")
   for country,capital := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",capital)
   }
}