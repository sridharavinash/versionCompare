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

// returns -1,0,1 for lesser,equal and greater
// returns -2 on error
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

	//should never get here
	return -2
}	

// returns -1,0,1 for lesser,equal and greater
// returns err if it is unable to parse the input.
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

// extracts a string in the form of 1.2.3-blah
// Returns a version major minor and patch if successful
// otherwise returns an error.
func extract(in string) (*Version, error){
	ver := new(Version)
	r := regexp.MustCompile(`^(\d+)?.?(\d+).?(.*)`)
	r2 := r.FindStringSubmatch(in)
	if r2 == nil {
		return ver, errors.New("Unable to parse input.")
	}
	ver.major,_ = strconv.Atoi(r2[1])
	ver.minor,_ = strconv.Atoi(r2[2])
	ver.patch = r2[3]
	return ver, nil
}
