# webcolor

Simple web server for colorful Kubernetes demos!

This is a web server that serves a single HTML page with
a background color that will basically be `${HOSTNAME%%-*}`
so if you create a deployment named `pink`, the pods
will be named `pink-xxxxyyyyzzzz-abcde` and therefore
the web server will serve a pink background.

If you create deployments named `blue` and `green`
they will serve web page with respectively
blue and green backgrounds, so you can do very
literal blue/green deployment demos.

There are two images on the Docker Hub using that code:
- `jpetazzo/webcolor` (listens on port 8000)
- `jpetazzo/color` (listens on port 80, shows extra info)
