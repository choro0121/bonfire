package model

import (
    "time"

    "gorm.io/gorm"
)

type (
    User struct {
        UserId          int             `gorm:"primary_key; auto_increment"`
        Username        string          `gorm:"unique"`
        Mail            string          `gorm:"unique"`
        Bio             string
        Company         string
        Location        string
        WebSite         string
        CreatedAt       time.Time       `gorm:"autoCreateTime"`
        UpdatedAt       time.Time       `gorm:"autoUpdateTime"`
        Deleted         gorm.DeletedAt
    }

    Follow struct {
        UserId          int             `gorm:"not null";`
        FollowUserId    int             `gorm:"not null";`
        CreateAt        time.Time       `gorm:"autoCreateTime"`
    }

    OAuth struct {
        UserId          int             `gorm:"primary_key;"`
        Provider        string          `gorm:"not null"`
        ExternId        string          `gorm:"not null"`
    }

    Password struct {
        UserId          int             `gorm:"primary_key;"`
        Hash            string
        UpdateAt        time.Time       `gorm:"autoUpdateTime"`
    }
)

func CreateUserWithPassword(user *User, password *Password) (*User, *Password, error) {
    // create user
    err := db.Create(user).Error
    if err != nil {
        return nil, nil, err
    }

    // write user_id
    password.UserId = user.UserId

    // create password
    err = db.Create(password).Error
    if err != nil {
        return nil, nil, err
    }

    return user, password, nil
}

func CreateUserWithOAuth(user *User, oauth *OAuth) (*User, *OAuth, error) {
    // create user
    err := db.Create(user).Error
    if err != nil {
        return nil, nil, err
    }

    // write user_id
    oauth.UserId = user.UserId

    // create oauth
    err = db.Create(oauth).Error
    if err != nil {
        return nil, nil, err
    }

    return user, oauth, nil
}

func GetUser(user *User) (*User, error) {
    err := db.Model(user).Where(user).First(user).Error

    return user, err
}

func GetUsers(offset int, user *User) ([]User, error) {
    var users []User

    err := db.Model(user).Where(user).Offset(offset).Limit(20).Find(&users).Error

    return users, err
}

func UpdateUser(user *User) (*User, error) {
    err := db.Model(user).Where("user_id = ?", user.UserId).Updates(user).Error

    return user, err
}

func DeleteUser(user *User) (*User, error) {
    err := db.Model(user).Where("user_id = ?", user.UserId).Delete(&User{}).Error

    return user, err
}


func GetOAuth(oauth *OAuth) (*OAuth, error) {
    err := db.Model(oauth).Where(oauth).First(oauth).Error

    return oauth, err
}


func GetPassword(password *Password) (*Password, error) {
    err := db.Model(password).Where(password).First(password).Error

    return password, err
}

func UpdatePassword(password *Password) (*Password, error) {
    err := db.Model(password).Where("user_id = ?", password.UserId).Updates(password).Error

    return password, err
}


func CreateFollow(follow *Follow) (*Follow, error) {
    _, err := CreateMany2ManyWidhoutDuplicate(follow)
    return follow, err
}

func GetFollows(offset int, follow *Follow) ([]Follow, error) {
    var follows []Follow

    err := db.Model(follow).Where(follow).Offset(offset).Limit(20).Find(&follows).Error

    return follows, err
}

func DeleteFollow(follow *Follow) (*Follow, error) {
    err := db.Model(follow).Where("user_id = ? AND follow_user_id = ?", follow.UserId, follow.FollowUserId).Delete(&Follow{}).Error

    return follow, err
}
