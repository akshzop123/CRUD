package main

import (
	//"errors"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("Error while mocking")
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "role"}).AddRow(1, "Jane", "Jane@gmail.com", "Lead")
	TC := []struct {
		id          int
		user        *emp
		mockQuery   []interface{}
		expectError error
	}{
		{
			id:          1,
			user:        &emp{1, "jane", "Jane@gmail.com", "Lead"},
			mockQuery:   []interface{}{mock.ExpectQuery("SELECT * FROM employee WHERE id=?").WithArgs(1).WillReturnRows(rows)},
			expectError: nil,
		},
		{
			id:          3,
			user:        &emp{3, "jane", "Jane@gmail.com", "Lead"},
			mockQuery:   []interface{}{mock.ExpectQuery("SELECT * FROM employee WHERE id=?").WithArgs(3).WillReturnError(errors.New("err"))},
			expectError: errors.New("err"),
		},
	}

	for _, testCase := range TC {
		t.Run("", func(t *testing.T) {
			user, err := GetById(db, testCase.id)
			if err != nil && err.Error() != testCase.expectError.Error() {
				t.Errorf("expected error: %v, got: %v", testCase.expectError, err)
			}
			if !reflect.DeepEqual(user, testCase.user) {
				t.Errorf("expected user: %v, got: %v", testCase.user, user)
			}
		})
	}
}

func TestRemoveById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("Error while mocking")
	}
	defer db.Close()

	TC := []struct {
		id          int
		mockQuery   interface{}
		expectError error
	}{
		//succ
		{
			id: 3,
			mockQuery: []interface{}{
				mock.ExpectExec("DELETE FROM employee WHERE id=?").WithArgs(3).WillReturnResult(sqlmock.NewResult(3, 1)),
			},
			expectError: nil,
		},
		//fail
		{
			id: 1,
			mockQuery: []interface{}{
				mock.ExpectExec("DELETE FROM employee WHERE id=?").WithArgs(1).WillReturnError(errors.New("err")),
			},
				
			expectError: errors.New("err"),
		},
		
	}


	for _, tc := range TC {
		err := RemoveById(db,tc.id)
		fmt.Println(err)

		if err != nil && err.Error() != tc.expectError.Error() {
			t.Errorf("expected error: %v, got: %v", tc.expectError, err)
		}

	}
	
}

func TestInsert(t *testing.T){
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("Error while mocking")
	}
	defer db.Close()

	TC:=[]struct{
      id int
	  name string
	  email string
	  role string
	  mockQuery   []interface{}
	  expectError error

	}{
       //succ
		{
          id:1,
		  name:"akshata",
		  email:"ak@gmail.com",
		  role:"intern",
		  mockQuery: []interface{}{
			mock.ExpectExec("INSERT INTO employee VALUES(?,?,?,?)").WithArgs(1,"akshata","ak@gmail.com","intern").WillReturnResult(sqlmock.NewResult(1,1)),
		  },
		  expectError: nil,


		},
		//fail
		{
			id:3,
			name:"akshata",
			email:"ak@gmail.com",
			role:"intern",
			mockQuery: []interface{}{
			  mock.ExpectExec("INSERT INTO employee VALUES(?,?,?,?)").WithArgs(1,"akshata","ak@gmail.com","intern").WillReturnError(errors.New("err")),
			},
			expectError: errors.New("err"),
  
  
		  },

	}

	for _,tc:=range TC{
		err:=Insert(db,emp{tc.id,tc.name,tc.email,tc.role})
		if err != nil && err.Error() != tc.expectError.Error() {
			t.Errorf("expected error: %v, got: %v", tc.expectError, err)
		}
	}
}


func TestUpdateById(t *testing.T){
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("Error while mocking")
	}
	defer db.Close()

    TC:=[]struct{
		id int
		name string
		mockQuery   []interface{}
		expectError error
  
	}{
		//succ
		{
			id:1,
			name:"akshata",
			mockQuery: []interface{}{
				mock.ExpectExec("UPDATE employee SET name=? WHERE id=?").WithArgs("akshata",1).WillReturnResult(sqlmock.NewResult(1,1)),
			},
			expectError: nil,

		},

		//fail
		{
			id:5,
			name:"aks",
			mockQuery: []interface{}{
				mock.ExpectExec("UPDATE employee SET name=? WHERE id=?").WithArgs("aks",5).WillReturnError(errors.New("err")),
			},
			expectError: errors.New("err"),

		},

	}
	for _,tc:=range TC{
		err:=UpdateById(db,tc.name,tc.id)
		if err != nil && err.Error() != tc.expectError.Error() {
			t.Errorf("expected error: %v, got: %v", tc.expectError, err)
		}
	}

}