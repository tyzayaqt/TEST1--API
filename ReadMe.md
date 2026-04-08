//CREATE a task <3
curl -X POST http://localhost:4000/tasks \
     -H "Content-Type: application/json" \
     -d '{"project_id": 1, "created_by": 1, "title": "Setup Routing", "description": "Configure all API endpoints on port 4000"}'

//READ
curl http://localhost:4000/tasks

//UPDATE
curl -X PUT "http://localhost:4000/tasks?id=1" \
     -H "Content-Type: application/json" \
     -d '{"status": "completed"}'

//DELETE
curl -X DELETE "http://localhost:4000/tasks?id=1"