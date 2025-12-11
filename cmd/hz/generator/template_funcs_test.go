package generator

import "testing"

func TestIsRedundantAlias(t *testing.T) {
	tests := []struct {
		alias       string
		packagePath string
		expected    bool
	}{
		// 原始包名和别名相同的情况
		{"system", "codeup.aliyun.com/gray/hd_admin/biz/model/system", true},
		{"admin", "codeup.aliyun.com/gray/hd_admin/biz/model/admin", true},
		
		// 驼峰别名和包名相同的情况
		{"systemStructs", "codeup.aliyun.com/gray/hd_admin/biz/model/system_structs", true},
		{"adminStructs", "codeup.aliyun.com/gray/hd_admin/biz/model/admin_structs", true},
		{"userModels", "codeup.aliyun.com/gray/hd_admin/biz/model/user_models", true},
		
		// 不冗余的情况
		{"adminService", "codeup.aliyun.com/gray/hd_admin/biz/service/admin", false},
		{"systemStructs", "codeup.aliyun.com/gray/hd_admin/biz/model/admin_structs", false},
		{"customAlias", "codeup.aliyun.com/gray/hd_admin/biz/model/system_structs", false},
	}

	for _, test := range tests {
		result := isRedundantAlias(test.alias, test.packagePath)
		if result != test.expected {
			t.Errorf("isRedundantAlias(%s, %s) = %v, expected %v", test.alias, test.packagePath, result, test.expected)
		}
	}
}
