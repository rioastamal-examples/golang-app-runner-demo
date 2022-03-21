package main

import (
	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(200, albums)
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(200)
		c.Writer.Write([]byte(`<!DOCTYPE html>
<html>
<body style="background-color:#007d9c;color:white;font-family:sans-serif">
<h2>Let's Go with AWS App Runner!</h2>
<ul><li><a href="/albums/">/albums/</a></li></ul>
<svg height="572" width="183" xmlns="http://www.w3.org/2000/svg"><mask id="a" fill="#fff"><path d="M0 .2h9.7V572H0z" fill-rule="evenodd"/></mask><mask id="b" fill="#fff"><path d="M.4.2h9.3V572H.4z" fill-rule="evenodd"/></mask><mask id="c" fill="#fff"><path d="M0 572h116V0H0z" fill-rule="evenodd"/></mask><mask id="d" fill="#fff"><path d="M.2.7h36.1v15H.3z" fill-rule="evenodd"/></mask><g fill="none"><g fill="#fff" opacity=".6"><path d="M9.7 572H0V.7h7.5l1.7 2.4V572z" mask="url(#a)"/><path d="M9.8 572H.3V2.6L2 .2h7.3V572z" mask="url(#b)" transform="translate(106)"/><path d="M0 373h115v-9H0zm0-82h115v-9H0zm0-164h115v-9H0zm0 82h115v-9H0zM0 46h115v-9H0zm0 408h115v-9H0zm0 82h115v-9H0z" mask="url(#c)"/></g><path d="M29.1 92.9c-3.8 17.4-3.6 35-2 52.6.5 6.6 3.4 12.5 6.4 18.2 1.4 2.5 2.1 4.4.7 7-1.4 2.5-2.3 5.4.3 7.7 2.5 2.4 5.3 1.7 8.2.4 2.7-1.2 5.5-3.4 8.6-.7 2.7-.4 2.4-2.6 2.7-4.4 0-1.1-.2-2.1-1.2-2.8-3.2-.9-5.5 1.5-8.2 2.4-1.7.5-4 2.1-5 .4-1.3-2 1.3-3.4 2.6-4.7 2-2 4.9-2.7 6.5-5 .8-2.1-.2-4-2-3.6-8.2 1.8-9.5-4.9-12.3-9.8C30 141.2 30.6 131 31 121c.3-7.7 1-15.5 2.7-23.1.6-3.1 1-6.2.8-9.4.2-5.7 3.2-10.5 4.6-15.9.6-2.2 2-4.4.2-6.7-2.2 1-2 3.2-2.8 4.7-.7 1.4-.4 3.8-2.6 3.6-2-.2-2.6-2.7-2.9-4.4l-.2-.5c-.4-1.7-4.3-3.2-3.7 1.4-.2 4.4 4 8 2.3 12.6-1.2 1.2-1.5 2.6-.5 4 1 1.8.7 3.6.3 5.5z" fill="#15293d"/><g fill="#132433"><path d="M80.6 22l-7.3.8c-4.9-.8-7 3.8-10.7 5.4-3.7 2.9-7 6.3-11.5 8.1-3.8 2.9-8.2 5.3-8.7 11.5-3.8-3.5-4.8-7.5-6-11.3-1-2.9 1.5-3.3 3.3-4.2 3.1-1.4 5.8-3.1 3.4-7.2-.5-1-3.8.3-2.2-2.9 1.4-2.7-.6-3.3-3-3.2-6.3.5-7.1 2.2-6.9 8.3.3 7.4 1 14.3 5.3 20.6 1.5 2.1 2.8 4.3 5 5.8 1.7 3.4 5 5.9 5 10.1 3.2-.1 2.9-3.3 4.7-4.4.7 0 1.2.3 1.5.9.1 2.9 1 4.8 4.3 3l2.3.2c3.7.3 3.5-1.4 1.9-3.9-2.7-3.4-5.2-7.2-9.5-9-1.8-.8-3-2.6-2-4.6 1.4-2.3 3-.6 4.2.4C57 49 61.3 49 65 50.7c1.2 0 2-.6 2.5-1.6.1-4.1 0-8.2-2.3-11.8-1.9-2.7-.8-4.4 1.5-6 3-1.9 6-3.4 9.6-4.2 2.3-.5 4.8-.8 5.7-3.6 0-.9-.3-1.6-1.4-1.4M131 152.9c1.5-11.6 2-23.5 11-32.6a83 83 0 0012.4-17.6 38.2 38.2 0 003.7-29.6c-.6-2.1-1.9-4.4 1.5-6 4-1.9 6.5-5.2 4.9-10-1.6-4.6-4.3-7.2-9.5-7-1.8.4-7.1-.8-9 0-.4.3 5.3 2.2 6 2.7 1.3 1 3 1.3 4.4 2 2.2 1.3 4.1 3 3.2 5.7-.8 2.5-3 3-5.4 2.4h-.3c-2.8.2-7.1-2.2-6.2 4 .3.6.7 1.5 1.3 1.6 6.4 1.2 6.7 6.2 7.2 11.3a36 36 0 01-6.4 24.6 184 184 0 01-11.5 14.7c-8.2 10.6-10.1 29.5-11.1 42-.1 1.4 2 .1 3-2.3l.9-6z"/><path d="M154.9 50.6c-8.1.8-11.8-6.5-17.7-9.9-3.2-1.8-6-4.4-9-6.4-2-1.2-2.3-2.4-1.6-4.6 1.6-5.2-1.9-9.5-6.8-7.3-7 3.2-13 .5-19.4-.6-6.6-1.1-13-1-19.4.3l.8 1.7c1.9 2 4.2 1.3 6.3 1 8.6-.8 16.3.8 23.7 5.7 7.8 5.2 16.9 8.4 24 14.7 2.8 2.4 5.7 4.3 6.4 8.2.4 2 1.7 2.8 3.6 2.6 2.6-.9 5.6-1 7.5-3.3-2.8-3.2 2.5-.4 1.6-2.1zm-24 101.4c-2.8.6-2.2 3.4-3.4 5.1h1.3c-1.2 4-2.4 8-4.4 11.7-8.7 11.5-19.1 19.4-34.5 18-.9-.2-5.1.2-6 .7-2.8 4.7 2.1 4.6 3.6 5.5 1.5 0 3.4.2 4.5 0 3.5-.9 10.1 5.3 12.5 2 4-1 4.3-5.2 6.6-7 4.5-3.3 8.8-7.6 11.7-12.6 1.5-5.8 5.5-5 9.7-3.7 3.3.7 5-1.6 5.7-4 .7-2.2-1-5.3-2.5-5.7-6.7-1.4-6.4-5.7-5.4-10.6"/></g><path d="M53.2 62.3c-.4-1.3.9-2.6-.2-3.8-1.2 0-1.6 1.4-2.6 2-.2-1.5.5-3-.4-4.2-2.6 1-2 3.8-3.8 6.2 2-6.1-3.1-8-6.1-9.8 1.3 2.3 1.8 4.8 0 7-2.4 2.8-3.7 6.3-5.3 9.6-.5 1.1-.4 3-1.9 3.2-2 .2-1.4-2.3-2.6-3-.8 2.6-.1 5.3 2.4 5.8 2.2.3 2.6-1.8 3.5-3.5.8-1.4-4.5 8.4-4.3 16.2 2.5-1.2 3-3.7 3.8-6 2.2-6.8 6.6-14.7 9.8-21.6-.8 5.3-.6 8.3 6 9l2.5-1.2c2.5-1.7 2.4-4 1.8-7.3-1.4-.5-1.5 1.3-2.6 1.4z" fill="#15293d"/><path d="M-.3 9.8c1.6-.5 4-1.3 5.2-2.2C14.8 0 25.5-1.5 36.3 6c-1 .8-4 0-2.5 2.3.9 1.3 3.5 1.7 2.2 3.9-1.4 2.3-3.3-1-5-.1-.5-.3-.6-.2-.7-.7 1.4-2.9-1.7-4.8-4.3-5.2a27 27 0 00-14.8 2.3c-4.6 2 .8 5-.7 7.3C6 15.2 4 10.4-.3 9.8z" fill="#132433" mask="url(#d)" transform="matrix(-1 0 0 1 73 13)"/><path d="M88 193c-.3-1.5-4.4.2-2.6-3.1-.4-1.5-1.6-2-2.8-2.4-6-1.1-11.7-3.2-17.1-6-4.8-2.1-8-6.7-13-8.5-.8 1.3 0 3.2-1.5 4.3A60.2 60.2 0 0088 193" fill="#15293d"/><path d="M170.9 85.3c1.2 1.1 2.3 2.5 3.6 3.4 2.6 1.5 5.5 2.1 7.6-.6 2-2.6.4-5-1.3-7.3-2.8-4-7.3-4-11-5.8-1.2.4-3-.7-3.3 1.2-.4 1.7.9 2.3 2.2 2.7 2.9 1 6 1.5 8.2 5-.3 5.3-5-3.7-7.6-4.3-1.4-.3-2.4-4.9-3-3-.5 1.5-.6 1.8 1.3 4.4.7 1 2.6 3 3.3 4.3" fill="#333"/><path d="M101.3 192.9c-3.1 2.6-6.2-2.3-9.3 0 3.2.4 5.4 2.1 7.2 5.2 1.6 3 4.2 6 7.9 4.5 3.5-1.5 3.2-5.4 2.4-8.7-.6-2.5.6-4 1.5-5.9-2.6 1.2-4 4.8-7.4 4.4.2 1.1 1.2 2 1 3.4-1.7-.2-2.5-1.5-3.3-3M28.4 86l2.6-2.2c.4-2.6-3.5-5.4-4-7.6-.6-2.1-1.3-3.8.5-5.7-1-2.5.4-2.9 2.5-2.1-3.5-3.2-6-3-7.9 1.1-3 7 1.3 13.5 6.7 18z" fill="#15293d"/><path d="M132.5 171.5c-7-3.3-9.6 1-9.7 4 2.3-4.5 6.1-4.8 10.2-4z" fill="#132433"/><path d="M45.7 60c-6.6 8.4-9 18.4-12.7 27.8-1.2 9.8-2.6 19.7-3.5 29.6-.8 8.9-1 17.8 1.3 26.5.6 2.3.7 4.8 2.8 6.4l.4-.3.2-.5a22 22 0 0013.9 9.4c3.6 1 6 3.8 6.9 8.4 1.6 8.7 8.6 12.8 16.3 15 4 1 8 1.9 11.8 3 15.4 5 26.3-1.5 36-12.9 5-5.7 5-12.5 6-19.1 1.5-9.5 3.3-19 5.3-28.3.6-2.6 2.8-4.3 4.4-6.4 3.2-4.4 2.9-8.9-.9-11.8-4.3-3.4-8.6-3-12.2 1.1a40 40 0 00-8.7 15.3c-1.1 4.1-5 7.7-4.8 12.6.1 3 .4 5.8.1 8.8-.2 3.7-1.5 4.4-5.4 3.6-6.4-1.3-9.7-6.3-14-10.4a3.5 3.5 0 01-.7-4.1c2.4-5.1 1.7-10.6 2-16 .3-3.6 1.4-6.7 4.7-8.9 3-2 4-5.3 4.5-8.6.6-3.7-1.3-6.1-4.7-7.2-2.3-.6-4.4 0-4.2 2.8.2 4.3-1.3 6.5-5.8 5.8-3.3-.6-6.9-5.6-5.4-8.5 4.2-8.6-4.7-12.7-6-19.3-.2-.6.2-.3-.4-.3-5.8 0-9.3-5.6-14.8-6.2-2-.6-3.5 1.2-5.5 1.2-4.2-1-8.6-2-6.9-8.5z" fill="#79d4fd"/><path d="M63.4 73.7c.5 1 1.7 1.7 2.9 1.2 5.4-2 6 2 7.2 5.5.5 1.6.8 3.8 1.9 4.5 4.2 2.7 3 4.6.8 8.4-1.9 3.2 1.3 7.5 4.4 9.7 3 2.2 6.8 2 9.9-.7 1-.9 1.1-2.2.9-3.6-.2-1.4-2.9-3.9.8-4 2.9-.2 5.2 2 5.2 4.8 0 2.9-.7 6.9-3.2 8-6.7 3.2-6.8 9-6.6 14.6a30 30 0 01-2.3 13c-.9 2.1-.5 3.3 1.6 5 4.8 4 8.6 9.3 15.4 11 4.3 1 6.7-.4 6.9-4.9 0-4.4-.7-8.7.2-13.2.5-2.5.6-4.7 3.3-5.8 1.2-.5 2.3-1.4 2-3-1.1-8.3 5.8-12.2 10.2-17.2 1.9-2.2 6.1.2 8.4 2.3 2.3 2.2 3.3 5.7.7 8-7 6.5-6.7 15.1-8.6 23.3-1.5 6.5-2.9 13-3 19.4-.4 11.5-8.9 17.8-16.1 22.7-5.3 3.7-11.7 4.9-17.5 3.7-7.8-1.6-17-4-23.2-6.5S55.7 173 55 165.6c-.6-5.2-5.3-6.4-8.6-7.5A23.6 23.6 0 0132 145.7v5.6c.4 4.2 3.6 6.9 5.8 10 1.5 2.1 3.2 1.2 4.9.6 1.9-.8 4-2 3.9 1.8.9 2.8 1.3 5.9 4.6 7.1 1 .8.8 1.8.4 2.9 3 4.2 7 7.3 11.9 9 6 2.9 12 5 18.4 6.5l3 1.2c16.8 2.9 28.3-5.1 37.4-18 4-6.8 4.8-8 5.6-12.4l.1-3.6c.6-8 1.3-14.3 3-20.7 1.2-4.8 2.8-9.7 6-14.8 3.9-5 6.6-10.5 9.2-15.4 4.8-9.6 10-19.2 6.8-30.7-.8-2.8-1.8-5.1-5.2-5.7-2.4-.4-1.8-1 .2-1.8-1.7-6.9 3.6-3.7 3.6-3.7-.4-1.3-1.9-1.5-2.2-2.9-.4-1.3 0-2.8-.7-4-.7-1.2-2-1.5-3.1-2-1.1-.3-3.5.5-2.9-1.4 1.8-5-2.8-6.2-5-8-10.6-8.7-23.2-14.4-35.3-20.7-5.2-2.8-11.3-1.1-17-1-1.7 0-3.3.2-5 .3-5.7 2.5-12.3 2.7-17 7.5-2.1 2-4.2 3.1-.9 5.8 3.3 2.8 1.8 7.5 3.5 11 .7 3.2.9 6.6 2.6 9.6 1.4.4 2.3-.5 2.4-1.5.6-3.5 3-5.4 5.5-7.3C86 39.4 99.4 37.2 108 45c11 9.7 15 20.9 6.7 33A19.2 19.2 0 0193.4 86a31.8 31.8 0 01-21.8-20c-1.6-2 .4-6-3.6-6.7-2.1.9-3 3-4.8 4.3-1.6 2.6-6.7.6-7 5 2.9 1.1 5.8 2 7.2 5.1z" fill="#79d4fd"/><path d="M46.7 163.4c-3.9 2-7.2 3.4-9.5 7.3-.6 1-1.8 2-.8 3.3.8 1.2 2.7 1.6 3.9 1.1 4-1.5 6.7-2.8 10.7-4.4-1-3.8-2.5-6.3-4.3-7.3zM24.9 77.5c.4 2.4 3 3.5 5.1 5.3 1.9-.6 2.9-3.6 2-4.8-2.3-3-3.3-4-4.7-7-5.2.9-2.8 4.4-2.4 6.5z" fill="#c9b8bc"/><path d="M62.9 28.7c.1-1.3 1.3-1.6 2.2-2.2 2.8-2 2.4-4-.5-5.3-6-2.6-12.2-4.1-18.7-2.6-3.1.7-5.4 2.1-2.7 5.6 1.9 2.8-.4 5.5 0 8.3 1.3 7.6 2 8 8.4 3.8 5-.8 7.6-5 11.3-7.6" fill="#f0f4fb"/><path d="M64.8 49.6a3.8 3.8 0 00-4.1-2.2c-2.8.4-4.5-1.1-6.4-2.5-1.9-1.3-3.6-5-6.2-1.3-2 2.7-1.4 6.4 1.7 7.5 4.7 1.7 7.3 5.2 10.5 8.3 4.4 1.5 6.4.3 6.7-4.1 0-2.2 0-4.4-2.2-5.7" fill="#c9b8bc"/><path d="M50.6 37c-2.3-.7-2.8 1.5-5.4 1.2-.4-3.8-2.3-7.5-1-11.7.3-1.3-.4-2-1.6-2.5h1c1.5 4.5-1.7 5.6-4.8 6-2.5.3-3.9.5-3.6 4.2.5 5 2.3 9 5.2 13 2 2.5 4.3 2.3 4.3.3-.2-6.3 4.5-8.7 7.7-12.2" fill="#f0f4fb"/><path d="M64.5 49c1.8 8.7 2.2 10.3-3.5 10.6 1.5 3-1.4 2.7-3 2.3 1 4.3 3.9 2.6 6.5 1.6 2.5-.3 2.7-2.1 3.5-4 .5-1.3 1-1.5.7-2.2.6-3.5-.9-7-3-9.8-.5 1 .3 2-1 2.1z" fill="#282b3a"/><path d="M65 63.7c-2-.8-5 .9-6.3-2.1L57 61c-1.7 2.4-1 6-4.1 7.6 1.7 1 3.2-.4 4.8-.5" fill="#67d3fd"/><path d="M59.5 44.4c1.4.8 3 1 3.4-.8.6-2.6-1-4.2-3.3-5-1.6-.7-4-1.2-4.5 1-.8 3.2 2.8 3.3 4.4 4.8" fill="#6b6e7f"/><path d="M50.9 57l-.2.5c-.3 2.6-3.7 6 0 7.4 2.8 1 1.8-3.9 3.3-5.6-.4-1.6-1.4-2.4-3.1-2.3" fill="#fff"/><path d="M142.7 60c-2.1 5.9-1 7.8 5.4 9.4 1.6.4 3.1 1.1 3.2 2.6.3 5.1 3.1 9.8 1 15.4-3.5 9-6.7 17.8-13.7 25-.8.8-3.6 3.8 0 6.4 4.7-4.3 7.5-9.4 11.2-14.4 7.1-9.5 7.5-20.8 5.3-31.5-1-4.9-5.6-4.3-8.3-6.5-3.9-.7.5-6.3-4.1-6.4zm9.4 3.8c4 2 8.1 1.8 9.6-2.6 1.4-4.5-2.1-6.7-6.4-7.6L153 53c-2.5.2-4.9.7-7.1 2 .3 1.4 1.7 1.7 2.8 2.2 1.4 1.8 3.8 1.2 6.4 2.6-2.8.4-5.3 0-6.8 1.7.6 1.6 2.4 1.6 3.7 2.3zm-28.6-35.3c.6-.3.7-1.1 0-1.6-2-1.5-4-.9-6.5.2 2.3 1.8 4.3 2.5 6.5 1.4z" fill="#79d4fd"/><path d="M129.2 164.1c-.7 2.8.6 3.7 2.3 4.4 1.7.7 3.3 1 4.2-1.2 1.2-2.9-1.2-2.8-2.5-3.4-1.4-.6-3-1.8-4 .2" fill="#c9b8bc"/><path d="M101.2 192.4c0 .8-.4 1.7 0 2.2a5.7 5.7 0 004.3 3.4c1 .1 1.3-1.2 1.4-2.1.4-2.6-1.2-3.4-3.2-3.9z" fill="#b5a4a8"/><path d="M83 189l-18-7a40 40 0 0018 7" fill="#67d3fd"/><path d="M70.7 66A31.6 31.6 0 0093 87.2c15.1 4.5 31.5-11 27.2-27-4.8-18-21.2-26.7-39.2-17.5A14 14 0 0076 48c-3.7 3.1-5.5 9.1-5.3 18z" fill="#132433"/><path d="M91.8 131.7c-2.9 1.4-1.6 4-.8 5.6 2 4.6 5.4 8.1 10.5 9.5 3.5 1 6.3-1.4 5.3-5-1.3-4.2-3.3-8.2-5-12.4-4.1-1.4-3.7 1.3-3.3 3.6.3 2.5 2.5 4.5 2.2 7.6-4.7-1.5-4.7-7.3-8.9-9" fill="#19364e"/><path d="M92 131.6c2.4 3.6 3.6 6.4 7 9.2 1.1.8 2.1 1.8 3.4.7 1-1 .5-2.2.1-3.2-.9-2.6-2.7-7-3.4-8.3-.2-.7-.3-.7-1-.7-.8 0-2.9.8-6.1 2.3z" fill="#c9b8bc"/><path d="M148.3 61.6h3.2c2.4-.2 6.2 1.9 6.5-1.8.2-2.6-3.9-3.4-6.5-3.8-1 0-1.9.2-2.8.4-1 1.6-.7 3.4-.4 5.2" fill="#132433"/><path d="M70.7 66A27.5 27.5 0 0176 48c-4.7 1.5-4.8 6.3-7.3 9.3l-.7 2.1c3.2 1.3.3 5 2.7 6.6" fill="#79d4fd"/><path d="M76.5 54.6c-5 13.8 2.8 22.7 15.7 29 9.5 4.5 21.4-2.2 25.1-12.8 3-8.7-5-24.3-14.2-27.6-10.4-3.8-23.1 1.7-26.6 11.4" fill="#f9fafa"/><path d="M84.4 54.8c-.8 3.1-.8 6.6 3.4 8.5 4 1.8 8.4 0 9.9-4 1.2-3.5-1.2-9-4.5-10-3.2-1-7.2 1.2-8.8 5.5" fill="#0e0c16"/></g></svg>
</body>
</html>`))
	})

	router.GET("/albums", getAlbums)
	router.Run()
}
