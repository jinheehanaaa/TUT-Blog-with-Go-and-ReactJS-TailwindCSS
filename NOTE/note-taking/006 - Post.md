# Objective 
- Dealing with Post

# Learn
## Pagination (GET)
- http://localhost:3000/api/allpost?page=1

## Single Post (GET)
- http://localhost:3000/api/allpost/2

## Update Post (PUT)
``` JSON
{
  "title": "Updated Post",
  "desc": "Updated Description",
  "image": "updated.png",
  "UserId": "1"
}
```

## Unique Post for particular user (GET)
- http://localhost:3000/api/uniquepost
