package crawler

import (
	"interphlix/lib/movies"
	"strings"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (Video *Setter)SetServers() {
	collector := colly.NewCollector()

    collector.OnHTML(".nav", func(element *colly.HTMLElement) {
        element.ForEach(".nav-item", func(index int, element *colly.HTMLElement) {
            var server movies.Server
			server.ID = primitive.NewObjectID()
            server.WatchID = element.ChildAttr("a", "data-linkid")
			if server.WatchID == "" {
				server.WatchID = element.ChildAttr("a", "data-id")
			}
            server.Name = element.ChildAttr("a", "title")
            server.Name = strings.ReplaceAll(server.Name, "Server ", "")
           	if !ServerExists(server, Video.Servers) {
				Video.Servers = append(Video.Servers, server)
			}
        })
    })
    collector.Visit(Video.Url)
	Video.SetID()
	Video.AddServer()
	Video.SetServer()
}


func (Video *Setter)SetID() {
	for index, server := range Video.Servers {
		url := "https://tinyzonetv.to/ajax/get_link/"+ server.WatchID
		data, _, err := GetRequest(url, false)
		HandleError(err)
		res, err := UnmarshalLinkResponse(data)
		HandleError(err)
        if server.Name == "Streamlare" {
			Video.Servers[index].Id = strings.ReplaceAll(res.Link, "https://streamlare.com/e/", "")
			Video.Servers[index].Url = "https://streamlare.com/v/" + Video.Servers[index].Id
		}else if server.Name == "Vidcloud"{
			Video.Servers[index].Id = strings.ReplaceAll(res.Link, "https://rabbitstream.net/embed-4/", "")
			Video.Servers[index].Id = strings.ReplaceAll(Video.Servers[index].Id, "?z=", "")
			Video.Servers[index].Url = "https://rabbitstream.net/embed/m-download/" + Video.Servers[index].Id
		}else if server.Name == "UpCloud" {
			Video.Servers[index].Id = strings.ReplaceAll(res.Link, "https://mzzcloud.life/embed-4/", "")
			Video.Servers[index].Id = strings.ReplaceAll(Video.Servers[index].Id, "?z=", "")
			Video.Servers[index].Url = "https://mzzcloud.life/embed/m-download/" + Video.Servers[index].Id
		}else {
			Video.Servers[index].Url = res.Link
		}
	}
}


func (Video *Setter)AddServer() {
    for _, server := range Video.Servers {
        if server.Name == "Vidcloud" || server.Name == "UpCloud" {
            collector := colly.NewCollector()

			collector.OnHTML(".download-list", Video.AddServers)
			collector.Visit(server.Url)
        }
    }
}


func (Video *Setter)AddServers(element *colly.HTMLElement) {
    element.ForEach(".dl-site", func(_ int, element *colly.HTMLElement) {
		var server movies.Server
		server.ID = primitive.NewObjectID()
		server.Name = element.ChildText(".site-name")
		server.Url = element.ChildAttr("a", "href")
		if ServerExists(server, Video.Servers) {
			Video.Servers = append(Video.Servers, server)
		}
	})
}


func (Video *Setter) SetServer() {
	for index := range Video.Servers {
		if Video.Servers[index].Name == "Streamlare" {
			Video.Available = true
			Video.Server = &Video.Servers[index]
		}
	}
}


func ServerExists(Server movies.Server, Servers []movies.Server) bool {
	for index := range Servers {
		if Servers[index].Name == Server.Name {
			return true
		}
	}
	return false
}