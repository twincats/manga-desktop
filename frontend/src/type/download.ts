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

export class EventPage implements EventPage {
  images: string
  index_page: number
  stat_error: string | null
  chapter: number

  static createFrom(source: any = {}) {
    return new EventPage(source)
  }

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source)
    this.images = source['images']
    this.index_page = source['index_page']
    this.stat_error = source['stat_error']
    this.chapter = source['errorchapter']
  }
}

export class EventChap implements EventChap {
  list_page: string[]
  index_chap: number
  total_chap: number
  total_page: number
  chapter: number

  static createFrom(source: any = {}) {
    return new EventChap(source)
  }

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source)
    this.list_page = source['list_page']
    this.index_chap = source['index_chap']
    this.total_chap = source['total_chap']
    this.total_page = source['total_page']
    this.chapter = source['chapter']
  }
}
