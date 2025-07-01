package repository

import "my-gin-app/internal/model"

func SeedRoles() {
	var count int64
	DB.Model(&model.Role{}).Count(&count)
	if count == 0 {
		roles := []model.Role{
			{ID: 1, Name: "Admin"},
			{ID: 2, Name: "User"},
			{ID: 3, Name: "Manager"},
			{ID: 4, Name: "Developer"},
			{ID: 5, Name: "Content Creator"},
			{ID: 6, Name: "Acountant"},
		}
		for _, r := range roles {
			DB.Create(&r)
		}
	}
}

func SeedDepartments() {
	var count int64
	DB.Model(&model.Department{}).Count(&count)
	if count == 0 {
		departments := []model.Department{
			{ID: 1, Name: "Default Department"},
			{ID: 2, Name: "Engineering"},
			{ID: 3, Name: "Marketing"},
			{ID: 4, Name: "Sales"},
			{ID: 5, Name: "Human Resources"},
			{ID: 6, Name: "Finance"},
			{ID: 7, Name: "Customer Support"},
			{ID: 8, Name: "IT"},
			{ID: 9, Name: "Legal"},
			{ID: 10, Name: "Operations"},
			{ID: 11, Name: "Product Management"},
			{ID: 12, Name: "Research and Development"},
			{ID: 13, Name: "Business Development"},
			{ID: 14, Name: "Supply Chain"},
			{ID: 15, Name: "Quality Assurance"},
		}
		for _, d := range departments {
			DB.Create(&d)
		}
	}
}

func SeedStacks() {
	var count int64
	DB.Model(&model.Stack{}).Count(&count)
	if count == 0 {
		stacks := []model.Stack{
			{ID: 1, Name: "Default Stack"},
			{ID: 2, Name: "JavaScript"},
			{ID: 3, Name: "Python"},
			{ID: 4, Name: "Java"},
			{ID: 5, Name: "Go"},
			{ID: 6, Name: "Ruby"},
			{ID: 7, Name: "C#"},
			{ID: 8, Name: "PHP"},
			{ID: 9, Name: "C++"},
			{ID: 10, Name: "Swift"},
			{ID: 11, Name: "Kotlin"},
			{ID: 12, Name: "Rust"},
			{ID: 13, Name: "TypeScript"},
			{ID: 14, Name: "Dart"},
			{ID: 15, Name: "Scala"},
			{ID: 16, Name: "Elixir"},
			{ID: 17, Name: "Haskell"},
			{ID: 18, Name: "Lua"},
			{ID: 19, Name: "Shell"},
			{ID: 20, Name: "Clojure"},
			{ID: 21, Name: "Erlang"},
			{ID: 22, Name: "Perl"},
			{ID: 23, Name: "Objective-C"},
			{ID: 24, Name: "R"},
			{ID: 25, Name: "MATLAB"},
			{ID: 26, Name: "Julia"},
			{ID: 27, Name: "Groovy"},
			{ID: 28, Name: "Visual Basic"},
			{ID: 29, Name: "Assembly"},
			{ID: 30, Name: "F#"},
			{ID: 31, Name: "Crystal"},
			{ID: 32, Name: "OCaml"},
			{ID: 33, Name: "ClojureScript"},
			{ID: 34, Name: "Elm"},
			{ID: 35, Name: "Svelte"},
			{ID: 36, Name: "Vue.js"},
			{ID: 37, Name: "Angular"},
			{ID: 38, Name: "React"},
			{ID: 39, Name: "Next.js"},
			{ID: 40, Name: "Nuxt.js"},
			{ID: 41, Name: "Flutter"},
			{ID: 42, Name: "React Native"},
			{ID: 43, Name: "Ionic"},
			{ID: 44, Name: "Cordova"},
			{ID: 45, Name: "Electron"},
			{ID: 46, Name: "ASP.NET"},
			{ID: 47, Name: "Spring"},
			{ID: 48, Name: "Django"},
			{ID: 49, Name: "Flask"},
			{ID: 50, Name: "Ruby on Rails"},
			{ID: 51, Name: "Laravel"},
			{ID: 52, Name: "Express.js"},
			{ID: 53, Name: "FastAPI"},
			{ID: 54, Name: "Gin"},
			{ID: 55, Name: "Phoenix"},
			{ID: 56, Name: "Play Framework"},
			{ID: 57, Name: "Symfony"},
			{ID: 58, Name: "CodeIgniter"},
			{ID: 59, Name: "CakePHP"},
			{ID: 60, Name: "Zend Framework"},
			{ID: 61, Name: "Yii"},
			{ID: 62, Name: "Slim"},
			{ID: 63, Name: "Lumen"},
			{ID: 64, Name: "ASP.NET Core"},
			{ID: 65, Name: "Spring Boot"},
			{ID: 66, Name: "Quarkus"},
			{ID: 67, Name: "Micronaut"},
			{ID: 68, Name: "Ktor"},
			{ID: 69, Name: "Vaadin"},
		}
		for _, s := range stacks {
			DB.Create(&s)
		}
	}
}

