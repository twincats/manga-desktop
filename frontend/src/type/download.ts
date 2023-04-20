export interface EventChap {
  list_page: string[]
  index_chap: number
  total_chap: number
  total_page: number
  chapter: number
}

export interface EventPage {
  images: string
  index_page: number
  stat_error: string | null
  chapter: number
}
