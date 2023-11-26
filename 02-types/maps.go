package types

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

type Server struct {
	name   string
	status int
}

func statusToString(status int) string {
	switch status {
	case 0:
		return "Online"
	case 1:
		return "Offline"
	case 2:
		return "Maintenance"
	case 3:
		return "Retired"
	}

	return ""
}

func getServerStatus(servers []Server) map[string]int {
	serverStatus := map[string]int{
		"total":       len(servers),
		"Online":      0,
		"Offline":     0,
		"Maintenance": 0,
		"Retired":     0,
	}
	for _, server := range servers {
		serverStatus[statusToString(server.status)]++
	}

	return serverStatus

}

func changeServerStatus(servers []Server, serverName string, newStatus int) {
	for i := range servers {
		if servers[i].name == serverName {
			servers[i].status = newStatus
			break
		}
	}
}

func Maps() {
	servers := []Server{
		{name: "darkstar", status: Online},
		{name: "aiur", status: Online},
		{name: "omicron", status: Online},
		{name: "w359", status: Online},
		{name: "baseline", status: Online},
	}

	fmt.Println("Status of phase 1", getServerStatus(servers))

	changeServerStatus(servers, "darkstar", Retired)
	changeServerStatus(servers, "aiur", Offline)
	changeServerStatus(servers, "w359", Maintenance)

	fmt.Println("Status of phase 2", getServerStatus(servers))

}