func SeedPositions() {
	var count int64
	DB.Model(&model.Position{}).Count(&count)
	if count == 0 {
		positions := []model.Position{
			{ID: 1, Name: "Default Position"},
			{ID: 2, Name: "Software Engineer"},
			{ID: 3, Name: "Senior Software Engineer"},
			{ID: 4, Name: "Lead Software Engineer"},
			{ID: 5, Name: "Software Architect"},
			{ID: 6, Name: "DevOps Engineer"},
			{ID: 7, Name: "Data Scientist"},
			{ID: 8, Name: "Machine Learning Engineer"},
			{ID: 9, Name: "Frontend Developer"},
			{ID: 10, Name: "Backend Developer"},
			{ID: 11, Name: "Full Stack Developer"},
			{ID: 12, Name: "Mobile Developer"},
			{ID: 13, Name: "QA Engineer"},
			{ID: 14, Name: "UI/UX Designer"},
			{ID: 15, Name: "Product Manager"},
			{ID: 16, Name: "Project Manager"},
			{ID: 17, Name: "Business Analyst"},
			{ID: 18, Name: "System Administrator"},
			{ID: 19, Name: "Network Engineer"},
			{ID: 20, Name: "Security Engineer"},
			{ID: 21, Name: "Database Administrator"},
			{ID: 22, Name: "Cloud Engineer"},
			{ID: 23, Name: "Site Reliability Engineer"},
			{ID: 24, Name: "Technical Support Engineer"},
			{ID: 25, Name: "Game Developer"},
			{ID: 26, Name: "Embedded Systems Engineer"},
			{ID: 27, Name: "Blockchain Developer"},
			{ID: 28, Name: "IoT Developer"},
			{ID: 29, Name: "AR/VR Developer"},
			{ID: 30, Name: "Robotics Engineer"},
			{ID: 31, Name: "Data Engineer"},
			{ID: 32, Name: "BI Developer"},
			{ID: 33, Name: "SEO Specialist"},
			{ID: 34, Name: "Content Strategist"},
			{ID: 35, Name: "Digital Marketing Specialist"},
			{ID: 36, Name: "Social Media Manager"},
			{ID: 37, Name: "Sales Engineer"},
			{ID: 38, Name: "Technical Writer"},
			{ID: 39, Name: "Customer Success Manager"},
			{ID: 40, Name: "Training Specialist"},
			{ID: 41, Name: "Business Development Manager"},
			{ID: 42, Name: "Account Manager"},
			{ID: 43, Name: "HR Manager"},
			{ID: 44, Name: "Recruiter"},
			{ID: 45, Name: "Talent Acquisition Specialist"},
			{ID: 46, Name: "Learning and Development Specialist"},
			{ID: 47, Name: "Compensation and Benefits Specialist"},
			{ID: 48, Name: "Organizational Development Specialist"},
			{ID: 49, Name: "Employee Relations Specialist"},
			{ID: 50, Name: "Diversity and Inclusion Specialist"},
			{ID: 51, Name: "Performance Management Specialist"},
			{ID: 52, Name: "Workforce Planning Specialist"},
			{ID: 53, Name: "Change Management Specialist"},
			{ID: 54, Name: "Compliance Specialist"},
			{ID: 55, Name: "Risk Management Specialist"},
			{ID: 56, Name: "Quality Assurance Specialist"},
			{ID: 57, Name: "Regulatory Affairs Specialist"},
			{ID: 58, Name: "Supply Chain Specialist"},
			{ID: 59, Name: "Logistics Specialist"},
			{ID: 60, Name: "Procurement Specialist"},
			{ID: 61, Name: "Inventory Specialist"},
			{ID: 62, Name: "Operations Specialist"},
			{ID: 63, Name: "Facilities Specialist"},
			{ID: 64, Name: "Maintenance Specialist"},
			{ID: 65, Name: "Health and Safety Specialist"},
			{ID: 66, Name: "Environmental Specialist"},
			{ID: 67, Name: "Sustainability Specialist"},
			{ID: 68, Name: "Corporate Social Responsibility Specialist"},
			{ID: 69, Name: "Public Relations Specialist"},
			{ID: 70, Name: "Communications Specialist"},
			{ID: 71, Name: "Brand Manager"},
			{ID: 72, Name: "Market Research Analyst"},
			{ID: 73, Name: "Pricing Analyst"},
			{ID: 74, Name: "Sales Analyst"},
			{ID: 75, Name: "Customer Insights Analyst"},
			{ID: 76, Name: "Business Intelligence Analyst"},
			{ID: 77, Name: "Data Analyst"},
			{ID: 78, Name: "Financial Analyst"},
			{ID: 79, Name: "Investment Analyst"},
			{ID: 80, Name: "Risk Analyst"},
			{ID: 81, Name: "Credit Analyst"},
			{ID: 82, Name: "Compliance Analyst"},
			{ID: 83, Name: "Internal Auditor"},
			{ID: 84, Name: "External Auditor"},
			{ID: 85, Name: "Tax Specialist"},
			{ID: 86, Name: "Accounting Specialist"},
			{ID: 87, Name: "Payroll Specialist"},
			{ID: 88, Name: "Treasury Specialist"},
		}
		for _, p := range positions {
			DB.Create(&p)
		}
	}
}

