package services

var userCache = make(map[int][]int)

func OrderQuery(userID int) []int {
	cache, e := userCache[userID]
	if !e {
		cache = []int{}
		userCache[userID] = cache
	}

	return cache
}
func OrderSubmit(userID int, submit interface{}) int {
	cache, e := userCache[userID]
	if !e {
		cache = []int{}
	}

	l := len(cache)
	orderID := userID*100 + l + 1
	cache = append(cache, orderID)
	userCache[userID] = cache
	return orderID
}
func OrderCancel(userID int, orderID int) bool {
	cache, e := userCache[userID]
	if !e {
		return false
	}
	i := 0
	l := len(cache)
	for ; i < l; i++ {
		if cache[i] == orderID {
			break
		}
	}
	if i == l {
		return false
	}

	var ncache []int
	ncache = append(ncache, cache[0:i]...)
	ncache = append(ncache, cache[i+1:l]...)
	userCache[userID] = ncache
	return true
}
