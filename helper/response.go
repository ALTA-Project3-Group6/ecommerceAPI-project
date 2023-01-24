package helper

func ErrorResponse(msg string) interface{} {
	resp := map[string]interface{}{}
	resp["meesage"] = msg

	return resp
}
