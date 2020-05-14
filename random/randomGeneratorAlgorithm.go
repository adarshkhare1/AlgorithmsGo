package random

type RandomGeneratorAlgorithm string

const(
	// Generate uniform distribution random numbers using Linear Congruential Method
	//Algorithm -
	//NumNext = (Multiplier*Num + increment) mod M
	// Num = Numnext
	Congruential RandomGeneratorAlgorithm = "Congruential"

	// Generate uniform distribution random numbers using Additive Number generator
	//Algorithm -
	//To output X[n], Y[j] = X[n-24], Y[k]=X[n-55]
	// Y[k] = Y[k]+Yj] mod 2^n
	// X[n] = Y[k]
	Additive = "Additive"
)
