package service

import (
	"my-gin-app/internal/model"
	"my-gin-app/internal/repository"
)

// ServiceContainer holds all service instances
type ServiceContainer struct {
	User        *UserService
	Department  *DepartmentService
	Position    *PositionService
	Stack       *StackService
	Role        *RoleService
	Permission  *PermissionService
	Resource    *ResourceService
	FeatureFlag *FeatureFlagService
}

// NewServiceContainer creates a new service container
func NewServiceContainer() *ServiceContainer {
	return &ServiceContainer{
		User:        &UserService{},
		Department:  &DepartmentService{},
		Position:    &PositionService{},
		Stack:       &StackService{},
		Role:        &RoleService{},
		Permission:  &PermissionService{},
		Resource:    &ResourceService{},
		FeatureFlag: &FeatureFlagService{},
	}
}

// UserService struct
type UserService struct{}

func (s *UserService) GetAll() ([]model.User, error) {
	return GetAllUsers()
}

func (s *UserService) GetByID(id int) (*model.User, error) {
	user, err := repository.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetByEmail(email string) (*model.User, error) {
	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Create(user *model.User) error {
	return AddUser(user)
}

func (s *UserService) Update(id int, user *model.User) error {
	user.ID = id
	return repository.UpdateUser(user)
}

func (s *UserService) Delete(id int) error {
	return repository.DeleteUser(id)
}

// DepartmentService struct
type DepartmentService struct{}

func (s *DepartmentService) GetAll() ([]model.Department, error) {
	return GetAllDepartments()
}

func (s *DepartmentService) GetByID(id int) (*model.Department, error) {
	dept, err := repository.FindDepartmentByID(id)
	if err != nil {
		return nil, err
	}
	return &dept, nil
}

func (s *DepartmentService) Create(dept *model.Department) error {
	return AddDepartment(dept)
}

func (s *DepartmentService) Update(id int, dept *model.Department) error {
	dept.ID = id
	return repository.UpdateDepartment(dept)
}

func (s *DepartmentService) Delete(id int) error {
	return repository.DeleteDepartment(id)
}

// PositionService struct
type PositionService struct{}

func (s *PositionService) GetAll() ([]model.Position, error) {
	return GetAllPositions()
}

func (s *PositionService) GetByID(id int) (*model.Position, error) {
	pos, err := repository.FindPositionByID(id)
	if err != nil {
		return nil, err
	}
	return &pos, nil
}

func (s *PositionService) Create(pos *model.Position) error {
	return AddPosition(pos)
}

func (s *PositionService) Update(id int, pos *model.Position) error {
	pos.ID = id
	return repository.UpdatePosition(pos)
}

func (s *PositionService) Delete(id int) error {
	return repository.DeletePosition(id)
}

// StackService struct
type StackService struct{}

func (s *StackService) GetAll() ([]model.Stack, error) {
	return GetAllStacks()
}

func (s *StackService) GetByID(id int) (*model.Stack, error) {
	stack, err := repository.FindStackByID(id)
	if err != nil {
		return nil, err
	}
	return &stack, nil
}

func (s *StackService) Create(stack *model.Stack) error {
	return AddStack(stack)
}

func (s *StackService) Update(id int, stack *model.Stack) error {
	stack.ID = id
	return repository.UpdateStack(stack)
}

func (s *StackService) Delete(id int) error {
	return repository.DeleteStack(id)
}

// RoleService struct
type RoleService struct{}

func (s *RoleService) GetAll() ([]model.Role, error) {
	return repository.FindAllRoles()
}

func (s *RoleService) GetByID(id int) (*model.Role, error) {
	role, err := repository.FindRoleByID(id)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *RoleService) Create(role *model.Role) error {
	return repository.CreateRole(role)
}

func (s *RoleService) Update(id int, role *model.Role) error {
	role.ID = id
	return repository.UpdateRole(role)
}

func (s *RoleService) Delete(id int) error {
	return repository.DeleteRole(id)
}

// PermissionService struct
type PermissionService struct{}

func (s *PermissionService) GetAll() ([]model.Permission, error) {
	return GetAllPermissions()
}

func (s *PermissionService) GetByID(id int) (*model.Permission, error) {
	perm, err := repository.FindPermissionByID(id)
	if err != nil {
		return nil, err
	}
	return &perm, nil
}

func (s *PermissionService) Create(perm *model.Permission) error {
	return AddPermission(perm)
}

func (s *PermissionService) Update(id int, perm *model.Permission) error {
	perm.ID = id
	return repository.UpdatePermission(perm)
}

func (s *PermissionService) Delete(id int) error {
	return repository.DeletePermission(id)
}

// ResourceService struct
type ResourceService struct{}

func (s *ResourceService) GetAll() ([]model.Resource, error) {
	return GetAllResources()
}

func (s *ResourceService) GetByID(id int) (*model.Resource, error) {
	res, err := repository.FindResourceByID(id)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *ResourceService) Create(res *model.Resource) error {
	return AddResource(res)
}

func (s *ResourceService) Update(id int, res *model.Resource) error {
	res.ID = id
	return repository.UpdateResource(res)
}

func (s *ResourceService) Delete(id int) error {
	return repository.DeleteResource(id)
}

// FeatureFlagService struct
type FeatureFlagService struct{}

func (s *FeatureFlagService) GetAll() ([]model.FeatureFlag, error) {
	return GetAllFeatureFlags()
}

func (s *FeatureFlagService) GetByID(id int) (*model.FeatureFlag, error) {
	flag, err := repository.FindFeatureFlagByID(id)
	if err != nil {
		return nil, err
	}
	return &flag, nil
}

func (s *FeatureFlagService) Create(flag *model.FeatureFlag) error {
	return AddFeatureFlag(flag)
}

func (s *FeatureFlagService) Update(id int, flag *model.FeatureFlag) error {
	flag.ID = id
	return repository.UpdateFeatureFlag(flag)
}

func (s *FeatureFlagService) Delete(id int) error {
	return repository.DeleteFeatureFlag(id)
}
