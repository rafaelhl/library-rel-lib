// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	books "github.com/rafaelhl/library-rel-lib/books"

	mock "github.com/stretchr/testify/mock"
)

// BookListFindEngine is an autogenerated mock type for the BookListFindEngine type
type BookListFindEngine struct {
	mock.Mock
}

// FindAllBooks provides a mock function with given fields: ctx
func (_m *BookListFindEngine) FindAllBooks(ctx context.Context) ([]books.Book, error) {
	ret := _m.Called(ctx)

	var r0 []books.Book
	if rf, ok := ret.Get(0).(func(context.Context) []books.Book); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]books.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
