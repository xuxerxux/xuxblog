package xuxblog
import "testing"

func TestHello( t *testing.T){
		t.Run("different input different outputs...", func(t *testing.T){
			got := Hello("hahah")
			want:= "hello hahah"
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
		t.Run("empty input how to deal", func(t *testing.T){
			got := Hello("")
			want:="hello world"
			if got != want {
				t.Errorf("want %q is not same as got %q", want, got)
				}
		})
}
