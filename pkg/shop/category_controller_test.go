package shop

import "testing"

func TestListRootCategories(t *testing.T) {
	ctx := MakeTestingCtx()
	var v []interface{}
	resp := PerformTestJSONRequest(ctx.r, "GET", "/api/shop/categories", nil, &v)
	if resp.Code != 200 {
		t.Fatalf("Bad code %v", resp.Code)
	}

}
