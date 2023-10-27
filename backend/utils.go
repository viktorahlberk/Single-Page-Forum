package backend

import (
	"fmt"
	"sort"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//var usr = getAllUsersFromDb()

type sortUsers []User

func sortAlphabetically(usr []User) []User {
	sort.Sort(sortUsers(usr))
	return usr
}

func (u sortUsers) Len() int      { return len(u) }
func (u sortUsers) Swap(i, j int) { u[i], u[j] = u[j], u[i] }
func (u sortUsers) Less(i, j int) bool {
	return strings.ToLower(u[i].NickName) < strings.ToLower(u[j].NickName)
}
