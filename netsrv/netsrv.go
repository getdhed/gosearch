package netsrv

import (
	"1dz/GoSearch/pkg/crawler"
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func Handler(conn net.Conn, idx map[string][]int, docsByID map[int]crawler.Document) {
	defer conn.Close()

	r := bufio.NewReader(conn)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			// клиент ушёл — это нормально
			return
		}

		query := normalizeWord(line)
		if query == "" {
			fmt.Fprintln(conn, "empty!")
			fmt.Fprintln(conn, "END")
			continue
		}

		ids, ok := idx[query]
		if !ok || len(ids) == 0 {
			fmt.Fprintln(conn, "empty!")
			fmt.Fprintln(conn, "END")
			continue
		}

		for _, id := range ids {
			doc, ok := docsByID[id]
			if !ok {
				continue
			}
			// Отдаём и ID, и URL (можешь добавить Title при желании)
			fmt.Fprintf(conn, "%d %s\n", doc.ID, doc.URL)
		}
		fmt.Fprintln(conn, "END")
	}
}

func normalizeWord(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.Trim(s, " \t\r\n.,:;!?\"'()[]{}<>|\\/+-=*&#@%^`~")
	return s
}

// Serve запускает accept-loop и размазывает клиентов по горутинам.
func Serve(addr string, idx map[string][]int, docsByID map[int]crawler.Document) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Println("listening on", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept:", err)
			continue
		}
		go Handler(conn, idx, docsByID)
	}
}