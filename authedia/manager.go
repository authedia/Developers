package authedia

import (
    "bytes"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "mime/multipart"
)


const URL = "https://authedia.uc.r.appspot.com/"

type ApiKey struct {
    ApiKey string
}

type Response struct {
    Success        bool
    Bytes          []byte
    Message        string
    // DataBlock      *data_packer.ExtDataBlock
    VerifiedBlocks [][]int
}


func Wrap(api_key, input_file_name, output_file_name string) *Response {
    /*
      PARAMS
          api_key : (str) : User's API key
          input_file_name : (str) : Path to input file
          output_file_name : (str) : Desired location of output file
      RETRUNS
          wrap_response : (dict) : See documentation for details
    */

    url_endpoint := URL + "wrap_request"

    file_bytes, err := ioutil.ReadFile(input_file_name)
    if err != nil {
        panic(err)
    }

    response := makeRequest(url_endpoint, api_key, file_bytes)

    if len(response.Bytes) > 0 {
        err = ioutil.WriteFile(output_file_name, response.Bytes, 0644)
        if err != nil {
            panic(err)
        }
    }

    return response
}


func Verify(api_key, input_file_name string) *Response {
    /*
      PARAMS
          api_key : (str) : User's API key
          input_file_name : (str) : Path to input file
      RETRUNS
          verify_request : (dict) : See documentation for details
    */
    url_endpoint := URL + "verify_request"

    file_bytes, err := ioutil.ReadFile(input_file_name)
    if err != nil {
        panic(err)
    }

    return makeRequest(url_endpoint, api_key, file_bytes)
}


func makeRequest(url_endpoint, api_key_str string, file_bytes []byte) *Response {
		// Buffer to store our request body as bytes
		var requestBody bytes.Buffer

		// Create a multipart writer
		multiPartWriter := multipart.NewWriter(&requestBody)

		// Initialize the file field
		fileWriter, err := multiPartWriter.CreateFormFile("data", "file_name")
		if err != nil {
			panic(err)
		}
		fileWriter.Write(file_bytes)

		// Populate other fields
		fieldWriter, err := multiPartWriter.CreateFormField("api_key")
		if err != nil {
				panic(err)
		}
		json.NewEncoder(fieldWriter).Encode(
        ApiKey{ApiKey : api_key_str},
    )

		// We completed adding the file and the fields, let's close the multipart writer
		// So it writes the ending boundary
		multiPartWriter.Close()

		// By now our original request body should have been populated, so let's just use it with our custom request
		req, err := http.NewRequest("POST", url_endpoint, &requestBody)
		if err != nil {
				panic(err)
		}

		// We need to set the content type from the writer, it includes necessary boundary as well
		req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

		// Do the request
		client := &http.Client{}
		http_response, err := client.Do(req)
		if err != nil {
				panic(err)
		}

		response := &Response{}
		json.NewDecoder(http_response.Body).Decode(response)
	  return response
}
