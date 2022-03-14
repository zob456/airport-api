# airport-api
Go code sample for airport API
Uses gRPC server & client along with a REST layer

**NOTE** You must login & set the JWT as a `Bearer token` before any of the other endpoints can be hit.
- Insomnia used instead of Postman but `Insomnia_Aiport_calls.json` should also work as an imported file into Postman
- Platform runs on Kubernetes with Tilt to run locally so please install [Tilt](https://docs.tilt.dev/install.html)

### Steps to run platform
1. Install [Tilt](https://docs.tilt.dev/install.html)
2. Make sure that Kubernetes is enabled on your local machine's Docker Desktop
3. start platform with `tilt up`
4. make API calls
