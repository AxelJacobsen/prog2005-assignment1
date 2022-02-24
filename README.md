# Assignment-one

# Overview
In this assignment, you are going to develop a REST web application in Golang that provides the client to retrieve information about universities that may be candidates for application based on their name, alongside useful contextual information pertaining to the country it is situated in. For this purpose, you will interrogate existing web services and return the result in a given output format.
The REST web services you will be using for this purposes are:


http://universities.hipolabs.com/

Documentation/Source under: https://github.com/Hipo/university-domains-list/


https://restcountries.com/

Documentation/Source under: https://gitlab.com/amatos/rest-countries



The first API focuses on the provision of university information, whereas the second one provides country information, both of which you will need to use in order to complete the assignment.
The API documentation is provided under the corresponding links, and both services vary vastly with respect to feature set and quality of documentation. Use Postman to explore the APIs, but be mindful of rate-limiting.
A general note: when you develop your services that interrogate existing services, try to find the most efficient way of retrieving the necessary information. This generally means reducing the number of requests to these services to a minimum by using the most suitable endpoint that those APIs provide. Consider mocking those services based on exemplary outputs that you can use to develop your service against locally, before invoking the actual APIs.
The final web service should be deployed on Heroku. The initial development should occur on your local machine. For the submission, you will need to provide both a URL to the deployed Heroku service as well as your code repository.
In the following, you will find the specification for the REST API exposed to the user for interrogation/testing.

# Specification
Note: Please post an issue if the specification is unclear - so we can clarify and refine it if needed.

# Endpoints
Your web service will have three resource root paths:

/unisearcher/v1/uniinfo/
/unisearcher/v1/neighbourunis/
/unisearcher/v1/diag/
Assuming your web service should run on localhost, port 8080, your resource root paths would look something like this:

http://localhost:8080/unisearcher/v1/uniinfo/
http://localhost:8080/unisearcher/v1/neighbourunis/
http://localhost:8080/unisearcher/v1/diag/
The supported request/response pairs are specified in the following.
For the specifications, the following syntax applies:


{:value} indicates mandatory input parameters specified by the user (i.e., the one using your service).

{value} indicates optional input specified by the user (i.e., the one using your service), where `value' can itself contain further optional input. The same notation applies for HTTP parameter specifications (e.g., {?param}).


Retrieve information for a given university
The initial endpoint focuses on return information about a country a particular university/ies (or universities containing a particular string in their name) is/are situated in, such as the official name of the country, spoken languages, and the OpenStreetMap link to the map.

Request

Method: GET
Path: uniinfo/{:partial_or_complete_university_name}/
Note: The name of the university can be partial or complete, and may return a single ("Cambridge") or multiple universities (e.g., "Middle").
Example request: uniinfo/norwegian%20university%20of%20science%20and%20technology/

Response

Content type: application/json

Status code: 200 if everything is OK, appropriate error code otherwise. Ensure to deal with errors gracefully.

Body (Example):

[
  {
      "name": "Norwegian University of Science and Technology", 
      "country": "Norway",
      "isocode": "NO",
      "webpages": ["http://www.ntnu.no/"],
      "languages": {"nno": "Norwegian Nynorsk",
                    "nob": "Norwegian Bokmål",
                    "smi": "Sami"},
      "map": "https://www.openstreetmap.org/relation/2978650"
  },
  ...
]

Retrieve universities with same name components in neighbouring countries
The second endpoint provides an overview of universities in neighbouring countries to a given country that have the same name component (e.g., "Middle") in their institution name.

Request

Method: GET
Path: neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}
{:country_name} refers to the English name for the country that is the basis (basis country) of the search of unis with the same name in neighbouring countries.
{:partial_or_complete_university_name} is the partial or complete university name, for which universities with similar name are sought in neighbouring countries
{?limit={:number}} is an optional parameter that limits the number of universities in bordering countries (number) that are reported.
Example request: neighbourunis/norway/science?limit=5

Response

Content type: application/json

Status code: 200 if everything is OK, appropriate error code otherwise. Ensure to deal with errors gracefully.

Body (Example):

[
  {
      "name": "Norwegian University of Science and Technology", 
      "country": "Norway",
      "isocode": "NO",
      "webpages": ["http://www.ntnu.no/"],
      "languages": {"nno": "Norwegian Nynorsk",
                    "nob": "Norwegian Bokmål",
                    "smi": "Sami"},
      "map": "https://www.openstreetmap.org/relation/2978650"
  },
  {
      "name": "Swedish University of Agricultural Sciences", 
      "country": "Sweden",
      "isocode": "SE",
      "webpages": ["http://www.slu.se/"],
      "languages": {"swe":"Swedish"},
      "map": "https://www.openstreetmap.org/relation/52822"
  },
  ...
]

Diagnostics interface
The diagnostics interface indicates the availability of individual services this service depends on. The reporting occurs based on status codes returned by the dependent services, and it further provides information about the uptime of the service.

Request

Method: GET
Path: diag/

Response

Content type: application/json

Status code: 200 if everything is OK, appropriate error code otherwise.

Body:

{
   "universitiesapi": "<http status code for universities API>",
   "countriesapi": "<http status code for restcountries API>",
   "version": "v1",
   "uptime": <time in seconds from the last service restart>
}
Note: <some value> indicates placeholders for values to be populated by the service.

# Deployment
The service is to be deployed on Heroku. You will need to provide the URL to the deployed service as part of the submission.

# General Aspects
As indicated during the initial sessions, ensure you work with professionalism in mind (see Course Rules). In addition to professionalism, you are at liberty to introduce further features into your service, as long it does not break the specification given above.
Please work in the provided workspace environment (see here - lodge an issue if you have trouble accessing it) for your user and create a project assignment-1 in this workspace.
Consider to review the example projects provided as part of the lectures and coding tutorials in order to develop understanding of concepts, rather (or in addition) to online resources. Chances are that you will have a better basic understanding, before you consult resources like StackOverflow for more specialised questions.
As mentioned above, be sensitive to rate limits of external services. If needed, consider mocking the remote services during development.
Where possible, avoid the use of third-party libraries, unless you wish to, since they can make problems we may not be able to support. The functionality of this assignment can be developed using the Golang standard API. If you plan to use third-party libraries, test the deployment early to ensure they are supported by Heroku.

# Submission
The assignment is an individual assignment. The submission deadline is provided on the course main wiki page. No extensions will be given for late submissions (unless the deadline is collectively extended, i.e., if we agree in class).
As part of the submission you will need to provide:

a link to your code repository (ensure it is public at that stage)
a link to the deployed Heroku service

In addition, we will provide you with an option to clarify aspects of your submission (e.g., aspects that don't quite work, or additional features).
The submission occurs via our submission system that not only facilitates the submission, but also the peer review of the assignment. Instructions for the submission system (submission, review) will be introduced in class, and linked here.

# Peer Review
After the submission deadline, there will be a second deadline during which you will review other students' submissions. To do this the system provides you with a checklist of aspects to assess. You will need to review at least two submissions to meet the mandatory requirements of peer review, but you can review as many submissions as you like, which counts towards your participation mark for the course. The peer-review deadline will be indicated closer to submission time and then listed on the main course wiki page.