package profiling

import (
	"os"
)

const (
	cpuProfTmpFilename   = "current_cpu_profile"
	traceProfTmpFilename = "current_trace_profile"
)

type dumper struct {
	c *Config
}

type heapDumper struct {
	dumper *dumper
}

type cpuDumper struct {
	c    *Config
	data *os.File
}

type traceDumper struct {
	c    *Config
	data *os.File
}

func (d *dumper) Prepare() error { _ = "STUB: not implemented"; return nil }

func (d *dumper) Dump(timestamp string, name string) error { _ = "STUB: not implemented"; return nil }

func (d *dumper) Release() error {
	_ = "STUB: not implemented"
	// Do nothing
	return nil
}

func (d *heapDumper) Prepare() error { _ = "STUB: not implemented"; return nil }

func (d *heapDumper) Dump(timestamp string, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *heapDumper) Release() error { _ = "STUB: not implemented"; return nil }

func (d *traceDumper) Prepare() error { _ = "STUB: not implemented"; return nil }

func (d *traceDumper) Dump(timestamp string, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *traceDumper) Release() error { _ = "STUB: not implemented"; return nil }

func (d *cpuDumper) Prepare() error { _ = "STUB: not implemented"; return nil }

func (d *cpuDumper) Dump(timestamp string, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *cpuDumper) Release() error { _ = "STUB: not implemented"; return nil }

func getTempFilename(tag, name string) string { _ = "STUB: not implemented"; return "" }

func getFilename(timestamp, tag, name string) string { _ = "STUB: not implemented"; return "" }

func createProfilesFolder() error { _ = "STUB: not implemented"; return nil }
