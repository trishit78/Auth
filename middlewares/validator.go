package middlewares

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"context"
	"fmt"
	"net/http"
)

func UserLoginRequestValidator(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){
		var payload dto.LoginUserRequestDTO
		if err:= utils.ReadJsonBody(r,&payload); err!=nil{
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Invalid request body",err)
			return
		}
		if err := utils.Validator.Struct(payload); err!= nil{
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Validation failed",err)
			return 
		}

		fmt.Println("Login user",payload)
		
		req_context := r.Context()

		ctx:=context.WithValue(req_context,"payload",payload)

		
		next.ServeHTTP(w,r.WithContext(ctx))
	})
}




func UserCreateRequestValidator(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){
		var payload dto.CreateUserDTO
		if err:= utils.ReadJsonBody(r,&payload); err!=nil{
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Invalid request body",err)
			return
		}
		if err := utils.Validator.Struct(payload); err!= nil{
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Validation failed",err)
			return 
		}

		fmt.Println("Create User",payload)

		req_context := r.Context()

		ctx:=context.WithValue(req_context,"payload",payload)

		
		next.ServeHTTP(w,r.WithContext(ctx))
		next.ServeHTTP(w,r)
	})
}
