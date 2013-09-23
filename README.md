goproxy-rubygems-cache
======================

an example caching proxy server for rubygems using goproxy

Run the proxy.

```
./goproxy-rubygems-cache -root cache -v
```

Run gem install in another terminal.

```
gem install rails -p http://localhost:8080
```

The proxy log.

```
2013/09/23 21:55:02 [001] INFO: Got request HEAD http://rubygems.org/latest_specs.4.8.gz
2013/09/23 21:55:02 [001] INFO: Sending request HEAD http://rubygems.org/latest_specs.4.8.gz
2013/09/23 21:55:08 [001] INFO: Received response 302 Moved Temporarily
2013/09/23 21:55:08 [001] INFO: Copying response to client 302 Moved Temporarily [302]
2013/09/23 21:55:08 [001] INFO: Copied 0 bytes to client error=<nil>
2013/09/23 21:55:08 [002] INFO: Got request HEAD http://production.s3.rubygems.org/latest_specs.4.8.gz
2013/09/23 21:55:08 [002] INFO: Sending request HEAD http://production.s3.rubygems.org/latest_specs.4.8.gz
2013/09/23 21:55:10 [002] INFO: Received response 200 OK
2013/09/23 21:55:10 [002] INFO: Copying response to client 200 OK [200]
2013/09/23 21:55:10 [002] INFO: Copied 0 bytes to client error=<nil>
2013/09/23 21:55:10 [003] INFO: Got request GET http://rubygems.org/latest_specs.4.8.gz
2013/09/23 21:55:10 [003] INFO: Sending request GET http://rubygems.org/latest_specs.4.8.gz
2013/09/23 21:55:10 [003] INFO: Received response 302 Moved Temporarily
2013/09/23 21:55:10 [003] INFO: Copying response to client 302 Moved Temporarily [302]
2013/09/23 21:55:10 [003] INFO: Copied 154 bytes to client error=<nil>
2013/09/23 21:55:10 [004] INFO: Got request GET http://production.s3.rubygems.org/latest_specs.4.8.gz
2013/09/23 21:55:10 [004] INFO: Sending request GET http://production.s3.rubygems.org/latest_specs.4.8.gz
2013/09/23 21:55:10 [004] INFO: Received response 200 OK
2013/09/23 21:55:11 [004] INFO: save cache to cache/latest_specs.4.8.gz
2013/09/23 21:55:11 [004] INFO: Copying response to client 200 OK [200]
2013/09/23 21:55:11 [004] INFO: Copied 519594 bytes to client error=<nil>
2013/09/23 21:55:11 [005] INFO: Got request HEAD http://rubygems.org/specs.4.8.gz
2013/09/23 21:55:11 [005] INFO: Sending request HEAD http://rubygems.org/specs.4.8.gz
2013/09/23 21:55:12 [005] INFO: Received response 302 Moved Temporarily
2013/09/23 21:55:12 [005] INFO: Copying response to client 302 Moved Temporarily [302]
2013/09/23 21:55:12 [005] INFO: Copied 0 bytes to client error=<nil>
2013/09/23 21:55:12 [006] INFO: Got request HEAD http://production.s3.rubygems.org/specs.4.8.gz
2013/09/23 21:55:12 [006] INFO: Sending request HEAD http://production.s3.rubygems.org/specs.4.8.gz
2013/09/23 21:55:12 [006] INFO: Received response 200 OK
2013/09/23 21:55:12 [006] INFO: Copying response to client 200 OK [200]
2013/09/23 21:55:12 [006] INFO: Copied 0 bytes to client error=<nil>
2013/09/23 21:55:12 [007] INFO: Got request GET http://rubygems.org/specs.4.8.gz
2013/09/23 21:55:12 [007] INFO: Sending request GET http://rubygems.org/specs.4.8.gz
2013/09/23 21:55:12 [007] INFO: Received response 302 Moved Temporarily
2013/09/23 21:55:12 [007] INFO: Copying response to client 302 Moved Temporarily [302]
2013/09/23 21:55:12 [007] INFO: Copied 154 bytes to client error=<nil>
2013/09/23 21:55:12 [008] INFO: Got request GET http://production.s3.rubygems.org/specs.4.8.gz
2013/09/23 21:55:12 [008] INFO: Sending request GET http://production.s3.rubygems.org/specs.4.8.gz
2013/09/23 21:55:12 [008] INFO: Received response 200 OK
2013/09/23 21:55:16 [008] INFO: save cache to cache/specs.4.8.gz
2013/09/23 21:55:16 [008] INFO: Copying response to client 200 OK [200]
2013/09/23 21:55:16 [008] INFO: Copied 1457624 bytes to client error=<nil>
2013/09/23 21:55:20 [009] INFO: Got request GET http://rubygems.org/gems/rails-4.0.0.gem
2013/09/23 21:55:20 [009] INFO: Sending request GET http://rubygems.org/gems/rails-4.0.0.gem
2013/09/23 21:55:21 [009] INFO: Received response 302 Moved Temporarily
2013/09/23 21:55:21 [009] INFO: Copying response to client 302 Moved Temporarily [302]
2013/09/23 21:55:21 [009] INFO: Copied 154 bytes to client error=<nil>
2013/09/23 21:55:21 [010] INFO: Got request GET http://tokyo-m.rubygems.org/gems/rails-4.0.0.gem
2013/09/23 21:55:21 [010] INFO: Sending request GET http://tokyo-m.rubygems.org/gems/rails-4.0.0.gem
2013/09/23 21:55:26 [010] INFO: Received response 200 OK
2013/09/23 21:55:26 [010] INFO: save cache to cache/gems/rails-4.0.0.gem
2013/09/23 21:55:26 [010] INFO: Copying response to client 200 OK [200]
2013/09/23 21:55:26 [010] INFO: Copied 1542656 bytes to client error=<nil>
```

