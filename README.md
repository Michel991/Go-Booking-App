This is a simple console-based Go application that allows users to book tickets for a conference. 
The app interacts with users by taking inputs such as first name, last name, email address, and the number of tickets they wish to book. 
The app then validates the input and processes the booking, sending a confirmation message. 
It is designed to run concurrently, sending email confirmations asynchronously.

Features
Allows users to input their details (first name, last name, email, and number of tickets).
Validates the user input (name length, email format, and ticket availability).
Books tickets and provides a confirmation message.
Sends an asynchronous ticket confirmation email (simulated using a time.Sleep for 10 seconds).
Displays the first names of all users who have booked tickets.
Stops taking bookings when tickets are sold out.

Technologies Used
Go Programming Language: The entire application is written in Go.
Concurrency: Uses Go routines to send tickets asynchronously.
