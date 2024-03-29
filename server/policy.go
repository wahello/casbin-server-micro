package server

import (
	"bufio"
	"context"
	"os"
	"strings"

	casbinpb "github.com/paysuper/casbin-server/casbinpb"
)

func (s *Server) LoadPolicyFile(path string) error {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	req := new(casbinpb.PolicyRequest)
	sc := bufio.NewScanner(f)
	ctx := context.Background()
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		fields := strings.Split(line, ",")
		if len(fields) < 2 {
			continue
		}

		for i, f := range fields {
			fields[i] = strings.TrimSpace(f)
		}

		req.PType = fields[0]
		req.Params = fields[1:]
		switch fields[0] {
		case "p": // policy line
			err = s.AddNamedPolicy(ctx, req, nil)
		case "g": // group line
			err = s.AddNamedGroupingPolicy(ctx, req, nil)
		}
		if err != nil {
			return err
		}
	}
	if err := sc.Err(); err != nil {
		return err
	}
	return nil
}
