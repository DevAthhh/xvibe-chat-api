xVibe-chat-api
___
This is a simple API for interactions with the database.

# Endpoints
___
1. POST `api/v1/u/sign-up` - return JSON-response { "user_id": id, "token": "token" }
2. GET `api/v1/u/:id` - return JSON-response {}
3. POST `api/v1/chat` - create chat
4. POST `api/v1/members` - return JSON-response { "members": [] }
