package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avPrepTime int) int {
	if avPrepTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * avPrepTime
}

// Define the function Quantities that takes an array of layers as parameter as a []string. The function will then determine the quantity of noodles and sauce needed to make your meal. The result should be returned as two values of noodles as an int and sauce as a float64.
func Quantities(layers []string) (int, float64) {
	var noodles, sauce int

	for i := 0; i < len(layers); i++ {
		layer := layers[i]
		if layer == "noodles" {
			noodles++
		} else if layer == "sauce" {
			sauce++
		}
	}
	return noodles * 50, float64(sauce) * 0.2
}

// Write a function AddSecretIngredient that accepts two arrays of ingredients of type []string as parameters. The first parameter is the list your friend sent you, the second is the ingredient list for your own recipe. The function should generate a new slice and add the last item from your friends list to the end of your list. Neither argument should be modified.

func AddSecretIngredient(friendList, myList []string) []string {
	return append(myList, friendList[len(friendList)-1])
}

// Implement a function ScaleRecipe that takes two parameters.
//     A slice of float64 amounts needed for 2 portions.
//     The number of portions you want to cook.
// The function should return a slice of float64 of the amounts needed for the desired number of portions. You want to keep the original recipe though. This means the amounts argument should not be modified in this function.

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var newquantities []float64
	for i := 0; i < len(quantities); i++ {
		newquantities = append(newquantities, quantities[i]*float64(portions)*0.5)
	}
	return newquantities
}
