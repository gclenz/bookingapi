# Booking API
This project is just a way to practice development with Go. The focus is on improving its quality over time.

## Use cases

### Create booking
  ```mermaid
    sequenceDiagram
        Customer->>+API: Creates a reservation
        API->>+Customer: Returns the payment link
        Customer->>+Payment_Gateway: Pays using the link
        Payment_Gateway->>+API: Sends payment webhook
        API->>+Customer: Notifies customer with payment success/failure
  ```