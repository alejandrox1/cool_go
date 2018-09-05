# JWTs

[Authentication in Golang with JWTs](https://auth0.com/blog/authentication-in-golang/)


## Resources
* [JWT Security Cheat Sheet](https://assets.pentesterlab.com/jwt_security_cheatsheet/jwt_security_cheatsheet.pdf)


## Routes
* `http://localhost:3000/` will take you to a dummy sign in page.

* `http://localhost:3000/status` will return the message `API is up and running`
as long as the server is indeed running.

* `localhost:3000/products` will return a message indicating `Required
authorization token not found`.

* `http://localhost:3000/get-token` will send a response ith nothing but a JWT
in it.

Now that we have a jwt:
```
 $ curl -H "Authorization: Bearer $jwt" http://localhost:3000/products
[
	{
		"id": 1,
		"name": "Hover Shooters",
		"slug": "hover-shooters",
		"description": "Different hoverboards"
	},
	{
		"id": 2,
		"name": "Ocean Explorer",
		"slug": "ocean-explorer",
		"description": "Explore the depths of the sea"
	},
	{
		"id": 3,
		"name": "Dinosaur Park",
		"slug": "dinosaur-park",
		"description": "Ride a T-Rex"
	},
	{
		"id": 4,
		"name": "Cars VR",
		"slug": "cars-vr",
		"description": "Get behind the wheel of the fastest cars"
	},
	{
		"id": 5,
		"name": "Robin Hood",
		"slug": "robin-hood",
		"description": "Master the art of archery"
	},
	{
		"id": 6,
		"name": "Real World VR",
		"slug": "real-world-vr",
		"description": "Explore the world"
	}
]
```

```
 $ curl -H "Authorization: Bearer $jwt" -d "" http://localhost:3000/products/hover-shooters/feedback
{
	"id": 1,
	"name": "Hover Shooters",
	"slug": "hover-shooters",
	"description": "Different hoverboards"
}
```

If you want to change a field:
```
 $ curl -H "Authorization: Bearer $jwt" -d '{"description": "aaaaa"}' http://localhost:3000/products/hover-shooters/feedback
{
	"id": 1,
	"name": "Hover Shooters",
	"slug": "hover-shooters",
	"description": "aaaaa"
}
```
