**GCI 2019**
*JBOSS Community*
*Rest API using GoLang*

**By MiskaKyto**

**How to use the REST API**

*Viewing all contents:*
```
GET http://localhost:10000/foods/

Response:
[
  {
    "Id": "1",
    "Title": "Banana",
    "desc": "🍌",
    "content": "Bananas are yellow"
  },
  {
    "Id": "2",
    "Title": "Apple",
    "desc": "🍏",
    "content": "Apples are green"
  },
  {
    "Id": "3",
    "Title": "Cherry",
    "desc": "🍒",
    "content": "Cherries are red"
  }
]
```

*Getting specific IDs:*
```
GET http://localhost/food/{id}
```

*Adding new objects:*
```
POST http://localhost/food

in request body include:
{
    "Id": "",
    "Title": "",
    "desc": "",
    "content": ""
}
```

*Deleting objects*
```
DELETE http://localhost/food/{id}
```