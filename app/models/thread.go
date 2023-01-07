package models

import "time"

type Thread struct {
	Author  string    `json:"author" binding:"required"`  // Пользователь, создавший данную тему.
	Created time.Time `json:"created"`                    // Дата создания ветки на форуме.
	Forum   string    `json:"forum"`                      // Форум, в котором расположена данная ветка обсуждения.
	ID      int       `json:"id"`                         // Идентификатор ветки обсуждения.
	Message string    `json:"message" binding:"required"` // Описание ветки обсуждения.
	Title   string    `json:"title" binding:"required"`   // Заголовок ветки обсуждения.
	Slug    string    `json:"slug" binding:"required"`    // Человекопонятный URL
	Votes   int       `json:"votes"`                      // Кол-во голосов непосредственно за данное сообщение форума.
}

type Vote struct {
	Nickname string `json:"nickname"` // Идентификатор пользователя.
	Voice    int32  `json:"voice"`    // Отданный голос.
}