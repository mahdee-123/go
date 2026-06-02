package main


func SumOfIntsOrFloats[K comparable , V int64 | float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}
func main() {

}