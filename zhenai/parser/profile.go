package parser

import (
	"crawler/engine"
	"regexp"
	"strconv"
	"crawler/models"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var aarriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var weightdRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var heightdRe = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">(\d+)CM</span></td>`)
var hourseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinZuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contens []byte, name string) engine.ParseResult {
	profile := models.Profile{}
	age, err := strconv.Atoi(extractString(contens, ageRe))
	if err == nil {
		profile.Age = age
	}
	weight, err := strconv.Atoi(extractString(contens, weightdRe))
	if err == nil {
		profile.Weight = weight
	}
	height, err := strconv.Atoi(extractString(contens, heightdRe))
	if err == nil {
		profile.Height = height
	}
	profile.Marriage = extractString(contens, aarriageRe)
	profile.Gender = extractString(contens, genderRe)
	if profile.Gender != "女" { //筛选条件
		return engine.NilParse(contens)
	}
	profile.Occupation = extractString(contens, occupationRe)
	profile.Education = extractString(contens, educationRe)
	profile.Name = name
	profile.Income = extractString(contens, incomeRe)
	profile.House = extractString(contens, hourseRe)
	profile.Car = extractString(contens, carRe)
	profile.Hokou = extractString(contens, hokouRe)
	profile.XinZuo = extractString(contens, xinZuoRe)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
