# HighPerf_Final

# ===============Num 1==================
# ======================================
run `go build -o myapp` <br></br>
run `./myapp -input input/rand-10M.txt -output output/rand-10M.txt -profile cpu -profile-path ./profile` <br></br>
run `go tool pprof -http=:8080 ./profile/cpu.pprof` <br></br>