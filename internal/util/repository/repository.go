package repository

type OptionsFetch struct {
	Offset int `exhaustruct:"optional"`
	Limit  int `exhaustruct:"optional"`
}
