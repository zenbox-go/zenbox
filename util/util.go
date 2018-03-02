package util

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(ctx context.Context, prog string, args ...string) ([]byte, error) {

	cmd := exec.CommandContext(ctx, prog, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("")
	}
	if out != nil && err == nil && len(out) != 0 {

	}

	return out, nil
}

func Prompt(ctx context.Context, query, defaultAnswer string) (string, error) {
	fmt.Printf("%s [%s]: ", query, defaultAnswer)

	type result struct {
		answer string
		err    error
	}

	ch := make(chan result, 1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		if !s.Scan() {
			ch <- result{"", s.Err()}
			return
		}
		answer := s.Text()
		if answer == "" {
			answer = defaultAnswer
		}
		ch <- result{answer, nil}
	}()

	select {
	case r := <-ch:
		return r.answer, r.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
