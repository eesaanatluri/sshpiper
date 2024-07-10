package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"net"
	"strconv"
)

// readGroupUpstreamFile reads the file and returns a map with the group as the key and the whole line as the value
func readGroupUpstreamFile(fileName string) (map[string]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	groups := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && line[0] != '#' {
			parts := strings.Fields(line)
			if len(parts) > 0 {
				groups[parts[0]] = line
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

// getUserGroups retrieves the Unix groups the user is part of
func getUserGroups(user string) ([]string, error) {
	cmd := exec.Command("groups", user)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	groupString := strings.TrimSpace(string(output))
	groupParts := strings.Split(groupString, ":")
	if len(groupParts) < 2 {
		return nil, fmt.Errorf("unexpected output from groups command: %s", groupString)
	}

	userGroups := strings.Fields(groupParts[1])
	return userGroups, nil
}

// findIntersection returns the lines from groupMap where the keys intersect with userGroups
func findIntersection(groupMap map[string]string, userGroups []string) []string {
	var intersection []string
	for _, group := range userGroups {
		if line, exists := groupMap[group]; exists {
			intersection = append(intersection, line)
		}
	}
	return intersection
}

func SplitHostPortForSSH(addr string) (host string, port int, err error) {
	host = addr
	h, p, err := net.SplitHostPort(host)
	if err == nil {
		host = h
		port, err = strconv.Atoi(p)

		if err != nil {
			return
		}
	} else if host != "" {
		// test valid after concat :22
		if _, _, err = net.SplitHostPort(host + ":22"); err == nil {
			port = 22
		}
	}

	if host == "" {
		err = fmt.Errorf("empty addr")
	}

	return
}


func parseUpstreamData(data string) (host string, port int, group string, err error) {
	r := bufio.NewReader(strings.NewReader(data))
	for {
		host, err = r.ReadString('\n')
		if err != nil {
			break
		}

		host = strings.TrimSpace(host)

		if host != "" && host[0] != '#' {
			break
		}
	}

	t := strings.SplitN(host, " ", 2)

	if len(t) > 1 {
		group = t[0]
		host = t[1]
	}

	host, port, err = libplugin.SplitHostPortForSSH(host)
	return
}


func main() {
	groupUpstreamFile := "/var/sshpiper_config/groupUpstreamMapFile" // Replace with the actual file path

	groupMap, err := readGroupUpstreamFile(groupUpstreamFile)
	if err != nil {
		fmt.Printf("Error reading group upstream file: %v\n", err)
		return
	}

	user := "alice" // Replace with the actual username
	userGroupList, err := getUserGroups(user)
	if err != nil {
		fmt.Printf("Error getting user groups: %v\n", err)
		return
	}

	intersection := findIntersection(groupMap, userGroupList)
	for _, line := range intersection {
			fmt.Println(line)
			parseUpstreamData(line)

}

