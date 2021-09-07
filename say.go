package gomod

import "fmt" 

// v1.0.0
func SayHiV100(name string) string {
   return fmt.Sprintf("Hi, v1.0.0 %s", name)
}

// v1.0.1
// func SayHiV101(name string) string {
//    return fmt.Sprintf("Hi, v1.0.1 %s", name)
// }