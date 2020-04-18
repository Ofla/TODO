package memdatabase_test

import (
	td "github.com/Ofla/TODO/TestData"
	"github.com/Ofla/TODO/config"
	"github.com/Ofla/TODO/database"
	"github.com/Ofla/TODO/database/memdatabase"
	utl "github.com/Ofla/TODO/helpers"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

var (
	db database.DbHandler
	log *logrus.Logger
)


func TestMain(m *testing.M) {
	// insert categories and return their IDs for test
	log = utl.GetLogger()
	conf, err := config.LoadConfiguration()
	if err!=nil {
		log.Fatalf("Can not load configuration, error : %v", err)
	}
	db = memdatabase.NewTodoRepo(&conf.Database, log)
	code := m.Run()
	os.Exit(code)
}

func Test_AddTodo(t *testing.T) {
	testCases := td.CreateTTAddTodo()
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			err := db.AddTodo(&testCase.Todo)
			if err != nil && !testCase.HasError {
				t.Errorf("expected success , got error: %v", err)
			}
			if err == nil && testCase.HasError {
				t.Error("expected error, got nil")
			}
		})
	}
}
