package main  

import (
"github.com/gin-gonic/gin"
"errors"
)
type tree struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Nativity string `json:"nativity"`
	Quantity int `json:"quantity"`

}
var trees=[]tree{
	{"1","mahogany","congo",20},
	{"2","ashoka","coast",50},
	{"3","mugumo","kenya",100},
}
func getTree(c *gin.Context){
	c.IndentedJSON(200,trees)
}
func treeById(c *gin.Context){
	id:=c.Param("id")
	tree,err:=getTreeById(id)
if err != nil{
	c.IndentedJSON(404, gin.H{"message": "tree not found."})
	return 
}
c.IndentedJSON(200,tree)
}
func checkoutTree(c *gin.Context){
	id, ok:=c.GetQuery("id")
    if !ok {
		c.IndentedJSON(400,gin.H{"message":"missing id parameter"})
		return
	}
	tree, err:=getTreeById(id)
	if err!=nil{
		c.IndentedJSON(404,gin.H{"message":"tree not found"})
		return
	}
	 if tree.Quantity<=0{
	c.IndentedJSON(404,gin.H{"message":"tree not available"})
	return
	 }
	 tree.Quantity-=1
	 c.IndentedJSON(200,tree)
}
func returnTree(c *gin.Context){
	id, ok:=c.GetQuery("id")
    if !ok {
		c.IndentedJSON(400,gin.H{"message":"missing id parameter"})
		return
	}
	tree, err:=getTreeById(id)
	if err!=nil{
		c.IndentedJSON(404,gin.H{"message":"tree not found"})
		return
	}
	tree.Quantity+=1
	c.IndentedJSON(200,tree)	
}
func getTreeById(id string) (*tree,error) {
for i , t := range trees {
	if t.Id==id {
		return &trees[i],nil
	}
}
return nil, errors.New("tree not found")
}
func createTree(c *gin.Context){
var newTree tree	
if err:=c.BindJSON(&newTree);err!=nil {
	return
}
trees=append(trees,newTree)
c.IndentedJSON(201,newTree)
}
func main(){
	r:=gin.Default()
	r.GET("/trees",getTree)
	r.GET("/trees/:id",treeById)
	r.POST("/trees",createTree)
	r.PATCH("/checkout",checkoutTree)
	r.PATCH("/return",returnTree)
	r.Run("localhost:8080")
}