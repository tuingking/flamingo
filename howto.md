# How To

load test
1. run server -> go run cmd/rest/main.go
2. load test using hey -> hey -m POST -n 100 -c 10 http://localhost:8080/api/products/seed/5000
3. profiling -> go tool pprof bin/flamingo-rest http://localhost:8080/debug/pprof/profile
3. profiling -> go tool pprof bin/flamingo-rest http://localhost:8080/debug/pprof/allocs
3. profiling -> go tool pprof bin/flamingo-rest http://localhost:8080/debug/pprof/heap