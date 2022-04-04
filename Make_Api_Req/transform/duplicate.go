package transform

import "Make_Api_Req/get_data"

func RemoveDuplicates(u []get_data.User) []get_data.User {
	k := make(map[get_data.User]struct{})
	var list []get_data.User
	for _, value := range u {
		if _, ok := k[value]; !ok {
			k[value] = struct{}{}
			list = append(list, value)
		}
	}
	return list
}
