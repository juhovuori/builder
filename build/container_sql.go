package build

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // SQLite driver used in database/sql
)

type sqlContainer struct {
	db *sql.DB // TODO: when to db.Close
}

const filename = "/tmp/builder.db"

func (c *sqlContainer) Init(purge bool) error {
	var err error
	os.Remove(filename)

	c.db, err = sql.Open("sqlite3", filename)
	if err != nil {
		return err
	}
	sqlStmt := `
	create table builds (
		id text not null primary key,
		projectid text not null,
		script text not null,
		executortype text not null,
		output text not null);
	create table stages (
		id integer not null primary key autoincrement,
		build integer not null,
		type text not null,
		timestamp int64 not null,
		name text not null,
		data text);
	`
	_, err = c.db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}

func (c *sqlContainer) Builds() []string {
	builds := []string{}

	rows, err := c.db.Query("select id from builds")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			log.Println(err)
			return nil
		}
		builds = append(builds, id)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil
	}
	return builds
}

func (c *sqlContainer) Build(ID string) (Build, error) {
	build := defaultBuild{BID: ID}
	stmt, err := c.db.Prepare("select projectid, script, executortype, output from builds where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&build.BProjectID, &build.BScript, &build.BExecutorType, &build.Boutput)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	stmt, err = c.db.Prepare("select type, timestamp, name, data from stages where build = ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		stage := Stage{}
		err = rows.Scan((*string)(&stage.Type), &stage.Timestamp, &stage.Name, &stage.Data)
		if err != nil {
			return nil, err
		}
		err := build.AddStage(stage)
		if err != nil {
			return nil, err
		}
	}

	return &build, nil
}

func (c *sqlContainer) New(b Buildable) (Build, error) {
	build, err := New(b)
	if err != nil {
		return nil, err
	}
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.Prepare("insert into builds(id, projectid, script, executortype, output) values(?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(build.ID(), build.ProjectID(), build.Script(), build.ExecutorType(), "")
	if err != nil {
		return nil, err
	}
	tx.Commit()

	return build, nil
}

func (c *sqlContainer) AddStage(buildID string, stage Stage) error {
	buf := Stage{}
	previous := &buf
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("select type, timestamp, name, data from stages where build = ? order by id desc limit 1")
	if err != nil {
		return err
	}
	t := ""
	err = stmt.QueryRow(buildID).Scan(&t, &buf.Timestamp, &buf.Name, &buf.Data)
	buf.Type = StageType(t)
	if err == sql.ErrNoRows {
		previous = nil
	} else if err != nil {
		return err
	}
	err = stage.ValidateWithPredecessor(previous)
	if err != nil {
		return err
	}
	stmt, err = tx.Prepare("insert into stages (build, type, timestamp, name, data) values (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(buildID, string(stage.Type), stage.Timestamp, stage.Name, stage.Data)
	if err != nil {
		return err
	}
	return tx.Commit()

}

func (c *sqlContainer) Output(buildID string, output []byte) error {
	return errors.New("moe")
}
