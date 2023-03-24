package repository

// Nothing special here, just helper GO functions to make the
// code as DRY as possible and placed in own file to keep the
// main code clean

func flatten(repos *[]map[string]interface{}) []interface{} {

	if repos != nil {
		ois := make([]interface{}, len(*repos), len(*repos))

		for i, m := range *repos {
			ois[i] = pick_bones(m)
		}
		return ois
	}

	return make([]interface{}, 0)
}

func extract(repos map[string]interface{}) interface{} {

	if repos != nil {
		ois := make([]interface{}, 1, 1)
		ois[0] = pick_bones(repos)
		return ois
	}

	return make([]interface{}, 0)
}

func pick_bones(data map[string]interface{}) map[string]interface{} {
	oi := make(map[string]interface{})
	oi["id"] = data["id"]
	oi["name"] = data["name"]

	return oi
}
