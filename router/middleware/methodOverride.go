package middleware

import (
	"github.com/gin-gonic/gin"
)

// see: https://www.alexedwards.net/blog/http-method-spoofing
//<form method="POST" action="/">
//	<input type="hidden" name="_method" value="PUT">
//	<label>Example field</label>
//	<input type="text" name="example">
//	<button type="submit">Submit</button>
//</form>
func MethodOverrideMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		// Only act on POST requests.
		if r.Method == "POST" {

			// Look in the request body and headers for a spoofed method.
			// Prefer the value in the request body if they conflict.
			method := r.PostFormValue("_method")
			if method == "" {
				method = r.Header.Get("X-HTTP-Method-Override")
			}

			// Check that the spoofed method is a valid HTTP method and
			// update the request object accordingly.
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				r.Method = method
			}
		}

		c.Next()
	}
}
