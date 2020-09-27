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
	db *sql.DB
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
	plugin.gravityDatabase = gravityDatabase
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
	if blockedNames == nil || pluginsState.sessionData["whitelisted"] != nil {
		return nil
	}
	sqlAdListSearch := `SELECT domain FROM gravity WHERE domain == ?`
	statement, err := plugin.db.Query(sqlAdListSearch)
	if err != nil {
		//TODO handle error
		
	}
	_, err := blockedNames.check(pluginsState, pluginsState.qName, nil)
	return err
}