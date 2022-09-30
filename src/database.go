package src

import (
	"context"
	"fmt"
	"log"
	"sync"

	"cloud.google.com/go/firestore"
)

// SyncGroup to ensure multiple clients are not generated
var syncOnce *sync.Once = new(sync.Once)
var syncOnce2 *sync.Once = new(sync.Once)

// Client to store the datastore Client.
var client *firestore.Client = &firestore.Client{}

func get_client() *firestore.Client {
	var err error
	clt := client // Client Derives from the global client
	syncOnce.Do(func() {
		ctx := context.Background()
		if clt, err = firestore.NewClient(ctx, "leafy-ai"); err != nil {
			log.Fatal("Could Not Connect to Database: ", err)
		}
	})
	fmt.Println("New Client: ", clt)
	return clt // Return the client
}

type BlogManager struct {
	col *firestore.CollectionRef
}

// Function to create a new BlogManager
func (b *BlogManager) new() {
	syncOnce2.Do(
		func() {
			b.col = get_client().Collection("Blogs")
		},
	)
	fmt.Println("Blog Manager Created: ", b.col)
}

// Function to Get All Blogs
func (b *BlogManager) GetAllBlogs() (*[]Blog, error) {
	blogs := &[]Blog{}
	var err error = nil
	docSnaps, err := b.col.Documents(context.Background()).GetAll()
	getMultiple(docSnaps, blogs)
	if err != nil {
		log.Println("Could Not Get All Blogs: ", err)
	}

	return blogs, err
}

func (b *BlogManager) SearchBlogs(key, value, operator string) (*[]Blog, error) {
	blogs := &[]Blog{}
	var err error = nil
	docSnaps, err := b.col.Where(key, operator, value).Documents(context.Background()).GetAll()
	getMultiple(docSnaps, blogs)
	if err != nil {
		log.Println("Could Not Get All Blogs: ", err)
	}

	return blogs, err
}

// Create Blog through Blogmanager
func (b *BlogManager) CreateBlog(blog *Blog) error {
	err := blog.Create(b)
	return err
}

type User struct {
	Username string                 `json:"username" firestore:"username"`
	ref      *firestore.DocumentRef `json:"-" firestore:",omitempty"`
}

type Blog struct {
	Createdby User                   `json:"createdby" firestore:"createdby"`
	ref       *firestore.DocumentRef `json:"" firestore:",omitempty"`
	Title     string                 `json:"title" firestore:"title"`
	Body      string                 `json:"body" firestore:"body"`
	Upvotes   int                    `json:"upvotes" firestore:"upvotes"`
	DownVotes int                    `json:"downvotes" firestore:"downvotes"`
	Creator   string                 `json:"creator" firestore:"creator"`
	deleted   bool                   `json:"deleted" firestore:"deleted"`
}

// Parse the data into the blogs
func getMultiple(docSnaps []*firestore.DocumentSnapshot, blogs *[]Blog) error {
	for _, snaps := range docSnaps {
		blog := &Blog{}
		if err := blog.Get(snaps); err != nil {
			return err
		}
		*blogs = append(*blogs, *blog)
	}
	return nil
}

// Function to Get a Blog
func (b *Blog) Get(docSnap *firestore.DocumentSnapshot) error {
	if err := docSnap.DataTo(b); err != nil {
		log.Println("Could Not Convert Data to Blog: ", err)
		return err
	}
	return nil
}

// Create Creates the Blog
func (b *Blog) Create(manager *BlogManager) error {
	var err error = nil

	b.ref, _, err = manager.col.Add(context.Background(), b)
	if err != nil {
		log.Println("Could Not Create Blog: ", err)
	}
	return err
}

// Update Updates the Blog
func (b *Blog) Update() error {
	var err error = nil
	_, err = b.ref.Set(context.Background(), b)
	if err != nil {
		log.Println("Could Not Update Blog: ", err)
	}
	return err
}

// Soft deletes the Blog
func (b *Blog) Delete() error {
	b.deleted = true
	return b.Update()
}

// Upvotes the Blog
func (b *Blog) Upvote() error {
	b.Upvotes++
	return b.Update()
}

// Downvotes the Blog
func (b *Blog) Downvote() error {
	b.DownVotes++
	return b.Update()
}

func (u *User) getById(id string) error {
	var err error = nil
	u.ref = get_client().Collection("Users").Doc(id)
	docSnap, err := u.ref.Get(context.Background())
	if err != nil {
		log.Println("Could Not Get User: ", err)
	}
	docSnap.DataTo(u)
	return err
}
