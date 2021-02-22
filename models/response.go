package models

/*
 * Response code
 * 0 : success complete
 *
 * 10 : internal server error
 * 11 : not found data (query result)
 * 12 : duplicate data (query result)
 */

/* Response format */
type ResponseFormat struct {
	Result		bool						`json:"result" xml:"result"`
	Code			int							`json:"code" xml:"code"`
	Message		[]interface{}		`json:"message" xml:"message"`
}