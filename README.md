# Using HTTP proxy in Go

This repo contains example code for using HTTP proxy in Go.

Once you have the Proxy URL from Enzoguard's [Cloud console](https://cloud.enzoguard.com), set the following environment variables:

```bash
HTTP_PROXY="https://test-app:snipped@pe2xqel.enzoproxy.com"
HTTPS_PROXY="https://test-app:snipped@pe2xqel.enzoproxy.com"
```

Once set, run the Go program:

```bash
go run main.go
```

You will see IP address of the proxy instead of the IP address of your workstation.

To skip the proxy, you can unset the environment variable or launch a new shell and run the program again.

## Using in your program

To use Enzoguard in your code, you need to tweak your HTTP clients to use configuration:

```go
client := &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	},
}
```

And now, all requests through the client will be routed via Enzoguard as long as the two environment variables are set.

## Configure proxy without environment variables

If you would like to not rely on environment variables, you can use the following function to create a proxy aware HTTP client:

```go
func httpClientWithProxy(proxyURL string) {
    proxy, err := url.Parse(proxyURL)
    if err != nil {
    	fmt.Println("Error parsing proxy URL:", err)
    	return
    }

    tr := &http.Transport{Proxy: http.ProxyURL(proxy)}
    client := &http.Client{Transport: tr}
    return client
}
```

## Questions

If you have any questions, please open a GitHub issue or reach out to Enzoguard support.
