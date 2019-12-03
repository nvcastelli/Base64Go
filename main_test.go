package main

import "testing"

func TestImplementedEncode(t *testing.T) {
    test1 := implementedEncode("")
    if test1 != "" {
       t.Errorf("implementedEncode was incorrect, got: %s.", test1)
	}
	
	test2 := implementedEncode("f")
    if test2 != "Zg==" {
       t.Errorf("implementedEncode was incorrect, got: %s.", test2)
	}
	
	test3 := implementedEncode("fo")
    if test3 != "Zm8=" {
       t.Errorf("implementedEncode was incorrect, got: %s.", test3)
	}
	
	test4 := implementedEncode("foo")
    if test4 != "Zm9v" {
       t.Errorf("implementedEncode was incorrect, got: %s.", test4)
	}
	
	test5 := implementedEncode("foob")
    if test5 != "Zm9vYg==" {
       t.Errorf("implementedEncode was incorrect, got: %s.", test5)
	}
	
	test6 := implementedEncode("fooba")
    if test6 != "Zm9vYmE=" {
       t.Errorf("implementedEncode was incorrect, got: %s.", test6)
	}
	
	test7 := implementedEncode("foobar")
    if test7 != "Zm9vYmFy" {
       t.Errorf("implementedEncode was incorrect, got: %s.", test7)
	}
	
	test8 := implementedEncode("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
    if test8 != "TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdCwgc2VkIGRvIGVpdXNtb2QgdGVtcG9yIGluY2lkaWR1bnQgdXQgbGFib3JlIGV0IGRvbG9yZSBtYWduYSBhbGlxdWEuIFV0IGVuaW0gYWQgbWluaW0gdmVuaWFtLCBxdWlzIG5vc3RydWQgZXhlcmNpdGF0aW9uIHVsbGFtY28gbGFib3JpcyBuaXNpIHV0IGFsaXF1aXAgZXggZWEgY29tbW9kbyBjb25zZXF1YXQuIER1aXMgYXV0ZSBpcnVyZSBkb2xvciBpbiByZXByZWhlbmRlcml0IGluIHZvbHVwdGF0ZSB2ZWxpdCBlc3NlIGNpbGx1bSBkb2xvcmUgZXUgZnVnaWF0IG51bGxhIHBhcmlhdHVyLiBFeGNlcHRldXIgc2ludCBvY2NhZWNhdCBjdXBpZGF0YXQgbm9uIHByb2lkZW50LCBzdW50IGluIGN1bHBhIHF1aSBvZmZpY2lhIGRlc2VydW50IG1vbGxpdCBhbmltIGlkIGVzdCBsYWJvcnVtLg==" {
       t.Errorf("implementedEncode was incorrect, got: %s.", test8)
    }

}