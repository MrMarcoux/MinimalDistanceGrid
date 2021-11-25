module main

go 1.15

require (
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	matchanalysis/distancegrid v0.0.0-00010101000000-000000000000
)

replace matchanalysis/distancegrid => ./distancegrid
