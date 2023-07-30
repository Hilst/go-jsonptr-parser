package model

type Description struct {
	Variables []Variable
	Parents   []Content
}

type Variable struct {
	Id         string
	Pointer    string
	Typing     string
	Formatting string
}

type Content struct {
	Structure string
	Labeling  string
	Children  []Content
}

/**
Description{
	Variables: [
		Variable{
			"Id": "invoice-total-value"
			"Pointer": "/wrappers/.../invoice/total_value"
			"Typing": "floating-point"
			"Formatting": "currency"
		}
	]
	Parents: [
		Content{
			"Structure": "page"
			"Labeling": "Invoice"
			"Children": [
				{
					"Structure": "light-label"
					"Labeling": "{{invoice-total-value}}"
					"Children": []
				}
			]
		}
	]
}
*/
