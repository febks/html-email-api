## How to Use in Postman

To test the email-sending functionality using Postman, follow these steps:

1. Open Postman and create a new request.
2. Set the request method to `POST`.
3. Use the following JSON in the request body:

```json
{
  "to": "recipient@example.com",
  "subject": "Test Email",
  "body": "This is a test email sent using the Go backend."
}
```

4. Set the `Content-Type` header to `application/json`.
5. Enter the API endpoint URL (e.g., `http://localhost:8080/send-email`).
6. Click `Send` to test the API.

### Example Response

```json
{
  "message": "Email sent successfully."
}
```