 TransformData Package

The `transformdata` package is designed to transform various JSON data files into a specific output format using predefined transformation functions. It reads data from multiple input JSON files, applies transformations, and writes the results to new output files.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Data Structure](#data-structure)
- [Transformation Functions](#transformation-functions)
- [Directory Structure](#directory-structure)
- [JSON Files](#json-files)


## Installation

To use this package, ensure you have Go installed on your machine. You can clone this repository and install the necessary dependencies by running:

```bash
go get github.com/victormazeli/sis_data_scrapper
````

## Usage
Place the required JSON files in the same directory as your Go code or specify the correct path in the InputFileName field of the Data structure.
Call the TransformJsonData() function to perform the transformations.

## Data Structure
The Data struct represents the input and output file names for the transformation process:

```
type Data struct {
InputFileName  string // The name of the input JSON file
OutputFileName string // The name of the output JSON file
bytesValue     []byte // Holds the byte data of the input file
}
```

## Directory Structure
/Users/tife/sis_data_scrapper/
├── transformed_files/                        
│   ├── new_bundle.json       
│   ├── new_courses.json   
│   ├── new_result_template.json     
│   ├── new_grade_scales_item.json   
│   ├── new_programs.json   
│   ├── new_course_enrollment.json  
│   ├── new_programme_intakes.json   
│   └── new_enrollment.json    
└── transformdata/   
├── data.go   
└── transformDataModel/   
├── bundle.go  
├── courseEnrollment.go   
├── courses.go   
├── enrollment.go   
├── generalData.go   
├── gradeScale.go   
├── programme.go   
├── programmeIntake.go    
└── resultTemplate.go


## JSON Files

The following JSON files are expected as input and will be transformed:

Input File Name Output File Name
1. **bundles.json**:    new_bundle.json
2. **courses.json**:    new_courses.json
3. **courseResultTemplates.json**:  new_result_template.json
4. **gradeScales.json**:    new_grade_scales_item.json 
5. **programs.json**:   new_programs.json
6. **courseEnrollments.json**:  new_course_enrollment.json 
7. **intakes.json**:    new_programme_intakes.json
8. **enrollmentBundle.json**:   new_enrollment.json

## Transformation Functions

The package includes various transformation functions that convert input data into the desired output format:
1. TransformBundle
2. TransformCourse
3. TransformResultTemplate
4. TransformGradeScaleItem
5. TransformProgrammeData
6. TransformCourseEnrollment
7. TransformProgrammeIntake
8. TransformEnrollment
Each transformation function takes a specific type of input and returns the transformed output type.