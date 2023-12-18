module github.com/diakovliev/2rooms-oak

go 1.21.5

require (
	github.com/oakmound/oak/v4 v4.0.0-00010101000000-000000000000
	github.com/pior/runnable v0.12.0
	github.com/rs/zerolog v1.31.0
	go.uber.org/fx v1.20.1
)

replace github.com/oakmound/oak/v4 => ../oak

require (
	dmitri.shuralyov.com/gpu/mtl v0.0.0-20201218220906-28db891af037 // indirect
	github.com/BurntSushi/xgb v0.0.0-20210121224620-deaf085860bc // indirect
	github.com/BurntSushi/xgbutil v0.0.0-20190907113008-ad855c713046 // indirect
	github.com/disintegration/gift v1.2.1 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220320163800-277f93cfa958 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/jfreymuth/pulse v0.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/oakmound/alsa v0.0.2 // indirect
	github.com/oakmound/libudev v0.2.1 // indirect
	github.com/oakmound/w32 v2.1.0+incompatible // indirect
	github.com/oov/directsound-go v0.0.0-20141101201356-e53e59c700bf // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/exp v0.0.0-20220414153411-bcd21879b8fd // indirect
	golang.org/x/exp/shiny v0.0.0-20220518171630-0b5c67f07fdf // indirect
	golang.org/x/image v0.5.0 // indirect
	golang.org/x/mobile v0.0.0-20220325161704-447654d348e3 // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.12.0 // indirect
)
