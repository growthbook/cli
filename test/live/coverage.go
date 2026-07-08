//go:build live

package live

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

// coverage tracks which commands from the generated command-surface golden the
// live suite actually exercised. The report it emits is the reviewer-facing
// confidence artifact: it makes the gaps (e.g. stateful revision workflows)
// explicit rather than letting "the tests pass" imply full coverage.
type coverage struct {
	mu      sync.Mutex
	hit     map[string]bool // command path -> exercised
	surface []string        // every command path from the golden, sorted
}

func newCoverage() *coverage {
	return &coverage{hit: map[string]bool{}}
}

// record marks the command path (the leading non-flag tokens) as exercised,
// e.g. runGB("projects","create","--name","x") -> "growthbook projects create".
func (c *coverage) record(args []string) {
	var path []string
	for _, a := range args {
		if strings.HasPrefix(a, "-") {
			break
		}
		path = append(path, a)
	}
	if len(path) == 0 {
		return
	}
	c.mu.Lock()
	c.hit["growthbook "+strings.Join(path, " ")] = true
	c.mu.Unlock()
}

// loadSurface reads the command-path column of the generated golden. Each line
// looks like "growthbook projects create  [--flag ...]"; we keep the path only.
func (c *coverage) loadSurface(goldenPath string) error {
	f, err := os.Open(goldenPath)
	if err != nil {
		return err
	}
	defer f.Close()

	seen := map[string]bool{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimRight(sc.Text(), "\n")
		if line == "" {
			continue
		}
		// Drop the flag list; keep the command path.
		if i := strings.Index(line, "  ["); i >= 0 {
			line = line[:i]
		}
		path := strings.TrimSpace(line)
		// Only count leaf commands (skip the bare "growthbook" root and pure
		// group parents that take no operation of their own).
		if path == "growthbook" {
			continue
		}
		if !seen[path] {
			seen[path] = true
			c.surface = append(c.surface, path)
		}
	}
	sort.Strings(c.surface)
	return sc.Err()
}

// report writes a human-readable coverage summary to w and returns the counts.
func (c *coverage) report(w *os.File) (covered, total int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var hitList, missList []string
	for _, path := range c.surface {
		if c.hit[path] {
			hitList = append(hitList, path)
		} else {
			missList = append(missList, path)
		}
	}
	// Anything exercised that isn't a known leaf (e.g. a group parent invoked
	// directly) is still worth surfacing so the report is honest.
	surfaceSet := map[string]bool{}
	for _, p := range c.surface {
		surfaceSet[p] = true
	}
	var extra []string
	for path := range c.hit {
		if !surfaceSet[path] {
			extra = append(extra, path)
		}
	}
	sort.Strings(extra)

	total = len(c.surface)
	covered = len(hitList)
	pct := 0.0
	if total > 0 {
		pct = 100 * float64(covered) / float64(total)
	}

	fmt.Fprintf(w, "\n===== Live CLI coverage =====\n")
	fmt.Fprintf(w, "Exercised %d / %d commands (%.1f%%)\n\n", covered, total, pct)
	fmt.Fprintf(w, "Uncovered (%d):\n", len(missList))
	for _, p := range missList {
		fmt.Fprintf(w, "  %s\n", p)
	}
	if len(extra) > 0 {
		fmt.Fprintf(w, "\nExercised but not in the surface golden (%d):\n", len(extra))
		for _, p := range extra {
			fmt.Fprintf(w, "  %s\n", p)
		}
	}
	fmt.Fprintf(w, "=============================\n")
	return covered, total
}
