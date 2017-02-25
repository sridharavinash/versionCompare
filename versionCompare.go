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

func (v *Version) compare(v1 *Version, v2 *Version) int{
	if v1.major > v2.major {
		return 1
	}
	
	if v1.major < v2.major {
		return -1
	}

	if v1.major == v2.major {
		if v1.minor > v2.minor {
			return 1
		}
		if v1.minor < v2.minor {
			return -1
		}
		return 0
	}
}	

func versionCompare(v1 string, v2 string, op string) bool{
	va := extract(v1)
	vb := extract(v2)
	
	fmt.Println(va,vb)
	
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
