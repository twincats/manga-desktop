import { MangaParser } from './parser'
import { manga, types, download } from '@wails/go/models'
import {
  GetChapter,
  GetPage,
  DownloadJS,
  SaveMangaCover,
  CheckChapterDB,
} from '@wails/go/download/Download'
import type { EventChap } from '@/type/download'
import { EventsEmit } from '@wails/runtime/runtime'

export class Download {
  #p: MangaParser
  constructor() {
    this.#p = new MangaParser()
  }
  GetChap(u: string, o: types.Option, s: manga.Server): Promise<types.Chapter> {
    if (s.id == 1) {
      return GetChapter(o)
    } else {
      return new Promise((res, rej) => {
        this.#p
          .getManga(u, '')
          .then(r => {
            r.server_name = s.name
            CheckChapterDB(r)
              .then(rr => {
                res(rr)
              })
              .catch(() => {
                res(r)
              })
          })
          .catch(e => {
            rej(e)
          })
      })
    }
  }

  GetPage(c: types.Chapter, s: manga.Server): Promise<download.PageReport> {
    if (s.id == 1) {
      return GetPage(c)
    } else {
      return new Promise(async (res, rej) => {
        const p = new download.PageReport()
        var manga_id = c.manga_id
        // save mangaDB
        if (manga_id == 0) {
          // save manga and cover
          const ret = await SaveMangaCover(c)
            .then(id => {
              if (id != 0) {
                manga_id = id
              }
            })
            .then(e => {
              rej(e)
              return true
            })
          // return if being rejected
          if (ret) {
            return
          }
        }
        p.manga = c.manga

        // loop download chapter
        for (let i = 0; i < c.chapter.length; i++) {
          const cc = c.chapter[i]
          //report failed chap
          const fc = new download.FiledChapReport()
          fc.chap_id = cc.id
          fc.chapter = cc.chapter

          this.#p
            .getPages(cc.id, '')
            .then(r => {
              // do event
              const evc: EventChap = {
                chapter: parseFloat(cc.chapter),
                index_chap: i,
                total_chap: c.chapter.length,
                total_page: r.length,
                list_page: r,
              }

              EventsEmit('dl_eventchap', evc)

              // do download
              const paramJs = new download.ParamJS()
              paramJs.manga_id = c.manga_id
              paramJs.manga = c.manga
              paramJs.chapter = cc
              paramJs.pages = r
              DownloadJS(paramJs)
                .then(rd => {
                  // finish complete
                  if (rd.length > 0) {
                    fc.error = 'Error some pages failed download'
                    fc.fail_page = rd
                    p.fail_chap.push(fc)
                    // add error retry list files
                  }
                })
                .catch(e => {
                  // error download all pages / single chapter
                  fc.error = e
                  p.fail_chap.push(fc)
                })

              console.log(r)
            })
            .catch(e => {
              // error getting list pages
              fc.error = e
              p.fail_chap.push(fc)
            })
        } // endloop

        // end return pageresult
        res(p)
      })
    }
  }
}
