# Dictionary Line Bot

This is a golang application using [LINE Messaging API](https://developers.line.me/en/docs/messaging-api/overview/) and [Oxford dictionary API](https://developer.oxforddictionaries.com/documentation#/) to demonstrate simple dictionary BOT.

# Prerequisites
* A [channel on LINE Developers](https://developers.line.me/en/docs/liff/getting-started/) for your application.
* A [LINE channel access token](https://developers.line.me/en/docs/liff/getting-started/#preparing-channel-access-token)
* An [application on Oxford Developer](https://developer.oxforddictionaries.com/admin/applications)

# Running the application

1. export these environment variable
    - `OXFORD_APP_ID` from Oxford app_id
    - `OXFORD_APP_KEY` from Oxford app_key
    - `LINE_ACCESS_TOKEN` from LINE channel_access_token
    - `LINE_SECRET` from LINE channel_secret
2. Build docker from Dockerfile
    ```
    docker build -t dict-bot .
    ```
3. Run docker 
    ```
    docker run -p 8090:8090 dict-bot
    ```
4. Register webhook URL in LINE bot channel using `https://{yoursecureurl}/api/webhook` (if you running on local you can use [ngrok](https://ngrok.com/) to create a secure URL and forward all the request to golang application)
5. Finish!

# How to use dict bot

1. Add bot as a friend
2. Send message to bot. It will use the first word from your message to search for a definition and synonyms
3. Bot will reply you a definition and up to 5 synonyms from Oxford dictionary.
