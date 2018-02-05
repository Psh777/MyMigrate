package postgres

import (
	"os"
	"fmt"
	"database/sql"
	"strconv"
	"../types"
)

func RunSql(sql []byte) error {

	rows, err := DB.Query(string(sql))
	if err != nil {
		fmt.Println("RunSql postgres query error", err)
		return err
	}
	defer rows.Close()

	return nil
}

func InitialTable() error {

	sql := "CREATE TABLE public.my_migrate_versions" +
		"(" +
		"id_version SERIAL PRIMARY KEY," +
		"current_version BOOLEAN DEFAULT true," +
		"subject VARCHAR(100)" +
		");"

	rows, err := DB.Query(sql)
	if err != nil {
		fmt.Println("postgres query error", err)
		return err
	}
	defer rows.Close()

	return nil
}


func Reset() error {

	sql := "DROP TABLE public.my_migrate_versions;"

	rows, err := DB.Query(sql)
	if err != nil {
		fmt.Println("Reset query error", err)
		return err
	}
	defer rows.Close()

	return nil
}

func Create(subject, path string) error {

	var version int
	err := DB.QueryRow("INSERT INTO public.my_migrate_versions (subject, current_version) VALUES ($1, $2) RETURNING id_version;", subject, false).Scan(&version)

	switch {

	case err == sql.ErrNoRows:

		return err

	case err != nil:

		return err

	}

	filenameup := path + "/" + strconv.Itoa(version) + "_" + subject + "_up.sql"
	filenamedown := path + "/" + strconv.Itoa(version) + "_" + subject + "_down.sql"

	f1, err := os.Create(filenameup)

	if err == nil {
		fmt.Println("File created:", filenameup, f1)
	}

	f2, err := os.Create(filenamedown)
	if err == nil {
		fmt.Println("File created:", filenamedown, f2)
	}
	return nil
}

func ClearCurrent() error {
	rows, err := DB.Query("UPDATE public.my_migrate_versions SET current_version = $1;", false)
	if err != nil {
		fmt.Println("ClearCurrent query error", err)
		return err
	}
	defer rows.Close()
	return nil
}

func GetVersions(current, version int) ([]*types.ListVersion, error) {

	items := make([]*types.ListVersion, 0)

	rows, err := DB.Query("SELECT id_version, current_version, subject FROM public.my_migrate_versions WHERE id_version <= $1 and id_version > $2 ORDER BY id_version ASC;", version, current)
	if err != nil {
		fmt.Println("GetVersions query error", err)
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		item := new(types.ListVersion)
		err := rows.Scan(&item.IdVersion, &item.CurrentVersion, &item.Subject)
		if err != nil {
			fmt.Println("problem scan table get version", err)
			return items, err
		}
		items = append(items, item)
	}

	return items, nil
}

func GetVersionsDown(current, version int) ([]*types.ListVersion, error) {

	items := make([]*types.ListVersion, 0)

	rows, err := DB.Query("SELECT id_version, current_version, subject FROM public.my_migrate_versions WHERE id_version > $1 and id_version < $2 ORDER BY id_version DESC;", version, current + 1)
	if err != nil {
		fmt.Println("GetVersions query error", err)
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		item := new(types.ListVersion)
		err := rows.Scan(&item.IdVersion, &item.CurrentVersion, &item.Subject)
		if err != nil {
			fmt.Println("problem scan table get version", err)
			return items, err
		}
		items = append(items, item)
	}

	return items, nil
}

func GetCurrentVersion() (int, error){
	row1 := DB.QueryRow("SELECT id_version FROM public.my_migrate_versions WHERE current_version = $1", true)

	var count int

	err := row1.Scan(&count)
	switch {

	case err == sql.ErrNoRows:
		return 0, nil
	case err != nil:
		fmt.Println("getcurrant", err)
		return 0, err
	}

	return count, nil
}

func SetCurrent(version int) {
	rows, err := DB.Query("UPDATE public.my_migrate_versions SET current_version = $1 WHERE id_version = $2;", true, version)
	if err != nil {
		println("SetCurrent query error", err)
		return
	}
	defer rows.Close()
	return
}

func Grab() ([]*types.ListVersion, error) {

	items := make([]*types.ListVersion, 0)

	rows, err := DB.Query("SELECT id_version, current_version, subject FROM public.my_migrate_versions;")
	if err != nil {
		fmt.Println("GetVersions query error", err)
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		item := new(types.ListVersion)
		err := rows.Scan(&item.IdVersion, &item.CurrentVersion, &item.Subject)
		if err != nil {
			fmt.Println("problem scan table get version", err)
			return items, err
		}
		items = append(items, item)
	}

	return items, nil
}

func DropMainTable() {
	rows, err := DB.Query("DROP TABLE my_migrate_versions;")
	if err != nil {
		println(err)
		return
	}
	defer rows.Close()
	return
}