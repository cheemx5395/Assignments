1. GET request to a public API   

![get requst](./GET%20request.png)

2. POST request to a public API

![post request](./POST%20request.png)

3. API response analysis

![api response analysis](./api%20response%20analysis.png)

- URL :- [pokeapi](https://pokeapi.co/api/v2/pokemon/ditto)

- Method :- `GET`

- Request Headers :- Accept, Accept-Encoding, Cache-Control, Cookie etc.

- Response Headers :- Access-Control-Allowed-Origin, Age etc

- Body: None

So basically when a request is hit in the browser then browser checks if the domain name is already resoluted in its cache for the requested IP address.    
Now if its not cached already then the browser hits the router for IP ddress if its absent there too then to the ISP if absent there too then the DNS is a dictionary of all the domains resoluted according to their ip addresses that ip address is resolved and request is hit on that ip's stated port.   
Now the since the server is running on this port it responds to the request and the response is handled on the browser and we see the result.