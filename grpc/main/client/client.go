package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"time"

	"shared" //Path to the package contains shared struct
)

func main() {

	var reply int
	// Tries to connect to localhost:1234 (The port on which rpc server is listening)
	conn, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connectiong:", err)
	}
	defer conn.Close()

	// Create result.csv file
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatalf("%s: %s", "Cannot create file", err)
	}
	defer file.Close()

	// Create a struct, that mimics all methods provided by interface.
	// It is not compulsory, we are doing it here, just to simulate a traditional method call.
	start := time.Now()

	for i := 0; i < 10000; i++ {

		writer := csv.NewWriter(file)
		defer writer.Flush()

		t1 := time.Now()

		args := &shared.Args{i, i}
		err := conn.Call("Arithmetic.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}

		fmt.Println(reply)
		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		str := strconv.FormatFloat(x, 'f', 5, 64)
		xxx := []string{str}
		writer.Write(xxx)

	}
	elapsed := time.Since(start)
	fmt.Printf("Tempo: %s \n", elapsed)

}
