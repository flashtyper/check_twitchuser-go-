package main

type JSON struct {
        Data []struct {
                ID           string    `json:"id"`
                UserID       string    `json:"user_id"`
                UserLogin    string    `json:"user_login"`
                UserName     string    `json:"user_name"`
                GameID       string    `json:"game_id"`
                GameName     string    `json:"game_name"`
                Type         string    `json:"type"`
                Title        string    `json:"title"`
                ViewerCount  int       `json:"viewer_count"`
                StartedAt    string `json:"started_at"`
                Language     string    `json:"language"`
                ThumbnailURL string    `json:"thumbnail_url"`
                TagIds       []string  `json:"tag_ids"`
                IsMature     bool      `json:"is_mature"`
        } `json:"data"`
        Pagination struct {
        } `json:"pagination"`
}
