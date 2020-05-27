package types_test

import "taylz.io/types"

func MakeDataSlice() types.Slice {
	return types.Slice{
		"hello world",
		"1",
		[]int{20, 19},
		types.Dict{
			"stringer": types.NewStringer("foobar stringer"),
		},
	}
}

func MakeDataDict1() types.Dict {
	return types.Dict{
		"k1":    "v1",
		"k2":    "v2",
		"k3":    "v3",
		"i1":    1,
		"i2":    2,
		"i3":    3,
		"b1":    true,
		"b2":    false,
		"map":   types.Dict{"english": "hello", "spanish": "hola"},
		"slice": types.Slice{"foo", "bar"},
		"dat1":  "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"dat2":  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"dat3":  "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
	}
}

func MakeDataDict2() types.Dict {
	// data blob from 7 elements patch payload version 1
	const data = `{"cards":[{"image":"/img/card/1.jpg","name":"Time Runner","text":"","type":"being","powers":[{"usesturn":true,"useskill":false,"id":1,"text":"Draw a Card","costs":{},"trigger":"","target":"self"}],"costs":{"1":1},"body":{"attack":0,"health":1},"id":1},{"image":"/img/card/2.jpg","name":"Ifrit","text":"","type":"being","powers":[{"trigger":"","target":"player-being","usesturn":false,"useskill":false,"id":1,"text":"Deal 1 damage to Target Being or Player","costs":{"2":1}}],"costs":{"2":1},"body":{"attack":1,"health":1},"id":2},{"powers":[],"costs":{"3":1},"body":{"attack":2,"health":2},"id":3,"image":"/img/card/4.jpg","name":"Zealot","text":"So where is the party?","type":"being"},{"name":"Vine Spirit","text":"","type":"being","powers":[{"usesturn":false,"useskill":false,"id":1,"text":"Gain 1 Attack","costs":{},"trigger":"sunrise","target":"self"}],"costs":{"4":1},"body":{"attack":1,"health":1},"id":4,"image":"/img/card/3.jpg"},{"powers":[{"text":"Target Being becomes Asleep","costs":{},"trigger":"","target":"being","usesturn":true,"useskill":false,"id":1}],"costs":{"5":1},"body":{"attack":1,"health":1},"id":5,"image":"/img/card/5.jpg","name":"Water Dancer","text":"","type":"being"},{"type":"being","powers":[{"id":1,"text":"Target Being or Player gains Health equal to my Health","costs":{},"trigger":"","target":"self","usesturn":false,"useskill":true}],"costs":{"6":1},"body":{"attack":1,"health":1},"id":6,"image":"/img/card/6.jpg","name":"Pixie","text":""},{"body":{"attack":0,"health":1},"id":7,"image":"/img/card/7.jpg","name":"Nightmare Ader","text":"","type":"being","powers":[{"costs":{},"trigger":"","target":"being","usesturn":true,"useskill":true,"id":1,"text":"Target Being deals damage to itself"}],"costs":{"7":1}},{"body":{},"id":8,"image":"/img/card/8.jpg","name":"New Element","text":"Prepare for tomorrow","type":"spell","powers":[{"text":"Create a new Element","costs":{},"trigger":"play","target":"self","usesturn":false,"useskill":false,"id":1}],"costs":{"1":1}},{"body":{},"id":9,"image":"/img/card/9.jpg","name":"Burn","text":"Boom! Roasted!","type":"spell","powers":[{"id":1,"text":"Deal 2 damage to Target Being","costs":{},"trigger":"play","target":"being","usesturn":false,"useskill":false}],"costs":{"2":1}},{"powers":[{"id":1,"text":"Target Being or Item becomes Awake","costs":{},"trigger":"play","target":"being-item","usesturn":false,"useskill":false}],"costs":{"3":1},"body":{},"id":10,"image":"img/card/10.jpg","name":"Energize","text":"Now is the time!","type":"spell"},{"text":"The best defense is a strong offense","type":"spell","powers":[{"text":"Target Being gains 1 Attack","costs":{},"trigger":"play","target":"being","usesturn":false,"useskill":false,"id":1}],"costs":{"4":1},"body":{},"id":11,"image":"/img/card/11.jpg","name":"Inspire Growth"},{"name":"Wormhole","text":"It's a trap!","type":"spell","powers":[{"useskill":false,"id":1,"text":"Add Target Being to its owners Hand, and remove it from the Present","costs":{},"trigger":"play","target":"being","usesturn":false}],"costs":{"5":1},"body":{},"id":12,"image":"/img/card/12.jpg"},{"costs":{"6":1},"body":{},"id":13,"image":"/img/card/13.jpg","name":"Grace","text":"Is that better?","type":"spell","powers":[{"trigger":"play","target":"being","usesturn":false,"useskill":false,"id":1,"text":"Target Being gains 2 Health","costs":{}}]},{"body":{},"id":14,"image":"/img/card/14.jpg","name":"Memorialize","text":"Never forget","type":"spell","powers":[{"costs":{},"trigger":"play","target":"mypast-being","usesturn":false,"useskill":false,"id":1,"text":"Add a Card to your Hand that is a copy of a Being Card from your Past"}],"costs":{"7":1}}],"packs":[{"cost":3,"image":"/img/pack/1.jpg","cards":[1,2,3,4,5,6,7],"id":1,"name":"Alpha Beings Pack (1)","size":1},{"cost":5,"image":"/img/pack/2.jpg","cards":[1,2,3,4,5,6,7],"id":2,"name":"Alpha Beings Pack (3)","size":3},{"cards":[1,2,3,4,5,6,7],"id":3,"name":"Alpha Beings Pack (5)","size":5,"cost":7,"image":"/img/pack/3.jpg"},{"cost":3,"image":"/img/pack/4.jpg","cards":[8,9,10,11,12,13,14],"id":4,"name":"Alpha Spells Pack (1)","size":1},{"id":5,"name":"Alpha Spells Pack (3)","size":3,"cost":5,"image":"/img/pack/5.jpg","cards":[8,9,10,11,12,13,14]},{"image":"/img/pack/6.jpg","cards":[8,9,10,11,12,13,14],"id":6,"name":"Alpha Spells Pack (5)","size":5,"cost":7}],"decks":[{"id":1,"name":"AllCards","cover":"/img/card/4.jpg","cards":{"8":3,"10":3,"3":3,"11":3,"1":3,"5":3,"2":3,"9":3,"14":3,"13":3,"12":3,"7":3,"4":3,"6":3}},{"id":2,"name":"WREG","cover":"/img/card/1.jpg","cards":{"3":3,"4":3,"8":3,"9":3,"10":3,"1":3,"2":3}},{"name":"GBVK","cover":"/img/card/7.jpg","cards":{"14":3,"4":3,"5":3,"6":3,"7":3,"12":3,"13":3},"id":3}],"users":25}`
	json := types.Dict{}
	_ = types.DecodeJSON(types.NewReader(data), &json)
	return json
}

