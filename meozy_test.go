package behaviour
import ("/home/enigma/my_scripts"
          "testing"
)

func TestMeozy(t *testing.T){
	cases := []struct {
		    want string 
	}{
		{want: "attitude"},
	}
for _, c := range cases {
	       result := meozy.attitude()
		   if result != c.want {
			t.Fatalf("[%s] is incorrect, we want [%s]", result, c.want)
		   }
     }
}
