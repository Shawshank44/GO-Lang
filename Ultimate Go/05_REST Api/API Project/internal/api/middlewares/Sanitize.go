package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"restapi/pkg/utils"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

func XSSMiddleWares(next http.Handler) http.Handler {
	fmt.Println("******** Initializing XSSMiddleWare ********")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("******** XSSMiddleWare executed ********")
		// Sanitize the URL Path:
		sanitizedPath, err := Clean(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("Original Path", r.URL.Path)
		fmt.Println("Sanitized Path", sanitizedPath)

		// Sanitize the Query params
		params := r.URL.Query()
		sanitizedQuery := make(map[string][]string)
		for key, values := range params {
			sanitizedKey, err := Clean(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			var SanitizedValues []string
			for _, value := range values {
				cleanValue, err := Clean(value)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				SanitizedValues = append(SanitizedValues, cleanValue.(string))
			}
			sanitizedQuery[sanitizedKey.(string)] = SanitizedValues
			fmt.Printf("Original query %s : %s \n", key, strings.Join(values, ", "))
			fmt.Printf("Sanitized query %s : %s \n", sanitizedKey, strings.Join(SanitizedValues, ", "))
		}

		r.URL.Path = sanitizedPath.(string)
		r.URL.RawQuery = url.Values(sanitizedQuery).Encode()
		fmt.Println("Updated URL : ", r.URL.String())

		// Sanitize request body
		if r.Header.Get("Content-Type") == "application/json" {
			if r.Body != nil {
				bodybytes, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, utils.ErrorHandler(err, "Error reading request body").Error(), http.StatusBadRequest)
					return
				}

				bodyString := strings.TrimSpace(string(bodybytes))
				fmt.Println("Original Body", bodyString)

				// Reset the request :
				r.Body = io.NopCloser(bytes.NewReader([]byte(bodyString)))

				if len(bodyString) > 0 {
					var inputData interface{}
					err := json.NewDecoder(bytes.NewReader([]byte(bodyString))).Decode(&inputData)
					if err != nil {
						http.Error(w, utils.ErrorHandler(err, "Invalid JSON body").Error(), http.StatusBadRequest)
						return
					}
					fmt.Println("Original Json body", inputData)

					// Sanitize the json body
					sanitizedData, err := Clean(inputData)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					fmt.Println("Sanitized data", sanitizedData)

					// Marshal the sanitized data back to the body
					sanitizedbody, err := json.Marshal(sanitizedData)
					if err != nil {
						utils.ErrorHandler(err, "")
						http.Error(w, utils.ErrorHandler(err, "Error Sanitizing body").Error(), http.StatusBadRequest)
						return
					}
					r.Body = io.NopCloser(bytes.NewReader(sanitizedbody))
					fmt.Println("Sanitized body", string(sanitizedbody))
				} else {
					log.Println("Request body is empty")
				}
			} else {
				log.Printf("No body in the request")
			}
		} else if r.Header.Get("Content-Type") != "" {
			log.Printf("Recieved request with unsupported Content-Type : %s. Expected application/json \n", r.Header.Get("Content-Type"))
			http.Error(w, "Unsupported Content-Type", http.StatusUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
		fmt.Println("Sending response from XSSMiddleWare")
	})
}

// Clean Sanitizes input data to prevent XSS attacks
func Clean(data interface{}) (interface{}, error) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			v[key] = SanitizeValue(value)
		}
		return v, nil
	case []interface{}:
		for i, value := range v {
			v[i] = SanitizeValue(value)
		}
		return v, nil
	case string:
		return SanitizeValue(v), nil
	default:
		return nil, utils.ErrorHandler(fmt.Errorf("unsupported type: %T", data), fmt.Sprintf("unsupported type: %T", data))
	}

}

func SanitizeValue(data interface{}) interface{} {
	switch v := data.(type) {
	case string:
		return SantizeString(v)
	case map[string]interface{}:
		for key, value := range v {
			v[key] = SanitizeValue(value)
		}
		return v
	case []interface{}:
		for i, value := range v {
			v[i] = SanitizeValue(value)
		}
		return v
	default:
		return v // Return v as it is unsupported
	}
}

func SantizeString(value string) string {
	return bluemonday.UGCPolicy().Sanitize(value)
}
