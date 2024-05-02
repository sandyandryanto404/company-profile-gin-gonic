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
 
 func PortfolioRoutes() []RouteSource {
	 routes := []RouteSource{
		 {
			 "/portfolio/list",
			 "GET",
			 false,
			 controllers.PortfolioList,
		 },
		 {
			"/portfolio/detail/:id",
			"POST",
			false,
			controllers.PortfolioDetail,
		},
	 }
	 return routes
 }