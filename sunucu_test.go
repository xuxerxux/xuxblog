package xuxblog
import "testing"

func TestHello( t *testing.T){
	want := "merhaba dunya"
	got  := Hello()
	if got != want {
		t.Errorf("hata oldu")
	}
	}
