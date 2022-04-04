package transform

import "Make_Api_Req/get_data"

func Group_DataByIndex(u []get_data.User) map[string][]get_data.User {
	index_group := make(map[string][]get_data.User)
	for _, value := range u {
		if len(value.First) != 0 {
			index_group[value.First[0:1]] = append(index_group[value.First[0:1]], value)
		}
	}
	return index_group
}
