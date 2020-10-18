This is originally derived from the gzip handler by NYTimes & a code snipper from gist:1956518 by the42. Credits to them.

It has been cleaned & stripped down for performance.

* Removed Writer Pool. May consider adding a single writer object.
* We should use CDN for media, so Content-Type check is remvoed.
* No min-size check: if we make the effort to compress, just send it.