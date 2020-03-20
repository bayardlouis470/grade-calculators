# Overview

It's a common grade calculator API.


# Grade Calculator

There're 3 kinds of grading scales supportd:

* calculate grade in percentage
* calculate grade in letters
* calculate grade in points

All these are weighted grade calculators. That means there're 2 factors for a assessment, grade and weight(in some context
it could be credit). 

Whatever the grading scale used, the result is equal to sum of weight multiplies grades. For example:

***Weighted Grade*** = (g<sub>1</sub>+g<sub>2</sub>+g<sub>3</sub>+...+g<sub>n</sub>) &divide;
                    (m<sub>1</sub>+m<sub>2</sub>+m<sub>3</sub>+...+m<sub>n</sub>)

In case you have:

* Math: grade 80%, weight 30%
* History: grade 70 wieght 30%

Then your current weight is (80x30% + 70x30%)&divide;(30%+30%) =75

A reference could be found at: [Quick Grade Calculator](https://10converters.com/calculators/grade-calculator)
