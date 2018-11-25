# v1.1.1
Increasing the ammount of good `/health` responses to avoid `CrashLoopBackOff`
statuses.

# v1.1.0
Add `/health` route to web app. This will return a 200 response up to 5 times.
After this, the route will start returning with an internal server error.

# v1.0.0
Basic web app seving requests on `/`. Response includes the client's address
and the hostname of the server running the app.
