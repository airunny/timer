package dao

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUser_InitAdminUser(t *testing.T) {
	pwd, err := bcrypt.GenerateFromPassword([]byte("timer"), bcrypt.DefaultCost)
	assert.Nil(t, err)
	fmt.Println(string(pwd))
}
