# Assignment

## Some background knowledge:

We are developing a product where farmers can input their daily operations into the application, such as how much feed they have given the fish today, then we summarize and analyze those numbers for the farmers.

Here are some background information to help you better understand the product (you don't need to understand everything):

- [About fish farming](AboutFishFarming.md)
- [What our product does](WhatOurProductDoes.md)

## Tasks:

1. Design the database models for "Fish pen", "Juvenile" and "Fish lot" mentioned in "[About fish farming](AboutFishFarming.md)". Please implement the models in backend.
2. Create a simple page that shows all of a farms fish pens. (The UI design doesn't need to be pretty). You can implement the page under the route /farms/:farmId, or you can add new routes as you like.
3. Allow users to create, update and delete fish pens. Please implement a simple frontend page and appropriate APIs. You can implement the page under the route /farms/:farmId, or you can add new routes as you like. (Again the UI design doesn't need to be pretty, only need to be functional).


# Note
- Create a new branch, commit your work and create pull request.
- The work load is assumed to be about one day.
- We will evaluate your results assuming it's achieved within a days time.
- If there are any points that could not be worked on due to time constraints but should be fixed, please describe them in the description of the pull request.
- It doesn't need to be perfect.
- Submissions will not be used for any purpose other than recruitment.


# How to work the server
```
$ docker compose up -d --build
```

then access http://localhost:3002/

![](/images/screenshot.png)
