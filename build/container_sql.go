package build

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // SQLite driver used in database/sql
)

type sqlContainer struct {
	db       *sql.DB
	filename string
}

func (c *sqlContainer) Purge() error {
	err := os.Remove(c.filename)
	if err != nil && os.IsNotExist(err) {
		return err
	}
	return nil
}

func (c *sqlContainer) Init() error {
	var err error
	var migrationLevel int

	if c.db, err = sql.Open("sqlite3", c.filename); err != nil {
		return err
	}

	_ = c.db.QueryRow("select level from migrationlevel").Scan(&migrationLevel)

	for i := migrationLevel + 1; i < len(migrations); i++ {
		if err := c.applyMigration(i, migrations[i]); err != nil {
			return err
		}
	}

	return nil
}

func (c sqlContainer) Close() error {
	return c.db.Close()
}

func (c *sqlContainer) Builds() []string {
	builds := []string{}

	rows, err := c.db.Query("select id from builds order by created desc")
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
	stmt, err := c.db.Prepare("select created, projectid, script, executortype from builds where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&build.BCreated, &build.BProjectID, &build.BScript, &build.BExecutorType)
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
		err = build.AddStage(stage)
		if err != nil {
			return nil, err
		}
	}

	stmt, err = c.db.Prepare("select data from output where build = ? order by id")
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		data := []byte{}
		err = rows.Scan(&data)
		if err != nil {
			return nil, err
		}
		build.output = append(build.output, data...)
	}

	return &build, nil
}

func (c *sqlContainer) New(b Buildable) (Build, error) {
	build, err := New(b)
	if err != nil {
		return nil, err
	}
	stmt, err := c.db.Prepare("insert into builds(id, projectid, created, script, executortype) values(?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(build.ID(), build.ProjectID(), build.Created(), build.Script(), build.ExecutorType())
	if err != nil {
		return nil, err
	}
	return build, nil
}

func (c *sqlContainer) AddStage(buildID string, stage Stage) error {
	buf := Stage{}
	previous := &buf
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

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
	stmt, err := c.db.Prepare("insert into output(build, data) values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(buildID, output)
	if err != nil {
		return err
	}
	return nil
}

func (c sqlContainer) applyMigration(level int, migration string) error {
	log.Printf("Applying migration %d: %s\n", level, migration)
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err = tx.Exec(migration); err != nil {
		log.Printf("%q: %s\n", err, migration)
		return err
	}
	update := `update migrationlevel set level = ?`
	stmt, err := tx.Prepare(update)
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(level); err != nil {
		log.Printf("%q: %s\n", err, update)
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

var migrations = []string{
	``, // initial migration level 0
	`create table if not exists migrationlevel (level int64);`,
	`insert into migrationlevel (level) values (1);`,
	`create table if not exists builds (
		id text not null primary key,
		projectid text not null,
		created int64 not null,
		script text not null,
		executortype text not null);`,
	`create table if not exists stages (
		id integer not null primary key autoincrement,
		build text not null,
		type text not null,
		timestamp int64 not null,
		name text not null,
		data text);`,
	`create table if not exists output (
		id integer not null primary key autoincrement,
		build text not null,
		data text);`,
}
