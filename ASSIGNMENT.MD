In order to assess your programming skills we kindly request you to provide a sample implementation
considering the following:

● Build an application - REST service which is able to process at least 10K requests per second.

● The service has one GET endpoint - /api/verve/accept which is able to accept an integer id as a
mandatory query parameter and an optional string HTTP endpoint query parameter. It should return
String “ok” if there were no errors processing the request and “failed” in case of any errors.

● Every minute, the application should write the count of unique requests your application received in
that minute to a log file - please use a standard logger. Uniqueness of request is based on the id
parameter provided.

● When the endpoint is provided, the service should fire an HTTP POST request to the provided
endpoint with count of unique requests in the current minute as a query parameter. Also log the HTTP
status code of the response.

Please build also the following extensions:

Extension 1: Make sure the id deduplication works also when the service is behind a Load Balancer and 2
instances of your application get the same id simultaneously. Note: Please don’t create another repository with
separate deployments for a load balancer or similar. This task is about your application itself and the concept
as such.

Extension 2: Instead of writing the count of unique received ids to a log file, send the count of unique received
ids to a distributed streaming service of your choice.

Please formulate a high level overview describing your implementation approach and design considerations in
an additional file named thought-process.md.

Please provide us with a link to your solution - both an executable jar file or go binary or docker container and
the source code - hosted on the repository service of your choice (Github, Bitbucket etc.)