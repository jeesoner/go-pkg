package serialize

import (
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	input := `a:2:{i:0;a:7:{s:12:"pc_image_url";s:69:"https://p99.pstatp.com/list/300x170/pgc-image/15404520275419b14c54a1a";s:13:"comment_count";i:520;s:16:"video_play_count";i:951011;s:21:"video_duration_format";s:5:"15:00";s:14:"video_duration";i:900;s:11:"display_url";s:27:"/group/6616191620721148423/";s:5:"title";s:12:"人与自然";}i:1;a:7:{s:5:"title";s:12:"动物世界";s:12:"pc_image_url";s:55:"https://p3.pstatp.com/list/300x170/cc390008433eac69241b";s:13:"comment_count";i:76;s:16:"video_play_count";i:1587177;s:21:"video_duration_format";s:5:"03:55";s:14:"video_duration";i:235;s:11:"display_url";s:27:"/group/6607288769219396103/";}}`

	out, err := Unmarshal([]byte(input))
	if err != nil {
		panic(err)
	}

	result, ok := out.([]interface{})
	if !ok {
		t.Errorf("Unmarshal incorrectly, have got %t\n", out)
	}
	fmt.Println(result)
}
