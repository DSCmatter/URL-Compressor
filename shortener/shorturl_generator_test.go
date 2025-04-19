package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.amazon.in/Cosmic-Byte-Firefly-Mechanical-Swappable/dp/B0BTD484S1/?_encoding=UTF8&pd_rd_w=xhsyN&content-id=amzn1.sym.2c8720d9-6d29-4ec6-934b-42b530789226&pf_rd_p=2c8720d9-6d29-4ec6-934b-42b530789226&pf_rd_r=RRW1M2S2C41B3V7G6KXV&pd_rd_wg=H7632&pd_rd_r=921ac202-be7d-41f9-9358-199bbf2d1747&ref_=pd_hp_d_atf_dealz_cs"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://www.amazon.in/Redragon-K617-Keyboard-Mechanical-Supported/dp/B09BVCVTBC?pd_rd_w=D4ne3&content-id=amzn1.sym.b5387062-d66f-4264-a2b3-7498b12700ed&pf_rd_p=b5387062-d66f-4264-a2b3-7498b12700ed&pf_rd_r=ZP7S4CFDM9J2YEGENZ5H&pd_rd_wg=T2cES&pd_rd_r=b4570cfa-44a0-4b56-8b76-600787cb8cb3&pd_rd_i=B09BVCVTBC&th=1"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	initialLink_3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink_3 := GenerateShortLink(initialLink_3, UserId)

	assert.Equal(t, shortLink_1, "jTa4L57P")
	assert.Equal(t, shortLink_2, "d66yfx7N")
	assert.Equal(t, shortLink_3, "dhZTayYQ")
}
