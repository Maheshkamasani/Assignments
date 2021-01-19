package main
import "fmt"
func main() {
}
func cat(e error) error {
return fmt.Errorf("the error at cat %w", e)
}
func moo(e error)error {
cat(e)
return fmt.Errorf("the error at cat %w", e)
}