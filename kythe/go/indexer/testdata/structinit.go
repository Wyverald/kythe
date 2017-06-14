package structinit

//- @Inky defines/binding Inky
//- Inky.node/kind record
//- Inky.subkind struct
type Inky struct {
	//- @Pinky defines/binding Pinky
	//- Pinky.node/kind variable
	Pinky string

	//- @Blinky defines/binding Blinky
	//- Blinky.node/kind variable
	Blinky []byte

	//- @Sue defines/binding Sue
	//- Sue.node/kind variable
	Sue int
}

func msPacMan() {
	// Verify that named initializers ref their fields.
	a := &Inky{
		//- @Pinky ref Pinky
		//- @"\"pink\"" ref/init Pinky
		Pinky: "pink",
		//- @Blinky ref Blinky
		//- @"[]byte{255, 0, 0}" ref/init Blinky
		Blinky: []byte{255, 0, 0},
		//- @Sue ref Sue
		//- @"0x84077e" ref/init Sue
		Sue: 0x84077e,
	}
	_ = a

	// Verify that unnamed initializers ref their fields from zero-width
	// anchors.
	b := &Inky{
		//- @"a.Pinky" ref/init Pinky
		a.Pinky,

		//- @"[]byte{255, 0, 0}" ref/init Blinky
		[]byte{255, 0, 0},

		//- @"0x84077e" ref/init Sue
		0x84077e,
	}
	_ = b
}
