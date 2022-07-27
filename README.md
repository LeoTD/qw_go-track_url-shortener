# Welcome to My URL Shortener
***

## Task
Build a mock up of a URL shortening service using a SQL database to store URL pairings.

## Description
URL Shortening is a technique used by services like `bit.ly` to create smaller, more easily readable, and more memorable URLs to replace long, or very specific URLs.

## Installation
Requires a working installation of golang. See go.dev for details. Run `make` command in the root of the repo. Compiled binary will be created at `bin/my_url_shortener`.

## Usage
```
./my_url_shortener [-a | --add] [-t | --translate] <url>
    -a, --add <url>
        Adds <url> to the the database, then creates a paired short url. The new short url is printed to stdout.

    -t, --translate <url>
        Translates a shortened url into it's original, long counterpart. Prints original url to stdout.
```

### The Core Team
Leo Tanenbaum-Diaz


<span><i>Made at <a href='https://qwasar.io'>Qwasar Silicon Valley</a></i></span>
<span><img alt='Qwasar Silicon Valley Logo' src='https://storage.googleapis.com/qwasar-public/qwasar-logo_50x50.png' width='20px'></span>
