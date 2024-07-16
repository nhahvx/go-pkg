package middleware

/*
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next() // execute all the handlers

		// at this point, all the handlers have been executed
		// check if there's an error
		if len(ctx.Errors) > 0 {
			// pick the first error
			err := ctx.Errors[0].Err

			// handle the error
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, router.ErrValidation(ctx, err))
		}
	}
}
*/
