func maximumWealth(accounts [][]int) int {
	richestCustomer := 0

	for _, customer := range accounts {
		if balance := sum(customer); balance > richestCustomer {
			richestCustomer = balance
		}
	}

	return richestCustomer
}

func sum(arr []int) int {
	total := 0
	for _, value := range arr {
		total += value
	}
	return total
}