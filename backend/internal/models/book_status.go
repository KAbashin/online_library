package models

const (
	StatusBookVisible    = "visible"    // Отображается пользователям
	StatusBookArchived   = "archived"   // Устаревшее издание, удалено из выдачи, но не удалено из БД
	StatusBookQuarantine = "quarantine" // Временная блокировка на публикацию
	StatusBookPrivate    = "private"    // Непубличный контент
)
