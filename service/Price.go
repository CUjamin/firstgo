package service

import (
	"firstgo/util"
)

func Price( priceA int,countA int, priceB int ,countB int)int{
	return util.SumNum(priceA*countA,priceB*countB)
}