package main
import "fmt"
func main(){
    
    myMap:=make(map[string]int)
    
    myMap["apple"]=1
    myMap["banana"]=2
    myMap["orange"]=3
   
    
    appleValue:=myMap["apple"]
    bananaValue:=myMap["banana"]
    fmt.Println("value of apple:",appleValue)
    fmt.Println("value of banana:",bananaValue)
   
    
    myMap["apple"]=5
    fmt.Println("updated value of apple",myMap["apple"])
   
    
    delete(myMap,"orange")
    fmt.Println("after deleting orange: ",myMap)
   
    
    value, exists:=myMap["banana"]
    if exists{
        fmt.Println("value of banana: ",value)
    }else{
        fmt.Println("banana not found in the map")
    }
   
    
    for key,value:=range myMap{
        fmt.Println("key: ",key,"value:",value)
    }
    
    fmt.Println("length of map:",len(myMap))
   
}
