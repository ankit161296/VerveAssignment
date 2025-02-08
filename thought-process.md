Hello viewers,

This application has only one endpoint which is exposed to the client "/api/verve/accept". This accepts requests from the client and returns a response. The response is "pushed successfully" or "request already exists" on the basis of the key "id" in the request body. If the key "id" is already present in the redis cache then the response is "request already exists" otherwise the response is "pushed successfully".

This is basically a check similar to the idempotent property of the HTTP methods. If the request is already present in the cache then the response is "request already exists" otherwise the request is pushed to the cache and the response is "pushed successfully".

