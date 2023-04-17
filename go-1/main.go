package main

import (
	"flag"
	"fmt"
	"runtime"

	usemergeconcurrent "github.com/TikhampornSky/HighPerfFinal/go-1/sorting/useMergeConcurrent"
	"github.com/pkg/profile"
)

func main() {
	cpu := flag.Int("Your logical cores: ", runtime.NumCPU(), "logical core")

	inputFile := flag.String("input", "", "input file")
	outputFile := flag.String("output", "", "output file")
	profilePath := flag.String("profile-path", "", "Path of profile")
	profileType := flag.String("profile", "", "Type of profile (cpu, goroutine, trace, threadcreation, clock, mem, mem-heap, mem-allocs)")

	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		fmt.Println("Please include your Input file (-input) and Output file (-output)")
		return
	}
	fmt.Println("Your logical cores: ", *cpu)
	fmt.Println("Input file :", *inputFile)
	fmt.Println("Output file :", *outputFile)

	if *profileType != "" {

		if *profilePath == "" {
			fmt.Println("Profile path is required to track information (-profile-path)")
			return
		}

		fmt.Println("Your profile: ", *profileType)
		switch *profileType {
		case "cpu":
			defer profile.Start(
				profile.CPUProfile,
				profile.ProfilePath(*profilePath)).Stop()
		case "goroutine":
			defer profile.Start(
				profile.GoroutineProfile,
				profile.ProfilePath(*profilePath)).Stop()
		case "trace":
			defer profile.Start(
				profile.TraceProfile,
				profile.ProfilePath(*profilePath)).Stop()
		case "threadcreation":
			defer profile.Start(
				profile.ThreadcreationProfile,
				profile.ProfilePath(*profilePath)).Stop()
		case "clock":
			defer profile.Start(
				profile.ClockProfile,
				profile.ProfilePath(*profilePath)).Stop()
		case "mem":
			defer profile.Start(
				profile.MemProfile,
				profile.ProfilePath(*profilePath)).Stop()
		case "mem-heap":
			defer profile.Start(
				profile.MemProfileHeap,
				profile.ProfilePath(*profilePath)).Stop()
		case "mem-allocs":
			defer profile.Start(
				profile.MemProfileAllocs,
				profile.ProfilePath(*profilePath)).Stop()
		default:
			fmt.Println("Unknow Type: ", *profileType)
		}

	} else {
		fmt.Println("No profile to track")
	}

	fmt.Printf("Sorting...")
	usemergeconcurrent.Init(*inputFile, *outputFile, int64(*cpu))
	usemergeconcurrent.Run()
}