func MakeDataDict3() types.Dict {
	// data blob from 7 elements patch payload version 2
	const data = `{"cards":[{"body":{"attack":0,"health":2},"id":1,"image":"/img/card/1.jpg","name":"Time Runner","text":"","type":"being","powers":[{"costs":{"1":1},"trigger":"","target":"self","usesturn":true,"useskill":false,"id":1,"text":"Draw a Card"}],"costs":{"1":1}},{"type":"being","powers":[{"costs":{"2":1},"trigger":"","target":"player-being","usesturn":true,"useskill":false,"id":1,"text":"Deal 1 damage to Target Being or Player"}],"costs":{"2":1},"body":{"attack":2,"health":1},"id":2,"image":"/img/card/2.jpg","name":"Ifrit","text":""},{"text":"So where is the party?","type":"being","powers":[],"costs":{"3":1},"body":{"attack":2,"health":2},"id":3,"image":"/img/card/4.jpg","name":"Zealot"},{"id":4,"image":"/img/card/3.jpg","name":"Vine Spirit","text":"","type":"being","powers":[{"id":1,"text":"Gain 1 Attack","costs":{},"trigger":"sunrise","target":"self","usesturn":false,"useskill":false}],"costs":{"4":1},"body":{"attack":0,"health":2}},{"text":"","type":"being","powers":[{"target":"being","usesturn":true,"useskill":false,"id":1,"text":"Target Being becomes Asleep","costs":{},"trigger":""}],"costs":{"5":1},"body":{"attack":1,"health":1},"id":5,"image":"/img/card/5.jpg","name":"Water Dancer"},{"costs":{"6":1},"body":{"attack":1,"health":2},"id":6,"image":"/img/card/6.jpg","name":"Pixie","text":"","type":"being","powers":[{"text":"Target Being or Player gains Health equal to my Health, remove me from the Present","costs":{},"trigger":"","target":"self","usesturn":true,"useskill":false,"id":1}]},{"image":"/img/card/7.jpg","name":"Nightmare Ader","text":"","type":"being","powers":[{"target":"being","usesturn":true,"useskill":true,"id":1,"text":"Target Being deals damage to itself","costs":{},"trigger":""}],"costs":{"7":1},"body":{"attack":0,"health":3},"id":7},{"image":"/img/card/8.jpg","name":"New Element","text":"Prepare for tomorrow","type":"spell","powers":[{"costs":{},"trigger":"play","target":"self","usesturn":false,"useskill":false,"id":1,"text":"Create a new Element"}],"costs":{"1":1},"body":null,"id":8},{"id":9,"image":"/img/card/9.jpg","name":"Burn","text":"Boom! Roasted!","type":"spell","powers":[{"costs":{},"trigger":"play","target":"being","usesturn":false,"useskill":false,"id":1,"text":"Deal 2 damage to Target Being"}],"costs":{"2":1},"body":null},{"costs":{"3":1},"body":null,"id":10,"image":"img/card/10.jpg","name":"Energize","text":"Now is the time!","type":"spell","powers":[{"costs":{},"trigger":"play","target":"being-item","usesturn":false,"useskill":false,"id":1,"text":"Target Being or Item becomes Awake"}]},{"type":"spell","powers":[{"trigger":"play","target":"being","usesturn":false,"useskill":false,"id":1,"text":"Target Being gains 1 Attack","costs":{}}],"costs":{"4":1},"body":null,"id":11,"image":"/img/card/11.jpg","name":"Inspire Growth","text":"The best defense is a strong offense"},{"powers":[{"usesturn":false,"useskill":false,"id":1,"text":"Target Being's Owner adds a Copy of it to their Hand, remove it from the Present","costs":{},"trigger":"play","target":"being"}],"costs":{"5":1},"body":null,"id":12,"image":"/img/card/12.jpg","name":"Wormhole","text":"It's a trap!","type":"spell"},{"costs":{"6":1},"body":null,"id":13,"image":"/img/card/13.jpg","name":"Grace","text":"Is that better?","type":"spell","powers":[{"usesturn":false,"useskill":false,"id":1,"text":"Target Being gains 3 Health","costs":{},"trigger":"play","target":"being"}]},{"id":14,"image":"/img/card/14.jpg","name":"Memorialize","text":"Never forget","type":"spell","powers":[{"text":"Add a Card to your Hand that is a copy of a Being Card from your Past","costs":{},"trigger":"play","target":"mypast-being","usesturn":false,"useskill":false,"id":1}],"costs":{"7":1},"body":null},{"name":"Banhammer","text":"shwing","type":"item","powers":[{"useskill":false,"id":1,"text":"Remove Target Being or Item from the Past","costs":{"1":1},"trigger":"","target":"being-item","usesturn":true}],"costs":{"1":2},"body":null,"id":15,"image":"/img/card/15.jpg"},{"name":"Burning Rage","text":"","type":"item","powers":[{"text":"Burning Rage does 1 damage to each enemy Player","costs":{},"trigger":"sunset","target":"self","usesturn":false,"useskill":false,"id":1}],"costs":{"2":2},"body":null,"id":16,"image":"/img/card/16.jpg"},{"image":"/img/card/17.jpg","name":"Call the Banners","text":"Join me!","type":"spell","powers":[{"target":"self","usesturn":false,"useskill":false,"id":1,"text":"Add 3 Beings to your Present, each with 2 Attack and 2 Health","costs":{},"trigger":"play"}],"costs":{"3":3},"body":null,"id":17},{"type":"item","powers":[{"id":1,"text":"Target Being gains 1 Attack","costs":{},"trigger":"sunrise","target":"","usesturn":false,"useskill":false}],"costs":{"4":1},"body":null,"id":18,"image":"/img/card/18.jpg","name":"Symbiosis","text":"Mycelliated Spiranthes"},{"text":"I see...","type":"item","powers":[{"id":1,"text":"Look at the next card in your Future, then you can choose to Shuffle your Future","costs":{},"trigger":"","target":"self","usesturn":true,"useskill":false}],"costs":{"5":2},"body":null,"id":19,"image":"/img/card/19.jpg","name":"Crystal Ball"},{"image":"/img/card/20.jpg","name":"Font of Life","text":"","type":"item","powers":[{"useskill":false,"id":1,"text":"You gain 1 Life","costs":{"6":1},"trigger":"","target":"self","usesturn":true},{"useskill":false,"id":2,"text":"Target Being gains 1 Health","costs":{"6":1},"trigger":"","target":"my-being","usesturn":true}],"costs":{"6":2},"body":null,"id":20},{"type":"item","powers":[{"costs":{"0":1,"7":2},"trigger":"","target":"mypast-being","usesturn":true,"useskill":false,"id":1,"text":"Add Target Being in your Past to your Present, set its' Health to 1, you lose 1 Life"}],"costs":{"7":3},"body":null,"id":21,"image":"/img/card/21.jpg","name":"Necromancy 101","text":""}],"packs":[{"cards":[1,2,3,4,5,6,7],"id":1,"name":"Alpha Beings Pack (1)","size":1,"cost":3,"image":"/img/pack/1.jpg"},{"id":3,"name":"Alpha Beings Pack (5)","size":5,"cost":7,"image":"/img/pack/3.jpg","cards":[1,2,3,4,5,6,7]},{"size":1,"cost":3,"image":"/img/pack/4.jpg","cards":[8,9,10,11,12,13,14],"id":4,"name":"Alpha Spells Pack (1)"},{"id":6,"name":"Alpha Spells Pack (5)","size":5,"cost":7,"image":"/img/pack/6.jpg","cards":[8,9,10,11,12,13,14]},{"size":1,"cost":3,"image":"/img/pack/7.jpg","cards":[15,16,17,18,19,20,21],"id":7,"name":"2020 Cards(1)"},{"size":5,"cost":7,"image":"/img/pack/8.jpg","cards":[15,16,17,18,19,20,21],"id":8,"name":"2020 Cards(7)"}],"decks":[{"id":1,"name":"AllCards","cover":"/img/card/0.jpg","cards":{"21":1,"9":1,"10":1,"2":1,"13":1,"19":1,"18":1,"12":1,"5":1,"14":1,"7":1,"17":1,"4":1,"6":1,"16":1,"20":1,"8":1,"11":1,"3":1,"1":1,"15":1}},{"id":2,"name":"WhiteGold","cover":"/img/card/8.jpg","cards":{"8":3,"9":3,"10":3,"15":3,"17":3,"1":3,"3":3}},{"cards":{"10":3,"11":3,"13":3,"17":3,"18":3,"3":3,"4":3},"id":3,"name":"GreenGold","cover":"/img/card/4.jpg"},{"cover":"/img/card/5.jpg","cards":{"19":3,"2":3,"5":3,"7":3,"9":3,"12":3,"16":3},"id":4,"name":"RedBlue"},{"id":5,"name":"ItemTech","cover":"/img/card/19.jpg","cards":{"3":3,"9":3,"18":3,"19":3,"20":3,"21":3,"1":3}},{"id":6,"name":"Sk8_H1","cover":"/img/card/21.jpg","cards":{"13":3,"14":3,"18":3,"20":3,"21":3,"6":3,"7":3}},{"cards":{"3":3,"9":3,"10":3,"14":3,"16":3,"17":3,"2":3},"id":7,"name":"KetchupMustard","cover":"/img/card/2.jpg"}],"users":27}`
	json := types.Dict{}
	_ = types.DecodeJSON(types.NewReader(data), &json)
	return json
}
