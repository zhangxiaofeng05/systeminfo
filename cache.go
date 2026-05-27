package main

import (
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

// CachedSystemInfo holds cached system information
type CachedSystemInfo struct {
	mu       sync.RWMutex
	hostInfo *host.InfoStat
	cpuInfo  []cpu.InfoStat
	memInfo  *mem.VirtualMemoryStat

	lastUpdate time.Time
	ttl        time.Duration
}

// NewCachedSystemInfo creates a new cached system info with specified TTL
func NewCachedSystemInfo(ttl time.Duration) *CachedSystemInfo {
	return &CachedSystemInfo{
		ttl: ttl,
	}
}

// RefreshAll updates all cached system information
func (c *CachedSystemInfo) RefreshAll() error {
	hostInfo, err := host.Info()
	if err != nil {
		return err
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		return err
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.hostInfo = hostInfo
	c.cpuInfo = cpuInfo
	c.memInfo = memInfo
	c.lastUpdate = time.Now()

	return nil
}

// GetHostInfo returns cached host info, refreshing if expired
func (c *CachedSystemInfo) GetHostInfo() (*host.InfoStat, error) {
	if c.needsRefresh() {
		if err := c.RefreshAll(); err != nil {
			return nil, err
		}
	}
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.hostInfo, nil
}

// GetCpuInfo returns cached CPU info, refreshing if expired
func (c *CachedSystemInfo) GetCpuInfo() ([]cpu.InfoStat, error) {
	if c.needsRefresh() {
		if err := c.RefreshAll(); err != nil {
			return nil, err
		}
	}
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cpuInfo, nil
}

// GetMemInfo returns cached memory info, refreshing if expired
func (c *CachedSystemInfo) GetMemInfo() (*mem.VirtualMemoryStat, error) {
	if c.needsRefresh() {
		if err := c.RefreshAll(); err != nil {
			return nil, err
		}
	}
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.memInfo, nil
}

// needsRefresh checks if cache is expired
func (c *CachedSystemInfo) needsRefresh() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lastUpdate.IsZero() || time.Since(c.lastUpdate) > c.ttl
}

// StartBackgroundRefresh starts a background goroutine to refresh cache periodically
func (c *CachedSystemInfo) StartBackgroundRefresh() {
	go func() {
		ticker := time.NewTicker(c.ttl / 2) // Refresh at half TTL to avoid expiration
		defer ticker.Stop()
		for range ticker.C {
			_ = c.RefreshAll() // Silently ignore errors in background refresh
		}
	}()
}
