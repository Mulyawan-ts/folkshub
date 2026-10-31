package main

import "testing"

func TestMenuAppMemilikiMenuDanItem(t *testing.T) {
	a := &App{}
	m := menuApp(a)

	if len(m.Items) != 2 {
		t.Fatalf("menu utama seharusnya 2 submenu, dapat %d", len(m.Items))
	}
	if m.Items[0].Label != "Berkas" || m.Items[1].Label != "Bantuan" {
		t.Fatalf("label menu tidak sesuai: %q, %q", m.Items[0].Label, m.Items[1].Label)
	}
	if m.Items[0].SubMenu == nil || len(m.Items[0].SubMenu.Items) < 3 {
		t.Fatalf("submenu Berkas seharusnya berisi Muat Ulang, separator, dan Keluar")
	}
}
