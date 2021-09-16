# Goto

This is an application we will be developing as the final project for the course **Better and Faster - with Go Language**.

## Introducing project **goto**: the url shortener

In this project, we will develop a complete program: goto, the url shortener web application.

- [Version 0](#goto): this initial version.

- [Version 1](versions/1.md): a map and a struct are used, together with a Mutex from the sync package and a struct factory.

- [Version 2](versions/2.md): the data is made persistent because it is written to a file in gob-format.

- [Version 3](versions/3.md): the application is rewritten with goroutines and channels.

### What is a url shortener?

You know that some addresses in the browser (called URLs) are (very) long and/or complex and that there are services on the web which turn these into a nice short URL, to be used instead. Our project is like that. It is a web service with two functionalities:

- **Add**: given a long URL, it returns a short version, e.g., http://maps.google.com/maps?f=q&source=s_q&hl=en&geocode=&q=tokyo&sll=37.0625,-95.677068&sspn=68.684234,65.566406&ie=UTF8&hq=&hnear=Tokyo,+Japan&t=h&z=9  (link A) becomes http://goto/UrcGq (link B) and stores this pair of data (all our short URL’s start with http://goto/).

- **Redirect**: whenever a shortened URL is requested, it redirects the user to the original, long URL. So, if you type (B) in a browser, it redirects you to the page of (A). For example, http://goto/a redirects to http://google.com/ if it was shortened to http://goto/a.

### Running the app

In a terminal, run:
```bash
$ go run .
```

Then go to a browser and navigate to https://localhost:3000/ it will redirect to the add form...
