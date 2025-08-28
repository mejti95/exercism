package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
    if timePerLayer == 0 {
        timePerLayer = 2
    }
    return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    var noodles int
    var sauce float64
    for i := range layers {
        switch layers[i] {
            case "noodles":
            	noodles += 50
            case "sauce":
            	sauce += 0.2
        }
    }
    return noodles, sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
    secretIngredient := friendList[len(friendList) - 1]
    myList = myList[:len(myList) -1]
    myList = append(myList, secretIngredient)
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    var newQuantities []float64

    for i := range quantities {
        q := quantities[i] / 2.0 * float64(portions)
        newQuantities = append(newQuantities, q)
    }
    return newQuantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
