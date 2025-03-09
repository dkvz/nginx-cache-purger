# Nginx Cache Purger
Simple API meant to make calls to Nginx to purge its cache.

I'm using a special location on the server for the purging instead of the PURGE HTTP verb that is often seen in the Nginx doc.

I should show an example of Nginx config someday.

**Goal**: queue the purge calls but only run one at a time with some pausing in between.

For now the app should die in case almost any error occurs so I can quickly be informed when it no longer works.

## TODO
- [] Add server port to config