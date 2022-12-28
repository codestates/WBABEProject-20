package model

import (
	"strings"

	"github.com/google/uuid"
)

func CreateUUID() string {
	uuidValue := uuid.New()
	/*
	uuid 값에서 -를 제거하는 특별한 이유가 있을까요?
	없다면 특별히 제거할 필요는 없다고 생각합니다.
	 */
	uuid := strings.Replace(uuidValue.String(), "-", "", -1)
	return uuid
}
