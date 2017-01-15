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

func (c *sqlContainer) Builds() []string {
	builds := []string{}
	return builds
}

func (c *sqlContainer) Build(ID string) (Build, error) {
	return nil, ErrNotFound
}

func (c *sqlContainer) New(b Buildable) (Build, error) {
	build, err := New(b)
	if err != nil {
		return nil, err
	}
	tx, err := c.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into builds(id, projectid, script, executortype, output) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(build.ID(), build.ProjectID(), build.Script(), build.ExecutorType(), "")
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()

	return build, nil
}

func (c *sqlContainer) AddStage(buildID string, stage Stage) error {
	return errors.New("moe")
}

func (c *sqlContainer) Output(buildID string, output []byte) error {
	return errors.New("moe")
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
		_id integer not null primary key autoincrement,
		id text not null,
		projectid text not null,
		script text not null,
		executortype text not null,
		output text not null);
	create table stages (
		id integer not null primary key,
		build integer not null,
		type text not null,
		timestamp text not null,
		name text not null,
		data text not null);
	`
	_, err = c.db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}

	/*

		rows, err := db.Query("select id, name from foo")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			err = rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, name)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		stmt, err = db.Prepare("select name from foo where id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		var name string
		err = stmt.QueryRow("3").Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)

		_, err = db.Exec("delete from foo")
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
		if err != nil {
			log.Fatal(err)
		}

		rows, err = db.Query("select id, name from foo")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			err = rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, name)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	*/
	return nil
}
