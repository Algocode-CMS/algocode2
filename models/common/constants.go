package common

import "github.com/Algocode-CMS/algocode2/pkg/admin/models/choices"

const (
	EJUDGE          = "ej"
	EJUDGE_UNCACHED = "ej_nocache" // Do we need it?
	EJUDGE_EXTERNAL = "ej_ext"
	CODEFORCES      = "cf"
	PCMS            = "pc"
	NONE            = "no"
)

var kContestJudges = []choices.SelectChoice{
	{"Ejudge", EJUDGE},
	{"Ejudge uncached", EJUDGE_UNCACHED},
	{"Ejudge external", EJUDGE_EXTERNAL},
	{"Codeforces", CODEFORCES},
	{"PCMS", PCMS},
	{"None", NONE},
}

const (
	ACM = "acm"
	IOI = "ioi"
)

var kContestTypes = []choices.SelectChoice{
	{"ACM", ACM},
	{"Olympiad", IOI},
}

const (
	OPEN   = "open"
	ANON   = "anon"
	HIDDEN = "hidden"
)

var kAccessRights = []choices.SelectChoice{
	{"Open", OPEN},
	{"Hide names", ANON},
	{"Hidden", HIDDEN},
}

type ContestJudge struct{}

func (*ContestJudge) ListValues() []choices.SelectChoice {
	return kContestJudges
}

func (*ContestJudge) AllowNil() bool {
	return false
}

type ContestType struct{}

func (*ContestType) ListValues() []choices.SelectChoice {
	return kContestTypes
}

func (*ContestType) AllowNil() bool {
	return false
}

type AccessRight struct{}

func (*AccessRight) ListValues() []choices.SelectChoice {
	return kAccessRights
}

func (*AccessRight) AllowNil() bool {
	return false
}

type AccessRightNil struct{}

func (*AccessRightNil) ListValues() []choices.SelectChoice {
	return kAccessRights
}

func (*AccessRightNil) AllowNil() bool {
	return true
}
