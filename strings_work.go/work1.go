package main

import (
	"fmt"
	"strings"
)

func main() {
	//check whether a substring present in main string it returns true or false 
	//we are using print function a lot a times to reduce the code we do like this
	var p = fmt.Println
	result1 := strings.Contains("i am a learner", "i")
	p(result1)

	result2 := strings.Contains("i am a learner","you")
	p(result2)

	//check either character present

	result3 := strings.ContainsAny("i am a learner","xy")
	p(result3)

	result4 := strings.ContainsAny("i am a learner","xya")
	p(result4)

	result5 := strings.ContainsAny("i am a learner","xyam")
	p(result5)

	//check count of the character in string
	result6 := strings.Count("success","c")
	p(result6)
	
	//if the substring is empty in string the count is no of the characters+1
	result7 := strings.Count("success","")
	p(result7)

	//convert the string letters into lowercase letters
	p(strings.ToLower("SUDDDEn"))

	//convert the string letters into uppercase letters
	p(strings.ToUpper("ssdeSFD"))

	//comparing the two strings but GO is case sensitive
	p("Go"=="go")
	p("GO"=="GO")

	//if you want to avoid the case sensitivity
	p(strings.EqualFold("GO","go"))

	//to repeat the strings into to N no of times
	re := strings.Repeat("ok",5)
	p(re)

	//to replace a substring in string with how many places
	df := strings.Replace("12:23:45:55",":",",",2)
	p(df)

	//-1 gives to replace all things
	dfg := strings.Replace("12:23:45:55",":",",",-1)
	p(dfg)

	//direct method to replace all
	as := strings.ReplaceAll("12:23:45:55",":",",")
	p(as)

	//split the string based on given value and it gives us slice
	www := strings.Split("us,fdf,dfds-dfdf",",")
	p(www)

	//if the value is empty it split the string each character

	oij := strings.Split("US,SDS-SDS,SSS","")
	p(oij)

	//how to join the substrings 
	dfdd := strings.Join(www,",")
	p(dfdd)

	//if you have newline escape characters and Space  in the string to split 
	sdsd := strings.Fields("SD SD\nsdsd sd\nsdsd")
	p(sdsd)	

	//to remove leading spaces and last in string
	shhh := strings.TrimSpace("   RR-DFDFDF    ")
	p(shhh)

	//to remove unwanted characters in leading and ending 
	hfg := strings.Trim("@#$lsldnfsjf^%^","@$#^%")
	p(hfg)


	
}
