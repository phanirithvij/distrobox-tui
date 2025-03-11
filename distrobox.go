package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type distroboxFinishedMsg struct{ err error }

type distroboxItem struct {
	id     string
	name   string
	status string
	image  string
}

type DockerCmdOutput struct {
	Id     string `json:"ID"`
	Image  string `json:"Image"`
	Labels string `json:"Labels"`
	Mounts string `json:"Mounts"`
	Names  string `json:"Names"`
	Status string `json:"Status"`
}

type PodmanCmdOutput struct {
	Id     string            `json:"Id"`
	Image  string            `json:"Image"`
	Labels map[string]string `json:"Labels"`
	Mounts []string          `json:"Mounts"`
	Names  []string          `json:"Names"`
	Status string            `json:"Status"`
}

func clearScreen() tea.Cmd {
	cmd := exec.Command("clear")
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return distroboxFinishedMsg{err}
	})
}

func enterDistroBox(name string) tea.Cmd {
	dbCmd := fmt.Sprintf("clear && distrobox enter %s", name)
	cmd := exec.Command("/bin/sh", "-c", dbCmd)
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return distroboxFinishedMsg{err}
	})
}

func removeDistroBox(name string) tea.Cmd {
	devnull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0o755)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := exec.Command("distrobox", "rm", name, "--force")
	cmd.Stdout = devnull
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return distroboxFinishedMsg{err}
	})
}

func stopDistroBox(name string) tea.Cmd {
	devnull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0o755)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := exec.Command("distrobox", "stop", name, "--yes")
	cmd.Stdout = devnull
	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return distroboxFinishedMsg{err}
	})
}

func getOCICmd() (string, string) {
	podmanExists := true
	dockerExists := true

	podmanCmd, err := exec.LookPath("podman")
	if err != nil {
		podmanExists = false
	}

	dockerCmd, err := exec.LookPath("docker")
	if err != nil {
		dockerExists = false
	}

	if podmanExists {
		return podmanCmd, "podman"
	} else if dockerExists {
		return dockerCmd, "docker"
	} else {
		return "", ""
	}
}

func getDistroboxItems() (items []distroboxItem) {
	ociCmd, ociKind := getOCICmd()
	if ociCmd == "" {
		log.Fatalln("Missing dependency: we need a container manager. Please install one of podman or docker.")
	}

	rawOutput, _ := exec.Command(ociCmd, "ps", "-a", "--format", "json", "--no-trunc").Output()
	if ociKind == "podman" {
		var outputs []PodmanCmdOutput
		if err := json.Unmarshal(rawOutput, &outputs); err != nil {
			log.Fatalln(err)
		}
		for _, item := range outputs {
			for _, mount := range item.Mounts {
				if strings.Contains(mount, "/distrobox-export") {
					box := distroboxItem{
						id:     item.Id[:12],
						name:   strings.Join(item.Names, ","),
						status: item.Status,
						image:  item.Image,
					}

					items = append(items, box)
					break
				}
			}
		}

		if len(items) == 0 {
			for _, jsonElem := range outputs {
				for _, label := range jsonElem.Labels {
					if label == "manager=distrobox" {
						box := distroboxItem{
							id:     jsonElem.Id[:12],
							name:   strings.Join(jsonElem.Names, ","),
							status: jsonElem.Status,
							image:  jsonElem.Image,
						}

						items = append(items, box)

						break
					}
				}
			}
		}
	} else {
		var outputs []DockerCmdOutput
		for _, line := range bytes.Split(rawOutput, []byte{'\n'}) {
			if string(line) == "" {
				continue
			}
			var ociCmdOutput DockerCmdOutput
			if err := json.Unmarshal(line, &ociCmdOutput); err != nil {
				log.Fatalln(err)
			}
			outputs = append(outputs, ociCmdOutput)

			for _, mount := range strings.Split(ociCmdOutput.Mounts, ",") {
				if strings.Contains(mount, "/distrobox-export") {
					box := distroboxItem{
						id:     ociCmdOutput.Id[:12],
						name:   ociCmdOutput.Names,
						status: ociCmdOutput.Status,
						image:  ociCmdOutput.Image,
					}

					items = append(items, box)
					break
				}
			}
		}

		if len(items) == 0 {
			for _, jsonElem := range outputs {
				labels := strings.Split(jsonElem.Labels, ",")
				for _, label := range labels {
					if label == "manager=distrobox" {
						box := distroboxItem{
							id:     jsonElem.Id[:12],
							name:   jsonElem.Names,
							status: jsonElem.Status,
							image:  jsonElem.Image,
						}

						items = append(items, box)

						break
					}
				}
			}
		}
	}

	return items
}
