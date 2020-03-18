package main

import (
    "encoding/json"
    "log"
    "os"
    //"fmt"
)

type output_file struct {

    F    *os.File

}


func ( f *output_file ) record( packet packet_metadata ) {

    out, _ := json.Marshal( packet )
    _,err = (f.F).WriteString( string(out) )
    if err != nil {
        f.F.Close()
        panic(err)
		log.Fatal(err)
	}
    _,err = (f.F).WriteString( "\n" )
    if err != nil {
        f.F.Close()
        panic(err)
		log.Fatal(err)
	}
    return
}


func initFile( fname string ) *output_file {
    f, err := os.OpenFile( fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644 )

    if err != nil {
        panic(err)
		log.Fatal(err)
    }

    o := &output_file{
        F: f,
    }

    return o
}

//TODO: need to figure out when to close
