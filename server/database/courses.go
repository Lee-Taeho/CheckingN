package database

import (
	"context"
	"log"
	"server/middleware"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) GetDepartments() []middleware.Department {
	ctx := context.TODO()
	collection := m.mongo.Database(COURSES_DATABASE).Collection(DEPARTMENTS_COLLECTION)
	result, err := collection.Find(ctx, bson.D{})
	if err != nil { log.Fatal(err)}
	var deps []middleware.Department
	if err = result.All(ctx, &deps); err != nil {
		log.Fatal(err)
	}
	return deps
}

func (m *MongoDB) GetCoursesByDepartment(department_name string) []middleware.Course {
	ctx := context.TODO()
	collection := m.mongo.Database(COURSES_DATABASE).Collection(COURSES_COLLECTION)
	result, err := collection.Find(ctx, bson.M{"department":department_name})
	if err != nil { log.Fatal(err)}
	var courses []middleware.Course
	if err = result.All(ctx, &courses); err != nil {
		log.Fatal(err)
	}
	return courses
}

func (m *MongoDB) GetCoursesGroupedByDepartments() map[string][]middleware.Course {
	deps := m.GetDepartments()
	coursesByDeps := make(map[string][]middleware.Course)
	for _, dep := range deps {
		if coursesInDep := m.GetCoursesByDepartment(dep.Name); coursesInDep != nil {
			coursesByDeps[dep.Name] = coursesInDep
		}
	}
	return coursesByDeps
}
