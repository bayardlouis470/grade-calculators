# Welcome to Grade Calculators

[Grade Calculator](https://github.com/bayardlouis470/grade-calculators) is a GO package with a group
of calculators for grade caclculating. 

## Features

It supports the following features:

* grade calculator
* final grade calculator
* semester grade calculator
* high school GPA calculator
* college GPA calculator
* ...

The [10converters calculator](https://10converters.com/calculators/grade-calculator) is an implementation
based on this package.

## Input and Output

For grade calculators, a CSV file is used as input of the calculator.

```
Quiz, 90, 20
Homework, 85, 20
Midterm, 92, 30
Final, 90, 30
```

The grade could be in percentage, letters or points. The output is a weighted average grade.

For GPA calculators, a CSV file contains grades of all classes is used as input.

```
Math, 92, 3
Physics, 85, 4
```

There's course name, grade and credit of each course in the input file. The overall GPA will be the output.

## Usage

Use the calculator in your app:
```
import "github.com/bayardlouis470/grade-calculators"


let grade = gradecalculators.CalculateGrade("percentage", assessments)
```

Build the document:
```
cd readthedocs
make html
```
## Stories

Many colleges and high schools provide students with similar calculators so they can calculate GPA, grade point average, etc. A weighted grade is calculated based on the weight of the grade received and the final grade. Such a calculator usually works with a specified class or course as an average of performance, taking into account a certain period of time. 

Use the calculator to calculate the grade of the course based on the weighted average. Use the calculators to get the grade you want for an ongoing course. Calculate the grades you need for the remaining assignments and calculate the grade you need in the final exam to get the desired grade in this course, or use a calculator to find out what grade you need for your final exam in terms of final points and total scores in a course if you get the grades you want. 
