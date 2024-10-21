package transformdata

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	transformdata "github.com/victormazeli/sis_data_scrapper/transformdata/transformDataModel"
)

type Data struct {
	InputFileName  string
	OutputFileName string
	bytesValue     []byte
}

func getJsonData() []Data {
	return []Data{
		{
			InputFileName:  "bundles.json",
			OutputFileName: "new_bundle.json",
		},
		{
			InputFileName:  "courses.json",
			OutputFileName: "new_courses.json",
		},
		{
			InputFileName:  "courseResultTemplates.json",
			OutputFileName: "new_result_template.json",
		},
		{
			InputFileName:  "gradeScales.json",
			OutputFileName: "new_grade_scales_item.json",
		},
		{
			InputFileName:  "programs.json",
			OutputFileName: "new_programs.json",
		},
		{
			InputFileName:  "courseEnrollments.json",
			OutputFileName: "new_course_enrollment.json",
		},
		{
			InputFileName:  "intakes.json",
			OutputFileName: "new_programme_intakes.json",
		},
		{
			InputFileName:  "enrollmentBundle.json",
			OutputFileName: "new_enrollment.json",
		},
	}

}

func transformType(data Data) interface{} {
	switch data.InputFileName {
	case "bundles.json":
		var bundles []transformdata.Bundle
		err := json.Unmarshal(data.bytesValue, &bundles)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %s", err)
		}
		var outputBundles []transformdata.OutputBundle
		for _, bundle := range bundles {
			outputBundles = append(outputBundles, transformdata.TransformBundle(bundle))

		}
		return outputBundles
	case "courses.json":
		var courses []transformdata.InputCourse
		err := json.Unmarshal(data.bytesValue, &courses)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %s", err)
		}
		var outputCourse []transformdata.OutputCourse
		for _, course := range courses {
			outputCourse = append(outputCourse, transformdata.TransformCourse(course))

		}
		return outputCourse

	case "courseResultTemplates.json":
		var resultTemplate []transformdata.CourseResultTemplate
		err := json.Unmarshal(data.bytesValue, &resultTemplate)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %s", err)
		}
		var outputResultTemplate []transformdata.OuputResultTemplate
		for _, template := range resultTemplate {
			outputResultTemplate = append(outputResultTemplate, transformdata.TransformResultTemplate(template))

		}
		return outputResultTemplate
	case "gradeScales.json":
		var grades []transformdata.Grade
		err := json.Unmarshal(data.bytesValue, &grades)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %s", err)
		}
		var outputGradeScales []transformdata.GradeScaleItem
		for _, gradeScale := range grades {
			outputGradeScales = append(outputGradeScales, transformdata.TransformGradeScaleItem(gradeScale))

		}
		return outputGradeScales

	case "programs.json":
		var programmes []transformdata.Program
		err := json.Unmarshal(data.bytesValue, &programmes)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %s", err)
		}
		var outputProgramme []transformdata.OuputProgramme
		for _, programme := range programmes {
			outputProgramme = append(outputProgramme, transformdata.TransformProgrammeData(programme))

		}
		return outputProgramme

	case "courseEnrollments.json":
		log.Println("am herer course  enrollment")
		var courseEnrollments []transformdata.CourseEnrollment
		err := json.Unmarshal(data.bytesValue, &courseEnrollments)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %s", err)
		}
		var outputCourseEnrollment []transformdata.OuputCourseEnrollment
		for _, courseEnrollment := range courseEnrollments {
			outputCourseEnrollment = append(outputCourseEnrollment, transformdata.TransformCourseEnrollment(courseEnrollment))

		}
		return outputCourseEnrollment

	case "intakes.json":
		log.Println("am herer programme intake")
		var programmeIntakes []transformdata.ProgramIntake
		err := json.Unmarshal(data.bytesValue, &programmeIntakes)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %s", err)
		}
		var outputProgrammeIntakes []transformdata.OuputProgrammeIntake
		for _, intake := range programmeIntakes {
			outputProgrammeIntakes = append(outputProgrammeIntakes, transformdata.TransformProgrammeIntake(intake))

		}
		return outputProgrammeIntakes

	case "enrollmentBundle.json":
		log.Println("am herer enrollment bundle")
		var enrollments []transformdata.Enrollment
		err := json.Unmarshal(data.bytesValue, &enrollments)
		if err != nil {
			log.Fatalf("Error unmarshalling JSON: %s", err)
		}
		var outputEnrollments []transformdata.OutputEnrollment
		for _, enrollment := range enrollments {
			outputEnrollments = append(outputEnrollments, transformdata.TransformEnrollment(enrollment))

		}
		return outputEnrollments

	}
	return nil
}

func TransformJsonData() {
	outputDir := "transformed_files"

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating directory: %s", err)
	}
	for _, v := range getJsonData() {

		jsonFile, err := os.Open(v.InputFileName)
		if err != nil {
			log.Fatalf("Error opening JSON file: %s", err)
		}
		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)
		v.bytesValue = byteValue
		output := transformType(v)
		transformedData, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			log.Fatalf("Error marshalling transformed data: %s", err)
		}

		outputFilePath := fmt.Sprintf("%s/%s", outputDir, v.OutputFileName)
		if v.OutputFileName != "" {
			err := ioutil.WriteFile(outputFilePath, transformedData, 0644)
			if err != nil {
				log.Fatalf("Error writing to output file: %s", err)
			}
		}
	}
}
