package versionCompare

import (
	"fmt"
	"regexp"
	"strconv"
)

type Version struct{
	major int
	minor int
	patch string
	
}

func versionCompare(v1 string, v2 string, op string) bool{
	v := extract(v1)
	fmt.Println(v)
	
	return false

}

func extract(in string) *Version{
	fmt.Println(in)
	ver := new(Version)
	r := regexp.MustCompile(`(\d+).(\d+).?(.*)`)
	r2 := r.FindStringSubmatch(in)
	ver.major,_ = strconv.Atoi(r2[1])
	ver.minor,_ = strconv.Atoi(r2[2])
	ver.patch = r2[3]
	return ver
}
