package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	clusterFieldNames          = builderx.RawFieldNames(&Cluster{})
	clusterRows                = strings.Join(clusterFieldNames, ",")
	clusterRowsExpectAutoSet   = strings.Join(stringx.Remove(clusterFieldNames, "`create_time`", "`update_time`"), ",")
	clusterRowsWithPlaceHolder = strings.Join(stringx.Remove(clusterFieldNames, "`name`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	ClusterModel interface {
		List() ([]Cluster, error)
		Insert(data Cluster) (sql.Result, error)
		FindOne(name string) (*Cluster, error)
		Update(data Cluster) error
		Delete(name string) error
	}

	defaultClusterModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Cluster struct {
		Id    string `db:"id"`    // 集群ID
		Name  string `db:"name"`  // 集群名称
		Url   string `db:"url"`   // 集群地址
		Token string `db:"token"` // 集群token
	}
)

func NewClusterModel(conn sqlx.SqlConn) ClusterModel {
	return &defaultClusterModel{
		conn:  conn,
		table: "`cluster`",
	}
}

func (m *defaultClusterModel) List() ([]Cluster, error) {
	query := fmt.Sprintf("select * from %s", m.table)
	var resp []Cluster
	err := m.conn.QueryRows(&resp, query)
	return resp, err
}

func (m *defaultClusterModel) Insert(data Cluster) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, clusterRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Id, data.Name, data.Url, data.Token)
	return ret, err
}

func (m *defaultClusterModel) FindOne(name string) (*Cluster, error) {
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", clusterRows, m.table)
	var resp Cluster
	err := m.conn.QueryRow(&resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultClusterModel) Update(data Cluster) error {
	query := fmt.Sprintf("update %s set %s where `name` = ?", m.table, clusterRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Id, data.Url, data.Token, data.Name)
	return err
}

func (m *defaultClusterModel) Delete(name string) error {
	query := fmt.Sprintf("delete from %s where `name` = ?", m.table)
	_, err := m.conn.Exec(query, name)
	return err
}