Uninstall and reinstall

```
gem uninstall rails 
gem install rails -p http://localhost:8080
```

This time all responses are served from local cache.

```
2013/09/23 21:57:38 [011] INFO: Got request HEAD http://rubygems.org/latest_specs.4.8.gz
2013/09/23 21:57:38 [011] INFO: return response from local cache/latest_specs.4.8.gz
2013/09/23 21:57:38 [011] INFO: Copying response to client  [200]
2013/09/23 21:57:38 [011] INFO: Copied 0 bytes to client error=http: request method or response status code does not allow body
2013/09/23 21:57:38 [012] INFO: Got request GET http://rubygems.org/latest_specs.4.8.gz
2013/09/23 21:57:38 [012] INFO: return response from local cache/latest_specs.4.8.gz
2013/09/23 21:57:38 [012] INFO: Copying response to client  [200]
2013/09/23 21:57:38 [012] INFO: Copied 519594 bytes to client error=<nil>
2013/09/23 21:57:38 [013] INFO: Got request HEAD http://rubygems.org/specs.4.8.gz
2013/09/23 21:57:38 [013] INFO: return response from local cache/specs.4.8.gz
2013/09/23 21:57:38 [013] INFO: Copying response to client  [200]
2013/09/23 21:57:38 [013] INFO: Copied 0 bytes to client error=http: request method or response status code does not allow body
2013/09/23 21:57:38 [014] INFO: Got request GET http://rubygems.org/specs.4.8.gz
2013/09/23 21:57:38 [014] INFO: return response from local cache/specs.4.8.gz
2013/09/23 21:57:38 [014] INFO: Copying response to client  [200]
2013/09/23 21:57:38 [014] INFO: Copied 1457624 bytes to client error=<nil>
2013/09/23 21:57:42 [015] INFO: Got request GET http://rubygems.org/gems/rails-4.0.0.gem
2013/09/23 21:57:42 [015] INFO: return response from local cache/gems/rails-4.0.0.gem
2013/09/23 21:57:42 [015] INFO: Copying response to client  [200]
2013/09/23 21:57:42 [015] INFO: Copied 1542656 bytes to client error=<nil>
```
