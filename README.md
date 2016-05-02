D3 + Python
===================

A simple visualisation of London's housing market data taken from http://data.london.gov.uk/dataset/london-borough-profiles<br>
To run locally, run londonhousing.py to generate the JSON needed for D3 to do its thing.<br>
Then, run a HTTP server locally via Python - <b>sudo python -m http.server <port> &</b> - and browse to londonhousingdata.html on your filesystem.<br>
You must have topojson installed for the conversion to work - see here for how to install. http://bost.ocks.org/mike/map/
