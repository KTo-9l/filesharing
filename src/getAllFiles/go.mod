module filesharing/getAllFiles

go 1.23.6

replace filesharing/mongoApi => ../mongoApi

replace github.com/okonma-violet/services => /home/andrey/go/pkg/mod/github.com/okonma-violet/services@v0.0.0-20250303113135-3d036f9a188f

replace github.com/big-larry/suckutils => /home/andrey/go/pkg/mod/github.com/big-larry/suckutils@v0.0.0-20231029230114-645d5d858694

replace github.com/big-larry/mgo => ../mgo@v1.0.0

require filesharing/mongoApi v0.0.0-00010101000000-000000000000

require (
	github.com/big-larry/mgo v1.0.0
	github.com/big-larry/suckhttp v0.0.0-20220801042759-6d5e1d8b45e4
	github.com/okonma-violet/services v0.0.0-00010101000000-000000000000
)

require (
	github.com/big-larry/suckutils v0.0.0-20231029230114-645d5d858694 // indirect
	github.com/okonma-violet/confdecoder v0.0.0-20230926094403-7e3eab7eff29 // indirect
	github.com/okonma-violet/dynamicworkerspool v0.0.0-20240317115954-810a6361f715 // indirect
)
