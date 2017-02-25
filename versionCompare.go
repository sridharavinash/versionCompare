package versionCompare

import (
	"regexp"
	"strconv"
	"errors"
)

type Version struct{
	major int
	minor int
	patch string
}


func (v1 *Version) compare(v2 *Version) int{
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
		if v1.minor == v2.minor {
			if v1.patch > v2.patch {
				return 1
			}
			if v1.patch < v2.patch {
				return -1
			}
			return 0
		}
	}

	return -2
}	

func versionCompare(v1 string, v2 string) (int, error){
	va, err := extract(v1)
	if err != nil {
		return -2, err
	}
	
	vb, err2 := extract(v2)

	if err2 != nil {
		return -2, err2
	}

	return va.compare(vb), nil
	
}

func extract(in string) (*Version, error){
	ver := new(Version)
	r := regexp.MustCompile(`(\d+).(\d+).?(.*)`)
	r2 := r.FindStringSubmatch(in)
	if r2 == nil {
		return ver, errors.New("Unable to parse input.")
	}
	ver.major,_ = strconv.Atoi(r2[1])
	ver.minor,_ = strconv.Atoi(r2[2])
	ver.patch = r2[3]
	return ver, nil
}
