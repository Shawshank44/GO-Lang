package utils

func ExtractImages(content map[string]interface{}) []string {
	var urls []string

	data, ok := content["Data"].([]interface{})
	if !ok {
		return urls
	}

	for _, item := range data {
		product, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		images, ok := product["images"].([]interface{})
		if !ok {
			continue
		}

		for _, i := range images {
			block, ok := i.(map[string]interface{})
			if !ok {
				continue
			}

			url, ok := block["url"].(string)
			if !ok {
				continue
			}
			urls = append(urls, url)
		}
	}

	return urls
}
