package routes

import (
	"api/src/router/routes/controllers"
	"net/http"
)

var UserRoutes = []Route{

	//creating a use route
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},

	//getting a user by id route

	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.GetUse,
		AuthRequired: false,
	},

	//getting all users route

	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		AuthRequired: false,
	},

	//updating a user route

	{
		URI:          "/users/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: false,
	},

	//deleting a user route

	{
		URI:          "/users/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: false,
	}}
