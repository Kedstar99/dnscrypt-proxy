package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
	"unicode"
	"database/sql"

	"github.com/jedisct1/dlog"
	"github.com/miekg/dns"
	"github.com/mattn/go-sqlite3"
)

type PluginBlockGravity struct {

}

func (plugin *PluginBlockGravity) Name() string {
	return "block_gravity_name"
}

func (plugin *PluginBlockGravity) Description() string {
	return "Block DNS queries found in pihole's gravity DB."
}

func (plugin *PluginBlockGravity) Init(proxy *Proxy) error {
	gravityDatabase, err := sql.Open("sqlite3", proxy.gravityDBLocation)
	if err != nil {
		return err
	}
	defer gravityDatabase.Close()
	return nil
}

func (plugin *PluginBlockGravity) Drop() error {
	return nil
}

func (plugin *PluginBlockGravity) Reload() error {
	return nil
}

func (plugin *PluginBlockName) Eval(pluginsState *PluginsState, msg *dns.Msg) error {
	return err
}