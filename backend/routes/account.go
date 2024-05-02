/**
 * This file is part of the Sandy Andryanto Company Profile Website.
 *
 * @author     Sandy Andryanto <sandy.andryanto404@gmail.com>
 * @copyright  2024
 *
 * For the full copyright and license information,
 * please view the LICENSE.md file that was distributed
 * with this source code.
 */
 
 package routes

 import "backend/controllers"
 
 func AccountRoutes() []RouteSource {
	 routes := []RouteSource{
		 {
			 "/account/profile",
			 "GET",
			 true,
			 controllers.ProfileDetail,
		 },
		 {
			"/account/profile",
			"POST",
			true,
			controllers.ProfileUpdate,
		},
		{
			"/account/password",
			"POST",
			true,
			controllers.PasswordUpdate,
		},
		{
			"/account/upload",
			"POST",
			true,
			controllers.ProfileUpload,
		},
	 }
	 return routes
 }