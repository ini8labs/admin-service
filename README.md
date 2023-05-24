# Admin-Service API's

- Admin Service provides various endpoints to access user information, event information and perform operations of adding and deleting events.

- Available API's are as follows:

    - Get information of all events.

    - Add a new event by providing event date, event name, event type, and winning numbers as input.

    - To get information by applying certain filters such as evnt type, date and events between a certain date range.

    - Delete an event by giving event ID as input.

    - To get information of all the users participated in a particular event by giving event id as input.

    - Get information of particular user by providing phone number, user id or government id as input.
    
    
![Alt text](images/Screenshot%202023-05-17%20130857.png)

# API Endpoints

## ```/events```

- Request Type: GET

- Description: Fetches information of all the events from database.

- Response: Information of all the events as response as shown below.

![Alt text](images/Screenshot%202023-05-17%20131942.png)

## ```/event```

1. ### Request Type: POST

    - Description: Generates a new event by providing event date, event type, event name and winnig numbers as input.

    - Response: If input is in correct format as shown in the image below we get status code 201 with a successful message.

![Alt text](images/Screenshot%202023-05-17%20142122.png)

2. ### Request Type: Get

    - Description: Fetches event information by query parameter event type, date or range of date(start date and end date).

    - Response:

        - If query is for event type we get all information of all event having that event type.

        - If query is for event date we get information of all events on that particular date.

        - If query is for range of dates we get information of all event between start date and end date.

        - If only start date is given, end date will be automaticcaly set to last day of the year i.e 31st Dec, and, if only end date is given start date will be automatially set to first day of year i.e 1st Jan.

![Alt text](images/Screenshot%202023-05-17%20143845.png)


## ```/event/{EventUID}```

- Request Type: DELETE

- Description: Deletes an event by givng event id as path parameter.

![Alt text](images/Screenshot%202023-05-17%20170027.png)

## ```/users```

- Request Type: GET

- Description: Fetches information of all the users who have participated in a particular event, accepts event id as input in form of query parameter.

![Alt text](images/Screenshot%202023-05-17%20170203.png)

## ```/user```

- Request Type: GET

- Description: Fetches information of a particular user, accepts user phone number or user id or user government id as query parameter.

![Alt text](images/Screenshot%202023-05-17%20171322.png)