func SeedUsers() {
	var count int64
	DB.Model(&model.User{}).Count(&count)
	if count == 0 {
		users := []model.User{
			{
				ID:           1,
				Email:        "admin123@gmail.com",
				Password:     "48ee5c0de13758158d77a3a8b27ff9b773348649ed85af990585554eb0349073",
				Name:         "Admin User",
				RoleID:       1, // Admin role
				DepartmentID: 1, // Default Department
				PositionID:   1, // Default Position
				StackID:      1, // Default Stack
			}}
		for _, u := range users {
			DB.Create(&u)
		}
	}
}

func SeedResources() {
	var count int64
	DB.Model(&model.Resource{}).Count(&count)
	if count == 0 {
		resources := []model.Resource{
			{ID: 1, Name: "users"},
			{ID: 2, Name: "roles"},
			{ID: 3, Name: "departments"},
			{ID: 4, Name: "stacks"},
			{ID: 5, Name: "positions"},
			{ID: 6, Name: "permissions"},
			{ID: 7, Name: "auth"},
			{ID: 8, Name: "resources"},
			{ID: 9, Name: "all"}, // special resource for admin
		}
		for _, r := range resources {
			DB.Create(&r)
		}
	}
}

func SeedPermissions() {
	var count int64
	DB.Model(&model.Permission{}).Count(&count)
	if count == 0 {
		permissions := []model.Permission{
			// users
			{ID: 1, ResourceID: 1, Action: "create"},
			{ID: 2, ResourceID: 1, Action: "read"},
			{ID: 3, ResourceID: 1, Action: "update"},
			{ID: 4, ResourceID: 1, Action: "delete"},
			// roles
			{ID: 5, ResourceID: 2, Action: "create"},
			{ID: 6, ResourceID: 2, Action: "read"},
			{ID: 7, ResourceID: 2, Action: "update"},
			{ID: 8, ResourceID: 2, Action: "delete"},
			// departments
			{ID: 9, ResourceID: 3, Action: "create"},
			{ID: 10, ResourceID: 3, Action: "read"},
			{ID: 11, ResourceID: 3, Action: "update"},
			{ID: 12, ResourceID: 3, Action: "delete"},
			// stacks
			{ID: 13, ResourceID: 4, Action: "create"},
			{ID: 14, ResourceID: 4, Action: "read"},
			{ID: 15, ResourceID: 4, Action: "update"},
			{ID: 16, ResourceID: 4, Action: "delete"},
			// positions
			{ID: 17, ResourceID: 5, Action: "create"},
			{ID: 18, ResourceID: 5, Action: "read"},
			{ID: 19, ResourceID: 5, Action: "update"},
			{ID: 20, ResourceID: 5, Action: "delete"},
			// permissions
			{ID: 21, ResourceID: 6, Action: "create"},
			{ID: 22, ResourceID: 6, Action: "read"},
			{ID: 23, ResourceID: 6, Action: "update"},
			{ID: 24, ResourceID: 6, Action: "delete"},
			// auth
			{ID: 25, ResourceID: 7, Action: "login"},
			{ID: 26, ResourceID: 7, Action: "register"},
			{ID: 27, ResourceID: 7, Action: "forgot_password"},
			{ID: 28, ResourceID: 7, Action: "reset_password"},
			{ID: 29, ResourceID: 7, Action: "change_password"},
			{ID: 30, ResourceID: 7, Action: "profile"},
			{ID: 31, ResourceID: 7, Action: "logout"},
			{ID: 32, ResourceID: 7, Action: "refresh_token"},
			{ID: 33, ResourceID: 7, Action: "verify_email"},
			{ID: 34, ResourceID: 7, Action: "send_verification_email"},
			{ID: 35, ResourceID: 7, Action: "resend_verification_email"},
			// resources
			{ID: 36, ResourceID: 8, Action: "create"},
			{ID: 37, ResourceID: 8, Action: "read"},
			{ID: 38, ResourceID: 8, Action: "update"},
			{ID: 39, ResourceID: 8, Action: "delete"},
			// special: manage::all for admin
			{ID: 40, ResourceID: 9, Action: "manage"},
		}
		for _, p := range permissions {
			DB.Create(&p)
		}
	}

}

func SeedData() {
	SeedRoles()
	SeedDepartments()
	SeedStacks()
	SeedPositions()
	SeedUsers()
	SeedResources()
	SeedPermissions()

	var manageAll model.Permission
	DB.Where("resource_id = ? AND action = ?", 9, "manage").First(&manageAll)
	var admin model.Role
	DB.First(&admin, 1)
	DB.Model(&admin).Association("Permissions").Append(&manageAll)
}
