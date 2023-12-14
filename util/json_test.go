package util

import (
	"testing"
)

func TestJson(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{`{
			"name": "Sneaker #172248099",
			"description": "NFT Sneaker, use it in STEPN to move2earn",
			"image": "https://arweave.net/SL46Z2ZmVLpm95VelmN2zyzIALOBDXEaLnmAb0amnC4",
			"external_url": "https://stepn.com",
			"attributes": [
			  {
				"trait_type": "Sneaker type",
				"value": "Jogger"
			  },
			  {
				"trait_type": "Sneaker quality",
				"value": "Common"
			  },
			  {
				"trait_type": "Level",
				"value": "0"
			  },
			  {
				"trait_type": "Optimal Speed",
				"value": "4.0-10.0km/h"
			  },
			  {
				"trait_type": "Shoe-minting Count",
				"value": "0/7"
			  },
			  {
				"trait_type": "Efficiency",
				"value": "7.6"
			  },
			  {
				"trait_type": "Luck",
				"value": "5.6"
			  },
			  {
				"trait_type": "Comfortability",
				"value": "2.1"
			  },
			  {
				"trait_type": "Resilience",
				"value": "8.1"
			  },
			  {
				"trait_type": "Durability",
				"value": "100/100"
			  },
			  {
				"trait_type": "Socket 1",
				"value": "Efficiency/unknown/empty"
			  },
			  {
				"trait_type": "Socket 2",
				"value": "Durability/unknown/empty"
			  },
			  {
				"trait_type": "Socket 3",
				"value": "Durability/unknown/empty"
			  },
			  {
				"trait_type": "Socket 4",
				"value": "Efficiency/unknown/empty"
			  },
			  {
				"trait_type": "Badge",
				"value": "None"
			  }]
		  }`, `{"name":"Sneaker #172248099","description":"NFT Sneaker, use it in STEPN to move2earn","image":"https://arweave.net/SL46Z2ZmVLpm95VelmN2zyzIALOBDXEaLnmAb0amnC4","external_url":"https://stepn.com","attributes":[{"trait_type":"Sneaker type","value":"Jogger"},{"trait_type":"Sneaker quality","value":"Common"},{"trait_type":"Level","value":"0"},{"trait_type":"Optimal Speed","value":"4.0-10.0km/h"},{"trait_type":"Shoe-minting Count","value":"0/7"},{"trait_type":"Efficiency","value":"7.6"},{"trait_type":"Luck","value":"5.6"},{"trait_type":"Comfortability","value":"2.1"},{"trait_type":"Resilience","value":"8.1"},{"trait_type":"Durability","value":"100/100"},{"trait_type":"Socket 1","value":"Efficiency/unknown/empty"},{"trait_type":"Socket 2","value":"Durability/unknown/empty"},{"trait_type":"Socket 3","value":"Durability/unknown/empty"},{"trait_type":"Socket 4","value":"Efficiency/unknown/empty"},{"trait_type":"Badge","value":"None"}]}`},
		{``, ``},
	}

	for _, c := range cases {
		got, _ := CompressJson(c.input)
		if got != c.want {
			t.Errorf("结果错误，got=%q, want=%q", got, c.want)
		}
	}
}
