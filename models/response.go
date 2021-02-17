package models

/*
 * Response code
 * 100 : success complete
 * 101 : not found data (query result)
 *
 * 200 : internal server error
 */

/* Response format */
type ResponseFormat struct {
	Result		bool						`json:"result" xml:"result"`
	Code			int							`json:"code" xml:"code"`
	Message		[]interface{}		`json:"message" xml:"message"`
}