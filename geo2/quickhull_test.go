package geo2

import (
	"github.com/luxengine/glm"
	"github.com/luxengine/glm/geo2/internal/qhull"
	"github.com/luxengine/math"
	"testing"
)

func TestQuickhull(t *testing.T) {
	// These tests we're visualised and verified by hand (using an svg debug
	// program), then the data was exported here. If the algorithm is changed
	// and this breaks it may not be that the algorithm is broken it may just be
	// that the data is now in a different order (this test only does slice
	// content equality).
	tests := []struct {
		points []glm.Vec2
		hull   []glm.Vec2
	}{
		{
			points: []glm.Vec2{{0, 0}, {0.5, -0.1}, {1, 0}, {1.1, 1}, {0, 1}, {0.5, 0.5}, {-0.4, 0.5}, {0.4, 0.4}, {0.5, 0.4}, {0.6, 0.6}},
			hull:   []glm.Vec2{{-0.4, 0.5}, {0, 1}, {1.1, 1}, {1, 0}, {0.5, -0.1}, {0, 0}},
		},
		{
			points: []glm.Vec2{{0.5, 0.5}, {-0.4, 0.5}, {0.4, 0.4}, {0.5, 0.4}, {0.6, 0.6}, {0, 0}, {0.5, -0.1}, {1, 0}, {1.1, 1}, {0, 1}},
			hull:   []glm.Vec2{{-0.4, 0.5}, {0, 1}, {1.1, 1}, {1, 0}, {0.5, -0.1}, {0, 0}},
		},
		{
			points: []glm.Vec2{{270, 319}, {373, 174}, {396, 92}, {321, 148}, {354, 416}, {290, 322}, {9, 457}, {361, 328}, {421, 49}, {486, 256}},
			hull:   []glm.Vec2{{9, 457}, {354, 416}, {486, 256}, {421, 49}, {321, 148}},
		},
		{
			points: []glm.Vec2{glm.Vec2{62, 482}, glm.Vec2{448, 208}, glm.Vec2{393, 395}, glm.Vec2{75, 180}, glm.Vec2{201, 136}, glm.Vec2{339, 22}, glm.Vec2{354, 491}, glm.Vec2{7, 369}, glm.Vec2{155, 108}, glm.Vec2{252, 352}},
			hull:   []glm.Vec2{glm.Vec2{7, 369}, glm.Vec2{62, 482}, glm.Vec2{354, 491}, glm.Vec2{393, 395}, glm.Vec2{448, 208}, glm.Vec2{339, 22}, glm.Vec2{155, 108}, glm.Vec2{75, 180}},
		},
		{
			points: []glm.Vec2{glm.Vec2{378, 112}, glm.Vec2{239, 65}, glm.Vec2{202, 239}, glm.Vec2{277, 240}, glm.Vec2{451, 269}, glm.Vec2{498, 127}, glm.Vec2{322, 490}, glm.Vec2{115, 442}, glm.Vec2{415, 119}, glm.Vec2{135, 496}},
			hull:   []glm.Vec2{glm.Vec2{115, 442}, glm.Vec2{135, 496}, glm.Vec2{322, 490}, glm.Vec2{451, 269}, glm.Vec2{498, 127}, glm.Vec2{239, 65}},
		},
		{
			points: []glm.Vec2{glm.Vec2{43, 429}, glm.Vec2{136, 262}, glm.Vec2{79, 17}, glm.Vec2{383, 374}, glm.Vec2{216, 39}, glm.Vec2{61, 347}, glm.Vec2{244, 343}, glm.Vec2{308, 172}, glm.Vec2{283, 422}, glm.Vec2{405, 252}},
			hull:   []glm.Vec2{glm.Vec2{43, 429}, glm.Vec2{283, 422}, glm.Vec2{383, 374}, glm.Vec2{405, 252}, glm.Vec2{216, 39}, glm.Vec2{79, 17}},
		},
		{
			points: []glm.Vec2{glm.Vec2{65, 427}, glm.Vec2{150, 399}, glm.Vec2{242, 55}, glm.Vec2{99, 136}, glm.Vec2{182, 339}, glm.Vec2{344, 185}, glm.Vec2{38, 109}, glm.Vec2{410, 332}, glm.Vec2{366, 41}, glm.Vec2{133, 254}},
			hull:   []glm.Vec2{glm.Vec2{38, 109}, glm.Vec2{65, 427}, glm.Vec2{410, 332}, glm.Vec2{366, 41}, glm.Vec2{242, 55}},
		},
		{
			points: []glm.Vec2{glm.Vec2{379, 227}, glm.Vec2{111, 451}, glm.Vec2{134, 225}, glm.Vec2{3, 441}, glm.Vec2{163, 196}, glm.Vec2{296, 452}, glm.Vec2{356, 46}, glm.Vec2{63, 196}, glm.Vec2{73, 410}, glm.Vec2{448, 214}},
			hull:   []glm.Vec2{glm.Vec2{3, 441}, glm.Vec2{111, 451}, glm.Vec2{296, 452}, glm.Vec2{448, 214}, glm.Vec2{356, 46}, glm.Vec2{63, 196}},
		},
		{
			points: []glm.Vec2{glm.Vec2{33, 349}, glm.Vec2{441, 345}, glm.Vec2{7, 98}, glm.Vec2{82, 46}, glm.Vec2{283, 303}, glm.Vec2{184, 39}, glm.Vec2{179, 210}, glm.Vec2{433, 58}, glm.Vec2{81, 376}, glm.Vec2{157, 492}},
			hull:   []glm.Vec2{glm.Vec2{7, 98}, glm.Vec2{33, 349}, glm.Vec2{157, 492}, glm.Vec2{441, 345}, glm.Vec2{433, 58}, glm.Vec2{184, 39}, glm.Vec2{82, 46}},
		},
		{
			points: []glm.Vec2{glm.Vec2{12, 60}, glm.Vec2{235, 399}, glm.Vec2{344, 439}, glm.Vec2{215, 225}, glm.Vec2{424, 141}, glm.Vec2{456, 471}, glm.Vec2{129, 421}, glm.Vec2{42, 31}, glm.Vec2{483, 490}, glm.Vec2{16, 90}},
			hull:   []glm.Vec2{glm.Vec2{12, 60}, glm.Vec2{16, 90}, glm.Vec2{129, 421}, glm.Vec2{483, 490}, glm.Vec2{424, 141}, glm.Vec2{42, 31}},
		},
		{
			points: []glm.Vec2{glm.Vec2{178, 453}, glm.Vec2{47, 375}, glm.Vec2{10, 147}, glm.Vec2{240, 111}, glm.Vec2{197, 135}, glm.Vec2{326, 258}, glm.Vec2{19, 44}, glm.Vec2{449, 350}, glm.Vec2{439, 201}, glm.Vec2{454, 409}},
			hull:   []glm.Vec2{glm.Vec2{10, 147}, glm.Vec2{47, 375}, glm.Vec2{178, 453}, glm.Vec2{454, 409}, glm.Vec2{439, 201}, glm.Vec2{240, 111}, glm.Vec2{19, 44}},
		},
		{
			points: []glm.Vec2{glm.Vec2{90, 352}, glm.Vec2{243, 287}, glm.Vec2{368, 105}, glm.Vec2{55, 476}, glm.Vec2{427, 44}, glm.Vec2{386, 288}, glm.Vec2{303, 100}, glm.Vec2{4, 47}, glm.Vec2{223, 172}, glm.Vec2{352, 267}},
			hull:   []glm.Vec2{glm.Vec2{4, 47}, glm.Vec2{55, 476}, glm.Vec2{386, 288}, glm.Vec2{427, 44}},
		},
		{
			points: []glm.Vec2{glm.Vec2{165, 27}, glm.Vec2{495, 340}, glm.Vec2{472, 28}, glm.Vec2{403, 158}, glm.Vec2{300, 277}, glm.Vec2{107, 231}, glm.Vec2{471, 462}, glm.Vec2{445, 38}, glm.Vec2{266, 495}, glm.Vec2{487, 180}},
			hull:   []glm.Vec2{glm.Vec2{107, 231}, glm.Vec2{266, 495}, glm.Vec2{471, 462}, glm.Vec2{495, 340}, glm.Vec2{487, 180}, glm.Vec2{472, 28}, glm.Vec2{165, 27}},
		},
		{
			points: []glm.Vec2{glm.Vec2{469, 72}, glm.Vec2{173, 361}, glm.Vec2{364, 48}, glm.Vec2{356, 466}, glm.Vec2{274, 199}, glm.Vec2{252, 418}, glm.Vec2{78, 309}, glm.Vec2{424, 181}, glm.Vec2{265, 45}, glm.Vec2{493, 162}},
			hull:   []glm.Vec2{glm.Vec2{78, 309}, glm.Vec2{252, 418}, glm.Vec2{356, 466}, glm.Vec2{493, 162}, glm.Vec2{469, 72}, glm.Vec2{364, 48}, glm.Vec2{265, 45}},
		},
		{
			points: []glm.Vec2{glm.Vec2{105, 414}, glm.Vec2{26, 328}, glm.Vec2{409, 182}, glm.Vec2{314, 144}, glm.Vec2{364, 30}, glm.Vec2{0, 180}, glm.Vec2{424, 112}, glm.Vec2{351, 19}, glm.Vec2{409, 353}, glm.Vec2{2, 223}},
			hull:   []glm.Vec2{glm.Vec2{0, 180}, glm.Vec2{2, 223}, glm.Vec2{26, 328}, glm.Vec2{105, 414}, glm.Vec2{409, 353}, glm.Vec2{424, 112}, glm.Vec2{364, 30}, glm.Vec2{351, 19}},
		},
		{
			points: []glm.Vec2{glm.Vec2{217, 311}, glm.Vec2{130, 34}, glm.Vec2{24, 381}, glm.Vec2{450, 78}, glm.Vec2{413, 409}, glm.Vec2{495, 231}, glm.Vec2{440, 109}, glm.Vec2{127, 282}, glm.Vec2{325, 468}, glm.Vec2{20, 59}},
			hull:   []glm.Vec2{glm.Vec2{20, 59}, glm.Vec2{24, 381}, glm.Vec2{325, 468}, glm.Vec2{413, 409}, glm.Vec2{495, 231}, glm.Vec2{450, 78}, glm.Vec2{130, 34}},
		},
		{
			points: []glm.Vec2{glm.Vec2{463, 367}, glm.Vec2{350, 449}, glm.Vec2{188, 476}, glm.Vec2{163, 67}, glm.Vec2{264, 323}, glm.Vec2{54, 129}, glm.Vec2{189, 265}, glm.Vec2{112, 195}, glm.Vec2{115, 471}, glm.Vec2{40, 492}},
			hull:   []glm.Vec2{glm.Vec2{40, 492}, glm.Vec2{188, 476}, glm.Vec2{350, 449}, glm.Vec2{463, 367}, glm.Vec2{163, 67}, glm.Vec2{54, 129}},
		},
		{
			points: []glm.Vec2{glm.Vec2{168, 343}, glm.Vec2{228, 411}, glm.Vec2{471, 430}, glm.Vec2{95, 473}, glm.Vec2{198, 9}, glm.Vec2{265, 194}, glm.Vec2{330, 408}, glm.Vec2{214, 309}, glm.Vec2{270, 261}, glm.Vec2{199, 60}},
			hull:   []glm.Vec2{glm.Vec2{95, 473}, glm.Vec2{471, 430}, glm.Vec2{198, 9}},
		},
		{
			points: []glm.Vec2{glm.Vec2{369, 122}, glm.Vec2{181, 433}, glm.Vec2{383, 388}, glm.Vec2{45, 37}, glm.Vec2{396, 201}, glm.Vec2{466, 22}, glm.Vec2{479, 154}, glm.Vec2{437, 168}, glm.Vec2{23, 462}, glm.Vec2{373, 351}},
			hull:   []glm.Vec2{glm.Vec2{23, 462}, glm.Vec2{181, 433}, glm.Vec2{383, 388}, glm.Vec2{479, 154}, glm.Vec2{466, 22}, glm.Vec2{45, 37}},
		},
		{
			points: []glm.Vec2{glm.Vec2{270, 176}, glm.Vec2{437, 60}, glm.Vec2{14, 438}, glm.Vec2{473, 487}, glm.Vec2{355, 147}, glm.Vec2{407, 152}, glm.Vec2{254, 396}, glm.Vec2{3, 294}, glm.Vec2{298, 375}, glm.Vec2{493, 220}},
			hull:   []glm.Vec2{glm.Vec2{3, 294}, glm.Vec2{14, 438}, glm.Vec2{473, 487}, glm.Vec2{493, 220}, glm.Vec2{437, 60}},
		},
		{
			points: []glm.Vec2{glm.Vec2{146, 389}, glm.Vec2{415, 404}, glm.Vec2{51, 325}, glm.Vec2{428, 298}, glm.Vec2{365, 441}, glm.Vec2{183, 201}, glm.Vec2{334, 235}, glm.Vec2{473, 374}, glm.Vec2{359, 261}, glm.Vec2{102, 405}},
			hull:   []glm.Vec2{glm.Vec2{51, 325}, glm.Vec2{102, 405}, glm.Vec2{365, 441}, glm.Vec2{473, 374}, glm.Vec2{428, 298}, glm.Vec2{334, 235}, glm.Vec2{183, 201}},
		},
		{
			points: []glm.Vec2{glm.Vec2{448, 395}, glm.Vec2{195, 479}, glm.Vec2{353, 300}, glm.Vec2{276, 496}, glm.Vec2{102, 110}, glm.Vec2{108, 233}, glm.Vec2{420, 329}, glm.Vec2{324, 82}, glm.Vec2{227, 423}, glm.Vec2{126, 49}},
			hull:   []glm.Vec2{glm.Vec2{102, 110}, glm.Vec2{108, 233}, glm.Vec2{195, 479}, glm.Vec2{276, 496}, glm.Vec2{448, 395}, glm.Vec2{324, 82}, glm.Vec2{126, 49}},
		},
		{
			points: []glm.Vec2{glm.Vec2{22, 61}, glm.Vec2{328, 449}, glm.Vec2{289, 59}, glm.Vec2{35, 82}, glm.Vec2{134, 40}, glm.Vec2{224, 490}, glm.Vec2{100, 319}, glm.Vec2{400, 146}, glm.Vec2{165, 133}, glm.Vec2{47, 145}},
			hull:   []glm.Vec2{glm.Vec2{22, 61}, glm.Vec2{47, 145}, glm.Vec2{100, 319}, glm.Vec2{224, 490}, glm.Vec2{328, 449}, glm.Vec2{400, 146}, glm.Vec2{289, 59}, glm.Vec2{134, 40}},
		},
		{
			points: []glm.Vec2{glm.Vec2{163, 90}, glm.Vec2{40, 319}, glm.Vec2{285, 289}, glm.Vec2{402, 328}, glm.Vec2{237, 114}, glm.Vec2{438, 218}, glm.Vec2{65, 146}, glm.Vec2{287, 115}, glm.Vec2{0, 357}, glm.Vec2{362, 466}},
			hull:   []glm.Vec2{glm.Vec2{0, 357}, glm.Vec2{362, 466}, glm.Vec2{438, 218}, glm.Vec2{287, 115}, glm.Vec2{163, 90}, glm.Vec2{65, 146}},
		},
		{
			points: []glm.Vec2{glm.Vec2{323, 76}, glm.Vec2{279, 49}, glm.Vec2{9, 300}, glm.Vec2{133, 166}, glm.Vec2{454, 451}, glm.Vec2{422, 225}, glm.Vec2{429, 35}, glm.Vec2{427, 348}, glm.Vec2{404, 477}, glm.Vec2{492, 153}},
			hull:   []glm.Vec2{glm.Vec2{9, 300}, glm.Vec2{404, 477}, glm.Vec2{454, 451}, glm.Vec2{492, 153}, glm.Vec2{429, 35}, glm.Vec2{279, 49}, glm.Vec2{133, 166}},
		},
		{
			points: []glm.Vec2{glm.Vec2{456, 199}, glm.Vec2{193, 409}, glm.Vec2{53, 446}, glm.Vec2{37, 71}, glm.Vec2{155, 137}, glm.Vec2{44, 83}, glm.Vec2{424, 356}, glm.Vec2{328, 292}, glm.Vec2{219, 169}, glm.Vec2{106, 449}},
			hull:   []glm.Vec2{glm.Vec2{37, 71}, glm.Vec2{53, 446}, glm.Vec2{106, 449}, glm.Vec2{424, 356}, glm.Vec2{456, 199}},
		},
		{
			points: []glm.Vec2{glm.Vec2{171, 438}, glm.Vec2{256, 371}, glm.Vec2{453, 118}, glm.Vec2{221, 397}, glm.Vec2{232, 203}, glm.Vec2{110, 345}, glm.Vec2{237, 74}, glm.Vec2{77, 17}, glm.Vec2{351, 430}, glm.Vec2{421, 204}},
			hull:   []glm.Vec2{glm.Vec2{77, 17}, glm.Vec2{110, 345}, glm.Vec2{171, 438}, glm.Vec2{351, 430}, glm.Vec2{453, 118}},
		},
		{
			points: []glm.Vec2{glm.Vec2{462, 225}, glm.Vec2{427, 212}, glm.Vec2{462, 34}, glm.Vec2{79, 453}, glm.Vec2{389, 7}, glm.Vec2{247, 405}, glm.Vec2{242, 160}, glm.Vec2{342, 323}, glm.Vec2{69, 164}, glm.Vec2{205, 199}},
			hull:   []glm.Vec2{glm.Vec2{69, 164}, glm.Vec2{79, 453}, glm.Vec2{247, 405}, glm.Vec2{462, 225}, glm.Vec2{462, 34}, glm.Vec2{389, 7}},
		},
		{
			points: []glm.Vec2{glm.Vec2{307, 332}, glm.Vec2{412, 83}, glm.Vec2{23, 426}, glm.Vec2{80, 103}, glm.Vec2{126, 27}, glm.Vec2{147, 130}, glm.Vec2{442, 292}, glm.Vec2{127, 93}, glm.Vec2{153, 238}, glm.Vec2{485, 398}},
			hull:   []glm.Vec2{glm.Vec2{23, 426}, glm.Vec2{485, 398}, glm.Vec2{412, 83}, glm.Vec2{126, 27}, glm.Vec2{80, 103}},
		},
		{
			points: []glm.Vec2{glm.Vec2{123, 26}, glm.Vec2{197, 334}, glm.Vec2{170, 349}, glm.Vec2{405, 354}, glm.Vec2{132, 359}, glm.Vec2{373, 356}, glm.Vec2{209, 3}, glm.Vec2{54, 84}, glm.Vec2{401, 401}, glm.Vec2{86, 398}},
			hull:   []glm.Vec2{glm.Vec2{54, 84}, glm.Vec2{86, 398}, glm.Vec2{401, 401}, glm.Vec2{405, 354}, glm.Vec2{209, 3}, glm.Vec2{123, 26}},
		},
		{
			points: []glm.Vec2{glm.Vec2{42, 257}, glm.Vec2{68, 185}, glm.Vec2{370, 469}, glm.Vec2{195, 153}, glm.Vec2{320, 401}, glm.Vec2{483, 302}, glm.Vec2{473, 441}, glm.Vec2{494, 457}, glm.Vec2{379, 89}, glm.Vec2{12, 92}},
			hull:   []glm.Vec2{glm.Vec2{12, 92}, glm.Vec2{42, 257}, glm.Vec2{370, 469}, glm.Vec2{494, 457}, glm.Vec2{483, 302}, glm.Vec2{379, 89}},
		},
		{
			points: []glm.Vec2{glm.Vec2{351, 81}, glm.Vec2{304, 171}, glm.Vec2{466, 139}, glm.Vec2{476, 289}, glm.Vec2{8, 409}, glm.Vec2{457, 295}, glm.Vec2{453, 136}, glm.Vec2{116, 494}, glm.Vec2{70, 273}, glm.Vec2{314, 144}},
			hull:   []glm.Vec2{glm.Vec2{8, 409}, glm.Vec2{116, 494}, glm.Vec2{476, 289}, glm.Vec2{466, 139}, glm.Vec2{351, 81}, glm.Vec2{70, 273}},
		},
		{
			points: []glm.Vec2{glm.Vec2{142, 55}, glm.Vec2{56, 427}, glm.Vec2{439, 65}, glm.Vec2{320, 23}, glm.Vec2{310, 464}, glm.Vec2{269, 445}, glm.Vec2{402, 446}, glm.Vec2{273, 61}, glm.Vec2{76, 172}, glm.Vec2{345, 225}},
			hull:   []glm.Vec2{glm.Vec2{56, 427}, glm.Vec2{310, 464}, glm.Vec2{402, 446}, glm.Vec2{439, 65}, glm.Vec2{320, 23}, glm.Vec2{142, 55}, glm.Vec2{76, 172}},
		},
		{
			points: []glm.Vec2{glm.Vec2{115, 422}, glm.Vec2{91, 125}, glm.Vec2{2, 249}, glm.Vec2{412, 66}, glm.Vec2{276, 247}, glm.Vec2{83, 386}, glm.Vec2{115, 221}, glm.Vec2{70, 117}, glm.Vec2{40, 390}, glm.Vec2{348, 282}},
			hull:   []glm.Vec2{glm.Vec2{2, 249}, glm.Vec2{40, 390}, glm.Vec2{115, 422}, glm.Vec2{348, 282}, glm.Vec2{412, 66}, glm.Vec2{70, 117}},
		},
		{
			points: []glm.Vec2{glm.Vec2{289, 372}, glm.Vec2{310, 338}, glm.Vec2{310, 86}, glm.Vec2{88, 352}, glm.Vec2{486, 99}, glm.Vec2{327, 304}, glm.Vec2{144, 226}, glm.Vec2{135, 146}, glm.Vec2{249, 352}, glm.Vec2{427, 99}},
			hull:   []glm.Vec2{glm.Vec2{88, 352}, glm.Vec2{289, 372}, glm.Vec2{486, 99}, glm.Vec2{310, 86}, glm.Vec2{135, 146}},
		},
		{
			points: []glm.Vec2{glm.Vec2{486, 217}, glm.Vec2{52, 24}, glm.Vec2{352, 56}, glm.Vec2{326, 368}, glm.Vec2{11, 363}, glm.Vec2{186, 205}, glm.Vec2{67, 406}, glm.Vec2{435, 45}, glm.Vec2{488, 350}, glm.Vec2{406, 279}},
			hull:   []glm.Vec2{glm.Vec2{11, 363}, glm.Vec2{67, 406}, glm.Vec2{488, 350}, glm.Vec2{486, 217}, glm.Vec2{435, 45}, glm.Vec2{52, 24}},
		},
		{
			points: []glm.Vec2{glm.Vec2{426, 443}, glm.Vec2{213, 7}, glm.Vec2{412, 5}, glm.Vec2{71, 466}, glm.Vec2{73, 315}, glm.Vec2{33, 290}, glm.Vec2{201, 263}, glm.Vec2{227, 222}, glm.Vec2{159, 129}, glm.Vec2{475, 441}},
			hull:   []glm.Vec2{glm.Vec2{33, 290}, glm.Vec2{71, 466}, glm.Vec2{475, 441}, glm.Vec2{412, 5}, glm.Vec2{213, 7}},
		},
		{
			points: []glm.Vec2{glm.Vec2{437, 148}, glm.Vec2{446, 304}, glm.Vec2{260, 198}, glm.Vec2{401, 167}, glm.Vec2{342, 111}, glm.Vec2{308, 497}, glm.Vec2{192, 326}, glm.Vec2{17, 443}, glm.Vec2{245, 305}, glm.Vec2{253, 490}},
			hull:   []glm.Vec2{glm.Vec2{17, 443}, glm.Vec2{253, 490}, glm.Vec2{308, 497}, glm.Vec2{446, 304}, glm.Vec2{437, 148}, glm.Vec2{342, 111}},
		},
		{
			points: []glm.Vec2{glm.Vec2{213, 18}, glm.Vec2{147, 496}, glm.Vec2{136, 377}, glm.Vec2{250, 461}, glm.Vec2{197, 346}, glm.Vec2{322, 0}, glm.Vec2{251, 39}, glm.Vec2{246, 476}, glm.Vec2{212, 238}, glm.Vec2{485, 217}},
			hull:   []glm.Vec2{glm.Vec2{136, 377}, glm.Vec2{147, 496}, glm.Vec2{246, 476}, glm.Vec2{485, 217}, glm.Vec2{322, 0}, glm.Vec2{213, 18}},
		},
		{
			points: []glm.Vec2{glm.Vec2{282, 88}, glm.Vec2{196, 126}, glm.Vec2{460, 81}, glm.Vec2{348, 341}, glm.Vec2{353, 454}, glm.Vec2{69, 451}, glm.Vec2{388, 386}, glm.Vec2{164, 180}, glm.Vec2{13, 410}, glm.Vec2{89, 350}},
			hull:   []glm.Vec2{glm.Vec2{13, 410}, glm.Vec2{69, 451}, glm.Vec2{353, 454}, glm.Vec2{388, 386}, glm.Vec2{460, 81}, glm.Vec2{282, 88}, glm.Vec2{196, 126}},
		},
		{
			points: []glm.Vec2{glm.Vec2{318, 274}, glm.Vec2{235, 93}, glm.Vec2{335, 253}, glm.Vec2{164, 256}, glm.Vec2{471, 398}, glm.Vec2{306, 432}, glm.Vec2{352, 202}, glm.Vec2{336, 124}, glm.Vec2{348, 289}, glm.Vec2{292, 99}},
			hull:   []glm.Vec2{glm.Vec2{164, 256}, glm.Vec2{306, 432}, glm.Vec2{471, 398}, glm.Vec2{336, 124}, glm.Vec2{292, 99}, glm.Vec2{235, 93}},
		},
		{
			points: []glm.Vec2{glm.Vec2{498, 338}, glm.Vec2{358, 187}, glm.Vec2{107, 196}, glm.Vec2{70, 433}, glm.Vec2{175, 313}, glm.Vec2{452, 73}, glm.Vec2{305, 363}, glm.Vec2{308, 339}, glm.Vec2{128, 409}, glm.Vec2{214, 383}},
			hull:   []glm.Vec2{glm.Vec2{70, 433}, glm.Vec2{498, 338}, glm.Vec2{452, 73}, glm.Vec2{107, 196}},
		},
		{
			points: []glm.Vec2{glm.Vec2{204, 131}, glm.Vec2{74, 370}, glm.Vec2{303, 294}, glm.Vec2{59, 292}, glm.Vec2{121, 149}, glm.Vec2{78, 457}, glm.Vec2{427, 26}, glm.Vec2{184, 497}, glm.Vec2{111, 277}, glm.Vec2{80, 377}},
			hull:   []glm.Vec2{glm.Vec2{59, 292}, glm.Vec2{78, 457}, glm.Vec2{184, 497}, glm.Vec2{303, 294}, glm.Vec2{427, 26}, glm.Vec2{121, 149}},
		},
		{
			points: []glm.Vec2{glm.Vec2{317, 77}, glm.Vec2{397, 139}, glm.Vec2{352, 197}, glm.Vec2{183, 23}, glm.Vec2{156, 75}, glm.Vec2{262, 451}, glm.Vec2{446, 399}, glm.Vec2{224, 343}, glm.Vec2{385, 122}, glm.Vec2{388, 273}},
			hull:   []glm.Vec2{glm.Vec2{156, 75}, glm.Vec2{224, 343}, glm.Vec2{262, 451}, glm.Vec2{446, 399}, glm.Vec2{397, 139}, glm.Vec2{385, 122}, glm.Vec2{317, 77}, glm.Vec2{183, 23}},
		},
		{
			points: []glm.Vec2{glm.Vec2{463, 245}, glm.Vec2{83, 249}, glm.Vec2{347, 466}, glm.Vec2{103, 100}, glm.Vec2{372, 405}, glm.Vec2{417, 421}, glm.Vec2{274, 166}, glm.Vec2{226, 169}, glm.Vec2{202, 467}, glm.Vec2{320, 401}},
			hull:   []glm.Vec2{glm.Vec2{83, 249}, glm.Vec2{202, 467}, glm.Vec2{347, 466}, glm.Vec2{417, 421}, glm.Vec2{463, 245}, glm.Vec2{274, 166}, glm.Vec2{103, 100}},
		},
		{
			points: []glm.Vec2{glm.Vec2{401, 359}, glm.Vec2{429, 366}, glm.Vec2{148, 455}, glm.Vec2{360, 334}, glm.Vec2{306, 200}, glm.Vec2{308, 339}, glm.Vec2{26, 93}, glm.Vec2{133, 470}, glm.Vec2{336, 435}, glm.Vec2{296, 260}},
			hull:   []glm.Vec2{glm.Vec2{26, 93}, glm.Vec2{133, 470}, glm.Vec2{336, 435}, glm.Vec2{429, 366}, glm.Vec2{306, 200}},
		},
		{
			points: []glm.Vec2{glm.Vec2{133, 438}, glm.Vec2{247, 95}, glm.Vec2{413, 423}, glm.Vec2{5, 237}, glm.Vec2{429, 427}, glm.Vec2{121, 33}, glm.Vec2{201, 279}, glm.Vec2{490, 355}, glm.Vec2{449, 56}, glm.Vec2{378, 102}},
			hull:   []glm.Vec2{glm.Vec2{5, 237}, glm.Vec2{133, 438}, glm.Vec2{429, 427}, glm.Vec2{490, 355}, glm.Vec2{449, 56}, glm.Vec2{121, 33}},
		},
		{
			points: []glm.Vec2{glm.Vec2{331, 389}, glm.Vec2{442, 199}, glm.Vec2{357, 155}, glm.Vec2{213, 238}, glm.Vec2{460, 269}, glm.Vec2{302, 320}, glm.Vec2{213, 385}, glm.Vec2{346, 9}, glm.Vec2{347, 75}, glm.Vec2{281, 419}},
			hull:   []glm.Vec2{glm.Vec2{213, 238}, glm.Vec2{213, 385}, glm.Vec2{281, 419}, glm.Vec2{331, 389}, glm.Vec2{460, 269}, glm.Vec2{442, 199}, glm.Vec2{346, 9}},
		},
		{
			points: []glm.Vec2{glm.Vec2{261, 201}, glm.Vec2{408, 98}, glm.Vec2{66, 343}, glm.Vec2{287, 5}, glm.Vec2{269, 430}, glm.Vec2{285, 311}, glm.Vec2{487, 32}, glm.Vec2{255, 448}, glm.Vec2{225, 318}, glm.Vec2{78, 305}},
			hull:   []glm.Vec2{glm.Vec2{66, 343}, glm.Vec2{255, 448}, glm.Vec2{269, 430}, glm.Vec2{487, 32}, glm.Vec2{287, 5}, glm.Vec2{78, 305}},
		},
		{
			points: []glm.Vec2{glm.Vec2{22, 449}, glm.Vec2{31, 210}, glm.Vec2{168, 440}, glm.Vec2{257, 206}, glm.Vec2{221, 426}, glm.Vec2{411, 421}, glm.Vec2{218, 399}, glm.Vec2{431, 68}, glm.Vec2{154, 319}, glm.Vec2{158, 326}},
			hull:   []glm.Vec2{glm.Vec2{22, 449}, glm.Vec2{168, 440}, glm.Vec2{411, 421}, glm.Vec2{431, 68}, glm.Vec2{31, 210}},
		},
		{
			points: []glm.Vec2{glm.Vec2{362, 76}, glm.Vec2{110, 332}, glm.Vec2{166, 194}, glm.Vec2{483, 373}, glm.Vec2{111, 74}, glm.Vec2{34, 26}, glm.Vec2{267, 366}, glm.Vec2{446, 311}, glm.Vec2{245, 443}, glm.Vec2{400, 235}},
			hull:   []glm.Vec2{glm.Vec2{34, 26}, glm.Vec2{110, 332}, glm.Vec2{245, 443}, glm.Vec2{483, 373}, glm.Vec2{362, 76}},
		},
		{
			points: []glm.Vec2{glm.Vec2{403, 316}, glm.Vec2{486, 168}, glm.Vec2{259, 294}, glm.Vec2{231, 466}, glm.Vec2{338, 440}, glm.Vec2{285, 106}, glm.Vec2{290, 490}, glm.Vec2{15, 249}, glm.Vec2{106, 244}, glm.Vec2{110, 171}},
			hull:   []glm.Vec2{glm.Vec2{15, 249}, glm.Vec2{231, 466}, glm.Vec2{290, 490}, glm.Vec2{338, 440}, glm.Vec2{486, 168}, glm.Vec2{285, 106}, glm.Vec2{110, 171}},
		},
		{
			points: []glm.Vec2{glm.Vec2{290, 21}, glm.Vec2{428, 49}, glm.Vec2{424, 377}, glm.Vec2{88, 268}, glm.Vec2{37, 41}, glm.Vec2{187, 275}, glm.Vec2{276, 317}, glm.Vec2{170, 330}, glm.Vec2{166, 184}, glm.Vec2{72, 449}},
			hull:   []glm.Vec2{glm.Vec2{37, 41}, glm.Vec2{72, 449}, glm.Vec2{424, 377}, glm.Vec2{428, 49}, glm.Vec2{290, 21}},
		},
	}

	for i, test := range tests {
		hull := Quickhull(test.points)
		for n, v := range hull.Vertices {
			if !v.Position.ApproxEqual(&test.hull[n]) {
				t.Errorf("[%d] %+v", i, test.points)
				t.Errorf("[%d]\n\thull %v\n\twant %v", i, hull, test.hull)
				break
			}
		}
	}
}

func TestQuickhullSupport(t *testing.T) {
	points := []glm.Vec2{{-0.1, -0.1}, {0.5, -0.1}, {1, 0}, {1.1, 1}, {0, 1}, {0.5, 0.5}, {-0.4, 0.5}, {0.4, 0.4}, {0.5, 0.4}, {0.6, 0.6}}
	hull := Quickhull(points)

	var pos []glm.Vec2
	for n := range hull.Vertices {
		pos = append(pos, hull.Vertices[n].Position)
	}
	t.Log("Vertices ", pos)
	t.Log("support dir ", qhull.SupportDirection)
	t.Log("support cache ", hull.Vertices[hull.bestSupport[0]].Position,
		hull.Vertices[hull.bestSupport[1]].Position,
		hull.Vertices[hull.bestSupport[2]].Position)

	const sep = 7
	for n := 0; n < sep; n++ {
		dir := glm.Vec2{math.Cos(float32(n) * 2.0 * math.Pi / float32(sep)), math.Sin(float32(n) * 2.0 * math.Pi / float32(sep))}
		s := hull.Support(&dir)
		t.Logf("[%d] %s %d", n, dir.String(), s)
	}
}
