package datamodel

import "github.com/hofstadter-io/hof/schema"

#User: schema.#Model & {
	Fields: {
		email: schema.#Email	
		persona: schema.#Enum & {
			vals: ["guest", "user", "admin", "owner"]
			default: "guest"
		}	
		password: schema.#Password
		active: schema.#Bool
	}
}

#UserProfile: schema.#Model & {
	Fields: {
		firstName: schema.#String
		middleName: schema.#String
		lastName: schema.#String
	}
}
