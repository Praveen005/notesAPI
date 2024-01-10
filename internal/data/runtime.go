//Named the file runtime.go cz , I want to chnage the data type of the Runtime field in the movie struct :p
//And To prevent our internal/data/movie.go file from getting cluttered
/*

when Go is encoding a particular type to JSON it looks to see if the type
satisfies the json.Marshaler interface, which looks like this:

	type Marshaler interface {
		MarshalJSON() ([]byte, error)
	}

	If the type does satisfy the interface, then Go will call its MarshalJSON() method and use the
	[]byte slice that it returns as the encoded JSON value.
		means, agar aapne, MarshalJSON(), jo ki user define karta hai, agar kar rakha hai, then, that field of 
		JSON will be encoded as per the logic you write inside this method.

		agar aapne nahi define kiya hai ye method, to encoding phir defautlt WAY ME KAREGI GO, jo internally defined hai. 
	
	If the type doesn’t have a MarshalJSON() method, then Go will fall back to trying to encode
	it to JSON based on its own internal set of rules

So, if you want to customize how something is encoded, all we need to do is implement a
MarshalJSON() method on it which returns a custom JSON representation of itself in a
[]byte slice.


*/

package data

import(
	"fmt"
	"strconv"
)


// Declare a custom Runtime type, which has the underlying type int32 (the same as our
// Movie struct field).
type Runtime int32


// Implement a MarshalJSON() method on the Runtime type so that it satisfies the
// json.Marshaler interface. This should return the JSON-encoded value for the movie
// runtime (in our case, it will return a string in the format "<runtime> mins").

func (r Runtime) MarshalJSON()([]byte, error){
	// Generate a string containing the movie runtime in the required format.
	jsonValue := fmt.Sprintf("%d mins", r)


	// Use the strconv.Quote() function on the string to wrap it in double quotes. It
	// needs to be surrounded by double quotes in order to be a valid *JSON string*.

	quotedJsonValue := strconv.Quote(jsonValue)

	//convert the quoted string into a byte of slice and return it.
	//Remember, we are customising the Json, and hence writing our very own MarshalJSON() function 
	//which is present in Marshaler interface

	return []byte(quotedJsonValue), nil
}