package Control

import (
	"errors"
	. "BJS_ChainDemo/Log"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/crypto/attr"
)

// consts associated with TCert
const (
	role        = "role"
	contactInfo = "contactInfo"
)

//CertHandler provides APIs used to perform operations on incoming TCerts
type certHandler struct {
}

// NewCertHandler creates a new reference to CertHandler
func NewCertHandler() *certHandler {
	return &certHandler{}
}

// isAuthorized checks if the transaction invoker has the appropriate role
// stub: chaincodestub
// requiredRole: required role; this function will return true if invoker has this role
func (t *certHandler) IsAuthorized(stub shim.ChaincodeStubInterface, requiredRole string) (bool, error) {
	//read transaction invoker's role, and verify that is the same as the required role passed in
	return stub.VerifyAttribute(role, []byte(requiredRole))
}

// getContactInfo retrieves the contact info stored as an attribute in a Tcert
// cert: TCert
func (t *certHandler) GetContactInfo(cert []byte) (string, error) {
	if len(cert) == 0 {
		return "", errors.New("cert is empty")
	}

	contactInfo, err := attr.GetValueFrom(contactInfo, cert)
	if err != nil {
		Logger.Errorf("system error %v", err)
		return "", errors.New("unable to find user contact information")
	}

	return string(contactInfo), err
}

// getAccountIDsFromAttribute retrieves account IDs stored in  TCert attributes
// cert: TCert to read account IDs from
// attributeNames: attribute names inside TCert that stores the entity's account IDs
func (t *certHandler) GetAccountIDsFromAttribute(cert []byte, attributeNames []string) ([]string, error) {
	if cert == nil || attributeNames == nil {
		return nil, errors.New("cert or accountIDs list is empty")
	}

	//decleare return object (slice of account IDs)
	var acctIds []string

	// for each attribute name, look for that attribute name inside TCert,
	// the correspounding value of that attribute is the account ID
	for _, attributeName := range attributeNames {
		Logger.Debugf("get value from attribute = v%", attributeName)
		//get the attribute value from the corresbonding attribute name
		accountID, err := attr.GetValueFrom(attributeName, cert)
		if err != nil {
			Logger.Errorf("system error %v", err)
			return nil, errors.New("unable to find user contact information")
		}

		acctIds = append(acctIds, string(accountID))
	}

	Logger.Debugf("ids = %v", acctIds)
	return acctIds, nil
}
