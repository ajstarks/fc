# Elections

Display US Presidential election results on a state-based grid, with encoded populations.

![elections](elections.png)

## Usage
	
	elections [options] file...

Specify at least one data file (format below).  Use the forward and back arrows to navigate.

To see election results from 1920 to 2020 (included in this repository):

	elections 1920.d 1924.d 1928.d 1932.d 1936.d 1940.d 1944.d 1948.d 1952.d 1956.d 1960.d 1964.d 1968.d 1972.d 1976.d 1980.d 1984.d 1988.d 1992.d 1996.d 2000.d 2004.d 2008.d 2012.d 2016.d 2020.d

To see the years FDR was president:

	elections 1932.d 1936.d 1940.d 1944.d

Notable landslides:

	elections 1936.d 1964.d 1972.d 1984.d

Close elections

	elections 1960.d 1968.d 2000.d

## Options
```
	-bgcolor string 
		background color (default "black")
	-colsize float
		column size (canvas %) (default 7)
	-height int
		canvas height (default 900)
	-left float
		map left value (canvas %) (default 7)
	-rowsize float
		rowsize (canvas %) (default 9)
	-textcolor string
		text color (default "white")
	-top float
		map top value (canvas %) (default 75)
	-width int
		canvas width (default 1200)
```

## Input File format

The input files have five tab-separated fields

* State, 
* Grid Column, 
* Grid Row, 
* Party (r=Republican, d=Democrat, i=Independent), 
* State population.

The line beginning with '#' specifies the election year, Democratic candidate, Republican Candidate, and optional third-party candidate.

For example:

	# 2020 Biden Trump
	AL      5       7       r       4903185
	AK      1       0       r       731545
	AZ      5       2       d       7278717
	AR      5       5       r       3017804
	CA      5       1       d       39512223
	...
	WA      2       1       d       7614893
	WV      4       8       r       1792147
	WI      2       6       d       5822434
	WY      3       3       r       578759

	