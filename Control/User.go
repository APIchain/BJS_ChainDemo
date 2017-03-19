package Control

import (
	"BJS_ChainDemo/Log"
	. "BJS_ChainDemo/Module/Role"
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"sync"
)

var DefaultUserMemory UserMemory

type UserMemory struct {
	lock sync.RWMutex
	list map[string]*Account
}

//系统启动以后将用户信息读入内存
func InitUserMemory(stub shim.ChaincodeStubInterface) {
	userMemory := new(UserMemory)
	userMemory.list = make(map[string]*Account)

	startKey := DEFAULT_USER_LIST + "000000"
	endKey := DEFAULT_USER_LIST + "999999"
	iter, err := stub.RangeQueryState(startKey, endKey)
	if err != nil {
		Log.Logger.Error("ReadUser From DB Failed.", err)
	}

	for iter.HasNext() {
		_, bytes, err := iter.Next()
		if err != nil {
			Log.Logger.Error("ReadUser From DB Failed.", err)
		}
		account := new(Account)
		err = json.Unmarshal(bytes, account)
		if err != nil {
			Log.Logger.Error("Unmarshal User Failed.", err)
		}
		userMemory.list[account.UserName] = account
	}
	DefaultUserMemory = *userMemory
}

//加一个用户到List，并存储到db
func (u *UserMemory) AddToUserMemory(stub shim.ChaincodeStubInterface, username string, postDataServer string, returnDataServer string) error {
	u.lock.Lock()
	defer u.lock.Unlock()
	if exist := u.CheckExist(username); exist {
		return errors.New("This user has already exist.")
	}
	account, err := NewAccount(stub, username, postDataServer, returnDataServer)
	if err != nil {
		return errors.New("Create Account failed..")
	}
	err = account.Put(stub)
	if err != nil {
		return err
	}
	u.list[account.UserName] = account
	return nil
}

//按照用户名取得User
func (u *UserMemory) GetByUserName(username string) (*Account, error) {
	if exist := u.CheckExist(username); !exist {
		return nil, errors.New("This user is not exist.")
	}
	return u.list[username], nil
}

//修改一个用户信息，并存储到DB
func (u *UserMemory) UpdateToUserMemory(stub shim.ChaincodeStubInterface, username string, postDataServer string, returnDataServer string) error {
	u.lock.Lock()
	defer u.lock.Unlock()
	if exist := u.CheckExist(username); !exist {
		return errors.New("This user does not Existed. Please Check.")
	}
	account, err := NewAccount(stub, username, postDataServer, returnDataServer)
	if err != nil {
		return errors.New("Create Account failed..")
	}
	err = account.Put(stub)
	if err != nil {
		return err
	}
	u.list[account.UserName] = account
	return nil
}

//从DB以及内存中删除指定用户
func (u *UserMemory) DeleteFromUserMemory(stub shim.ChaincodeStubInterface, username string) error {
	u.lock.Lock()
	defer u.lock.Unlock()
	if exist := u.CheckExist(username); !exist {
		return errors.New("This user does not Existed. Please Check.")
	}

	err := u.list[username].Del(stub)
	if err != nil {
		return err
	}
	delete(u.list, username)
	return nil
}

//检查用户是否存在
func (u *UserMemory) CheckExist(username string) bool {
	if _, exist := u.list[username]; exist {
		return true
	}
	return false
}

//增加请求次数
func (u *UserMemory) AddRequest(stub shim.ChaincodeStubInterface, username string) error {
	u.lock.Lock()
	defer u.lock.Unlock()
	if exist := u.CheckExist(username); !exist {
		return errors.New("This user does not Existed. Please Check.")
	}
	u.list[username].RequestTime++
	u.list[username].Put(stub)
	return nil
}

//增加返答次数
func (u *UserMemory) AddResponse(stub shim.ChaincodeStubInterface, username string) error {
	u.lock.Lock()
	defer u.lock.Unlock()
	if exist := u.CheckExist(username); !exist {
		return errors.New("This user does not Existed. Please Check.")
	}
	u.list[username].ResponseTime++
	u.list[username].Put(stub)
	return nil
}

//增加超时次数
func (u *UserMemory) AddTimeOut(stub shim.ChaincodeStubInterface, username string) error {
	u.lock.Lock()
	defer u.lock.Unlock()
	if exist := u.CheckExist(username); !exist {
		return errors.New("This user does not Existed. Please Check.")
	}
	u.list[username].TimeoutTime++
	u.list[username].Put(stub)
	return nil
}
