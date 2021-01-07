package model

import (
    "time"
    "gorm.io/gorm"
)

type (
    Post struct {
        PostId          int             `gorm:"primary_key; auto_increment" json:"post_id"`
        UserId          int             `gorm:"not null" json:"user_id"`
        Title           string          `gorm:"not null" json:"title"`
        Description     string          `json:"description"`
        Code            string          `gorm:"not null" json:"code"`
        CreatedAt       time.Time       `gorm:"autoCreateTime"`
        UpdatedAt       time.Time       `gorm:"autoUpdateTime"`
        Deleted         gorm.DeletedAt
    }

    Bookmark struct {
        UserId          int             `gorm:"not null"`
        PostId          int             `gorm:"not null"`
        CreateAt        time.Time       `gorm:"autoCreateTime"`
    }

    Good struct {
        UserId          int             `gorm:"not null"`
        PostId          int             `gorm:"not null"`
        CreateAt        time.Time       `gorm:"autoCreateTime"`
    }

    Tag struct {
        TagId           int             `gorm:"primary_key; auto_increment"`
        PostId          int             `gorm:"not null"`
        Name            string          `gorm:"not null"`
    }

    Comment struct {
        CommentId       int             `gorm:"primary_key; auto_increment" json:"comment_id"`
        PostId          int             `gorm:"not null"`
        UserId          int             `gorm:"not null"`
        Body            string          `gorm:"not null" json:"body"`
        CreateAt        time.Time       `gorm:"autoCreateTime"`
        UpdateAt        time.Time       `gorm:"autoUpdateTime"`
        Deleted         gorm.DeletedAt
    }
)

func CreatePost(post *Post) (*Post, error) {
    err := db.Create(post).Error

    return post, err
}

func GetPost(post *Post) (*Post, error) {
    err := db.Model(post).Where(post).First(post).Error

    return post, err
}

func GetPosts(offset int, post *Post) ([]Post, error) {
    var posts []Post

    err := db.Model(post).Where(post).Offset(offset).Limit(20).Find(&posts).Error

    return posts, err
}

func UpdatePost(post *Post) (*Post, error) {
    err := db.Model(post).Where("post_id = ?", post.PostId).Updates(post).Error

    return post, err
}

func DeletePost(post *Post) (*Post, error) {
    err := db.Model(post).Where("post_id = ?", post.PostId).Delete(&Post{}).Error

    return post, err
}


func CreateBookmark(bookmark *Bookmark) (*Bookmark, error) {
    _, err := CreateMany2ManyWidhoutDuplicate(bookmark)
    return bookmark, err
}

func GetBookmarks(offset int, bookmark *Bookmark) ([]Bookmark, error) {
    var bookmarks []Bookmark

    err := db.Model(bookmark).Where(bookmark).Offset(offset).Limit(20).Find(&bookmarks).Error

    return bookmarks, err
}

func DeleteBookmark(bookmark *Bookmark) (*Bookmark, error) {
    err := db.Model(bookmark).Where("user_id = ? AND post_id = ?", bookmark.UserId, bookmark.PostId).Delete(&Bookmark{}).Error

    return bookmark, err
}


func CreateGood(good *Good) (*Good, error) {
    _, err := CreateMany2ManyWidhoutDuplicate(good)
    return good, err
}

func GetGoods(offset int, good *Good) ([]Good, error) {
    var goods []Good

    err := db.Model(good).Where(good).Offset(offset).Limit(20).Find(&goods).Error

    return goods, err
}

func DeleteGood(good *Good) (*Good, error) {
    err := db.Model(good).Where("user_id = ? AND post_id = ?", good.UserId, good.PostId).Delete(&Good{}).Error

    return good, err
}


func CreateTag(tag *Tag) (*Tag, error) {
    err := db.Create(tag).Error

    return tag, err
}

func GetTags(tag *Tag) ([]Tag, error) {
    var tags []Tag

    err := db.Model(tag).Where(tag).Find(&tags).Error

    return tags, err
}

func UpdateTag(tag *Tag) (*Tag, error) {
    err := db.Model(tag).Where("tag_id = ?", tag.TagId).Updates(tag).Error

    return tag, err
}

func DeleteTag(tag *Tag) (*Tag, error) {
    err := db.Model(tag).Where("tag_id = ?", tag.TagId).Delete(&Tag{}).Error

    return tag, err
}


func CreateComment(comment *Comment) (*Comment, error) {
    err := db.Create(comment).Error

    return comment, err
}

func GetComments(offset int, comment *Comment) ([]Comment, error) {
    var comments []Comment

    err := db.Model(comment).Where(comment).Offset(offset).Limit(20).Find(&comments).Error

    return comments, err
}

func UpdateComment(comment *Comment) (*Comment, error) {
    err := db.Model(comment).Where("comment_id = ?", comment.CommentId).Updates(comment).Error

    return comment, err
}

func DeleteComment(comment *Comment) (*Comment, error) {
    err := db.Model(comment).Where("comment_id = ?", comment.CommentId).Delete(&Comment{}).Error

    return comment, err
}
