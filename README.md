# rss
<p align="left">
  <a href="https://github.com/toorosan/rss"><img alt="RSS status" src="https://github.com/toorosan/rss/workflows/Go/badge.svg"></a>
</p>

## RSS reader application.
Allows to gather RSS items from several feeds to the single JSON file.
Available to be run using command line arguments and with JSON configuration file.

#### Installation:
```
# go install -race -v ./...
```

#### Command line arguments:
- `-version` - print version and exit.
- `-silent` - do not print the results, only errors.
- `-config` - path to configuration file, overrides values passed directly as arguments.
- `-feed` - URL(s) of RSS feed(s) to read, allows to pass multiple values (for more details see first usage example). 
- `-output` - path to output json file where all RSS parsing results will be stored.

#### Usage example:
```
#  ~/go/bin/app -feed https://www.feedforall.com/sample.xml -feed https://www.feedforall.com/sample-feed.xml
  "RSS Reader" initiated
  output file was not passed, using default one: "output.json"
  FeedForAll Sample Feed: 2004-10-19 11:09:11 -0400 -0400 - "RSS Solutions for Restaurants" - http://www.feedforall.com/restaurant.htm
  FeedForAll Sample Feed: 2004-10-19 11:09:09 -0400 -0400 - "RSS Solutions for Schools and Colleges" - http://www.feedforall.com/schools.htm
  FeedForAll Sample Feed: 2004-10-19 11:09:07 -0400 -0400 - "RSS Solutions for Computer Service Companies" - http://www.feedforall.com/computer-service.htm
  FeedForAll Sample Feed: 2004-10-19 11:09:05 -0400 -0400 - "RSS Solutions for Governments" - http://www.feedforall.com/government.htm
  FeedForAll Sample Feed: 2004-10-19 11:09:03 -0400 -0400 - "RSS Solutions for Politicians" - http://www.feedforall.com/politics.htm
  FeedForAll Sample Feed: 2004-10-19 11:09:01 -0400 -0400 - "RSS Solutions for Meteorologists" - http://www.feedforall.com/weather.htm
  FeedForAll Sample Feed: 2004-10-19 11:08:59 -0400 -0400 - "RSS Solutions for Realtors & Real Estate Firms" - http://www.feedforall.com/real-estate.htm
  FeedForAll Sample Feed: 2004-10-19 11:08:57 -0400 -0400 - "RSS Solutions for Banks / Mortgage Companies" - http://www.feedforall.com/banks.htm
  FeedForAll Sample Feed: 2004-10-19 11:08:56 -0400 -0400 - "RSS Solutions for Law Enforcement" - http://www.feedforall.com/law-enforcement.htm
  Sample Feed - Favorite RSS Related Software & Resources: 2004-10-26 14:01:01 -0500 -0500 - "RSS Resources" - http://www.feedforall.com
  Sample Feed - Favorite RSS Related Software & Resources: 2004-10-26 14:03:25 -0500 -0500 - "Recommended Desktop Feed Reader Software" - http://www.feedforall.com/feedforall-partners.htm
  Sample Feed - Favorite RSS Related Software & Resources: 2004-10-26 14:06:44 -0500 -0500 - "Recommended Web Based Feed Reader Software" - http://www.feedforall.com/feedforall-partners.htm
  "RSS Reader" is stopped
```

#### Usage example (config file):
```
# ~/go/bin/app -config app/config.json
  "RSS Reader" initiated
  FeedForAll Sample Feed: 2004-10-19 11:09:11 -0400 -0400 - "RSS Solutions for Restaurants" - http://www.feedforall.com/restaurant.htm
  FeedForAll Sample Feed: 2004-10-19 11:09:09 -0400 -0400 - "RSS Solutions for Schools and Colleges" - http://www.feedforall.com/schools.htm
  FeedForAll Sample Feed: 2004-10-19 11:09:07 -0400 -0400 - "RSS Solutions for Computer Service Companies" - http://www.feedforall.com/computer-service.htm
  FeedForAll Sample Feed: 2004-10-19 11:09:05 -0400 -0400 - "RSS Solutions for Governments" - http://www.feedforall.com/government.htm
  FeedForAll Sample Feed: 2004-10-19 11:09:03 -0400 -0400 - "RSS Solutions for Politicians" - http://www.feedforall.com/politics.htm
  FeedForAll Sample Feed: 2004-10-19 11:09:01 -0400 -0400 - "RSS Solutions for Meteorologists" - http://www.feedforall.com/weather.htm
  FeedForAll Sample Feed: 2004-10-19 11:08:59 -0400 -0400 - "RSS Solutions for Realtors & Real Estate Firms" - http://www.feedforall.com/real-estate.htm
  FeedForAll Sample Feed: 2004-10-19 11:08:57 -0400 -0400 - "RSS Solutions for Banks / Mortgage Companies" - http://www.feedforall.com/banks.htm
  FeedForAll Sample Feed: 2004-10-19 11:08:56 -0400 -0400 - "RSS Solutions for Law Enforcement" - http://www.feedforall.com/law-enforcement.htm
  Sample Feed - Favorite RSS Related Software & Resources: 2004-10-26 14:01:01 -0500 -0500 - "RSS Resources" - http://www.feedforall.com
  Sample Feed - Favorite RSS Related Software & Resources: 2004-10-26 14:03:25 -0500 -0500 - "Recommended Desktop Feed Reader Software" - http://www.feedforall.com/feedforall-partners.htm
  Sample Feed - Favorite RSS Related Software & Resources: 2004-10-26 14:06:44 -0500 -0500 - "Recommended Web Based Feed Reader Software" - http://www.feedforall.com/feedforall-partners.htm
  "RSS Reader" is stopped
```

