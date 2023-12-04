#Assessment Task GoLang

Project Folders:
Connection Folder: This directory organizes classes related to connections.
Question 1: This folder encompasses the solution to question 1 of the assignment.
Question 2: Within this folder, you'll find our response to question 2 of the assignment.
Question 3: The contents of this folder address question 3 of the assignment.
Question 4: Contained in this folder is our answer to question 4 of the assignment.
Routes: This section encompasses all the routes associated with question 1.
.env File: This file houses system variables crucial for the connection string.

Question No 1:
Objective Accomplished:
I have successfully developed a RESTful API using Golang, leveraging the Gin framework, PostgreSQL as the database, and pgx as the database driver. The API is equipped with features to create a user, generate OTP for the user, and verify the OTP.
Issue:
However, there is a challenge regarding the management of SQL queries, as sqlc is not currently utilized.
I previously reached out to you about encountering compatibility issues with sqlc and PostgreSQL 16 on Windows. Consequently, the implementation of sqlc has been deferred. Nevertheless, I am well-versed in how sqlc maps SQL queries into aliases, which can be utilized later in Go code. This approach enhances code robustness by effectively segregating SQL from code logic.
Running the API:
1.	To execute the first question, modify the .env file to align with your system and PostgreSQL environment.
2.	Execute the SQL command provided in the createtable.sql file at the root in your PostgreSQL environment.
3.	Run the following command in the terminal: 
go run questionNo1/question1.go
4.	Access and utilize the API.


Question No 2:

The provided string, denoted as’s’, has been skillfully reorganized to guarantee the absence of consecutive identical characters. The resulting rearrangement adheres to the stipulated criterion, augmenting the variety of adjacent characters.

To execute:

1.	Navigate to the questionNo2/question2 directory, where the main function resides. Inside, you'll find the variable’s,' representing the string, which is manipulated within the code. Feel free to modify its value according to your preference.

2.	Run the following command in the terminal:
go run questionNo2/question2.go


Question No 3:
The implementation of a solution to interchange the seat IDs of each pair of consecutive students has been executed successfully. In instances where the number of students is an odd count, the ID of the last student remains unchanged. The resulting table, sorted in ascending order based on student IDs, has been generated and is now available for utilization.
To execute:
1.	Modify the .env file to align with your system and PostgreSQL environment.
2.	Create a table named "seat" with columns "id" and "name."
3.	Run the following command in the terminal:
go run questionNo3/question3.go




Question No 4:
The potential deadlock and race conditions that could arise in a scenario where M goroutines read from a shared buffer and N goroutines write into it have been effectively addressed. Strategies have been implemented to mitigate the risk of deadlock and race conditions, all accomplished without the need for wait groups. The goroutines operate seamlessly, continuously running without encountering synchronization issues. This ensures a stable and efficient execution of the program.
To execute:
1.	Navigate to the questionNo4/question4.go file, which contains the main function. Inside, you will find the variables mreaders and nwriters. Set their values as per your preference.
2.	Run the following command in the terminal:
	go run questionNo4/question4.go
