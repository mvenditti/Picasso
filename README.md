# Picasso

Basic demo involving Open AI Dall-E and Whisper services

## Getting Started

First retrieve the corresponding API Key from Open AI [website]("https://platform.openai.com/overview").
Then set up the `API_KEY` environment variable pointing to your Open AI key, application boot up will fail otherwise.

## Usage

In order to consume this basic example API we must execute a Get CURL as follows:

```
curl --location --request GET 'http://localhost:8080/image'
```

This will send our API a signal to read a predefined `audio.mp3` file, get it transcripted in `es` ISO language and finally generate an image with the resulting transcription.

Here is an example response containing the URL of the hosted result:

```
{
    "data": {
        "created": 1679419015,
        "data": [
            {
                "url": "some URL"
            }
        ]
    }
```
