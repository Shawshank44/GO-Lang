package utils

func ExtractImages(content map[string]interface{}) []string {
	var urls []string

	blocks, ok := content["blocks"].([]interface{})
	if !ok {
		return urls
	}

	for _, b := range blocks {
		block, ok := b.(map[string]interface{})
		if !ok {
			continue
		}

		data, ok := block["data"].(map[string]interface{})
		if !ok {
			continue
		}

		file, ok := data["file"].(map[string]interface{})
		if !ok {
			continue
		}

		url, ok := file["url"].(string)
		if !ok {
			continue
		}

		urls = append(urls, url)
	}

	return urls
}
