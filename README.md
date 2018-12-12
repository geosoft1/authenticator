![caution](https://user-images.githubusercontent.com/6298396/39394204-c8a60516-4ad7-11e8-8f07-5c28de190586.png)

Authenticator return some informations which can unique identify your machine. For now only [mac](https://gist.github.com/rucuriousyet/ab2ab3dc1a339de612e162512be39283) address and machine [id](https://github.com/denisbrodbeck/machineid) is used. **This project is an experiment and should not be used for production yet.**

Useful to link a web application to an specified machine (like a hardware token). Work on Linux, Mac and Windows.

Basically is a web service which return a [JSONP](https://en.wikipedia.org/wiki/JSONP) structure. This technique is used to avoid [CORS](https://en.wikipedia.org/wiki/Cross-origin_resource_sharing) restriction to accessing localhost from web application with other origins. This structure can be requested by others applications at login to identify the machine and 	further to decide if the host machine is authorised or not to use the application.

To run the program use:

	./authenticator [-port=value]

Use any available port you wish using optional `-port=value` flag. Default (if not specified) is `8080`. The authenticated application must use the same port.

To check your machine identifiers you can use:

	curl localhost:8080/authenticator

Short example of use in Angular 1.4.8:

	var url = "http://localhost:8080/authenticator?callback=JSON_CALLBACK";
	$http.jsonp(url).success(function(data){
        console.log(data.mac,data.id);
    }).error(function(data){
		console.log('forbidden');
	})


New versions are different, take a look [here](https://stackoverflow.com/questions/12066002/parsing-jsonp-http-jsonp-response-in-angular-js) for various Angular versions. However you must change the code for various frameworks and libraries. Feel free to do that.
