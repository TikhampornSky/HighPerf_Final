# HighPerf_Final

# ===============Num 1==================
# ======================================
run `go build .` <br></br>
run `./go-1 -input input/rand-10M.txt -output output/rand-10M.txt -profile cpu -profile-path ./profile` <br></br>
run `go tool pprof -http=:8080 ./profile/cpu.pprof` <br></br>

# ======================================
Another option <br></br>
run `./go-1 -input input/rand-10.txt -output output/rand-10.txt -profile trace -profile-path ./profile` <br></br>
run `go tool trace ./profile/trace.out` <br></br>