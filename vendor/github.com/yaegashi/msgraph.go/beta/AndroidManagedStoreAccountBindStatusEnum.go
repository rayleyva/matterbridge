// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidManagedStoreAccountBindStatus undocumented
type AndroidManagedStoreAccountBindStatus int

const (
	// AndroidManagedStoreAccountBindStatusVNotBound undocumented
	AndroidManagedStoreAccountBindStatusVNotBound AndroidManagedStoreAccountBindStatus = 0
	// AndroidManagedStoreAccountBindStatusVBound undocumented
	AndroidManagedStoreAccountBindStatusVBound AndroidManagedStoreAccountBindStatus = 1
	// AndroidManagedStoreAccountBindStatusVBoundAndValidated undocumented
	AndroidManagedStoreAccountBindStatusVBoundAndValidated AndroidManagedStoreAccountBindStatus = 2
	// AndroidManagedStoreAccountBindStatusVUnbinding undocumented
	AndroidManagedStoreAccountBindStatusVUnbinding AndroidManagedStoreAccountBindStatus = 3
)

// AndroidManagedStoreAccountBindStatusPNotBound returns a pointer to AndroidManagedStoreAccountBindStatusVNotBound
func AndroidManagedStoreAccountBindStatusPNotBound() *AndroidManagedStoreAccountBindStatus {
	v := AndroidManagedStoreAccountBindStatusVNotBound
	return &v
}

// AndroidManagedStoreAccountBindStatusPBound returns a pointer to AndroidManagedStoreAccountBindStatusVBound
func AndroidManagedStoreAccountBindStatusPBound() *AndroidManagedStoreAccountBindStatus {
	v := AndroidManagedStoreAccountBindStatusVBound
	return &v
}

// AndroidManagedStoreAccountBindStatusPBoundAndValidated returns a pointer to AndroidManagedStoreAccountBindStatusVBoundAndValidated
func AndroidManagedStoreAccountBindStatusPBoundAndValidated() *AndroidManagedStoreAccountBindStatus {
	v := AndroidManagedStoreAccountBindStatusVBoundAndValidated
	return &v
}

// AndroidManagedStoreAccountBindStatusPUnbinding returns a pointer to AndroidManagedStoreAccountBindStatusVUnbinding
func AndroidManagedStoreAccountBindStatusPUnbinding() *AndroidManagedStoreAccountBindStatus {
	v := AndroidManagedStoreAccountBindStatusVUnbinding
	return &v
}
